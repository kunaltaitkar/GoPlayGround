package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.RouterGroup) {
	g.POST("/shutdown", shutdown)
}

func shutdown(c *gin.Context) {
	fmt.Println("QUITE CHanal value", Name)
}
