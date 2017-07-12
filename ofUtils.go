package ofutils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/astaxie/beego/orm"
	"github.com/axgle/mahonia"
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
func ChangeJson(data []byte, key string, value interface{}) ([]byte, error) {
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	m[key] = value
	return json.Marshal(m)
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
func ToJson(datas interface{}) string {
	jsonString, _ := json.Marshal(datas)
	return string(jsonString)
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
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
func ZeroBefore(i int) string {
	if i < 10 {
		return "0" + ToString(i)
	}
	return ToString(i)
}
func GetTimeStamp() string {
	return time.Now().Format("20060102150405")
}
func Utf8ToGBK(text string) string {
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(text)
}
func GBKFileToUtf8(filePath string) string {
	// Read UTF-8 from a GBK encoded file.
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	r := transform.NewReader(f, simplifiedchinese.GBK.NewDecoder())

	// Read converted UTF-8 from `r` as needed.
	// As an example we'll read line-by-line showing what was read:
	sc := bufio.NewScanner(r)
	result := ""
	for sc.Scan() {
		result += string(sc.Bytes()) + "\n"
	}
	return result
}
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
func TrimPrefix(s, suffix string) string {
	if strings.HasPrefix(s, suffix) {
		s = s[len(suffix):]
	}
	return s
}
func Copy(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

func GetStructName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
