package hashing

import (
	"crypto/md5"
	"fmt"
)

func HashURLAddr(url []byte) string {
	h := md5.New()
	h.Write(url)
	result := fmt.Sprintf("%x", h.Sum(nil))
	return result[:6]
}
