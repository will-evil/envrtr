// Copyright © 2021 Alexey Konovalenko
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
func NewEnvRetractor(envValues map[string]string) Retractor {
	r := &EnvRetractor{}
	r.values = make(map[string]envSet)

	for name, value := range envValues {
		originalVal, present := os.LookupEnv(name)
		r.values[name] = envSet{original: originalVal, tmp: value, present: present}
	}

	return r
}

// PullOut sets provided values for environment variables.
func (r *EnvRetractor) PullOut() Retractor {
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
