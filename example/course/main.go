package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/jonathannavas/go_course_sdk/course"
)

func main() {
	courseTrans := courseSdk.NewHttpClient("http://localhost:8082", "")

	course, err := courseTrans.Get("684cd79d-037b-42b6-a706-f133426fcfeb")

	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal server error", err.Error())
		os.Exit(1)
	}
	fmt.Println(course)
}
