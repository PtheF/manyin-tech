package utils

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
)

// HASH_BYTES 这个数x2是随机盐长度和加密后的密码长度
const HASH_BYTES = 32

// Encode 不带随机盐加密
func Encode(s string) (string, error) {
	return EncodeWithSalt(s, "")
}

// randomSalt 生成随机盐
func randomSalt() (string, error) {
	buf := make([]byte, HASH_BYTES)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

// EncodeWithSalt 指定随机盐加密
func EncodeWithSalt(s string, salt string) (string, error) {
	h, err := scrypt.Key([]byte(s), []byte(salt), 16384, 8, 1, HASH_BYTES)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

// EncodeWithRS 自己生成随机盐加密并返回盐
func EncodeWithRS(s string) (enc string, salt string, err error) {
	salt, _ = randomSalt()
	enc, err = EncodeWithSalt(s, salt)
	return enc, salt, err
}
