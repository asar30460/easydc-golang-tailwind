package util

import (
	"os"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func ParseJWT(ctx *gin.Context)(jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	tokenString, err := ctx.Cookie("jwt")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "無效的 JWT token"})
        return nil, err
    }
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't parse claims"})
		return nil, err
	}

	return claims, nil
}