package tserror

var (
	ERR_UNKNOWN                         = &TSError{ErrorTypeInternal, "ERR_UNKNOWN", ""}
	ERR_SALES_PERSON_ALREADY_REGISTERED = &TSError{ErrorTypeInvalidArgument, "ERR_SALES_PERSON_ALREADY_REGISTERED", ""}
)
