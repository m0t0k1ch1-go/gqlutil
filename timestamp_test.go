package gqlutil_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/m0t0k1ch1-go/timeutil/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/gqlutil"
)

func TestTimestampMarshalGQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   gqlutil.Timestamp
			out  []byte
		}{
			{
				"valid",
				gqlutil.Timestamp(timeutil.NewTimestamp(time.Unix(1231006505, 0))),
				[]byte(`"1231006505"`),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var buf bytes.Buffer
				tc.in.MarshalGQL(&buf)

				assert.Equal(t, tc.out, buf.Bytes())
			})
		}
	})
}

func TestTimestampUnmarshalGQL(t *testing.T) {
	t.Run("failure", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
		}{
			{
				"int",
				0,
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var g gqlutil.Timestamp
				{
					err := g.UnmarshalGQL(tc.in)
					assert.Error(t, err)
				}
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
			out  gqlutil.Timestamp
		}{
			{
				name: "valid",
				in:   "1231006505",
				out:  gqlutil.Timestamp(timeutil.NewTimestamp(time.Unix(1231006505, 0))),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var g gqlutil.Timestamp
				{
					err := g.UnmarshalGQL(tc.in)
					require.NoError(t, err)
				}

				require.Equal(t, tc.out, g)
			})
		}
	})
}
