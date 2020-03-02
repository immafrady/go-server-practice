package logger

import (
	"log"
)

func CheckError(errMsg string, errLike ...interface{}) {
	for i := 0; i < len(errLike); i++ {
		switch v := errLike[i].(type) {
		case error:
			log.Printf("Error Occur[%s]: %v \n", errMsg, v)
		}
	}
}
