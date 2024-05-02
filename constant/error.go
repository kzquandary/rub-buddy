package constant

import "errors"

var ErrInsertDatabase error = errors.New("Invalid Add Data in Database")
var ErrEmptyInput error = errors.New("name, email and password cannot be empty")

