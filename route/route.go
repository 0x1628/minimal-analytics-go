package route

import (
	"github.com/fragment0/minimal-analytics-go/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(str string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: false,
		AllowHeaders:     []string{"content-type"},
	}))

	r.POST("/:namespace/events", controllers.Events)
	r.POST("/:namespace/register", controllers.Register)
	// r.POST("/:namespace/batch", controllers.BatchLog)

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(200, "OK")
	})
}
