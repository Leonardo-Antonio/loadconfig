# Load Config

### Description:
- **English:** “Library for loading environment variables from the specified file (.env, .env.local, etc).”
- **Español:** “Biblioteca para cargar las variables de entorno desde el archivo especificado (.env, .env.local, config.json, config.yml, etc).”

### Install
```bash
go get github.com/Leonardo-Antonio/loadconfig
```

### Use 

#### .env
```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/loadconfig"
)

func main() {
	// Example .env
	err := loadconfig.New(loadconfig.TYPE_FILE_DOT_ENV, loadconfig.ENV_LOCAL).Load() // loadconfig.ENV_LOCAL = .env.local
	if err != nil {
		panic(err)
	}
	
	fmt.Println(loadconfig.GetEnv("HELLO").String())                          // out: WORLD
	/*
	* output:
	* ✅ Environments: .env.local
	* WORLD
	 */
}
```

#### json
```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/loadconfig"
)

func main() {
	err := loadconfig.New(loadconfig.TYPE_FILE_JSON, loadconfig.JSON_LOCAL).Load() // loadconfig.JSON_LOCAL = env.local.json
	if err != nil {
		panic(err)
	}
	
	fmt.Println(loadconfig.GetEnv("HELLO").String())                        // out: WORLD
	/*
	* output:
	* ✅ Environments: env.local.json
	* WORLD
	 */

	loadconfig.New(loadconfig.TYPE_FILE_YML, loadconfig.YAML_LOCAL).Load() // loadconfig.JSON_LOCAL = env.local.yml
	fmt.Println(loadconfig.GetEnv("HELLO").String())                       // out: WORLD
	/*
	* output:
	* ✅ Environments: env.local.yml
	* WORLD
	 */
}
```


#### yml
```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/loadconfig"
)

func main() {
	err := loadconfig.New(loadconfig.TYPE_FILE_YAML, loadconfig.YAML_LOCAL).Load() // loadconfig.JSON_LOCAL = env.local.yml
	if err != nil {
		panic(err)
	}

	fmt.Println(loadconfig.GetEnv("HELLO").String())                       // out: WORLD
	/*
	* output:
	* ✅ Environments: env.local.yml
	* WORLD
	 */
}
```