package main

import (
	"fmt"
	"mygoproject/types"
	"mygoproject/util"
)

func main() {
	user := types.User{
		Username: util.GetUserName(),
		Age:      util.GetAge(),
	}
	fmt.Printf("the number is  %+v", user)
}
