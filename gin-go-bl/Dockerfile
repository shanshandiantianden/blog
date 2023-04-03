FROM alpine:latest
ENV VERSION 1.0
MAINTAINER realy
# 设置编码
ENV LANG C.UTF-8


#RUN apk update \
#    && apk add tzdata \
#    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /docker_go
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o docker_go

# 暴露端口
EXPOSE 8088

CMD ["./docker_go"]


