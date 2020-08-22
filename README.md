# nats-streaming-cli
CLI for Nats streaming service.

The CLI can be used to publish messages and subscribe to messages from a nats streaming service cluster.

## Get Started: Install the CLI

You can install the CLI via go modules or pull a docker image.  

GO Module
```sh
go install github.com/sourabp/nats-streaming-cli
```

Docker pull
```sh
docker pull sourabp/nats-streaming-cli
```

## Run the CLI
The main commands supported by the CLI are:

* `publish`: publishes message to a channel
* `subscribe`: subscribe to messages from a channel, terminates on `SIGKILL` or `SIGTERM`

### Conection parameters
Default url of nats-streaming-service: `localhost:4222` and cluster id: `test-cluster`

### Publish
```sh
$ nats-streaming-cli publish -q "foo_channel" "Hello World"

// Publishing to non-default cluster
$ nats-streaming-cli publish -s "<nats_server_ip>" -p "<nats_server_port>" -c "prod_cluster" -q "foo_channel" "Hello World"
```

### Subscribe
```sh
$ nats-streaming-cli subscribe -q "foo_channel"

// Add multiple subscribers to listen to a single queue group
$ nats-streaming-cli subscribe -q "foo_channel" --queue-group "baz_group" // consumer 1
$ nats-streaming-cli subscribe -q "foo_channel" --queue-group "baz_group" // consumer 2
```

