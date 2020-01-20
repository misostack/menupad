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

## Common Errors

1. Set Header does not working

- https://github.com/golang/go/issues/17083

> // header should be set first at all
> 
> w.Header().Set("Content-Type", "application/json; charset=utf-8")