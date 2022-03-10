package scalars

import (
	"fmt"
	"io"

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
	_, err := w.Write([]byte(p))
	if err != nil {
		return
	}
}
