package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	os.Mkdir("uploads", os.ModeDir)

	r.POST("/", func(ctx *gin.Context) {
		file, err := ctx.FormFile("upload")
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return
		}
		err = ctx.SaveUploadedFile(file, "uploads/"+file.Filename)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		defer removeLater(file.Filename)
		log.Println(file.Filename)
		ctx.Status(http.StatusOK)
	})

	r.SetTrustedProxies(nil)
	log.Fatal(r.Run(":8080"))
}

func removeLater(filename string) {
	go func() {
		timer := time.NewTimer(5 * time.Minute)
		<-timer.C
		timer.Stop()
		os.Remove(filename)
	}()
}
