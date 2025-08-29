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
	t.Run("failure", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
			want string
		}{
			{
				"func",
				func() {},
				"",
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				_, err := gqlutil.EncodeToCursor(tc.in)
				require.ErrorContains(t, err, tc.want)
			})
		}
	})

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
				cursor, err := gqlutil.EncodeToCursor(tc.in)
				require.NoError(t, err)
				require.Equal(t, tc.want, cursor)
			})
		}
	})
}

func TestMustEncodeToCursor(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
		}{
			{
				"func",
				func() {},
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				require.Panics(t, func() {
					gqlutil.MustEncodeToCursor(tc.in)
				})
			})
		}
	})

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
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				cursor := gqlutil.MustEncodeToCursor(tc.in)
				require.Equal(t, tc.want, cursor)
			})
		}
	})
}

func TestDecodeCursor(t *testing.T) {
	t.Run("failure", func(t *testing.T) {
		tcs := []struct {
			name string
			in   string
			want string
		}{
			{
				"invalid base64",
				"a",
				"",
			},
			{
				"invalid json",
				"eyJvZmZzZXQiOiIifQ",
				"",
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				_, err := gqlutil.DecodeCursor[CursorPayload](tc.in)
				require.ErrorContains(t, err, tc.want)
			})
		}
	})

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

func TestMustDecodeCursor(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		tcs := []struct {
			name string
			in   string
		}{
			{
				"invalid base64",
				"a",
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				require.Panics(t, func() {
					gqlutil.MustDecodeCursor[CursorPayload](tc.in)
				})
			})
		}
	})

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
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				cp := gqlutil.MustDecodeCursor[CursorPayload](tc.in)
				require.Equal(t, tc.want, cp)
			})
		}
	})
}
