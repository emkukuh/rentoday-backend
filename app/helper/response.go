package response

import (
	"strings"
)

type Response struct {	
	Status	bool		`json:"status"`
	Message	string		`json:"message"`
	Errors	interface{}	`json:"errors,omitempty"`
	Data	interface{}	`json:"data"`
}

type EmptyObj struct {}

func BuildResponse(status bool, data interface{}) Response {
	res := Response {
		Status: status,
		Message: "OK",
		Errors: nil,
		Data: data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status: false,
		Message: message,
		Errors: splittedError,
		Data: data,
	}
	return res
}

func BuildBadRequestError(err string) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status: false,
		Message: "nil",
		Errors: splittedError,
		Data: nil,
	}
	return res
}