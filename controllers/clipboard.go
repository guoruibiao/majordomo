package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/majordomo/models/data/elasticsearch"
	"net/http"
)

func SearchClipboard(ctx *gin.Context) {
	content := ctx.DefaultQuery("content", "")
	ext     := ctx.DefaultQuery("ext", "")

	list, err := elasticsearch.Search(content, ext)
	res := map[string]interface{}{
		"data": list,
		"err": err,
	}

	ctx.JSON(http.StatusOK, res)
}
