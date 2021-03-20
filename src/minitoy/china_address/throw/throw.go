package throw

import (
	"fmt"
	"os"
)

func Throw(err error, msg string) {
	if msg != "" {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(msg)
				os.Exit(-1)
			}
		}()
	}

	if err != nil {
		panic(err)
	}
}
