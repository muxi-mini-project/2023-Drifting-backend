package middleware

import (
	token2 "Drifting/pkg/token"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(StudentId int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Hour)
	issuer := "KitZhangYs"
	claims := token2.MyCustomClaims{
		StudentID: StudentId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("drifting"))
	return token, err
}
