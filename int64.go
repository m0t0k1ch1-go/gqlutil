package gqlutil

import (
	"io"
	"strconv"

	"github.com/samber/oops"
)

// Int64 is a custom scalar for int64.
type Int64 int64

// Unwrap returns the int64.
func (gi Int64) Unwrap() int64 {
	return int64(gi)
}

// String implements the fmt.Stringer interface.
func (gi Int64) String() string {
	return gi.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (gi Int64) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(gi.string()))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (gi *Int64) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.Errorf("must be a string")
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*gi = Int64(i)

	return nil
}

func (gi Int64) string() string {
	return strconv.FormatInt(int64(gi), 10)
}
