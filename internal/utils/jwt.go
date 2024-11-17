package utils

import (
	"fmt"
	"time"

	"github.com/oriastanjung/grpc_server/internal/config"
	"github.com/oriastanjung/grpc_server/internal/entities"

	"github.com/golang-jwt/jwt/v5"
	"github.com/segmentio/ksuid"
)

// Buat struktur untuk klaim JWT (payload)
type JWTClaims struct {
	ID       ksuid.KSUID `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Role     string      `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(payload entities.User) (string, error) {
	cfg := config.LoadEnv()
	jwtSecret := cfg.JWTSecretKey

	// Buat claim JWT
	claims := JWTClaims{
		ID:       payload.ID,
		Username: payload.Username,
		Email:    payload.Email,
		Role:     payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	}

	// Membuat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Mengenerate token string JWT
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	// Enkripsi token string JWT dengan AES-128
	encryptedToken, err := EncryptAES(tokenString)
	if err != nil {
		return "", err
	}

	// Return token yang sudah terenkripsi
	return encryptedToken, nil
}

func VerifyTokenJWT(tokenString string) (*JWTClaims, error) {
	cfg := config.LoadEnv()
	jwtSecret := cfg.JWTSecretKey

	// Dekripsi token string yang sudah terenkripsi dengan AES-128
	decryptedToken, err := DecryptAES(tokenString)
	if err != nil {
		return nil, err
	}

	// Parsing token JWT setelah didekripsi
	token, err := jwt.ParseWithClaims(decryptedToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Memastikan bahwa algoritma yang digunakan adalah HMAC dengan SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// Validasi token dan klaim
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}