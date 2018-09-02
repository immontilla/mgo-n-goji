#!/usr/bin/env bash
export MGONGOJI_MONGODB=127.0.0.1:27017
export MGONGOJI_DBNAME=test_phonebook
export MGONGOJI_IPPORT=:9889
$GOPATH/bin/app &
newman run mgo-n-goji-tests.postman_collection.json --reporters cli
TESTPID=`ps -elf | grep /bin/app | grep -v grep | awk '{print $4}'`
kill $TESTPID
unset MGONGOJI_MONGODB
unset MGONGOJI_DBNAME
unset MGONGOJI_IPPORT
