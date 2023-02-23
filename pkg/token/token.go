package token

import "github.com/dgrijalva/jwt-go"

type MyCustomClaims struct {
	StudentID int64 `json:"student"`
	jwt.StandardClaims
}

func ParseToken(token string) (*MyCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("drifting"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
