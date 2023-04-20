package resource

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"math/big"
)

// RandomStr generate random string,exclude 0,i,l
func RandomStr(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	var result string
	for i := 0; i < n; i++ {
		randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		randomChar := charset[randomInt.Int64()]
		result += string(randomChar)
	}
	return result
}

func SHA256(s string) string {
	sha := sha256.New()
	sha.Write([]byte(s))
	return hex.EncodeToString(sha.Sum(nil))
}

// SaltSecret salt secret
func SaltSecret(ori, salt string) string {
	sha := sha256.New()
	sha.Write([]byte(ori))
	sha.Write([]byte(salt))
	given := hex.EncodeToString(sha.Sum(nil))
	return given
}

// GeneralMFASecret gen a secret for google authenticator
func GeneralMFASecret() string {
	randomBytes := make([]byte, 10)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	secret := base32.StdEncoding.EncodeToString(randomBytes)
	return secret
}

// RemoveSliceElement 移除数组指定元素
func RemoveSliceElement[T int | int64 | string | float32 | float64](a []T, el T) []T {
	j := 0
	for _, v := range a {
		if v != el {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// UpdateSliceElement 更新数组指定元素
func UpdateSliceElement[T int | int64 | string | float32 | float64](a []T, newEl T, oldEl T) []T {
	for i, v := range a {
		if v == oldEl {
			a[i] = newEl
		}
	}
	return a
}
