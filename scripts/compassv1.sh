#!/bin/bash

date=$(date +"%F")
path=$GOPATH/src/catchall/logs/$date
mkdir $path
mkdir $path/localdb
mkdir $path/compass
cd $GOPATH/src/catchall/app

compassPing(){
  hey -m GET -c 5 -n 1000 "http://localhost:5000/ping" >> $path/compass/ping.log
}
compassDelivered(){
  hey -m PUT -c 5 -n 1000 "http://localhost:5000/events/foobar/delivered" >> $path/compass/delivered.log
}
compassBounced(){
  hey -m PUT -c 5 -n 1000 "http://localhost:5000/events/foobar/bounced" >> $path/compass/bounced.log
}
compassCheckDomain(){
  hey -m GET -c 5 -n 1000 "http://localhost:5000/domains/foobar" >> $path/compass/domain.log
}

compassPing && compassDelivered && compassBounced && compassCheckDomain

