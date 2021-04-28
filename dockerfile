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