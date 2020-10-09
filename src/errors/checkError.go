package errors

import "fmt"

func CheckError(err error){
	defer func() {
		if ins,ok := recover().(error);ok{
			fmt.Println("程序出现异常： ",ins.Error())
		}
	}()
	if err  != nil{
		 panic(err)
	}
}