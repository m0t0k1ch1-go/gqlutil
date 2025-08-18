package gqlutil_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/gqlutil"
)

func TestInt64MarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   gqlutil.Int64
			out  []byte
		}{
			{
				"min",
				gqlutil.Int64(math.MinInt64),
				[]byte(`"-9223372036854775808"`),
			},
			{
				"zero",
				gqlutil.Int64(0),
				[]byte(`"0"`),
			},
			{
				"max",
				gqlutil.Int64(math.MaxInt64),
				[]byte(`"9223372036854775807"`),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var buf bytes.Buffer
				tc.in.MarshalGQL(&buf)

				require.Equal(t, tc.out, buf.Bytes())
			})
		}
	})
}

func TestInt64UnmarshalGQL(t *testing.T) {
	t.Run("failure", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
		}{
			{
				"int",
				0,
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var g gqlutil.Int64
				{
					err := g.UnmarshalGQL(tc.in)
					require.Error(t, err)
				}
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   string
			out  gqlutil.Int64
		}{
			{
				"min",
				"-9223372036854775808",
				gqlutil.Int64(math.MinInt64),
			},
			{
				"zero",
				"0",
				gqlutil.Int64(0),
			},
			{
				"max",
				"9223372036854775807",
				gqlutil.Int64(math.MaxInt64),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var g gqlutil.Int64
				{
					err := g.UnmarshalGQL(tc.in)
					require.NoError(t, err)
				}

				require.Equal(t, tc.out, g)
			})
		}
	})
}
