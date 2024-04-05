package main

import (
	"fmt"
	"os"

	"github.com/Leonardo-Antonio/loadconfig"
)

func main() {
	loadconfig.New(loadconfig.TYPE_FILE_DOT_ENV, loadconfig.ENV_LOCAL).Load()
	fmt.Println(os.Getenv("HELLO"))
	/*
	* output:
	* ✅ Environments: .env.local
	* WORLD
	 */
}
