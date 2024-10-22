# Find Index API REST Service

Find Index service finds the index of a given value, and if that's not possible, index of a value that is within range of +/- 10% of the given value.
If there is no number with 10% tolerance, the server returns error message.

Installation:
git clone https://github.com/marcin-mc/find-index
go mod tidy


Install golangci-lint for linting:
`https://golangci-lint.run/welcome/install/`

Setup:
Set port and logging level in config.yaml file:
port:
  3000
log_level:
  DEBUG


Usage:
GET /endpoint/{target}

Responses:
Success:
{
    "Target": 0
    "Index": 0,
    "Value": 0,
    "Message": "Value found"
}

Error:
{
    "Message": "Value could not be found"
}

Linting:
make lint

Testing:
make test

Usage:
GET http://<host>:<port>/endpoint/<target:int>
GET http://localhost:3000/endpoint/100



-----------------------

We would like to be able to call designed endpoint with http `GET` method and send `value` that should be found in input file.
As a response we should get `index` number for given value and the corresponding value and optional message.

For example, we are sending GET for /endpoint/100 and as result we should receive 3. 

``Remark: `As a requirement`, we want to load that file into `slice` once service starts.
So all search operations should be optimized for that particular slice.``

- In case you’re not able to find index for given value, you can return `index` for any other existing value, assuming that conformation is at `10% level`, (for example, you were looking for `index` for value = `1150`, but in input file you have `1100` and `1200`, so in that case you can return index for `1100` or `1200`).
 
- In case you were not able to find valid `index` number, `error message` should be added into response.

`To summarize`:
- Design API for http `GET` method
- Implement functionality for searching `index` for `given` value (it should be the most efficient algorithm) 
- Add logging
- Add possibility to use configuration file where you can specify service port and log level (you should be able to choose between Info, Debug, Error)
- Add `unit tests` for created components
- Add `README.md` to describe your service
- Automate running tests with `make` file
- Remember that code structure matters
- Upload solution into `GitHub` account and share the link

Check coverage with html display:
go test ./... -coverprofile=c.out
go tool cover -html="c.out"