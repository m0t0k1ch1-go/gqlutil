package gqlutil

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalUint64(i uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(strconv.FormatUint(i, 10)))
	})
}

func UnmarshalUint64(v any) (uint64, error) {
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

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid graphql string: %w", err)
	}

	return i, nil
}
