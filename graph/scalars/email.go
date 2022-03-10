package scalars

import (
	"fmt"
	"io"

	"github.com/go-ozzo/ozzo-validation/is"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Email string

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (e *Email) UnmarshalGQL(v interface{}) error {
	email, ok := v.(string)
	if !ok {
		return fmt.Errorf("email must be a string")
	}

	err := validation.Validate(email,
		validation.Required,
		is.Email,
	)

	if err != nil {
		return err
	}

	*e = Email(email)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (e Email) MarshalGQL(w io.Writer) {
	_, err := w.Write([]byte(e))
	if err != nil {
		return
	}
}
