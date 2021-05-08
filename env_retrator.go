// Package envrtr provides functional to pull out (apply) and retract (rollback) changes for environment variables.
package envrtr

import "os"

type envSet struct {
	original string
	tmp      string
	present  bool
}

// EnvRetractor provides functional for pull out and retracts environment variables.
type EnvRetractor struct {
	values map[string]envSet
}

// NewEnvRetractor constructor for EnvRetractor.
func NewEnvRetractor(envValues map[string]string) *EnvRetractor {
	r := &EnvRetractor{}
	r.values = make(map[string]envSet)

	for name, value := range envValues {
		originalVal, present := os.LookupEnv(name)
		r.values[name] = envSet{original: originalVal, tmp: value, present: present}
	}

	return r
}

// PullOut sets provided values for environment variables.
func (r *EnvRetractor) PullOut() *EnvRetractor {
	for name, set := range r.values {
		os.Setenv(name, set.tmp)
	}

	return r
}

// Retract rolls back the values of an environment variables to its original state.
func (r *EnvRetractor) Retract() {
	for name, set := range r.values {
		if set.present {
			os.Setenv(name, set.original)
		} else {
			os.Unsetenv(name)
		}
	}
}
