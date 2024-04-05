package loadconfig

type driver string

const TYPE_FILE_DOT_ENV driver = "ENVIRONMENT"
const TYPE_FILE_JSON driver = "JSON"
const TYPE_FILE_YML driver = "YML"

func New(typefile driver, filename string) iLoadConfig {
	switch typefile {
	case TYPE_FILE_DOT_ENV:
		return newEnv(filename)
	}

	panic("error: driver not implemented")
}
