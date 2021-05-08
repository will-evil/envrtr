# Environment Retractor

`envrtr` is a pure package for applying and rolling back the state of environment variables.

# For what?

This package is really useful when you need to test a piece of code whose behavior depends on the environment variables.
The package will help bring the variables to the desired state, and at the end of the test, return them to their original state.

## Installation

```
go get github.com/will-evil/envrtr
```

## Examples

```
import (
    "testing"

    "github.com/will-evil/envrtr"
)

### Using only envrtr

func TestSome(t *testing) {
    // map of environment variable names and new values for this variables
    envValues := map[string]string{
		"TEST_ENV":  "value",
		"TEST_ENV2": "value2",
		"TEST_ENV3": "value3",
	}
    // create new retractor instance and set new values for provided environment variables
    r := envrtr.NewEnvRetractor(envValues).PullOut()
    // rolling back environment variables to their original state
    defer r.Retract()

    res := Some()
    // check result of function
}
```

### Using with testify's suite

```
import (
	"testing"

  	"github.com/will-evil/envrtr"
	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
    retractor *envrtr.EnvRetractor
}

func (suite *ExampleSuite) SetupSuite() {
  	envValues := map[string]string{
		"TEST_ENV":  "value",
		"TEST_ENV2": "value2",
		"TEST_ENV3": "value3",
	}
    suite.retractor = envrtr.NewEnvRetractor(envValues)
}

func (suite *ExampleSuite) BeforeTest(_, _ string) {
  	suite.retractor.PullOut()
}

func (suite *ExampleSuite) AfterTest(_, _ string) {
  	suite.retractor.Retract()
}

func (suite *ExampleSuite) TestSomething() {
  	suite.Equal(true, true)
}

func TestRunSuite(t *testing.T) {
  	suite.Run(t, new(ExampleSuite))
}
```
