package gqlutil

import (
	"encoding/base64"
	"encoding/json"
)

// EncodeToCursor encodes any value as a Relay-style cursor string.
func EncodeToCursor[T any](v T) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

// DecodeCursor decodes a cursor string produced by EncodeToCursor.
func DecodeCursor[T any](cursor string) (v T, err error) {
	var b []byte
	{
		if b, err = base64.RawURLEncoding.DecodeString(cursor); err != nil {
			return
		}
	}

	err = json.Unmarshal(b, &v)

	return
}
