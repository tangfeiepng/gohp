FROM golang:1.18

WORKDIR /var/www/html

RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY start-container /usr/local/bin/start-container
RUN chmod -R 777 /usr/local/bin/start-container && cd /usr/local/bin
#启动脚本文件
ENTRYPOINT ["start-container"]