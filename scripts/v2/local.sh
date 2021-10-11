#!/bin/bash

date=$(date +"%F")
path=$GOPATH/src/catchall/logs/v2/$date
mkdir -p $path
mkdir -p $path/localdb
mkdir -p $path/atlas

localPing(){
  hey -m GET -c 5 -n 1000 "http://localhost:5000/ping" >> $path/localdb/ping.log
}
localDelivered(){
  hey -m PUT -c 5 -n 1000 "http://localhost:5000/events/foobar/delivered" >> $path/localdb/delivered.log
}
localBounced(){
  hey -m PUT -c 5 -n 1000 "http://localhost:5000/events/foobar/bounced" >> $path/localdb/bounced.log
}
localCheckDomain(){
  hey -m GET -c 5 -n 1000 "http://localhost:5000/domains/foobar" >> $path/localdb/domain.log
}

localPing && localDelivered # && localBounced && localCheckDomain
