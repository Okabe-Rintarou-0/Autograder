package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(input string) string {
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(input)))
	return md5String
}
