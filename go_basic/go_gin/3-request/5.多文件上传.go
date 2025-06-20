package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)

			return
		}
		for _, headers := range form.File {
			for _, header := range headers {
				c.SaveUploadedFile(header, "uploads/"+header.Filename)
			}
		}
	})
	r.Run(":8080")
}
