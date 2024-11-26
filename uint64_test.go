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
				"min",
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
				var buf = &bytes.Buffer{}
				tc.in.MarshalGQL(buf)

				require.Equal(t, tc.out, buf.Bytes())
			})
		}
	})
}

func TestUint64UnmarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   string
			out  gqlutil.Uint64
		}{
			{
				"min",
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
				require.Nil(t, g.UnmarshalGQL(tc.in))

				require.Equal(t, tc.out, g)
			})
		}
	})
}
