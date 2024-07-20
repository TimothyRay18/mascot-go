package controller

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("zxcvbnm")
var tokenName = "token"

type Claims struct {
	Username string `json:"username"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}

func GetUsernameFromToken(context *fiber.Ctx) string {
	if cookie := context.Cookies(tokenName); cookie != "" {
		accessToken := cookie
		accessClaims := &Claims{}
		fmt.Println("existing accessToken")
		fmt.Println(tokenName)
		fmt.Println(accessToken)
		fmt.Println("XXX")
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		fmt.Println("AAA")
		if err == nil && parsedToken.Valid {
			fmt.Println("BBB")
			fmt.Println("Access token: ", accessClaims.Username)
			fmt.Println("CCC")
			return accessClaims.Username
		}
	}
	return ""
}

func GenerateToken(context *fiber.Ctx, username string, userType int) {
	tokenExpiryTime := time.Now().Add(20 * time.Minute)

	claims := &Claims{
		Username: username,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}
	fmt.Println(tokenName)
	fmt.Println(signedToken)

	cookie := &fiber.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HTTPOnly: true,
	}

	// Set the cookie
	context.Cookie(cookie)
}

func ResetUserToken(context *fiber.Ctx) {

	cookie := &fiber.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HTTPOnly: true,
	}

	context.Cookie(cookie)
}
