package loadconfig

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// env.yml
const YAML_DEFAULT = "env.yml"

// env.local.yml
const YAML_LOCAL = "env.local.yml"

// env.dev.yml
const YAML_DEV = "env.dev.yml"

// env.qa.yml
const YAML_QA = "env.qa.yml"

// env.uat.yml
const YAML_UAT = "env.uat.yml"

type yml struct {
	filename string
}

func newYML(filename string) *yml {
	parts := strings.Split(filename, ".")
	dotfile := parts[len(parts)-1]
	if dotfile != "yaml" && dotfile != "yml" {
		panic(fmt.Sprintf("❌ error: The file entered is not valid, the file name must end with %s\n", ".yaml or .yml"))
	}

	return &yml{filename: filename}
}

// Spain: Carga las variables de entorno del archivo especificado
// English: Load environment variables from the specified file
func (e *yml) Load() error {
	buff, err := os.ReadFile(e.filename)
	if err != nil {
		return err
	}

	var config map[string]interface{}
	err = yaml.Unmarshal(buff, &config)
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
