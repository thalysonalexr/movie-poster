package errors

import "errors"

// FailedGetResourceMovies custom error
var FailedGetResourceMovies = errors.New("Failed to get resource movies")

// FailedToReadResponse custom error
var FailedToReadResponse = errors.New("Failed to read response")
