package gqlutil

import (
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

// MarshalEthAddress returns a [graphql.Marshaler] that encodes address as the quoted string returned by [ethcommon.Address.Hex].
func MarshalEthAddress(address ethcommon.Address) graphql.Marshaler {
	return graphql.MarshalString(address.Hex())
}

// UnmarshalEthAddress decodes a hexadecimal string representing an Ethereum address into an [ethcommon.Address].
// The string must be accepted by [ethcommon.IsHexAddress].
func UnmarshalEthAddress(v any) (ethcommon.Address, error) {
	if v == nil {
		return ethcommon.Address{}, errors.New("unsupported input: nil")
	}

	s, ok := v.(string)
	if !ok {
		return ethcommon.Address{}, fmt.Errorf("unsupported input type: %T", v)
	}
	if len(s) == 0 {
		return ethcommon.Address{}, errors.New("invalid string input: empty")
	}
	if ok := ethcommon.IsHexAddress(s); !ok {
		return ethcommon.Address{}, errors.New("invalid string input: must be an eth address")
	}

	return ethcommon.HexToAddress(s), nil
}
