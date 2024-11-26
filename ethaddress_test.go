package gqlutil_test

import (
	"bytes"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/gqlutil"
)

func TestEthAddressMarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   gqlutil.EthAddress
			out  []byte
		}{
			{
				"valid",
				gqlutil.EthAddress(ethcommon.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")),
				[]byte(`"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"`),
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

func TestEthAddressUnmarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   string
			out  gqlutil.EthAddress
		}{
			{
				"valid",
				"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
				gqlutil.EthAddress(ethcommon.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var g gqlutil.EthAddress
				require.Nil(t, g.UnmarshalGQL(tc.in))

				require.Equal(t, tc.out, g)
			})
		}
	})
}
