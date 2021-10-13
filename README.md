# catchall

## Description

Service to determine if a email domain is a catch-all domain.

## Set Up
Move to the `app` directory and run `go build ./<version>`
Then move to `<version>` directory and run the binary.

To connect to a MongoDB database, either pass in a `-uri` flag or set a `.env` file with `MONGO_URI` as the property
To specify a API port, pass in a `-addr` flag
To specify a pprof port, pass in a `-pprof` flag

## Usage

* PUT requests to `http://localhost:5000/events/<domain_name>/delivered` will add to the number of delivered emails associated with that domain name.
* PUT requests to `http://localhost:5000/events/<domain_name>/bounced` will add to the number of bounced emails associated with that domain name.
* GET requests to `http://localhost:5000/domains/<domain_name>` will return the status of the domain whether it is **not a catch-all**, a **catch-all**, or an **unknown** domain.

When running, you can run `./scripts/<version>/atlas.sh` or `./scripts/<version>/localdb.sh` to generate logs for a specific version on a specific day connecting to either local MongoDB or MongoDB Atlas.

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
### V2

V2 was a short lived attempt to improve the speed of the PUT requests by creating two collections: `delivered_domains` and `bounced_domains`

PUT requests to `/events/<domain_name>/delivered` added a new item with a `domain_name` key that matched the domain name provided in the request.

I abandoned this after getting worse performance than I had with V1 for `/events/<domain_name>/delivered`.

My rationale for this approach was to increase the performance of the PUT requests with some sacrifice to the GET request as we would need to aggregate the records in each collection that matched the domain name.

## Next Steps

### Long-Term
* Add integration tests
* Containerize V3
  * Run multiple instances of catchall simultaneously to make sure that does not introduce any integrity issues: 
    * Essentially, to make sure that 4 instances each seeing 500 requests to `/events/foobar/delivered` would result in an item with the name of "foobar" in the domains collection on MongoDB having a delivered property of 2000.
* Upgrade my MongoDB cluster to have access to the profile MongoDB Atlas provides to look for areas for improvement
  * Refactor v1 domain models to use [MongoDB transactions](https://www.mongodb.com/developer/quickstart/golang-multi-document-acid-transactions/) to ensure atomic operations and make sure this is still faster than v2's approach. Current work on branch [use_mongo_transactions](https://github.com/cpustejovsky/catchall/tree/use_mongo_transactions) is causing the error `IllegalOperation: Transaction numbers are only allowed on a replica set member or mongos`
* Create a V3 server following Ardan Lab's [`service` starter kit](https://github.com/ardanlabs/service/) using the pattern developed in V1. Current work is on branch [v3_convert_to_ardanlabs_service_pattern](https://github.com/cpustejovsky/catchall/tree/v3_convert_to_ardanlabs_service_pattern)
