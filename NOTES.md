# Go REST API assignment

## Potential areas for imporvement:

- DB migrations are done in a very primitive way and we're relying on `gorm` to do
  auto migrations. It's better to use better solutions for this such as `go-migrate`
  or `sql-migrate`

- Functions `GetEmployee/Contractor` (along with their plural versions)
  could be candidates to be consolidated into a single call to `GetMember/Members`,
  passing in the `type`

- Request validation is very limited for the sake of simplicity!

- For successful operations such as delete, we're returning an empty 200 reponse;
  it would be nice to return a message of some kind.

- Endpoint and routing setup is a bit convoluded. Most probably, it can use some
  refactoring and extracting some similar functions into a single place.

- Errors are sent back as text and not in JSON format, which is something that can
  be improved.

- Did not create a PUT endpoint to do an update on entities :)

- The stub repository used for tests mostly covers the happy paths (again for
  simplicity), so the tests also mostly cover the happy path.

- There's only tests for a single endpoint (`role_test.go`); the other endpoints
  would also use the same pattern for tests.

## Running the application

- In root directory of the application:
  1. Run `make build`
  2. Run `make up`
  3. By default, `docker-compose` tries to expose the application on the port `8887`.
     If the port is already being used, please change it in `docker-compose.yml` and
     in the `ports` section for the `api` service.
  4. There are two options to test the application:
     - Use `curl` or a similar tool to call the API endpoints.
     - Go to `localhost:8887/swagger/` (NOTE that you need the trailing slash) and
       the swagger UI should be enought to test the application
  5. Run `make down` to bring down and clean up your Docker environment

- Note: there's a `/reset` endpoint for convenience to reset the DB!
