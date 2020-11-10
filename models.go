package main

import "errors"

var (
	// ErrCouldNotUnmarshalRequest is thrown when an invalid request object is provided
	ErrCouldNotUnmarshalRequest = errors.New("request unable to be unmarshalled")
	// ErrFailedToInitializeDependencies is thrown when handler fails to initialize dependencies
	ErrFailedToInitializeDependencies = errors.New("failed to initialize dependencies")
	// ErrFailedToInitializeLogger is thrown when logger fails to initialize
	ErrFailedToInitializeLogger = errors.New("failed to initialize logger")
)

// ExampleRequest is an example of a request object that might be passed in as part of an API request
type ExampleRequest struct {
	LoanID     float64 `json:"loanID"`                      // sensitivity flag ommitted on purpose (sensitive:"false" is default)
	LoanAmount float64 `json:"loanAmount" sensitive:"true"` // sensitivity flag declared on purpose to omit this entry
}
