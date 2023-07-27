package route
import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"todolist/controller"
)

var Name = "api ------------------ route-----------------------init"

func init()  {
	fmt.Println(Name)
}

func ApiInit()  {
	R := gin.Default()
	//S.Use(middlware.Logger())	//注册日志中间件

	//要渲染模板----告诉gin模板文件存放位置
	R.LoadHTMLGlob("templates/*")
	R.Static("/static","./static")
	R.GET("/", func(c *gin.Context) {
		//渲染index.html
		c.HTML(http.StatusOK,"index.html",gin.H{"title":"清单"})
	})


	R.GET("/users", controller.Users)
	R.GET("/usersIndex", controller.UserIndex)
	// 注册日志中间件

	//设定路由组
	v1Group := R.Group("v1")
	{
		//添加
		v1Group.POST("/todo", controller.AddTodo)

		//查看所有代办事项
		v1Group.GET("/todo",controller.GetTodo)

		//修改某个待办事项
		v1Group.PUT("/todo/:id", controller.EditTodo)

		//删除某个代办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	err := R.Run(":8181")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}



