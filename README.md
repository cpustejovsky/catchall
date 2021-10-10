# catchall

## Description

Service to determine if a email domain is a catch-all domain.

## Set Up
Move to the `app` directory and run `go build ./<version>`

To connect to a MongoDB database, either pass in a `-uri` flag or set a `.env` file with `MONGO_URI` as the property

## Usage

* PUT requests to `http://localhost:5000/events/<domain_name>/delivered` will add to the number of delivered emails associated with that domain name.
* PUT requests to `http://localhost:5000/events/<domain_name>/bounced` will add to the number of bounced emails associated with that domain name.
* GET requests to `http://localhost:5000/domains/<domain_name>` will return the status of the domain whether it is **not a catch-all**, a **catch-all**, or an **unknown** domain.

## Versions

### V1

V1 was my first attempt at this. I used the same monolithic pattern I had used for [estuary](https://github.com/cpustejovsky/estuary) which is based on the pattern set forth by Alex Edwards in [Let's Go](https://lets-go.alexedwards.net/).

With first-hand experience of how painful queries of very large collections can be, I built the Domain model methods to update a single item within the domain database. The model looks like this:
```go
type Domain struct {
	Name      string ``
	Bounced   int    ``
	Delivered int    ``
}
```

PUT requests to `/events/<domain_name>/delivered` and `/events/<domain_name>/bounced` incremented the delivered and bounced properties of this item.

the GET request to `/domains/<domain_name>` only needed to find one item in the database and do simple conditional logic on the delivered and bounced properties.

I tested this for correctness on both my local mongodb and the cloud. The performance results I recorded for v1 are located in the `logs` directory

## Next Steps