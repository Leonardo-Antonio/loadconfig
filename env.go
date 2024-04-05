package loadconfig

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// .env
const ENV_DEFAULT = ".env"

// .env.local
const ENV_LOCAL = ".env.local"

// .env.dev
const ENV_DEV = ".env.dev"

// .env.qa
const ENV_QA = ".env.qa"

// .env.uat
const ENV_UAT = ".env.uat"

type env struct {
	filename string
}

func newEnv(filename string) *env {
	return &env{filename: filename}
}

// Spain: Carga las variables de entorno del archivo especificado
// English: Load environment variables from the specified file
func (e *env) Load() {
	buff, err := os.ReadFile(e.filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(buff)))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.Contains(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		if err := os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])); err != nil {
			panic(err)
		}
	}

	fmt.Println("✅ Environments: " + e.filename)
}
