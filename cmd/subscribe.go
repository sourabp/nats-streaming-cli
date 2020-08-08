/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	"github.com/spf13/cobra"
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe a channel and receive messages published to it",
	Long:  `Subscribe to a channel and receive the messages posted to the channel as per the arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		subscribeMessages()
	},
}

var (
	queueGroup              string
	durableSubscriptionName string
	deliverSince            string
	// autoAcknowledge         bool
)

func init() {
	rootCmd.AddCommand(subscribeCmd)

	subscribeCmd.Flags().StringVar(&queueGroup, "queue-group", "", "Queue group to which the client needs to subscribe to")
	subscribeCmd.Flags().StringVar(&durableSubscriptionName, "durable-name", "", "Durable name for the subscription")
	subscribeCmd.Flags().StringVar(&deliverSince, "deliver-since", "", "Time difference with reference to current time for which messages are to be delivered")
}

func subscribeMessages() {
	sigs := make(chan os.Signal, 1)
	exitFlag := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	natsURL := fmt.Sprintf("nats://%s:%d", natServer, natsPort)
	fmt.Printf("Subscribing to channel %s at %s \n\n", channel, natsURL)
	nc, err := nats.Connect(natsURL)

	if err != nil {
		log.Fatalf("Error connecting to nats server %s", err)
	}

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))

	if err != nil {
		log.Fatalf("Error connecting to nats server %s", err)
	}

	// Sub routine to listen to termination request and close connections
	go func() {
		receivedSignal := <-sigs
		fmt.Printf("Received termination signal %s. Closing connections", receivedSignal)
		sc.Close()
		nc.Close()
		exitFlag <- true
	}()

	defer nc.Close()
	defer sc.Close()

	if len(queueGroup) > 0 {
		sc.QueueSubscribe(channel, queueGroup, handleReceivedMessage, subscriptionOptions())
	} else {
		sc.Subscribe(channel, handleReceivedMessage, subscriptionOptions())
	}

	<-exitFlag
}

func subscriptionOptions() stan.SubscriptionOption {
	return func(o *stan.SubscriptionOptions) error {
		if len(durableSubscriptionName) > 0 {
			o.DurableName = durableSubscriptionName
		}

		if len(deliverSince) > 0 {
			duration, err := time.ParseDuration(deliverSince)
			if err != nil {
				log.Fatal("Incorrect duration value in deliver since")
			}
			o.StartAt = pb.StartPosition_TimeDeltaStart
			o.StartTime = time.Now().Add(-duration)
		}
		return nil
	}

}

func handleReceivedMessage(msg *stan.Msg) {
	fmt.Printf("Received a Message: %s \n", string(msg.Data))
}
