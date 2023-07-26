package middlware

import "fmt"
import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 打印请求信息
		fmt.Printf("[%s] %s %s\n", c.Request.Method, c.Request.URL.Path, c.ClientIP())
		// 继续处理请求
		c.Next()
	}
}
