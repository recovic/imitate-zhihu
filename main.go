package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"imitate-zhihu/cache"
	"imitate-zhihu/controller"
	"imitate-zhihu/middleware"
	"imitate-zhihu/tool"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)


func main() {
	if pid := syscall.Getpid(); pid != 1 {
		ioutil.WriteFile("server.pid", []byte(strconv.Itoa(pid)), 0777)
		defer os.Remove("server.pid")
	}

	tool.InitConfig("./config")
	tool.InitLogger()
	tool.InitDatabase("zhihu")
	tool.InitRedis()

	c := cron.New(cron.WithSeconds())
	c.AddFunc(tool.Cfg.RedisSyncTime, cache.SyncCount)
	c.AddFunc(tool.Cfg.RedisSyncTime, cache.SyncAnswerVote)
	c.Start()

	gin.SetMode(tool.Cfg.Mode)
	engine := gin.Default()

	if tool.Cfg.LogFile != "" {
		engine.Use(middleware.LoggerToFile)
	}

	engine.Use(middleware.CORSMiddleware)

	controller.RouteQuestionController(engine)
	controller.RouteUserController(engine)
	controller.RouteAnswerController(engine)
	controller.RouteVoteController(engine)
	controller.RouteHotController(engine)

	err := engine.Run(":" + tool.Cfg.Port)
	if err != nil {
		panic(err.Error())
	}

}
