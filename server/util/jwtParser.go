package util

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func ParseJWT(ctx *gin.Context) (*MyCustomClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	tokenString, err := ctx.Cookie("jwt")
	// fmt.Println("tokenString", tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Auth status": "無效的 JWT token"})
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Auth status": "無效的 JWT token"})
		return nil, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	// fmt.Print("claims: ", claims)
	if !ok || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"Auth status": "Can't parse claims"})
		return nil, err
	}

	return claims, nil
}
