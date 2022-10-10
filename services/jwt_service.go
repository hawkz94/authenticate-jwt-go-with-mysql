package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	secret_key := os.Getenv("APP_SECRET_KEY")
	return &jwtService{
		secretKey: secret_key,
		issure:    "tecmentor-api",
	}
}

type Clain struct {
	sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	clain := &Clain{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)

	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token! %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}
