package gqlutil_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/gqlutil"
)

func TestMarshalInt64(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   int64
			want string
		}{
			{
				"min",
				math.MinInt64,
				`"-9223372036854775808"`,
			},
			{
				"zero",
				0,
				`"0"`,
			},
			{
				"max",
				math.MaxInt64,
				`"9223372036854775807"`,
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var buf bytes.Buffer
				gqlutil.MarshalInt64(tc.in).MarshalGQL(&buf)
				require.Equal(t, tc.want, buf.String())
			})
		}
	})
}

func TestUnmarshalInt64(t *testing.T) {
	t.Run("failure", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
			want string
		}{
			{
				"nil",
				nil,
				"invalid graphql value: nil",
			},
			{
				"int",
				int(0),
				"unsupported graphql value type: int",
			},
			{
				"string: empty",
				"",
				"invalid graphql string: empty",
			},
			{
				"string: invalid",
				"invalid",
				"invalid graphql string",
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				_, err := gqlutil.UnmarshalInt64(tc.in)
				require.ErrorContains(t, err, tc.want)
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
			want int64
		}{
			{
				"min",
				"-9223372036854775808",
				math.MinInt64,
			},
			{
				"zero",
				"0",
				0,
			},
			{
				"max",
				"9223372036854775807",
				math.MaxInt64,
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				i, err := gqlutil.UnmarshalInt64(tc.in)
				require.NoError(t, err)
				require.Equal(t, tc.want, i)
			})
		}
	})
}
