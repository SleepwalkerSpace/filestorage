package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FileUpload 文件上传
func (srv *Server) FileUpload(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload"]

	for _, file := range files {
		log.Println(file.Filename)

		// 上传文件至指定目录
		c.SaveUploadedFile(file, srv.config.StoragePath)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
