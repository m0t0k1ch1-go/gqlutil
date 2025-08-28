package gqlutil_test

import (
	"testing"

	"github.com/m0t0k1ch1-go/gqlutil/v2"
	"github.com/stretchr/testify/require"
)

type CursorPayload struct {
	Offset int32 `json:"offset"`
}

func TestEncodeToCursor(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   CursorPayload
			want string
		}{
			{
				"zero value",
				CursorPayload{},
				"eyJvZmZzZXQiOjB9",
			},
			{
				"offset: 1",
				CursorPayload{Offset: 1},
				"eyJvZmZzZXQiOjF9",
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				s, err := gqlutil.EncodeToCursor(tc.in)
				require.NoError(t, err)
				require.Equal(t, tc.want, s)
			})
		}
	})
}

func TestDecodeCursor(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   string
			want CursorPayload
		}{
			{
				"zero value",
				"eyJvZmZzZXQiOjB9",
				CursorPayload{},
			},
			{
				"offset: 1",
				"eyJvZmZzZXQiOjF9",
				CursorPayload{Offset: 1},
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				cp, err := gqlutil.DecodeCursor[CursorPayload](tc.in)
				require.NoError(t, err)
				require.Equal(t, tc.want, cp)
			})
		}
	})
}
