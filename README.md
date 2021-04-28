#使用docker部署一个带配置文件的golang项目

  
    

## 首先看下我这的目录结构
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190418152549676.png)  

我这的gopath为 gowork目录  

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190418152606107.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FmMjI1MTMzNw==,size_16,color_FFFFFF,t_70)

## 编写dockerfile

首先编译main.go 生成二进制文件，该二进制文件可以直接在相应的linux服务器下运行。
我这里使用如下指令，编译后会多出一个main文件

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

可以根据自己需要编译的平台更改
使用from加入母镜像 这里使用scratch空镜像，因为编译后的main是可以直接运行的

```
FROM scratch
```
MAINTAINER指定维护者信息  
WORKDIR .这里使用当前目录作为工作目录，可以修改  
将main 与 test.toml 配置文件 放入当前目录  
EXPOSE 这是对方开发的端口，可修改，我这使用8082  
CMD 附带配置文件test.toml 运行当前目录下的main  

```
MAINTAINER  "hcf"

WORKDIR .
ADD main .
ADD test.toml .

EXPOSE 8082
CMD ["./main","-config=./test.toml"]
```

容器的配置就完成了  

## 生成镜像

在dockerfile目录运行
```
docker build -t dockertest .
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190418154534753.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FmMjI1MTMzNw==,size_16,color_FFFFFF,t_70)
successfully built 构建成功

使用
```
docker images
```
查看刚生成的镜像
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190418154734211.png)
运行镜像  
```
docker run -p 8082:8082 dockertest
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190418163658573.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190418163706270.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FmMjI1MTMzNw==,size_16,color_FFFFFF,t_70)
成功运行，页面输入http://localhost:8082/  成功访问

如果需要容器后台运行,指令加入 -d 就行了
```
docker run -p 8082:8082 -d dockertest
```


更新后的dockerfile 使用多阶段构建 加入go mod,在docker镜像制作时进行编译
```
FROM golang:1.16 as build

ENV GOPROXY https://goproxy.cn/
ENV GO111MODULE on
WORKDIR /go/src/docker-golang-demo
COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM scratch as prod
COPY --from=build /go/src/docker-golang-demo/main /
COPY --from=build /go/src/docker-golang-demo/test.toml /
EXPOSE 8082
CMD ["./main","-config=./test.toml"]
```