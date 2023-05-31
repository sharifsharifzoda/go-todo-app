package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) tokenAuthMiddleware(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not found any cookie",
		})
		return
	}

	userId, err := h.Auth.ParseToken(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Set("userId", userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "userId not found",
		})
		return 0, errors.New("userId not found")
	}

	idInt, ok := id.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid type of userId",
		})
		return 0, errors.New("invalid type of userId")
	}

	return idInt, nil
}
