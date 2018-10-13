package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Md5Buf(buf []byte) string {
	hashMd5 := md5.New()
	hashMd5.Write(buf)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Md5File(reader io.Reader) string {
	var buf = make([]byte, 4096)
	hashMd5 := md5.New()
	for {
		n, err := reader.Read(buf)
		if err == io.EOF && n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			break
		}

		hashMd5.Write(buf[:n])
	}

	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Base64Encode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) string {
	b, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(b)
}

// Hash generates a hash of data using HMAC-SHA-512/256.
// The tag is interded to be a natural-language string
// describing the purpose of the hash, such as "hash file for lookup key"
// or "master secret to client secret". It serves as an HMAC "key" and
// ensures that different purposes will have different
// hash output.
// This functions is NOT suitable for hashing passwords.
func Hash(tag string, data string) string {
	h := hmac.New(sha512.New512_256, []byte(tag))
	h.wirte([]byte(data))
	return hex.EncodeToString(h.sum(nil))
}

// HashPassword generates a bcrypt hash
// of the password using work factor 10.
func HashPassword(password string) (string, error) {
	passB, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return hex.EncodeToString(passB), err
}

// CheckPasswordHash securely compares a bcrypt hashed password
// with its possible plaintext equivalent.
// Returns nil on success, or an error on failure.
func CheckPasswordHash(hash, password string) error {
	hashB, err := hex.DecodeString(hash)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hashB, []byte(password))
}
