FROM scratch

# 将运行程序和日志复制到容器中
COPY log test ./

# 声明服务端口
EXPOSE 9090

#加权限
RUN chmod +x ./test

# 启动容器时运行的命令

CMD ["/test"]