package loadconfig

import (
	encoding_json "encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// env.json
const JSON_DEFAULT = "env.json"

// env.local.json
const JSON_LOCAL = "env.local.json"

// env.dev.json
const JSON_DEV = "env.dev.json"

// env.qa.json
const JSON_QA = "env.qa.json"

// env.uat.json
const JSON_UAT = "env.uat.json"

type json struct {
	filename string
}

func newJSON(filename string) *json {
	parts := strings.Split(filename, ".")
	dotfile := parts[len(parts)-1]
	if dotfile != "json" {
		panic(fmt.Sprintf("❌ error: The file entered is not valid, the file name must end with %s\n", ".json"))
	}

	return &json{filename: filename}
}

// Spain: Carga las variables de entorno del archivo especificado
// English: Load environment variables from the specified file
func (e *json) Load() error {
	buff, err := os.ReadFile(e.filename)
	if err != nil {
		return err
	}

	var config map[string]interface{}
	err = encoding_json.Unmarshal(buff, &config)
	if err != nil {
		return err
	}

	for key, value := range config {
		if err := os.Setenv(strings.TrimSpace(key), strings.TrimSpace(fmt.Sprint(value))); err != nil {
			return err
		}
	}

	log.Println("✅ Environments: " + e.filename)
	return nil
}
