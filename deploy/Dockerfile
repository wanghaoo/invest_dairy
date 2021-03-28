FROM alpine:3.6
RUN echo $'http://mirrors.aliyun.com/alpine/v3.6/main\n\
http://mirrors.aliyun.com/alpine/v3.6/community' > /etc/apk/repositories
RUN apk add --update ca-certificates
RUN update-ca-certificates
RUN apk add --update tzdata
RUN apk add --update curl
ENV TZ=Asia/Shanghai
RUN mkdir -p /usr/local/go/lib/time/
COPY zoneinfo.zip /usr/local/go/lib/time/
RUN  mkdir /swagger
COPY swagger /swagger
COPY invest-dairy /
CMD ["/invest-dairy"]
