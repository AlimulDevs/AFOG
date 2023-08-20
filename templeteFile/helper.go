package templetefile

const HashPasswordTemplete = `package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
`

const MakeUuidTemplete = `package helper

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.NewString()
	uuid := strings.ReplaceAll(id, "-", "")
	return uuid
}
`

const UtilTemplete = `package helper

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	return viper.GetString(key)
}
`
