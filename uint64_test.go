package gqlutil_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/gqlutil"
)

func TestUint64MarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   gqlutil.Uint64
			out  []byte
		}{
			{
				"zero",
				gqlutil.Uint64(0),
				[]byte(`"0"`),
			},
			{
				"max",
				gqlutil.Uint64(math.MaxUint64),
				[]byte(`"18446744073709551615"`),
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

func TestUint64UnmarshalGQL(t *testing.T) {
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
				var g gqlutil.Uint64
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
			out  gqlutil.Uint64
		}{
			{
				"zero",
				"0",
				gqlutil.Uint64(0),
			},
			{
				"max",
				"18446744073709551615",
				gqlutil.Uint64(math.MaxUint64),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var g gqlutil.Uint64
				{
					err := g.UnmarshalGQL(tc.in)
					require.NoError(t, err)
				}

				require.Equal(t, tc.out, g)
			})
		}
	})
}
