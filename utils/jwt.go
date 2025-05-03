package utils

import(
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtkey = []byte("your-secret-key")

type Claims struct{
	Email string `json:"email"`
	jwt.RegisteredClaims
}


func GenerateJWT(email string) (string, error){
	expirationTime:= time.Now().Add(15 * time.Minute)
	claims:= &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpirsAt: jwt.NewNumericDate(expirationTime)
		},
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)
}
