package initializers

import (
	"github.com/gofiber/fiber/v3/extractors"

	jwtWare "github.com/gofiber/contrib/v3/jwt"
)

var JWTConfig *jwtWare.Config

func SetUpJwt() {
	JWTConfig = &jwtWare.Config{
		SigningKey: jwtWare.SigningKey{Key: []byte("secret")},
		Extractor:  extractors.FromAuthHeader("Bearer"),
	}
}
