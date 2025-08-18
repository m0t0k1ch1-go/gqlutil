package gqlutil

import (
	"io"
	"strconv"

	"github.com/samber/oops"
)

// Int64 is a custom scalar for int64.
type Int64 int64

// Unwrap returns the int64.
func (g Int64) Unwrap() int64 {
	return int64(g)
}

// String implements the fmt.Stringer interface.
func (g Int64) String() string {
	return g.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (g Int64) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(g.string()))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (g *Int64) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.New("v must be string")
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return oops.Wrap(err)
	}

	*g = Int64(i)

	return nil
}

func (g Int64) string() string {
	return strconv.FormatInt(int64(g), 10)
}
