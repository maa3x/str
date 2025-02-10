package str

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

// var _ json.Marshaler = (*Str)(nil)

type Str string

func New(v any) Str {
	switch v := v.(type) {
	case Str:
		return v
	case string:
		return Str(v)
	case byte:
		return Str(v)
	case []byte:
		return Str(v)
	case rune:
		return Str(v)
	case []rune:
		return Str(v)
	case []string:
		return Str(strings.Join(v, " "))
	case int:
		return Str(strconv.FormatInt(int64(v), 10))
	case int8:
		return Str(strconv.FormatInt(int64(v), 10))
	case int16:
		return Str(strconv.FormatInt(int64(v), 10))
	case int64:
		return Str(strconv.FormatInt(int64(v), 10))
	case uint:
		return Str(strconv.FormatUint(uint64(v), 10))
	case uint16:
		return Str(strconv.FormatUint(uint64(v), 10))
	case uint32:
		return Str(strconv.FormatUint(uint64(v), 10))
	case uint64:
		return Str(strconv.FormatUint(uint64(v), 10))
	case float32:
		return Str(strconv.FormatFloat(float64(v), 'g', -1, 32))
	case float64:
		return Str(strconv.FormatFloat(float64(v), 'g', -1, 64))
	case bool:
		return Str(strconv.FormatBool(v))
	case fmt.Stringer:
		if v != nil {
			return Str(v.String())
		}
		return ""
	default:
		rv := reflect.ValueOf(v)
		if !rv.IsValid() || rv.IsNil() {
			return ""
		}
		if rv.Kind() == reflect.Ptr {
			if v = rv.Elem().Interface(); v != nil {
				return New(v)
			}
			return ""
		}

		return Str(fmt.Sprint(v))
	}
}

func (s Str) String() string {
	return string(s)
}

func (s Str) Append(values ...string) Str {
	for i := range values {
		s += Str(values[i])
	}
	return s
}

func (s Str) AppendRune(values ...rune) Str {
	return s + Str(values)
}

func (s Str) At(index int) Str {
	return Str(s[index])
}

func (s Str) Clone() Str {
	return Str(strings.Clone(string(s)))
}

func (s Str) Compare(value string) int {
	return strings.Compare(string(s), value)
}

func (s Str) Contains(values ...string) bool {
	for i := range values {
		if strings.Contains(string(s), values[i]) {
			return true
		}
	}
	return false
}

func (s Str) ContainsAny(chars string) bool {
	return strings.ContainsAny(string(s), chars)
}

func (s Str) Count(search string) int {
	return strings.Count(string(s), search)
}

func (s Str) Cut(search string) (before, after Str, found bool) {
	b, a, f := strings.Cut(string(s), search)
	return Str(b), Str(a), f
}

func (s Str) CutPrefix(prefix string) (Str, bool) {
	a, f := strings.CutPrefix(string(s), prefix)
	return Str(a), f
}

func (s Str) CutSuffix(suffix string) (Str, bool) {
	b, f := strings.CutSuffix(string(s), suffix)
	return Str(b), f
}

func (s Str) Fields() Array {
	return NewArray(strings.Fields(string(s)))
}

func (s Str) FieldsFunc(fn func(rune) bool) Array {
	return NewArray(strings.FieldsFunc(string(s), fn))
}

func (s Str) Find(pattern string) Str {
	if re, err := regexp.Compile(pattern); err == nil {
		return s.FindRegex(re)
	}
	return ""
}

func (s Str) FindAll(pattern string) Array {
	if re, err := regexp.Compile(pattern); err == nil {
		return s.FindAllRegex(re)
	}
	return nil
}

func (s Str) FindAllIndex(pattern string) [][]int {
	if re, err := regexp.Compile(pattern); err == nil {
		return s.FindAllIndexRegex(re)
	}
	return nil
}

func (s Str) FindIndex(pattern string) []int {
	if re, err := regexp.Compile(pattern); err == nil {
		return s.FindIndexRegex(re)
	}
	return nil
}

func (s Str) FindRegex(re *regexp.Regexp) Str {
	return Str(re.FindString(string(s)))
}

func (s Str) FindAllRegex(re *regexp.Regexp) Array {
	return NewArray(re.FindAllString(string(s), -1))
}

func (s Str) FindAllIndexRegex(re *regexp.Regexp) [][]int {
	return re.FindAllStringIndex(string(s), -1)
}

func (s Str) FindIndexRegex(re *regexp.Regexp) []int {
	return re.FindStringIndex(string(s))
}

func (s Str) HasPrefix(prefixes ...string) bool {
	for i := range prefixes {
		if strings.HasPrefix(string(s), prefixes[i]) {
			return true
		}
	}
	return false
}

func (s Str) HasSuffix(suffixes ...string) bool {
	for i := range suffixes {
		if strings.HasSuffix(string(s), suffixes[i]) {
			return true
		}
	}
	return false
}

func (s Str) In(values ...string) bool {
	return slices.Contains(values, string(s))
}

func (s Str) Index(search string) int {
	return strings.Index(string(s), search)
}

func (s Str) IndexAny(chars string) int {
	return strings.IndexAny(string(s), chars)
}

func (s Str) Join(values []string) Str {
	return Str(strings.Join(values, string(s)))
}

func (s Str) LastIndex(search string) int {
	return strings.LastIndex(string(s), search)
}

func (s Str) LastIndexAny(chars string) int {
	return strings.LastIndexAny(string(s), chars)
}

func (s Str) Len() int {
	return len(s)
}

func (s Str) Map(fn func(rune) rune) Str {
	return Str(strings.Map(fn, string(s)))
}

func (s Str) Match(patterns ...string) bool {
	for i := range patterns {
		if re, err := regexp.Compile(patterns[i]); err == nil {
			return s.MatchRegex(re)
		}
	}
	return false
}

func (s Str) MatchRegex(regexes ...*regexp.Regexp) bool {
	for i := range regexes {
		if regexes[i].MatchString(string(s)) {
			return true
		}
	}
	return false
}

func (s Str) PadEnd(length int, pad string) Str {
	if s.RuneCount() >= length {
		return s
	}
	padLen := Str(pad).RuneCount()
	if padLen == 0 {
		return s
	}

	length -= s.RuneCount()
	return s + Str(strings.Repeat(pad, length/padLen)) + Str(pad).SliceRunesTo(length%padLen)
}

func (s Str) PadStart(length int, pad string) Str {
	if s.RuneCount() >= length {
		return s
	}
	padLen := Str(pad).RuneCount()
	if padLen == 0 {
		return s
	}

	length -= s.RuneCount()
	return Str(strings.Repeat(pad, length/padLen)) + Str(pad).SliceRunesTo(length%padLen) + s
}

func (s Str) ParseFloat() (float64, error) {
	return strconv.ParseFloat(string(s), 64)
}

func (s Str) ParseInt() (int64, error) {
	return strconv.ParseInt(string(s), 10, 64)
}

func (s Str) ParseUint() (uint64, error) {
	return strconv.ParseUint(string(s), 10, 64)
}

func (s Str) PartOf(values ...string) bool {
	for i := range values {
		if strings.Contains(values[i], string(s)) {
			return true
		}
	}
	return false
}

func (s Str) PopEnd() (rune, Str) {
	if runes := []rune(s); len(runes) > 0 {
		last := len(runes) - 1
		return runes[last], Str(runes[:last])
	}
	return -1, ""
}

func (s Str) PopStart() (rune, Str) {
	if runes := []rune(s); len(runes) > 0 {
		return runes[0], Str(runes[1:])
	}
	return -1, ""
}

func (s Str) Prepend(values ...string) Str {
	v := ""
	for i := range values {
		v += values[i]
	}
	return Str(v) + s
}

func (s Str) PrependRune(values ...rune) Str {
	return Str(values) + s
}

func (s Str) Reader() *strings.Reader {
	return strings.NewReader(string(s))
}

func (s Str) Repeat(count int) Str {
	return Str(strings.Repeat(string(s), count))
}

func (s Str) Replace(old, new string) Str {
	return Str(strings.ReplaceAll(string(s), old, new))
}

func (s Str) ReplaceEach(old, new string) Str {
	src, dst := []rune(old), []rune(new)
	if len(src) != len(dst) {
		return s
	}

	for i := range src {
		s = Str(strings.ReplaceAll(string(s), string(src[i]), string(dst[i])))
	}
	return s
}

func (s Str) ReplaceLast(old, new string) Str {
	if old == "" {
		return s
	}
	lastIdx := s.LastIndex(old)
	if lastIdx == -1 {
		return s
	}
	return Str(string(s[:lastIdx]) + new + string(s[lastIdx+len(old):]))
}

func (s Str) ReplaceN(old, new string, n int) Str {
	return Str(strings.Replace(string(s), old, new, n))
}

func (s Str) ReplacePattern(pattern, new string) Str {
	if re, err := regexp.Compile(pattern); err == nil {
		return s.ReplaceRegex(re, new)
	}
	return s
}

func (s Str) ReplaceRegex(regex *regexp.Regexp, new string) Str {
	return Str(regex.ReplaceAllString(string(s), new))
}

func (s Str) RuneAt(index int) Str {
	runes := []rune(s)
	if len(runes) > index {
		return Str(runes[index])
	}
	return ""
}

func (s Str) RuneCount() int {
	return utf8.RuneCountInString(string(s))
}

func (s Str) RuneIndex(r rune) int {
	return strings.IndexRune(string(s), r)
}

func (s Str) Runes() Array {
	return NewArray([]rune(s))
}

func (s Str) Slice(start, end int) Str {
	if start < 0 {
		start = s.Len() + start
	}
	if end < 0 {
		end = s.Len() + end
	}
	if start < 0 || end < 0 || start >= s.Len() || end > s.Len() || start >= end {
		return ""
	}

	return s[start:end]
}

func (s Str) SliceFrom(index int) Str {
	if index < 0 {
		index = s.Len() + index
	}
	if index < 0 || index >= s.Len() {
		return ""
	}

	return s[index:]
}

func (s Str) SliceRunes(start, end int) Str {
	runes := []rune(s)
	if start < 0 {
		start = len(runes) + start
	}
	if end < 0 {
		end = len(runes) + end
	}
	if start < 0 || end < 0 || start >= len(runes) || end > len(runes) || start >= end {
		return ""
	}

	return Str(runes[start:end])
}

func (s Str) SliceRunesFrom(index int) Str {
	runes := []rune(s)
	if index < 0 {
		index = len(runes) + index
	}
	if index < 0 || index >= len(runes) {
		return ""
	}

	return Str(runes[index:])
}

func (s Str) SliceRunesTo(index int) Str {
	runes := []rune(s)
	if index < 0 {
		index = len(runes) + index
	}
	if index < 0 || index > len(runes) {
		return ""
	}

	return Str(runes[:index])
}

func (s Str) SliceTo(index int) Str {
	if index < 0 {
		index = s.Len() + index
	}
	if index < 0 || index >= s.Len() {
		return ""
	}

	return s[:index]
}

func (s Str) Split(delim string) Array {
	return s.SplitN(delim, -1)
}

func (s Str) SplitN(delim string, count int) Array {
	return NewArray(strings.SplitN(string(s), delim, count))
}

func (s Str) SplitPattern(pattern string) Array {
	if re, err := regexp.Compile(pattern); err == nil {
		return s.SplitRegex(re)
	}
	return Array{s}
}

func (s Str) SplitRegex(regex *regexp.Regexp) Array {
	return NewArray(regex.Split(string(s), -1))
}

func (s Str) ToCamel() Str {
	if first, remaining := s.ToPascal().PopStart(); first != -1 {
		return Str(strings.ToLower(string(first)) + string(remaining))
	}
	return s
}

func (s Str) ToGoCamel() Str {
	if first, remaining := s.ToGoPascal().PopStart(); first != -1 {
		return Str(strings.ToLower(string(first)) + string(remaining))
	}
	return s
}

func (s Str) ToGoPascal() Str {
	words := splitWords(string(s))
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return Str(goPascalWords(words))
}

func (s Str) ToKebab() Str {
	words := splitWords(string(s))
	return Str("-").Join(words).ToLower()
}

func (s Str) ToLower() Str {
	return Str(strings.ToLower(string(s)))
}

func (s Str) ToPascal() Str {
	return Str(pascalWords(splitWords(string(s))))
}

func (s Str) ToSnake() Str {
	words := splitWords(string(s))
	return Str("_").Join(words).ToLower()
}

func (s Str) ToUpper() Str {
	return Str(strings.ToUpper(string(s)))
}

func (s Str) ToUpperSnake() Str {
	return s.ToSnake().ToUpper()
}

func (s Str) Trim() Str {
	return Str(strings.TrimSpace(string(s)))
}

func (s Str) TrimFunc(fn func(rune) bool) Str {
	return Str(strings.TrimFunc(string(s), fn))
}

func (s Str) TrimPrefix(prefixes ...string) Str {
	for i := range prefixes {
		s = Str(strings.TrimPrefix(string(s), prefixes[i]))
	}
	return s
}

func (s Str) TrimSuffix(suffixes ...string) Str {
	for i := range suffixes {
		s = Str(strings.TrimSuffix(string(s), suffixes[i]))
	}
	return s
}

func (s Str) TrimValue(values ...string) Str {
	for i := range values {
		s = Str(strings.Trim(string(s), values[i]))
	}
	return s
}

func (s Str) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

func (s *Str) UnmarshalJSON(v []byte) error {
	return json.Unmarshal(v, (*string)(s))
}

func (s Str) MarshalText() ([]byte, error) {
	return []byte(s), nil
}

func (s *Str) UnmarshalText(v []byte) error {
	*s = Str(v)
	return nil
}

func (s Str) MarshalBinary() ([]byte, error) {
	return []byte(s), nil
}

func (s *Str) UnmarshalBinary(v []byte) error {
	*s = Str(v)
	return nil
}
