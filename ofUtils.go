package ofutils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ToString(str interface{}) string {
	return fmt.Sprintf("%v", str)
}

func ToInt(val interface{}) int {
	s, ok := val.(string)
	if ok {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return i
	}
	i, ok := val.(int)
	if ok {
		return i
	} else {
		return 0
	}
}
