package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

type AuthApiKey struct {
	// Set the prefix size. The default is 4
	PrefixSize uint `json:"prefix_size"`

	// Set the salt size. The default is 16
	SaltSize uint `json:"salt_size"`
}

// Generates a api key in the following format: "prefix.hash.salt"
// with a random prefix and salt
func (auth *AuthApiKey) Generate(value string) (string, error) {
	prefixSize := auth.PrefixSize
	saltSize := auth.SaltSize

	if reflect.ValueOf(auth.PrefixSize).IsZero() {
		prefixSize = 4
	}

	if reflect.ValueOf(auth.SaltSize).IsZero() {
		saltSize = 16
	}

	prefix, err := genSalt(prefixSize)
	if err != nil {
		return "", errors.Wrap(
			errors.Wrap(err, "error on generate prefix"),
			"error on generate api key",
		)
	}

	salt, err := genSalt(saltSize)
	if err != nil {
		return "", errors.Wrap(err, "error on generate api key")
	}

	key := genKey(prefix, value, salt)
	return key, nil
}

func (auth *AuthApiKey) Verify(value, key string) bool {
	peaces := strings.Split(key, ".")
	prefix := peaces[0]
	salt := peaces[2]

	newKey := genKey(prefix, value, salt)
	return key == newKey
}

func genSalt(size uint) (string, error) {
	salt := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", errors.Wrap(err, "error on generate salt")
	}

	return hex.EncodeToString(salt), nil
}

func genKey(prefix, value, salt string) string {
	key := prefix + value + salt

	sha := sha256.New()
	sha.Write([]byte(key))

	hash := hex.EncodeToString(sha.Sum(nil))
	return fmt.Sprint(prefix, ".", hash, ".", salt)
}
