package clg

import (
	"math"
	"strings"

	"github.com/xh3b4sd/anna/id"
)

func (i *clgIndex) ContainsString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	substr, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	contains := strings.Contains(s, substr)

	return []interface{}{contains}, nil
}

func (i *clgIndex) CountCharacterString(args ...interface{}) ([]interface{}, error) {
	str, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newMap := map[string]int{}
	for _, s := range strings.Split(str, "") {
		newMap[s]++
	}

	return []interface{}{newMap}, nil
}

// The following code is a golang port of the optimized C version of
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C.
func editDistance(s1, s2 string) int {
	var cost, lastdiag, olddiag int
	ls1 := len([]rune(s1))
	ls2 := len([]rune(s2))

	column := make([]int, ls1+1)

	for i := 1; i <= ls1; i++ {
		column[i] = i
	}

	for i := 1; i <= ls2; i++ {
		column[0] = i
		lastdiag = i - 1

		for j := 1; j <= ls1; j++ {
			olddiag = column[j]

			cost = 0
			if s1[j-1] != s2[i-1] {
				cost = 1
			}

			column[j] = editDistanceMin(column[j]+1, column[j-1]+1, lastdiag+cost)
			lastdiag = olddiag
		}
	}

	return column[ls1]
}

func editDistanceMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func (i *clgIndex) EditDistanceString(args ...interface{}) ([]interface{}, error) {
	s1, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	s2, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	newInt := editDistance(s1, s2)

	return []interface{}{newInt}, nil
}

func (i *clgIndex) LongerString(args ...interface{}) ([]interface{}, error) {
	s1, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	s2, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	longer := len(s1) > len(s2)

	return []interface{}{longer}, nil
}

func (i *clgIndex) NewIDString(args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 0 got %d", len(args))
	}

	newID := string(id.NewObjectID(id.Hex128))

	return []interface{}{newID}, nil
}

func (i *clgIndex) RepeatString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	count, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if count < 0 {
		return nil, maskAnyf(negativeIntError, "integer must not be negative")
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	repeated := strings.Repeat(s, count)

	return []interface{}{repeated}, nil
}

func (i *clgIndex) ReverseString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	newString := string(chars)

	return []interface{}{newString}, nil
}

func (i *clgIndex) ShorterString(args ...interface{}) ([]interface{}, error) {
	s1, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	s2, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	shorter := len(s1) < len(s2)

	return []interface{}{shorter}, nil
}

func (i *clgIndex) SplitString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	sep, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	newStringSlice := strings.Split(s, sep)

	return []interface{}{newStringSlice}, nil
}

func (i *clgIndex) SplitEqualString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	n, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if n < 0 {
		return nil, maskAnyf(negativeIntError, "integer must not be negative")
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if n > len(s) {
		return nil, maskAnyf(invalidDividerError, "cannot divider %s into %d parts", s, n)
	}

	l := float64(len(s))
	size := l / float64(n)
	c := math.Ceil(size)
	isInt := c == size
	if !isInt {
		size = c
	}
	var newStringSlice []string
	start := float64(0)
	end := size
	for i := 1; i <= n; i++ {
		newStringSlice = append(newStringSlice, s[int(start):int(end)])
		start = end
		if !isInt {
			start -= 1
		}
		end = start + size
		if end >= l {
			end = l
		}
	}

	return []interface{}{newStringSlice}, nil
}

func (i *clgIndex) ToLowerString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newString := strings.ToLower(s)

	return []interface{}{newString}, nil
}

func (i *clgIndex) ToUpperString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newString := strings.ToUpper(s)

	return []interface{}{newString}, nil
}
