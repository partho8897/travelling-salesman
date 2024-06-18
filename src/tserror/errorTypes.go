package tserror

type TSErrorTypes uint32

const (
	ErrorTypeInternal         TSErrorTypes = 0
	ErrorTypeInvalidArgument  TSErrorTypes = 1
	ErrorTypeNotSupported     TSErrorTypes = 2
	ErrorTypeInvalidOperation TSErrorTypes = 3
)
