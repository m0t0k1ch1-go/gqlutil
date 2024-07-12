package gqlutil

import (
	"io"
	"strconv"

	"github.com/samber/oops"
)

// Uint64 is a custom scalar for uint64.
type Uint64 uint64

// String implements the fmt.Stringer interface.
func (gi Uint64) String() string {
	return gi.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (gi Uint64) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(gi.string())))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (gi *Uint64) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.Errorf("Uint64 must be a string")
	}

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}

	*gi = Uint64(i)

	return nil
}

func (gi Uint64) string() string {
	return strconv.FormatUint(uint64(gi), 10)
}
