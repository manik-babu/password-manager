package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JwtSecretKey         = "h82bslk6hvGHs78mwWof7"
	defaultTokenDuration = 24 * time.Hour
)

type JWTClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(id uint, email string) (string, error)
	ValidateToken(jwtToken string) (*JWTClaims, error)
}

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTService(secretKey string) JWTService {
	if secretKey == "" {
		secretKey = JwtSecretKey
	}
	return &jwtService{
		secretKey:     secretKey,
		tokenDuration: defaultTokenDuration,
	}
}
func (s *jwtService) GenerateToken(id uint, email string) (string, error) {
	claims := JWTClaims{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "pwd-manager",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
func (s *jwtService) ValidateToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid singing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Invalid singing method: %v", token.Header["alg"])
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
