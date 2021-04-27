package main

import (
	"flag"
	"net/http"

	"docker-golang-demo/config"

	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
)

var (
	tomlFile = flag.String("config", "test.toml", "config file")
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "docker test",
	})
}

func main() {
	flag.Parse()
	// 解析配置文件
	tomlConfig, err := config.UnmarshalConfig(*tomlFile)
	if err != nil {
		log.Errorf("UnmarshalConfig: err:%v\n", err)
		return
	}
	router := gin.New()
	router.GET("/", indexHandler)
	router.Run(tomlConfig.GetListenAddr())
}
