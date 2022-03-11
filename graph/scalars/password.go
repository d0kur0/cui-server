package scalars

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Password string

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (p *Password) UnmarshalGQL(v interface{}) error {
	password, ok := v.(string)
	if !ok {
		return fmt.Errorf("password must be a string")
	}

	err := validation.Validate(password,
		validation.Required,
		validation.Length(6, 128),
	)

	if err != nil {
		return err
	}

	*p = Password(password)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (p Password) MarshalGQL(w io.Writer) {
	asJson, err := json.Marshal(p)
	if err != nil {
		log.Printf("error on Marshal password scalar type; %s", err)
		return
	}

	if _, err = w.Write(asJson); err != nil {
		log.Printf("error on write password scalar type; %s", err)
		return
	}
}
