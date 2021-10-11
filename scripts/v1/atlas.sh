#!/bin/bash

date=$(date +"%F")
path=$GOPATH/src/catchall/logs/v1/$date
mkdir -p $path
mkdir -p $path/localdb
mkdir -p $path/atlas

atlasPing(){
  hey -m GET -c 5 -n 1000 "http://localhost:5000/ping" >> $path/atlas/ping.log
}
atlasDelivered(){
  hey -m PUT -c 5 -n 1000 "http://localhost:5000/events/foobar/delivered" >> $path/atlas/delivered.log
}
atlasBounced(){
  hey -m PUT -c 5 -n 1000 "http://localhost:5000/events/foobar/bounced" >> $path/atlas/bounced.log
}
atlasCheckDomain(){
  hey -m GET -c 5 -n 1000 "http://localhost:5000/domains/foobar" >> $path/atlas/domain.log
}

atlasPing && atlasDelivered && atlasBounced && atlasCheckDomain

