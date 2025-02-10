package str

type Array []Str

func NewArray[T ~[]E, E ~string | ~rune | ~byte | []byte](in T) Array {
	if in == nil {
		return nil
	}

	out := make(Array, len(in))
	for i := range in {
		out[i] = Str(in[i])
	}
	return out
}

func (s Array) Strings() []string {
	out := make([]string, len(s))
	for i := range s {
		out[i] = string(s[i])
	}
	return out
}
