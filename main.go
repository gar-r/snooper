package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

const uploadDir = "uploads"
const cleanup = 5 * time.Minute

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
