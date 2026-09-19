package gqlutil

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalInt64 returns a [graphql.Marshaler] that encodes i as a quoted decimal string.
func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.MarshalString(strconv.FormatInt(i, 10))
}

// UnmarshalInt64 decodes a decimal string into an int64.
// The string may be signed; leading zeros are allowed and ignored.
func UnmarshalInt64(v any) (int64, error) {
	if v == nil {
		return 0, errors.New("unsupported value: nil")
	}

	s, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("unsupported value type: %T", v)
	}
	if len(s) == 0 {
		return 0, errors.New("invalid decimal string: empty")
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid decimal string: %w", err)
	}

	return i, nil
}
