package graphql

import (
	"fmt"
	"io"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"gopkg.in/guregu/null.v4"
)

func MarshalDecimal(v null.String) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		b, err := v.MarshalJSON()
		if err != nil {
			log.Fatalf("failed to marshal Decimal value: %v", err)
		}
		w.Write(b)
	})
}

func UnmarshalDecimal(v interface{}) (null.String, error) {
	switch v := v.(type) {
	case string:
		return null.StringFrom(v), nil
	case *string:
		return null.StringFromPtr(v), nil
	default:
		return null.String{}, fmt.Errorf("%T is not a string or *string and cannot be unmarshalled into Decimal", v)
	}
}
