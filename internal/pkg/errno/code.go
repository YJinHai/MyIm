package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

	//activity errors
	ErrUserIncorrect = &Errno{Code: 20201, Message: "The user was incorrect."}
	ErrMemberExist = &Errno{Code: 20202, Message: "The members was already exist ."}


	//team errors
	ErrTeamNotFound = &Errno{Code: 20301, Message: "The team was not found."}
	ErrTeamNotCreator = &Errno{Code: 20302, Message: "This user is not the creator of the team."}
	ErrTeamMemberFull = &Errno{Code: 20303, Message: "The team is full."}

	// disciple errors
	ErrDiscipleNotFound = &Errno{Code: 20401, Message: "The disciple was not found."}
	ErrDiscipleExisting = &Errno{Code: 20402, Message: "This disciple was existing."}

)