package errors

const (
	Unknown             = "UNKNOWN"
	TokenExpired        = "TOKEN_EXPIRED"
	TokenInvalid        = "TOKEN_INVALID"
	FieldMissing        = "FIELD_MISSING"
	FieldInvalid        = "FIELD_INVALID"
	CredentialsNotFound = "CREDENTIALS_NOT_FOUND"
	NotFound            = "NOT_FOUND"
	RelatedItemsExist   = "RELATED_ITEMS_EXIST"
	CannotBeDeleted     = "CANNOT_BE_DELETED"
	Forbidden           = "FORBIDDEN"
)

var messages = map[string]string{
	Unknown:             "Any unknown error that this list has not covered",
	TokenExpired:        "Access or refresh token is expired",
	TokenInvalid:        "Access or refresh token is invalid",
	FieldMissing:        "One or more fields are missing",
	FieldInvalid:        "One or more fields are invalid",
	CredentialsNotFound: "Provided credentials does not match any user",
	NotFound:            "",
	Forbidden:           "You don't have permissions",

	RelatedItemsExist: "Item to delete has constrained items, " +
		"delete them before",
	CannotBeDeleted: "At least one address must remain. To delete this address," +
		" first create a new address.",
}