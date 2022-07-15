package hashing

import (
	"crypto/md5"
	"fmt"
)

func HashURLAddr(url string) string {
	result := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	return result[:6]
}
