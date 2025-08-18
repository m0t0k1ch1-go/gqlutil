package gqlutil

import (
	"io"
	"strconv"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/samber/oops"
)

// EthAddress is a custom scalar for github.com/ethereum/go-ethereum/common.Address.
type EthAddress ethcommon.Address

// Unwrap returns the github.com/ethereum/go-ethereum/common.Address.
func (g EthAddress) Unwrap() ethcommon.Address {
	return ethcommon.Address(g)
}

// String implements the fmt.Stringer interface.
func (g EthAddress) String() string {
	return g.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (g EthAddress) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(g.string()))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (g *EthAddress) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.New("v must be string")
	}
	if ok := ethcommon.IsHexAddress(s); !ok {
		return oops.New("v must be valid eth address")
	}

	*g = EthAddress(ethcommon.HexToAddress(s))

	return nil
}

func (g EthAddress) string() string {
	return ethcommon.Address(g).Hex()
}
