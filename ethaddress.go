package gqlutil

import (
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

// MarshalEthAddress returns a [graphql.Marshaler] that encodes address as a quoted EIP-55 compliant hexadecimal string.
func MarshalEthAddress(address ethcommon.Address) graphql.Marshaler {
	return graphql.MarshalString(address.Hex())
}

// UnmarshalEthAddress decodes a hexadecimal string representing an Ethereum address into an [ethcommon.Address].
// The string must be accepted by [ethcommon.IsHexAddress].
func UnmarshalEthAddress(v any) (ethcommon.Address, error) {
	if v == nil {
		return ethcommon.Address{}, errors.New("unsupported value: nil")
	}

	s, ok := v.(string)
	if !ok {
		return ethcommon.Address{}, fmt.Errorf("unsupported value type: %T", v)
	}
	if len(s) == 0 {
		return ethcommon.Address{}, errors.New("invalid eth address string: empty")
	}
	if ok := ethcommon.IsHexAddress(s); !ok {
		return ethcommon.Address{}, errors.New("invalid eth address string")
	}

	return ethcommon.HexToAddress(s), nil
}
