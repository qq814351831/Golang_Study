package utils

import (
	"fmt"
	"os"
)

//打印错误
func HandlerError(err error,when string){
	if err != nil {
		fmt.Println(when,err)
		os.Exit(1)
	}
}
