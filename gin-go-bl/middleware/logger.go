package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Loger() gin.HandlerFunc {
	filePath := "log/log.log"
	linkname := "latest_log.log"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("文件打开失败", err)
	} //及时关闭file句柄

	log := logrus.New()
	log.Out = file
	log.SetLevel(logrus.DebugLevel)
	logWrite, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(linkname),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWrite,
		logrus.FatalLevel: logWrite,
		logrus.DebugLevel: logWrite,
		logrus.WarnLevel:  logWrite,
		logrus.ErrorLevel: logWrite,
		logrus.PanicLevel: logWrite,
		logrus.TraceLevel: logWrite,
	}
	hock := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000000000",
	})
	log.AddHook(hock)

	return func(c *gin.Context) {
		//startime, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05.000000000"), time.Local)
		startTime := time.Now()

		c.Next()

		endTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(endTime.Nanoseconds())/1000000.0)))
		host, err := os.Hostname()
		if err != nil {
			host = "unknown"
		}
		status := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := log.WithFields(logrus.Fields{
			"host":      host,
			"status":    status,
			"spendTime": spendTime,
			"clientIp":  clientIp,
			"method":    method,
			"userAgent": userAgent,
			"path":      path,
			"dataSize":  dataSize,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if status >= 500 {
			entry.Error()
		} else if status >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
