package utils

import (
	"fmt"
	"lexium-utility/config"
	"lexium-utility/redis_repository"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(username string, sid string) (string, error) {

	token_lifespan, err := strconv.Atoi(config.TOKEN_HOUR_LIFESPAN)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["session_id"] = sid
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.API_SECRET))

}

func TokenValid(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.API_SECRET), nil
	})
	if err != nil {
		return &jwt.Token{}, err
	}
	return token, nil
}
func TokenValidate(c *gin.Context) (*jwt.Token, error) {
	token, err := TokenValid(c)
	if err != nil {
		return &jwt.Token{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		sessionId := claims["session_id"].(string)
		v := claims["username"].(string)
		sid, err := redis_repository.GetSessionId(v)
		if err != nil {
			return &jwt.Token{}, err
		}
		if sessionId == sid {
			return token, nil
		} else {
			return &jwt.Token{}, fmt.Errorf("session expired")
		}
	} else {
		return &jwt.Token{}, fmt.Errorf("error in geting claims")
	}
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenUsername(c *gin.Context) (string, error) {
	token, err := TokenValid(c)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var v string = (claims["username"]).(string)
		return v, nil
	}
	return "", nil
}
func ExtractTokenSessionID(c *gin.Context) (string, error) {
	token, err := TokenValid(c)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var v string = (claims["session_id"]).(string)
		return v, nil
	}
	return "", nil
}
