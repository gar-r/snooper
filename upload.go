package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func handleUpload(ctx *gin.Context) {

	file, err := ctx.FormFile("upload")
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	filename := path(file.Filename)
	err = ctx.SaveUploadedFile(file, filename)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	removeLater(filename)
	ctx.Status(http.StatusOK)
}

func removeLater(filename string) {
	go func() {
		timer := time.NewTimer(cleanup)
		<-timer.C
		timer.Stop()
		log.Println(os.Remove(filename))
	}()
}
