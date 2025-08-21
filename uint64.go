package gqlutil

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Uint64 represents a GraphQL custom scalar for unsigned 64-bit integers.
type Uint64 uint64

// Unwrap returns the value as a uint64.
func (g Uint64) Unwrap() uint64 {
	return uint64(g)
}

// String implements fmt.Stringer.
// It returns the value as a decimal string.
func (g Uint64) String() string {
	return strconv.FormatUint(uint64(g), 10)
}

// MarshalGQL implements graphql.Marshaler.
// It writes the value as a quoted decimal string.
func (g Uint64) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(g.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler.
// It accepts a decimal string.
func (g *Uint64) UnmarshalGQL(v any) error {
	if v == nil {
		return errors.New("invalid graphql value: nil")
	}

	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("unsupported graphql value type: %T", v)
	}
	if len(s) == 0 {
		return errors.New("invalid graphql string: empty")
	}

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid graphql string: %w", err)
	}

	*g = Uint64(i)

	return nil
}
