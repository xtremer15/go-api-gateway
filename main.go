package main

import (
	common_utils "api-gateway/pkg/common-utils"
	"fmt"
)

func main() {

	common_utils.ReadFile("assets/schemas.json")
	fmt.Println("\n")
	common_utils.ReadJsonFile("envs/dev_env.json")
}
