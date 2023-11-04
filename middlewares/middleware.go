package middlewares

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/manishbadgotra/authentication/helpers"
)

var TokenString string

func JwtSecret(email, id string) (string, error) {
	// expire time for jwt
	expTime := time.Now().Add(time.Hour * 12).Unix()

	// creating jwt with email and expire time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":   email,
		"Id":      id,
		"expTime": expTime,
	})

	TokenString, err := token.SignedString([]byte(helpers.GetString("JwtSecretKey")))
	if err != nil {
		return "", err
	}
	return TokenString, nil

}

func ValidateToken(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT token missing"})
		c.Abort()
		return
	}

	type Claims struct {
		jwt.StandardClaims
	}
	var err error
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.GetString("JwtSecretKey")), err
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT token missing"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		expTime := time.Unix(claims.ExpiresAt, 0)
		currTime := time.Now()
		if currTime.Compare(expTime) >= 0 {
			c.Next()
		}
	}

	if err != nil && !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Jwt Token is not Valid."})
		c.Abort()
	}
}

func GenerateResetToken(username string) (string, error) {
	// expire time for jwt
	expTime := time.Now().Add(time.Minute * 20).Unix()

	// creating jwt with email and 30 minute expiry time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":   username,
		"expTime": expTime,
	})

	TokenString, err := token.SignedString([]byte(helpers.GetString("JwtSecretKey")))
	if err != nil {
		return "", err
	}
	return TokenString, nil
}
