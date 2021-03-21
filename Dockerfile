FROM scratch


COPY log /log

# 把test拷贝到当前目录
ADD test /test

# 声明服务端口
EXPOSE 9090


# 需要运行的命令
ENTRYPOINT ["/test"]