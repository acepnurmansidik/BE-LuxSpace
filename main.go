package main

import (
	"LuxSpace/configs"
	"fmt"
)

func main() {
	db, err := configs.Connection()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(db)
}
