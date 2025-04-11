package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {
		path := c.Request.URL.Path                   
		method := c.Request.Method                  

		c.Next()
		status := c.Writer.Status()             

		fmt.Printf("[%s] ~ %s Status:%v\n", method, path, status)
	}
}