package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/majordomo/controllers"
)
func Init(engine *gin.Engine) {
  
    indexRouter := engine.Group("/")
    {
        indexRouter.GET("/", controllers.HomeIndex)
    }

    clipboardRouter := engine.Group("/clipboard")
    {
    	clipboardRouter.GET("/search", controllers.SearchClipboard)
	}
}
