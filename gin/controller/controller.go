package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct {
}

type Test1Params struct {
	UserId int `form:"user_id" binding:"required"`
}

func (t *Test) Test1(c *gin.Context) {
	var params Test1Params
	if err := c.ShouldBind(&params); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, "gin ok")
}
