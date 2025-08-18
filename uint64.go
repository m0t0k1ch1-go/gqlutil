package gqlutil

import (
	"io"
	"strconv"

	"github.com/samber/oops"
)

// Uint64 is a custom scalar for uint64.
type Uint64 uint64

// Unwrap returns the uint64.
func (g Uint64) Unwrap() uint64 {
	return uint64(g)
}

// String implements the fmt.Stringer interface.
func (g Uint64) String() string {
	return g.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (g Uint64) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(g.string()))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (g *Uint64) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.New("v must be string")
	}

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return oops.Wrap(err)
	}

	*g = Uint64(i)

	return nil
}

func (g Uint64) string() string {
	return strconv.FormatUint(uint64(g), 10)
}
