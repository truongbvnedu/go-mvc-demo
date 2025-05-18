package routes

import (
	controllers "go-mvc-demo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	postCtl := controllers.NewPostController()
	commentCtl := controllers.NewCommentController()

	r.POST("/posts", postCtl.CreatePost)
	r.GET("/posts", postCtl.GetPosts)
	r.PUT("/posts/:id", postCtl.UpdatePost)
	r.DELETE("posts/:id", postCtl.DeletePost)
	r.GET("/posts/:id", postCtl.GetOne)

	r.POST("/comments", commentCtl.CreateComment)
	r.PUT("/comments/:id", commentCtl.UpdateComment)
	r.DELETE("/comments/:id", commentCtl.DeleteComment)

	return r
}
