package gqlutil

import (
	"io"
	"strconv"
	"time"

	"github.com/m0t0k1ch1-go/timeutil/v5"
	"github.com/samber/oops"
)

// Timestamp is a custom scalar for github.com/m0t0k1ch1-go/timeutil/v4.Timestamp.
type Timestamp timeutil.Timestamp

// Unwrap returns the github.com/m0t0k1ch1-go/timeutil/v4.Timestamp.
func (g Timestamp) Unwrap() timeutil.Timestamp {
	return timeutil.Timestamp(g)
}

// String implements the fmt.Stringer interface.
func (g Timestamp) String() string {
	return g.string()
}

// MarshalGQL implements the github.com/99designs/gqlgen/graphql.Marshaler interface.
func (g Timestamp) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(g.string()))
}

// UnmarshalGQL implements the github.com/99designs/gqlgen/graphql.Unmarshaler interface.
func (g *Timestamp) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return oops.New("v must be string")
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return oops.Wrap(err)
	}

	*g = Timestamp(timeutil.NewTimestamp(time.Unix(i, 0)))

	return nil
}

func (g Timestamp) string() string {
	return timeutil.Timestamp(g).String()
}
