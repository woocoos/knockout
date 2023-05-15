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
	randomBytes := make([]byte, 20)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	secret := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
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

// DiffArrays 查找两数组新增及删除的元素: a:新数组  b:旧数组
func DiffArrays[T int | int64 | string | float32 | float64](a []T, b []T) ([]T, []T) {
	var added, removed []T
	hash := make(map[T]bool)
	for _, num := range a {
		hash[num] = true
	}
	// 删除的元素
	for _, num := range b {
		if _, ok := hash[num]; !ok {
			removed = append(removed, num)
		} else {
			// 如果元素存在，则删除该元素
			delete(hash, num)
		}
	}
	// 新增的元素
	for num := range hash {
		added = append(added, num)
	}
	return added, removed
}
