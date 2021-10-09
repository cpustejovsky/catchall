#!/bin/bash

date=$(date +"%F")
path=$GOPATH/src/catchall/logs/$date
mkdir $path
mkdir $path/localdb
mkdir $path/compass
cd $GOPATH/src/catchall/app

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

localPing && localDelivered && localBounced && localCheckDomain

