// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1Language v1 language
//
// swagger:model v1Language
type V1Language string

func NewV1Language(value V1Language) *V1Language {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1Language.
func (m V1Language) Pointer() *V1Language {
	return &m
}

const (

	// V1LanguageLANGENGLISH captures enum value "LANG_ENGLISH"
	V1LanguageLANGENGLISH V1Language = "LANG_ENGLISH"

	// V1LanguageLANGRUSSIAN captures enum value "LANG_RUSSIAN"
	V1LanguageLANGRUSSIAN V1Language = "LANG_RUSSIAN"

	// V1LanguageLANGESPANOL captures enum value "LANG_ESPANOL"
	V1LanguageLANGESPANOL V1Language = "LANG_ESPANOL"

	// V1LanguageLANGITALIAN captures enum value "LANG_ITALIAN"
	V1LanguageLANGITALIAN V1Language = "LANG_ITALIAN"
)

// for schema
var v1LanguageEnum []interface{}

func init() {
	var res []V1Language
	if err := json.Unmarshal([]byte(`["LANG_ENGLISH","LANG_RUSSIAN","LANG_ESPANOL","LANG_ITALIAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1LanguageEnum = append(v1LanguageEnum, v)
	}
}

func (m V1Language) validateV1LanguageEnum(path, location string, value V1Language) error {
	if err := validate.EnumCase(path, location, value, v1LanguageEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 language
func (m V1Language) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1LanguageEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 language based on context it is used
func (m V1Language) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
