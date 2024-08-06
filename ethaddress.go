package gqlutil

import (
	"io"
	"strconv"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/samber/oops"
)

// EthAddress is a custom scalar for ethcommon.Address.
type EthAddress ethcommon.Address

// String implements the fmt.Stringer interface.
func (gea EthAddress) String() string {
	return gea.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (gea EthAddress) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(gea.string()))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (gea *EthAddress) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.Errorf("must be a string")
	}

	if ok := ethcommon.IsHexAddress(s); !ok {
		return oops.Errorf("invalid eth address")
	}

	*gea = EthAddress(ethcommon.HexToAddress(s))

	return nil
}

func (gea EthAddress) string() string {
	return ethcommon.Address(gea).Hex()
}
