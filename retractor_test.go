package envrtr

import "os"

func envApplyValues(envValues map[string]string) {
	for name, val := range envValues {
		if val != "" {
			os.Setenv(name, val)
		}
	}
}

func envRollbackValues(envValues map[string]string) {
	for name := range envValues {
		os.Unsetenv(name)
	}
}
