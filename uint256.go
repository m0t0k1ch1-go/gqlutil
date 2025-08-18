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
func (gx Uint256) Unwrap() bigutil.Uint256 {
	return bigutil.Uint256(gx)
}

// String implements the fmt.Stringer interface.
func (gx Uint256) String() string {
	return gx.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (gx Uint256) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(gx.string()))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (gx *Uint256) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.Errorf("must be a string")
	}

	x, err := bigutil.NewUint256FromHex(s)
	if err != nil {
		return err
	}

	*gx = Uint256(x)

	return nil
}

func (gx Uint256) string() string {
	return bigutil.Uint256(gx).String()
}
