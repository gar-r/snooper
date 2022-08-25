package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/", func(ctx *gin.Context) {
		file, err := ctx.FormFile("upload")
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return
		}
		err = ctx.SaveUploadedFile(file, file.Filename)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		defer os.Remove(file.Filename)
		data, err := os.ReadFile(file.Filename)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		log.Println(string(data))
		ctx.Status(http.StatusOK)
	})

	r.SetTrustedProxies(nil)
	log.Fatal(r.Run(":8080"))

}
