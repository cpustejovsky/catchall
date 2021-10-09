# catchall

## Description

Service to determine if a email domain is a catch-all domain.

* PUT requests to `/events/<domain_name>/delivered` will add to the number of delivered emails associated with that domain name.
* PUT requests to `/events/<domain_name>/bounced` will add to the number of bounced emails associated with that domain name.
* GET requests to `/domains/<domain_name>` will return the status of the domain whether it is **not a catch-all**, a **catch-all**, or an **unknown** domain.

## Set Up

## Usage

## Next Steps