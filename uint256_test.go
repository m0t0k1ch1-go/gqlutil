package gqlutil_test

import (
	"bytes"
	"testing"

	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/m0t0k1ch1-go/bigutil/v2"

	"github.com/m0t0k1ch1-go/gqlutil"
	"github.com/m0t0k1ch1-go/gqlutil/internal/testutil"
)

func TestUint256MarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   gqlutil.Uint256
			out  []byte
		}{
			{
				"min",
				gqlutil.Uint256(bigutil.MustHexToUint256("0x0")),
				[]byte(`"0x0"`),
			},
			{
				"max",
				gqlutil.Uint256(bigutil.MustBigIntToUint256(ethmath.MaxBig256)),
				[]byte(`"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"`),
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

func TestUint256UnmarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   string
			out  gqlutil.Uint256
		}{
			{
				"min",
				"0x0",
				gqlutil.Uint256(bigutil.MustHexToUint256("0x0")),
			},
			{
				"max",
				"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
				gqlutil.Uint256(bigutil.MustBigIntToUint256(ethmath.MaxBig256)),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var gx gqlutil.Uint256

				if err := gx.UnmarshalGQL(tc.in); err != nil {
					t.Fatal(err)
				}

				testutil.Equal(t, gx.Unwrap().BigInt().Cmp(tc.out.Unwrap().BigInt()), 0)
			})
		}
	})
}
