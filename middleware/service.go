package middleware

import (
	"LuxSpace/app/v1/user"
	"LuxSpace/auth"
	"LuxSpace/helper"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	authService auth.Service
	userService user.Service
}

func NewAuthMiddleware(authService auth.Service, userService user.Service) *authMiddleware {
	return &authMiddleware{authService, userService}
}

func (h *authMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil token di header
		authHeader := c.GetHeader("Authorization")
		// cek konten tokennya, jika tidak kirim respon unauthorized
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ambil token
		stringToken := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			// set token ke var string
			stringToken = arrayToken[1]
		}

		// validasi token yang dikirim dari client
		token, err := h.authService.ValidateToken(stringToken)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ubah token ke jwt mapClaims
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ambil user_id di dapat dari client, lalu ubah ke dalam int
		userID := int(claim["user_id"].(float64))

		// komparasi id user yang terdaftar di databse
		newUser, err := h.userService.GetDetailUser(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// mapping datanya ke dalam struct
		userData := user.FormatUserHeader{
			ID:       newUser.ID,
			Email:    newUser.Email,
			Username: newUser.Username,
			Role:     newUser.Role,
		}
		// simpan ke request
		c.Set("currentUser", userData)
	}
}
