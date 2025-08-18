package gqlutil

import (
	"io"
	"strconv"

	"github.com/m0t0k1ch1-go/bigutil/v3"
	"github.com/samber/oops"
)

// Uint256 is a custom scalar for bigutil.Uint256.
type Uint256 bigutil.Uint256

// Unwrap returns the bigutil.Uint256.
func (g Uint256) Unwrap() bigutil.Uint256 {
	return bigutil.Uint256(g)
}

// String implements the fmt.Stringer interface.
func (g Uint256) String() string {
	return g.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (g Uint256) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(g.string()))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (g *Uint256) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.New("v must be string")
	}

	x256, err := bigutil.NewUint256FromHex(s)
	if err != nil {
		return oops.Wrap(err)
	}

	*g = Uint256(x256)

	return nil
}

func (g Uint256) string() string {
	return bigutil.Uint256(g).String()
}
