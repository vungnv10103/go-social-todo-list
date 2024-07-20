package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppErr struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"root_error"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewFullErrRes(statusCode int, root error, msg, log, key string) *AppErr {
	return &AppErr{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrRes(root error, msg, log, key string) *AppErr {
	return &AppErr{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorizedRes(root error, msg, log, key string) *AppErr {
	return &AppErr{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func (e *AppErr) RootError() error {
	var err *AppErr
	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppErr) Error() string {
	return e.RootError().Error()
}

func NewCustomError(root error, msg, key string) *AppErr {
	if root != nil {
		return NewErrRes(root, msg, root.Error(), key)
	}
	return NewErrRes(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppErr {
	return NewFullErrRes(http.StatusInternalServerError, err, "somethings went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppErr {
	return NewErrRes(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternalRequest(err error) *AppErr {
	return NewFullErrRes(http.StatusInternalServerError, err, "somethings went wrong in the server", err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrrCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrrCannotDelete%s", entity),
	)
}
func ErrCannotUpdateEntity(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrrCannotUpdate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrrCannotGet%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Errr%sDeleted", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Errr%sAlreadyExists", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Errr%sNotFound", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrrCannotCreate%s", entity),
	)
}

func ErrNoPermission(entity string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrNoPermission"),
	)
}

var RecordNotFound = errors.New("record not found")
