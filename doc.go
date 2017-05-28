// THIS FILE IS GENERATED BY GO GENERATE.
// DO NOT EDIT THIS FILE BY HAND.

/*
kafka-firehose-nozzle is Cloud Foundry nozzle which forwards logs from
the loggeregagor firehose to Apache kafka (http://kafka.apache.org/).

Usage:

    kafak-firehose-nozzle [options]

Available options:

    -config PATH          Path to configuraiton file
    -username NAME        username to grant access token to connect firehose
    -password PASS        password to grant access token to connect firehose
    -worker NUM           Number of producer worker. Default is number of CPU core
    -subscription ID      Subscription ID for firehose. Default is 'kafka-firehose-nozzle'
    -stats-interval TIME  How often display stats info to console
    -debug                Output event to stdout instead of producing message to kafka
    -log-level LEVEL      Log level. Default level is INFO (DEBUG|INFO|ERROR)

*/
package main
