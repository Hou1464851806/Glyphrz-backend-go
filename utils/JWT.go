package utils

import (
	"Glyphrz-go/global"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type Claims struct {
	ID   uint
	Name string
	Role uint
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(global.Settings.JWTKey.SigningKey)}
}

func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
