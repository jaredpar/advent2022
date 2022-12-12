package util

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/exp/constraints"
)

func ParseLines(text string) []string {
	reader := strings.NewReader(text)
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadLines(f embed.FS, name string) ([]string, error) {
	file, err := f.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func MustReadLines(f embed.FS, name string) []string {
	lines, err := ReadLines(f, name)
	if err != nil {
		panic(err)
	}

	return lines
}

func ReadLinesAsInt(f embed.FS, name string) ([]int, error) {
	lines, err := ReadLines(f, name)
	if err != nil {
		return nil, err
	}

	values := make([]int, 0, len(lines))
	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

func ReadAsSingleLine(f embed.FS, name string) (string, error) {
	file, err := f.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", nil
	}

	line := scanner.Text()
	if scanner.Scan() {
		return "", errors.New("file had multiple lines")
	}

	return line, nil
}

func ParseCommaSepInt(line string) ([]int, error) {
	parts := strings.Split(line, ",")
	numbers := make([]int, 0, len(parts))
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func SplitOnWhiteSpace(line string) []string {
	startIndex := -1
	items := make([]string, 0)
	for index, r := range line {
		if unicode.IsSpace(r) {
			if startIndex >= 0 {
				item := line[startIndex:index]
				items = append(items, item)
			}
			startIndex = -1
		} else if startIndex < 0 {
			startIndex = index
		}
	}

	if startIndex >= 0 {
		item := line[startIndex:]
		items = append(items, item)
	}

	return items
}

// Convert a rune value between '0' and '9' to a int value
func RuneToInt(r rune) (int, error) {
	value := int(r - '0')
	if value >= 0 && value <= 9 {
		return value, nil
	}

	return value, errors.New("invalid value")
}

// Convert a rune value between '0' and '9' to a int value
func ByteToInt(b byte) (int, error) {
	value := int(b - '0')
	if value >= 0 && value <= 9 {
		return value, nil
	}

	return value, errors.New("invalid value")
}

// Return the first rune in the string. Will panic on a zero length string
func FirstRune(text string) rune {
	for _, r := range text {
		return r
	}

	panic("zero length string")
}

func SetAll[T any](values []T, value T) {
	for i := 0; i < len(values); i++ {
		values[i] = value
	}
}

func DigitToRune(d int) rune {
	return rune('0' + d)
}

func StringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("%s is not a number", s))
	}

	return value
}

func Min[T constraints.Ordered](x T, y T) T {
	if x < y {
		return x
	}
	return y
}

func MinSlice[T constraints.Ordered](data []T) T {
	if len(data) == 0 {
		panic("must be at least one element")
	}

	min := data[0]
	for i := 1; i < len(data); i++ {
		min = Min(min, data[i])
	}

	return min
}

func Max[T constraints.Ordered](x T, y T) T {
	if x > y {
		return x
	}
	return y
}

func MaxSlice[T constraints.Ordered](data []T) T {
	if len(data) == 0 {
		panic("must be at least one element")
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		max = Max(max, data[i])
	}

	return max
}

func Abs(x int) int {
	if x < 0 {
		x *= -1
	}

	return x
}

func Require(cond bool) {
	if !cond {
		panic("failed assert")
	}
}

func ReverseSlice[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func IsWhitespace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

func StartsWith(s string, r rune) bool {
	found, _ := utf8.DecodeRuneInString(s)
	return found == r
}

func InsertAt[T any](data []T, index int, value T) []T {
	length := len(data)
	if length == index {
		return append(data, value)
	}

	var dummy T
	data = append(data, dummy)

	for i := length; i > index; i-- {
		data[i] = data[i-1]
	}

	data[index] = value
	return data
}

func InsertSortedF[T any](data []T, value T, less func(left, right T) bool) []T {
	length := len(data)
	if length == 0 {
		return append(data, value)
	}

	index := sort.Search(length, func(i int) bool {
		return less(value, data[i])
	})

	return InsertAt(data, index, value)
}

func InsertSorted[S ~[]E, E constraints.Ordered](data S, value E) S {
	return InsertSortedF(data, value, func(left, right E) bool {
		return left < right
	})
}
