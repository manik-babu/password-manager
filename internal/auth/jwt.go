package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JwtSecretKey         = "sldneixowmbgi"
	defaultTokenDuration = 24 * time.Hour
)

type JWTClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(id uint, email string) (string, error)
	// ValidateToken(jwtToken string) (*JWTClaims, error)
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
