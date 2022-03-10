package scalars

import (
	"fmt"
	"io"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Username string

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *Username) UnmarshalGQL(v interface{}) error {
	username, ok := v.(string)
	if !ok {
		return fmt.Errorf("username must be a string")
	}

	err := validation.Validate(username,
		validation.Required,
		validation.Length(2, 32),
	)

	if err != nil {
		return err
	}

	*u = Username(username)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (u Username) MarshalGQL(w io.Writer) {
	_, err := w.Write([]byte(u))
	if err != nil {
		return
	}
}
