package loadconfig

type iLoadConfig interface {
	Load() error
}
