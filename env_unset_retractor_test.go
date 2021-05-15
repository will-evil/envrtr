package envrtr

import (
	"fmt"
	"os"
	"testing"
)

func Test_NewEnvUnsetRetractor(t *testing.T) {
	envValues := map[string]string{
		"TEST_ENV":  "value",
		"TEST_ENV2": "value2",
	}
	envApplyValues(envValues)
	defer envRollbackValues(envValues)

	r := NewEnvUnsetRetractor([]string{"TEST_ENV", "TEST_ENV2", "TEST_ENV3"})

	t.Run("when check len of values map", func(t *testing.T) {
		realLen := len(r.values)
		expectedLen := len(envValues)
		if realLen != expectedLen {
			t.Errorf("len of values map not correct. Expected %d, got %d", expectedLen, realLen)
		}
	})

	for name, value := range envValues {
		t.Run(fmt.Sprintf("check %s variable", name), func(t *testing.T) {
			realVal, ok := r.values[name]
			if !ok {
				t.Errorf("variable '%s' not exists in values map", name)
				return
			}
			if value != realVal {
				t.Errorf("value for variable '%s' not correct. Expected '%s', got '%s'", name, value, realVal)
			}
		})
	}
}

func Test_EnvUnsetRetractor_PullOut(t *testing.T) {
	envValues := map[string]string{
		"TEST_ENV":  "value",
		"TEST_ENV2": "value2",
		"TEST_ENV3": "value3",
	}
	envApplyValues(envValues)
	defer envRollbackValues(envValues)

	r := NewEnvUnsetRetractor([]string{"TEST_ENV", "TEST_ENV2", "TEST_ENV3"})
	r.PullOut()

	for name := range envValues {
		t.Run(fmt.Sprintf("when test %s variable", name), func(t *testing.T) {
			if _, present := os.LookupEnv(name); present {
				t.Errorf("variable '%s' still present, must be unset", name)
			}
		})
	}
}

func Test_EnvUnsetRetractor_Retract(t *testing.T) {
	envValues := map[string]string{
		"TEST_ENV":  "value",
		"TEST_ENV2": "value2",
		"TEST_ENV3": "value3",
	}
	envApplyValues(envValues)
	defer envRollbackValues(envValues)

	r := NewEnvUnsetRetractor([]string{"TEST_ENV", "TEST_ENV2", "TEST_ENV3"}).PullOut()
	r.Retract()

	for name, value := range envValues {
		t.Run(fmt.Sprintf("when test %s variable", name), func(t *testing.T) {
			res, present := os.LookupEnv(name)
			if !present {
				t.Errorf("variable '%s' must present", name)
				return
			}
			if res != value {
				t.Errorf("not correct value for variable '%s', Expected '%s', got '%s'", name, value, res)
			}
		})
	}
}
