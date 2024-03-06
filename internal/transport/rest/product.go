package transport

import (
	"encoding/json"
	"fmt"
	"miras/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct() gin.HandlerFunc {

	return func(c *gin.Context) {

		cookie, err := c.Request.Cookie("Token")

		if err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}

		var product models.Product

		err = json.NewDecoder(c.Request.Body).Decode(&product)

		if err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}

		err = h.Service.ProductService.CreateProduct(product, cookie.Value)
		if err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}

		c.JSON(200, gin.H{"status": "product successfuly created"})
	}
}

func (h *Handler) updateProduct() gin.HandlerFunc {

	return func(c *gin.Context) {

		cookie, err := c.Request.Cookie("Token")

		if err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}

		var product models.Product

		err = json.NewDecoder(c.Request.Body).Decode(&product)

		if err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}

		idStr := c.Request.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Wrong id wormat"})
			return
		}
		err = h.Service.ProductService.UpdateProduct(product, cookie.Value, id)
		if err != nil {
			c.JSON(400, gin.H{"error": fmt.Sprintf("Bad request %s", err.Error())})
			return
		}

		c.JSON(200, gin.H{"status": "product successfuly updated"})
	}
}
