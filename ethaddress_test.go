package gqlutil_test

import (
	"bytes"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/gqlutil/v2"
)

func TestMarshalEthAddress(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   ethcommon.Address
			want string
		}{
			{
				"vitalik.eth",
				ethcommon.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"),
				`"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"`,
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var buf bytes.Buffer
				gqlutil.MarshalEthAddress(tc.in).MarshalGQL(&buf)
				require.Equal(t, tc.want, buf.String())
			})
		}
	})
}

func TestUnmarshalEthAddress(t *testing.T) {
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
				"invalid graphql string: invalid eth address",
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				_, err := gqlutil.UnmarshalEthAddress(tc.in)
				require.ErrorContains(t, err, tc.want)
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
			want ethcommon.Address
		}{
			{
				"vitalik.eth",
				"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
				ethcommon.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				address, err := gqlutil.UnmarshalEthAddress(tc.in)
				require.NoError(t, err)
				require.Equal(t, tc.want, address)
			})
		}
	})
}
