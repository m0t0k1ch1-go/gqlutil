package gqlutil_test

import (
	"bytes"
	"testing"

	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/m0t0k1ch1-go/bigutil/v3"
	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/gqlutil"
)

func TestUint256MarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   gqlutil.Uint256
			out  []byte
		}{
			{
				"zero",
				gqlutil.Uint256(bigutil.NewUint256FromUint64(0)),
				[]byte(`"0x0"`),
			},
			{
				"max",
				gqlutil.Uint256(bigutil.MustNewUint256(ethmath.MaxBig256)),
				[]byte(`"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"`),
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

func TestUint256UnmarshalGQL(t *testing.T) {
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
				var g gqlutil.Uint256
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
			in   any
			out  gqlutil.Uint256
		}{
			{
				"zero",
				"0x0",
				gqlutil.Uint256(bigutil.NewUint256FromUint64(0)),
			},
			{
				"max",
				"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
				gqlutil.Uint256(bigutil.MustNewUint256(ethmath.MaxBig256)),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var g gqlutil.Uint256
				{
					err := g.UnmarshalGQL(tc.in)
					require.NoError(t, err)
				}

				require.Zero(t, g.Unwrap().BigInt().Cmp(tc.out.Unwrap().BigInt()))
			})
		}
	})
}
