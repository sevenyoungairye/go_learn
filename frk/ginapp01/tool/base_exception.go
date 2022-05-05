package tool

import (
	"errors"
	"fmt"
)

func CatchErr(catch func(errString string)) error {
	r := recover()
	if r != nil {
		// 打印错误
		errMsg := fmt.Sprintln(r)
		// 调用catch
		catch(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func TryCatch(Try func(), Catch func(errString string)) error {
	var err error = nil
	func() {
		Try()
		defer func() {
			err = CatchErr(Catch)
		}()
	}()
	return err
}
