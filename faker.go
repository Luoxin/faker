package faker

import (
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type I18nLanguage int

const (
	I18nLanguageEnUs I18nLanguage = 1
	I18nLanguageZhCn I18nLanguage = 2
	I18nLanguageJaJp I18nLanguage = 3
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
	Generator *rand.Rand
	Language  I18nLanguage
}

func (f *Faker) Choice(itemList interface{}) interface{} {
	ref := reflect.ValueOf(itemList)

	if ref.Kind() != reflect.Slice {
		panic("itemList is not slice")
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
	maxU := ^uint(0) >> 1
	max := int(maxU)
	min := -max - 1
	return f.IntBetween(min, max)
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

func (f *Faker) Bool() bool {
	return f.IntBetween(0, 100) > 50
}

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
	return &Faker{
		Generator: rand.New(rand.NewSource(time.Now().UnixNano())),
		Language:  I18nLanguageEnUs,
	}
}