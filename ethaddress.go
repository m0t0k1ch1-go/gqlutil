package gqlutil

import (
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

// MarshalEthAddress returns a graphql.Marshaler that encodes a go-ethereum/common.Address as a quoted checksummed hex string.
func MarshalEthAddress(address ethcommon.Address) graphql.Marshaler {
	return graphql.MarshalString(address.Hex())
}

// UnmarshalEthAddress decodes a GraphQL String (0x/0X-prefixed hex) into a go-ethereum/common.Address.
func UnmarshalEthAddress(v any) (ethcommon.Address, error) {
	if v == nil {
		return ethcommon.Address{}, errors.New("invalid graphql value: nil")
	}

	s, ok := v.(string)
	if !ok {
		return ethcommon.Address{}, fmt.Errorf("unsupported graphql value type: %T", v)
	}
	if len(s) == 0 {
		return ethcommon.Address{}, errors.New("invalid graphql string: empty")
	}
	if ok := ethcommon.IsHexAddress(s); !ok {
		return ethcommon.Address{}, errors.New("invalid graphql string: invalid eth address")
	}

	return ethcommon.HexToAddress(s), nil
}
