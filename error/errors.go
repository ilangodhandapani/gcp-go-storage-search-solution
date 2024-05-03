package error

import (
	"errors"
)

var ErrNoAuthorizationHeaderFound = errors.New("no authorization header passed for request")
var ErrTokenExpired = errors.New("token expired")
var ErrUnauthorized = errors.New("user unauthorized")
var ErrFileNotFound = errors.New("file not found")
var ErrInvalidJson = errors.New("invalid json string")
var ErrNoUsername = errors.New("no username passed for request")
var ErrInvalidAudience = errors.New("invalid audeince")
var ErrInvalidIssuer = errors.New("invalid issuer")
var ErrUnmarshalByteArray = errors.New("cant unmarshal the byte array")
