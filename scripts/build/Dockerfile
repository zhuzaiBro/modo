FROM alpine:latest
MAINTAINER zhufuyi "g.zhufuyi@gmail.com"

# set the time zone to Shanghai
RUN apk add tzdata  \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# add curl, used for http service checking, can be installed without it if deployed in k8s
RUN apk add curl

COPY configs/ /app/configs/
COPY cyy /app/cyy
RUN chmod +x /app/cyy

# http port
EXPOSE 8080


WORKDIR /app

CMD ["./cyy", "-c", "configs/cyy.yml"]
# if you use the Configuration Center, cyy.yml is changed to the Configuration Center configuration.
#CMD ["./cyy", "-c", "configs/cyy.yml", "-enable-cc"]
