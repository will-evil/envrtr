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

package envrtr

import "os"

// EnvUnsetRetractor provides functional for unseting environment variables and return their original values.
type EnvUnsetRetractor struct {
	values map[string]string
}

// NewEnvUnsetRetractor constructor for EnvUnsetRetractor.
func NewEnvUnsetRetractor(envVars []string) Retractor {
	r := &EnvUnsetRetractor{}
	r.values = make(map[string]string)

	for _, name := range envVars {
		value, present := os.LookupEnv(name)
		if !present {
			continue
		}
		r.values[name] = value
	}

	return r
}

// PullOut sets provided values for environment variables.
func (r *EnvUnsetRetractor) PullOut() Retractor {
	for name := range r.values {
		os.Unsetenv(name)
	}

	return r
}

// Retract rolls back the values of an environment variables to its original state.
func (r *EnvUnsetRetractor) Retract() {
	for name, value := range r.values {
		os.Setenv(name, value)
	}
}
