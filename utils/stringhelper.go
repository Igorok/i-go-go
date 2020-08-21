package utils

import (
	"crypto/sha1"
	"fmt"
	"regexp"
)

func ReplaceWhitespace(s, r string) string {
	space := regexp.MustCompile(`\s+`)
	s = space.ReplaceAllString(s, r)
	return s
}

func HashPwd(password, hashSalt string) string {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	return password
}
