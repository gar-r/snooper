package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const uploadDir = "uploads"
const cleanup = 6 * time.Hour

func main() {

	err := os.MkdirAll(uploadDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/:file", handleDownload)
	r.POST("/", handleUpload)

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(r.Run(":8080"))
}

func path(name string) string {
	return fmt.Sprintf("%s/%s", uploadDir, name)
}
