package router


import (
	"github/MatiasCermak/go-meli-integration/pkg/controller"
	"github.com/gin-gonic/gin"
)

func Run(){
	router := gin.Default()
	router.GET("/auth", controller.Auth)
	router.GET("/items/all", controller.ItemsAll)
	router.POST("/items/publish", controller.PublishItem)
	router.POST("/items/questions/ans", controller.AnswerQuestion)
	router.Run()
}
