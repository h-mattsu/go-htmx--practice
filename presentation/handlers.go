package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler はHTTPハンドラーの構造体
type Handler struct {
	// サービス層の依存関係をここに追加
	// 例: userService *services.UserService
}

// NewHandler はHandlerの新しいインスタンスを作成
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes はルーティング設定を行う
func (h *Handler) SetupRoutes(router *gin.Engine) {
	router.GET("/index", h.GetIndex)
	router.GET("/", h.GetHome)
}

// GetIndex はインデックスページを返す
func (h *Handler) GetIndex(c *gin.Context) {
	// サービス層を呼び出してデータを取得する予定
	// 例: data, err := h.userService.GetAll(c.Request.Context())

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

// GetHome はホームページを返す
func (h *Handler) GetHome(c *gin.Context) {
	// サービス層を呼び出してデータを取得する予定

	c.HTML(http.StatusOK, "pages/home", gin.H{
		"title": "Home",
	})
}
