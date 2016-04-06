package ofutils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
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
	if str == nil {
		return ""
	}
	return fmt.Sprintf("%v", str)
}

func ToInt(val interface{}) int {
	if val == nil {
		return 0
	}
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

func GetWeekFirstDay() string {
	week := time.Now().Weekday().String()
	var day time.Duration
	switch week {
	case "Sunday":
		day = 6
		break
	case "Monday":
		day = 0
		break
	case "Tuesday":
		day = 1
		break
	case "Wednesday":
		day = 2
		break
	case "Thursday":
		day = 3
		break
	case "Friday":
		day = 4
		break
	case "Saturday":
		day = 5
		break
	}
	date := time.Now().Add(-day * 24 * time.Hour)
	return date.Format("2006-01-02")
}

type ByKey struct {
	Key  string
	List []orm.Params
}

func (a ByKey) Len() int {
	return len(a.List)
}
func (a ByKey) Swap(i, j int) {
	a.List[i], a.List[j] = a.List[j], a.List[i]
}
func (a ByKey) Less(i, j int) bool {
	return ToInt(a.List[i][a.Key]) > ToInt(a.List[j][a.Key])
}

func Sort(list []orm.Params, key string) []orm.Params {
	byKey := ByKey{List: list, Key: key}
	sort.Sort(byKey)
	return byKey.List
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

func ToFloat(str interface{}) float64 {
	if str == nil {
		return 0
	}
	tf, err := strconv.ParseFloat(ToString(str), 64)
	if err != nil {
		return 0
	}
	return tf
}
func GetEncryptPhone(phone interface{}) string {
	temp := ToString(phone)
	if len(temp) == 11 {
		start := SubString(temp, 0, 3)
		end := SubString(temp, 7, 11)
		temp = start + "****" + end
	}
	return temp
}
