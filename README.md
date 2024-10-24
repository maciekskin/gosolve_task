## Number's index search API
Application exposes single endpoint returning index of an exact value or if exact value is not found then it returns other value matching conformation level (max 10% difference of provided value).

## Endpoint
Endpoint is exposed on `/numbers/{value}` path using GET method.
Value in endpoint's URI is expected to be a number.

### Returned values
- index - found index, -1 when not found
- value - found value
- error_message - error message if any error occured

## Errors
In case of any errors application returns `error_message` field describing the error.

## HTTP Codes
- 400 - invalid request, e.g. provided value is not a number
- 404 - index for provided value not found
- 200 - index found

## Build, run, tests
There is a Makefile providing basic commands to build, run and test code.

```make build```

Executes the unit tests and builds the app as `build/numbers` executable.

```make test```

Executes the unit tests.

```make run```

Executes `go run ./cmd/main.go` command.
