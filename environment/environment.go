package environment

import (
	"fmt"
	"github.com/oledakotajoe/codenvi-core/types"
	"os"
)

func WithEnv(env map[string]string, body *types.Closure) {
	var initialEnvironment = make(map[string]string)
	var customEnvironment =  make(map[string]string)
	fmt.Println("Printing a test line")
	for key, value := range env {
		initialValue := os.Getenv(key)
		if initialValue != "" {
			initialEnvironment[key] = initialValue
		} else {
			customEnvironment[key] = value
		}
		err := os.Setenv(key, value)
		if err != nil {
			fmt.Println("Error setting environment")
		}
	}

	body.Mutator(body)

	for key, _ := range customEnvironment {
		err := os.Unsetenv(key)
		if err != nil {
			fmt.Println("Error unsetting environment")
		}
	}

	for key, value := range initialEnvironment {
		err := os.Setenv(key, value)
		if err != nil {
			fmt.Println("Error resetting environment to default")
		}
	}
}