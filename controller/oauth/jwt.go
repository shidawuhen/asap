package oauth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func JWT(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")

	tokenString := creatJWT(mySigningKey)
	checkJWT(tokenString, mySigningKey)
	c.String(http.StatusOK, "ok")
}

func checkJWT(tokenString string, mySigningKey []byte) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	switch {
	case token.Valid:
		fmt.Printf("正确的token %+v", token)
	case errors.Is(err, jwt.ErrTokenMalformed):
		fmt.Println("That's not even a token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		fmt.Println("Invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	default:
		fmt.Println("Couldn't handle this token:", err)
	}
}

func creatJWT(mySigningKey []byte) string {
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		"hello",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "asap",
			Subject:   "com",
			ID:        "1",
			Audience:  []string{"pzq"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(ss, err)
	return ss
}
