# Requirements

- Create an http server
- Create REST API
- Storing data with MongoDB
- Writing test cases
- Build and deploy to server
- Create API Doc with Swagger UI

## Tools

- Source code management : github.com
- Code editor: visual code
- Local Enviroment: ubuntu

## Milestones

### Milestone 1

- Reading or writing headers
- Examining the type of a request
- Fetching data from a database
- Parsing request data
- Authentication

## Testing

### Unit Tests

> "If I give the function X these values, I should expect this value to be returned"

```go
func sum(x int, b int) int {
  return x + y
}
```

### Integration Tests

> "Integration tests typically test how the various parts of an application work
together"

### Functional Tests

> Functional tests are often known as end-to-end tests and outside-in tests
>
> For users of software, functional tests are perhaps
the most important tests that can be run. Examples of functional tests include:

- Testing that a command line tool responds to certain inputs with certain
outputs.
- Running automated tests on a web page.
- Running outside-in tests against an API to check response codes and
headers.

### Testing & Benchmark

```bash
go test -run=XXX # run the specific test
go test -bench=.
```

> Each benchmark is run for a minimum of 1 second by default. If the second has not elapsed when the Benchmark function returns, the value of b.N is increased in the sequence 1, 2, 5, 10, 20, 50, … and the function run again.


**Example

```bash
BenchmarkStringFromAssignment-8           200000              7393 ns/op
BenchmarkStringFromAppendJoin-8           300000              5579 ns/op
BenchmarkStringFromStringBuffer-8        1000000              1637 ns/op

==> BenchmarkStringFromStringBuffer for string concenate is the fastest!
```

## Common Errors

1. Set Header does not working

- https://github.com/golang/go/issues/17083

> // header should be set first at all
> 
> w.Header().Set("Content-Type", "application/json; charset=utf-8")

## References

- https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
