package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AccessToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		// Expires in 5 minutes
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	// Add Signature
	ss, err := token.SignedString([]byte("==signature=="))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}
