package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateTokenService(NasabahID string) string
	ValidateTokenService(token string) (*jwt.Token, error)
	RefreshTokenService(token string) (string, error)
}

type jwtCustomClaim struct {
	NasabahID string `json:"nasabah_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() JwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "rahmattullah",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "rahmattullah"
	}
	return secretKey
}

func (j *jwtService) GenerateTokenService(NasabahID string) string {
	claims := &jwtCustomClaim{
		NasabahID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateTokenService(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error signing method: %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) RefreshTokenService(token string) (string, error) {
	t, err := j.ValidateTokenService(token)
	if err != nil {
		return "", err
	}
	claims, ok := t.Claims.(*jwtCustomClaim)
	if !ok {
		return "", fmt.Errorf("unexpected claims type")
	}
	claims.ExpiresAt = time.Now().AddDate(1, 0, 0).Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := newToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
