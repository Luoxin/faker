package faker

import (
	"bytes"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type I18nLanguage int

var (
	I18nLanguageNil  I18nLanguage = 0 // 无关语言的默认
	I18nLanguageEnUs I18nLanguage = 1
	I18nLanguageZhCn I18nLanguage = 2
)

// 字符串集合
var whitespace = []byte(" \\t\\n\\r\\v\\f")
var asciiLowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var asciiUppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var asciiLetters = append(asciiLowercase, asciiUppercase...)
var digits = []byte("0123456789")
var hexigits = append(digits, []byte("abcdefABCDEF")...)
var octdigits = []byte("01234567")
var punctuation = []byte("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
var printable []byte

func init() {
	printable = append(digits, asciiLetters...)
	printable = append(printable, punctuation...)
	printable = append(printable, whitespace...)
}

type Faker struct {
	Generator   *rand.Rand
	Language    I18nLanguage
	ProviderMap map[I18nLanguage]*Provider
}

// 基础的随机选择
func (f *Faker) Choice(itemList interface{}) interface{} {
	ref := reflect.ValueOf(itemList)

	if ref.Kind() != reflect.Slice {
		panic("itemList is not slice")
	}

	if ref.Len() == 0 {
		ref = reflect.MakeSlice(ref.Type(), 1, 1)
		return ref.Index(0).Interface()
	}

	return ref.Index(f.Generator.Intn(ref.Len())).Interface()
}

func (f *Faker) RandomDigit() int {
	return f.Generator.Int() % 10
}

func (f *Faker) RandomDigitNot(ignore ...int) int {
	inSlice := func(el int, list []int) bool {
		for i := range list {
			if i == el {
				return true
			}
		}

		return false
	}

	for {
		current := f.RandomDigit()
		if inSlice(current, ignore) {
			return current
		}
	}
}

func (f *Faker) RandomDigitNotNull() int {
	return f.Generator.Int()%8 + 1
}

func (f *Faker) RandomNumber(size int) int {
	if size == 1 {
		return f.RandomDigit()
	}

	var min = int(math.Pow10(size - 1))
	var max = int(math.Pow10(size)) - 1

	return f.IntBetween(min, max)
}

func (f *Faker) FloatBetween(min, max float32) float32 {
	if min == 0 && max == 0 {
		return 0
	}

	if min >= max {
		return max
	}

	return min + (f.Generator.Float32() * (max - min))
}

func (f *Faker) Float32() float32 {
	max := float32(1 << 24)
	min := -max - 1
	return f.FloatBetween(min, max)
}

func (f *Faker) Float64() float64 {
	return float64(f.Float32())
}

func (f *Faker) Int() int {
	return f.Generator.Int()
}

func (f *Faker) Uint() uint {
	return uint(f.Int())
}

func (f *Faker) Int64() int64 {
	return int64(f.Int())
}

func (f *Faker) Uint64() uint64 {
	return uint64(f.Int())
}

func (f *Faker) Int32() int32 {
	return int32(f.Int())
}

func (f *Faker) Uint32() uint32 {
	return uint32(f.Int())
}

func (f *Faker) IntBetween(min, max int) int {
	if min == 0 && max == 0 {
		return 0
	}

	if min >= max {
		return max
	}

	return min + (f.Generator.Intn(max - min))
}

func (f *Faker) Int64Between(min, max int64) int64 {
	return int64(f.IntBetween(int(min), int(max)))
}

func (f *Faker) Int32Between(min, max int32) int32 {
	return int32(f.IntBetween(int(min), int(max)))
}

func (f *Faker) Unt32Between(min, max uint32) uint32 {
	return uint32(f.IntBetween(int(min), int(max)))
}

func (f *Faker) Unt64Between(min, max uint64) uint64 {
	return uint64(f.IntBetween(int(min), int(max)))
}

func (f *Faker) Letter() string {
	return f.RandomLetter()
}

func (f *Faker) RandomLetter() string {
	return string(f.Choice(asciiLetters).(byte))
}

func (f *Faker) RandomLowLetter() string {
	return string(f.Choice(asciiLowercase).(byte))
}

func (f *Faker) RandomUpperLetter() string {
	return string(f.Choice(asciiUppercase).(byte))
}

func (f *Faker) RandomLetterString(length int) string {
	return f.RandomString(asciiLetters, length)
}

func (f *Faker) RandomLowLetterString(length int) string {
	return f.RandomString(asciiLowercase, length)
}

func (f *Faker) RandomUpperLetterString(length int) string {
	return f.RandomString(asciiUppercase, length)
}

func (f *Faker) RandomString(dataset []byte, length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = dataset[r.Intn(len(dataset))]
	}

	return string(bytes)
}

func (f *Faker) RandomBytesElement(s []byte) byte {
	return f.Choice(s).(byte)
}

func (f *Faker) RandomStringElement(s []string) string {
	return f.Choice(s).(string)
}

func (f *Faker) RandomIntElement(a []int) int {
	return f.Choice(a).(int)
}

func (f *Faker) ShuffleString(s string) string {
	orig := strings.Split(s, "")
	dest := make([]string, len(orig))

	for i := 0; i < len(orig); i++ {
		dest[i] = orig[len(orig)-i-1]
	}

	return strings.Join(dest, "")
}

func (f *Faker) Bool() bool {
	return f.IntBetween(0, 100) > 50
}

// 创建新的
func (f *Faker) InitGenerator() {
	f.SetSeed(time.Now().UnixNano())
}

func (f *Faker) SetSeed(seed int64) {
	f.Generator = rand.New(rand.NewSource(seed))
}

func (f *Faker) SetLanguage(i18n I18nLanguage) {
	f.Language = i18n
}

func New() *Faker {
	f := &Faker{
		Generator: rand.New(rand.NewSource(time.Now().UnixNano())),
		Language:  I18nLanguageEnUs,
	}
	f.InitProviderMap()
	return f
}

func NewWithLanguage(language I18nLanguage) *Faker {
	f := New()
	f.SetLanguage(language)
	return f
}

// 格式化数据
func (f *Faker) Format(fmt string, args map[string]interface{}) (out string) {
	var msg bytes.Buffer

	tmpl, err := template.New("").Parse(fmt)
	if err != nil {
		return fmt
	}
	err = tmpl.Execute(&msg, args)
	if err != nil {
		return fmt
	}

	out = msg.String()
	out = f.Bothify(out)
	out = f.Asciify(out)
	return out
}

func (f *Faker) Numerify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "#" {
			c = strconv.Itoa(f.RandomDigit())
		}

		out = out + c
	}

	return
}

func (f *Faker) Lexify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "?" {
			c = f.RandomLetter()
		}

		out = out + c
	}

	return
}

func (f *Faker) Bothify(in string) (out string) {
	out = f.Lexify(in)
	out = f.Numerify(out)
	return
}

func (f *Faker) Asciify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "*" {
			c = string(f.IntBetween(97, 126))
		}

		out = out + c
	}

	return
}
