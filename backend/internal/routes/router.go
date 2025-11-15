package routes

import (
	"github.com/AnthonySJHenry/desktop-tool-gui/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/health", handlers.Health)
	return r
}
