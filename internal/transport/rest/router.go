package transport

import (
	servies "miras/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Gin     *gin.Engine
	Service *servies.Service
}

func NewHandler(service *servies.Service) *Handler {
	return &Handler{
		Gin:     gin.Default(),
		Service: service,
	}
}

func (h *Handler) Router() {

	h.Gin.POST("/sign-in", h.signIn())
	h.Gin.POST("/sign-up", h.signUp())
	h.Gin.POST("/product/create", h.createProduct())
	h.Gin.POST("/product/update", h.updateProduct())

}
