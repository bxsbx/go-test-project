package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type Test struct {
}

func (t *Test) Test1(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.JSON(http.StatusCreated, "gin 1")
	}()
	go func() {
		defer wg.Done()
		c.JSON(http.StatusAccepted, "gin 2")
	}()
	wg.Wait()
	c.JSON(http.StatusOK, "gin ok")
}
