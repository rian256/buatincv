package helper

import "log"

func ErrorHelper(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
