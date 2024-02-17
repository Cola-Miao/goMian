package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"offerBook/config/inner"
	"time"
)

type Claim struct {
	UID int
	*jwt.RegisteredClaims
}

func GenerateJWT(id int) (string, error) {
	claim := &Claim{
		UID: id,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "goMain",
			Subject:   "auth",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(inner.JWTExpiresTime)),
			NotBefore: nil,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	s, err := t.SignedString(inner.JWTSecret)
	return s, err
}

func ParseJWT(s string) (*Claim, error) {
	t, err := jwt.ParseWithClaims(s, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return inner.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("jwt valid failed")
	}
	claim, ok := t.Claims.(*Claim)
	if !ok {
		return nil, errors.New("parse claim failed")
	}
	return claim, nil
}
