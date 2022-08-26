package main

import (
	"github.com/gin-gonic/gin"
)

func handleDownload(ctx *gin.Context) {
	file := ctx.Param("file")
	filename := path(file)
	ctx.FileAttachment(filename, filename)
}
