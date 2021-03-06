package controller

import (
	"github.com/gin-gonic/gin"
	"imitate-zhihu/dto"
	"imitate-zhihu/service"
	"imitate-zhihu/tool"
	"net/http"
)

func RouteHotController(engine *gin.Engine) {
	group := engine.Group("/hot")
	group.GET("", GetHotQuestions)
}

func GetHotQuestions(c *gin.Context) {
	cursor, err := tool.StrToInt(c.Query("cursor"))
	if err != nil {
		cursor = 0
	}
	size, err := tool.StrToInt(c.DefaultQuery("size", "10"))
	if err != nil {
		size = 10
	}
	q, res := service.GetHotQuestions(cursor, size)
	if q == nil {
		q = []dto.HotQuestionDto{}
	}
	c.JSON(http.StatusOK, res.WithData(gin.H{
		"next_cursor": tool.IntToStr(cursor + size),
		"questions": q,
	}))
}
