package gqlutil

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalUint64 returns a [graphql.Marshaler] that encodes i as a quoted decimal string.
func MarshalUint64(i uint64) graphql.Marshaler {
	return graphql.MarshalString(strconv.FormatUint(i, 10))
}

// UnmarshalUint64 decodes a decimal integer string into a uint64.
// The string must not be signed; leading zeros are allowed and ignored.
func UnmarshalUint64(v any) (uint64, error) {
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

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid decimal string: %w", err)
	}

	return i, nil
}
