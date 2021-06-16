package envrtr

import (
	"fmt"
	"os"
	"testing"
)

func TestNewEnvRetractor(t *testing.T) {
	systemEnvValues := systemEnvValues()
	envApplyValues(systemEnvValues)
	defer envRollbackValues(systemEnvValues)

	envValues := envValues()
	r := NewEnvRetractor(envValues).(*EnvRetractor)
	for name, tmpVal := range envValues {
		t.Run(fmt.Sprintf("state_for_var_%s_with_val_%s", name, tmpVal), func(t *testing.T) {
			valSet, ok := r.values[name]
			if !ok {
				t.Errorf("variable '%s' not exists in result", name)
				return
			}
			if valSet.tmp != tmpVal {
				t.Errorf(
					"not correct tmp value for variable '%s'. Expected '%s', got '%s'",
					name,
					tmpVal,
					valSet.original,
				)
			}
			originalVal, ok := systemEnvValues[name]
			if !ok {
				t.Fatalf("variable '%s' not exists in systemEnvValues", name)
			}
			if valSet.original != originalVal {
				t.Errorf(
					"not correct original value for variable '%s'. Expected '%s', got '%s'",
					name,
					originalVal,
					valSet.original,
				)
			}
			if originalVal == "" && valSet.present != false {
				t.Errorf("not correct preset value for variable '%s'. Expected false, got true", name)
			}
		})
	}
}

func TestEnvRetractor_PullOut(t *testing.T) {
	systemEnvValues := systemEnvValues()
	envApplyValues(systemEnvValues)
	defer envRollbackValues(systemEnvValues)

	envValues := envValues()
	NewEnvRetractor(envValues).PullOut()
	for name, expectedVal := range envValues {
		t.Run(fmt.Sprintf("var_%s_expected_val_%s", name, expectedVal), func(t *testing.T) {
			val := os.Getenv(name)
			if val != expectedVal {
				t.Errorf("not correct value for variable '%s'. Expected '%s', got '%s'", name, expectedVal, val)
			}
		})
	}
}

func TestEnvRetractor_Retract(t *testing.T) {
	systemEnvValues := systemEnvValues()
	envApplyValues(systemEnvValues)
	defer envRollbackValues(systemEnvValues)

	envValues := envValues()
	r := NewEnvRetractor(envValues).PullOut()
	r.Retract()

	for name, expectedVal := range systemEnvValues {
		t.Run(fmt.Sprintf("var_%s_expected_val_%s", name, expectedVal), func(t *testing.T) {
			val, present := os.LookupEnv(name)
			if expectedVal == "" && present != false {
				t.Errorf("variable '%s' must be unset but present", name)
			}
			if val != expectedVal {
				t.Errorf("not correct value for variable '%s'. Expected '%s', got '%s'", name, expectedVal, val)
			}
		})
	}
}

func envValues() map[string]string {
	return map[string]string{
		"TEST_ENV":  "value",
		"TEST_ENV2": "value2_new",
		"TEST_ENV4": "value4_new",
	}
}

func systemEnvValues() map[string]string {
	return map[string]string{
		"TEST_ENV":  "value",
		"TEST_ENV2": "value2",
		"TEST_ENV3": "value3",
		"TEST_ENV4": "",
		"TEST_ENV5": "",
	}
}
