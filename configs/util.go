package configs

import (
	"fmt"
	"os"
)

func EnvLookup(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("[ENV CONFIG]: %s not found in environment variable", key))
	}
	return val
}
