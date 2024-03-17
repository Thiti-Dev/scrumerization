package time

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTime(b time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(fmt.Sprintf(`"%s"`, b.Format(time.RFC3339))))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	timeAsISO, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("time must be a valid ISO string")
	}

	parsedTime, err := time.Parse(time.RFC3339, timeAsISO)
	if err != nil {
		return time.Time{}, fmt.Errorf("time must be a valid ISO string")
	}

	return parsedTime, nil
}
