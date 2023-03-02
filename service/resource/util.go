package resource

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap/buffer"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var lettersAndDigits = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomStrByLetterAndDigit(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = lettersAndDigits[r.Intn(len(lettersAndDigits))]
	}
	return string(b)
}

func SHA256(s string) string {
	sha := sha256.New()
	sha.Write([]byte(s))
	return hex.EncodeToString(sha.Sum(nil))
}

func GeneralMFA() (string, error) {
	var buf buffer.Buffer
	err := binary.Write(&buf, binary.BigEndian, time.Now().UnixNano()/1000/30)
	if err != nil {
		return "", err
	}
	h := hmac.New(sha1.New, buf.Bytes())
	code := strings.ToUpper(base32.StdEncoding.EncodeToString(h.Sum(nil)))
	return code, nil
}

// String2Int string数组转int数组
func String2Int(strArr []string) []int {
	if strArr == nil {
		return nil
	}
	res := make([]int, len(strArr))
	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}
	return res
}

// RemoveDuplicateElement 去重
func RemoveDuplicateElement[T int | int64 | string | float32 | float64](arr []T) []T {
	if arr == nil {
		return nil
	}
	temp := make(map[string]bool)
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		key := fmt.Sprint(v)
		if _, ok := temp[key]; !ok {
			temp[key] = true
			result = append(result, v)
		}
	}
	return result
}
