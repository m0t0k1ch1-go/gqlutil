package gqlutil_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/m0t0k1ch1-go/gqlutil"
	"github.com/m0t0k1ch1-go/gqlutil/internal/testutil"
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
				buf := &bytes.Buffer{}

				tc.in.MarshalGQL(buf)

				testutil.Equal(t, tc.out, buf.Bytes())
			})
		}
	})
}

func TestInt64UnmarshalGQL(t *testing.T) {
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
				var gi gqlutil.Int64

				if err := gi.UnmarshalGQL(tc.in); err != nil {
					t.Fatal(err)
				}

				testutil.Equal(t, tc.out, gi)
			})
		}
	})
}
