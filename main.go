package main

import (
	"go-htmx-practice/presentation"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("presentation/templates/*")

	// ハンドラーの初期化
	handler := presentation.NewHandler()

	// ルーティング設定
	handler.SetupRoutes(router)

	router.Run(":8080")
}
