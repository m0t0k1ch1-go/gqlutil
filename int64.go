package gqlutil

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalInt64 returns a graphql.Marshaler that encodes an int64 as a quoted decimal string.
func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.MarshalString(strconv.FormatInt(i, 10))
}

// UnmarshalInt64 decodes a GraphQL String (decimal) into an int64.
func UnmarshalInt64(v any) (int64, error) {
	if v == nil {
		return 0, errors.New("invalid graphql value: nil")
	}

	s, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("unsupported graphql value type: %T", v)
	}
	if len(s) == 0 {
		return 0, errors.New("invalid graphql string: empty")
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid graphql string: %w", err)
	}

	return i, nil
}
