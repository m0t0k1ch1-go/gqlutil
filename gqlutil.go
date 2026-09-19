package gqlutil

import (
	"encoding/base64"
	"encoding/json/v2"
)

// EncodeToCursor encodes v as a cursor string for [Relay-style pagination].
//
// [Relay-style pagination]: https://relay.dev/graphql/connections.htm
func EncodeToCursor[T any](v T) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

// MustEncodeToCursor is like [EncodeToCursor] but panics if the input is invalid.
func MustEncodeToCursor[T any](v T) string {
	cursor, err := EncodeToCursor(v)
	if err != nil {
		panic(err)
	}

	return cursor
}

// DecodeCursor decodes a cursor string produced by [EncodeToCursor].
func DecodeCursor[T any](cursor string) (T, error) {
	var v T

	b, err := base64.RawURLEncoding.DecodeString(cursor)
	if err != nil {
		return v, err
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return v, err
	}

	return v, nil
}

// MustDecodeCursor is like [DecodeCursor] but panics if the input is invalid.
func MustDecodeCursor[T any](cursor string) T {
	v, err := DecodeCursor[T](cursor)
	if err != nil {
		panic(err)
	}

	return v
}
