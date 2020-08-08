# nats-streaming-cli
CLI for Nats streaming service

## Flags

### Publish
* server, s
* port, p
* clusterId
* channel
* message
* file containing message

### Subscribe
* server, s 
* port, p
* channel
* cluster-id
* queue-name (optional)
* cliend-id (optional)
* message-starting-index
* handle connection closure on unix signals

## TODO

### Metrics
* Number of messages published per channel
* Number of messages outbound per channel
* Number of active clients/ subscriptions
* Number of channels
* List of channels

