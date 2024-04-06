package loadconfig

import (
	"log"
	"os"
	"strconv"
)

type Value interface {
	bool | string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func logIfError[T Value](key string, val T, err error) T {
	if err != nil {
		log.Printf("❌ loadconfig => error getting config environments (%s) %s\n", key, err.Error())
		return val
	}

	return val
}

type environment struct{ key string }

func GetEnv(key string) environment { return environment{key: key} }

func (e environment) Bool() bool {
	env := os.Getenv(e.key)
	value, err := strconv.ParseBool(env)
	return logIfError(e.key, value, err)
}

func (e environment) String() string {
	return os.Getenv(e.key)
}

func (e environment) Int() int {
	env := os.Getenv(e.key)
	value, err := strconv.Atoi(env)
	return logIfError(e.key, value, err)
}

func (e environment) Int16() int16 {
	env := os.Getenv(e.key)
	value, err := strconv.Atoi(env)
	return logIfError(e.key, int16(value), err)
}

func (e environment) Int32() int32 {
	env := os.Getenv(e.key)
	value, err := strconv.Atoi(env)
	return logIfError(e.key, int32(value), err)
}

func (e environment) Int64() int64 {
	env := os.Getenv(e.key)
	value, err := strconv.Atoi(env)
	return logIfError(e.key, int64(value), err)
}

// uint
func (e environment) Uint() uint {
	env := os.Getenv(e.key)
	value, err := strconv.ParseUint(env, 10, 64)
	return logIfError(e.key, uint(value), err)
}

func (e environment) Uint8() uint8 {
	env := os.Getenv(e.key)
	value, err := strconv.ParseUint(env, 10, 64)
	return logIfError(e.key, uint8(value), err)
}

func (e environment) Uint16() uint16 {
	env := os.Getenv(e.key)
	value, err := strconv.ParseUint(env, 10, 64)
	return logIfError(e.key, uint16(value), err)
}

func (e environment) Uint32() uint32 {
	env := os.Getenv(e.key)
	value, err := strconv.ParseUint(env, 10, 64)
	return logIfError(e.key, uint32(value), err)
}

func (e environment) Uint64() uint64 {
	env := os.Getenv(e.key)
	value, err := strconv.ParseUint(env, 10, 64)
	return logIfError(e.key, uint64(value), err)
}

// float
func (e environment) Float32() float32 {
	env := os.Getenv(e.key)
	value, err := strconv.ParseFloat(env, 32)
	return logIfError(e.key, float32(value), err)
}

func (e environment) Float64() float64 {
	env := os.Getenv(e.key)
	value, err := strconv.ParseFloat(env, 64)
	return logIfError(e.key, float64(value), err)
}
