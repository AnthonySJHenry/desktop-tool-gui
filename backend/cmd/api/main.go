package main

import (
	"os"

	"github.com/AnthonySJHenry/desktop-tool-gui/backend/internal/config"
	"github.com/AnthonySJHenry/desktop-tool-gui/backend/internal/routes"
)

func main() {
	_ = config.Load()
	port := config.Port()
	if port == "" {
		port = "8080"
	}
	r := routes.NewRouter()
	r.Run(":" + port)
}
