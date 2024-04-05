# Load Config

### Description:
- **English:** “Library for loading environment variables from the specified file (.env, .env.local, etc).”
- **Español:** “Biblioteca para cargar las variables de entorno desde el archivo especificado (.env, .env.local, config.json, config.yml, etc).”

### Install
```bash
go get github.com/Leonardo-Antonio/loadconfig
```

### Use 
[Example](/examples/env.go)
```go
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

```
