package route
import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"todolist/controller"
	"todolist/db"
	"todolist/model"
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
		//代办事项

		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo model.Todo
			c.BindJSON(&todo)
			if err := db.DB.Create(&todo).Error;err != nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else {
				 c.JSON(http.StatusOK,todo)
			}
		})

		//查看所有代办事项
		v1Group.GET("/todo",func(c *gin.Context){
			//查询todo表的所有数据
			var todoList []model.Todo
			if err:=db.DB.Find(&todoList).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else{
				c.JSON(http.StatusOK,todoList)
			}
		})

		//修改某个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id,_ := c.Params.Get("id")
			var todo model.Todo
			if err := db.DB.Where("id=?",id).First(&todo).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
				return
			}

			c.BindJSON(&todo)
			if err := db.DB.Save(&todo).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else {
				c.JSON(http.StatusOK,todo)
			}
		})

		//删除某个代办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id,ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK,gin.H{"error":"无效id"})
			}
			if err:=db.DB.Where("id=?",id).Delete(model.Todo{}).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else {
				c.JSON(http.StatusOK,gin.H{id:"deleted"})
			}
		})
	}

	err := R.Run(":8181")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}



