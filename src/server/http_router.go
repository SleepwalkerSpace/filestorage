package server

import (
	"github.com/gin-gonic/gin"
)

func (srv *Server) router(root string, list bool) *gin.Engine {
	router := gin.Default()
	router.StaticFS("/static", gin.Dir(root, list))

	v1 := router.Group("/v1/api", Cors())
	{
		file := v1.Group("/file")
		{
			file.POST("/upload", srv.FileUpload)
		}

	}
	return router
}
