// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldIntro holds the string denoting the intro field in the database.
	FieldIntro = "intro"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLastLoginedAt holds the string denoting the last_logined_at field in the database.
	FieldLastLoginedAt = "last_logined_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldName,
	FieldGender,
	FieldIntro,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLastLoginedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMALE   Gender = "MALE"
	GenderFEMALE Gender = "FEMALE"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMALE, GenderFEMALE:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for gender field: %q", ge)
	}
}
