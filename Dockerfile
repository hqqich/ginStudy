# 引入基础运行环境
FROM alpine:latest
#FROM ubuntu:18.04

# 设置维护者信息
LABEL org.opencontainers.image.authors="hqqich1314@outlook.com"

# 设置工作目录：/app
WORKDIR /app

# 卷标挂载
VOLUME ~/files /app/static

# 将本机文件移动到docker工作目录
COPY ./ginStudy-linux-amd64-64 /app
COPY ./config.ini /app

# 修改可执行文件权限
RUN chmod 777 /app/ginStudy-linux-amd64-64

# 指定端口，docker对外的端口
EXPOSE 8888

# 设置工作目录：/app
WORKDIR /app

#需要运行的命令
CMD ["/app/ginStudy-linux-amd64-64"]