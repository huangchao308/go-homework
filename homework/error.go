package homework

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func OperateDB() error {
	return sql.ErrNoRows
}

func WrapError() error {
	err := OperateDB()
	if err != nil {
		return errors.Wrap(err, "WrapError.")
	}
	return nil
}

func HandleError() error {
	err := WrapError()
	if err != nil {
		fmt.Printf("WrapError has an error[%T]: %v.\nstack: %+v", errors.Cause(err), errors.Cause(err), err)
	}

	return nil
}
