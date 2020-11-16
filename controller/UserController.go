package controller

import (
	"github.com/gin-gonic/gin"
	"imitate-zhihu/dto"
	"imitate-zhihu/result"
	"imitate-zhihu/service"
	"net/http"
)

func RouteUserController(engine *gin.Engine) {
	group := engine.Group("/user")
	group.POST("/login", UserLogin)
	group.POST("/register", UserRegister)
}


func UserLogin(c *gin.Context) {
	userDto := dto.UserLoginDto{}
	err := c.BindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ShowBadRequest(err.Error()))
		return
	}
	res := service.UserLogin(&userDto)
	c.JSON(http.StatusOK, res.Show())
}


func UserRegister(c *gin.Context) {
	registerDto := dto.UserRegisterDto{}
	err := c.BindJSON(&registerDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ShowBadRequest(err.Error()))
		return
	}
	res := service.UserRegister(&registerDto)
	c.JSON(http.StatusOK, res.Show())
}