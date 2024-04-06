package loadconfig

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// .env
const ENV_DEFAULT = ".env"

// .env.local
const ENV_LOCAL = ENV_DEFAULT + ".local"

// .env.dev
const ENV_DEV = ENV_DEFAULT + ".dev"

// .env.qa
const ENV_QA = ENV_DEFAULT + ".qa"

// .env.uat
const ENV_UAT = ENV_DEFAULT + ".uat"

type env struct {
	filename string
}

func newEnv(filename string) *env {
	if !strings.Contains(filename, ENV_DEFAULT) {
		panic(fmt.Sprintf("❌ error: the file entered is not valid, the file name must start with %s\n", ENV_DEFAULT))
	}

	return &env{filename: filename}
}

// Spain: Carga las variables de entorno del archivo especificado
// English: Load environment variables from the specified file
func (e *env) Load() error {
	buff, err := os.ReadFile(e.filename)
	if err != nil {
		return err
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
			return err
		}
	}

	log.Println("✅ Environments: " + e.filename)
	return nil
}
