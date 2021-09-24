package week2

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// OperateDB 操作 DB，这里模拟查询操作
// q 查询的 SQL 语句
func OperateDB(q string) error {
	fmt.Printf("sql: %s\n", q)
	return sql.ErrNoRows
}

func WrapError() error {
	err := OperateDB("select user_id,user_name from user where id=1")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			notFoundErr := NewErr(NotFound, "Not found")
			return errors.Wrap(notFoundErr, err.Error())
		}
		databaseErr := NewErr(Database, "Database error")
		return errors.Wrap(databaseErr, err.Error())
	}

	return nil
}

func HandleError() error {
	err := WrapError()
	if err != nil {
		fmt.Printf("WrapError has an error[%T]: %v.\nstack: %+v", errors.Cause(err), errors.Cause(err), err)
		if ErrCode(errors.Cause(err)) == NotFound {
			return nil
		}
		return err
	}
	return nil
}
