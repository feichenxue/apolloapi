package main

import (
	"fmt"

	"github.com/feichenxue/apolloapi"
)

func main() {

	url := "http://10.191.20.21:9500"
	user := "apollo"
	passwds := "1qaz2wsx"

	apollo := apolloapi.NewApollo(url, user, passwds)

	ok, err := apollo.ReleaseApollo("go-test4", "prd", "test", "default", "application")

	if err != nil {
		fmt.Println(err)
	}

	if ok {
		fmt.Println("发布成功!!!")
	}

}
