package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/jonathannavas/go_course_sdk/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("0dbf99c6-504e-4254-a89a-7ff6726219a2")

	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal server error", err.Error())
		os.Exit(1)
	}
	fmt.Println(user)
}
