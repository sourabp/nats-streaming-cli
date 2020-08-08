/*
Copyright © 2020 Sourab Pareek <sourab.pareek21@gmail.com>

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
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish [OPTIONS] <MESSAGE>",
	Short: "Publish a message to a nats streaming server channel",
	Long: `Publish a message to a nats streaming server channel. For example:

	$ nats-streaming-cli publish -s nats.my-cluster.com -p 4224 -cl prod-cluster -c test-channel Test Message
	`,
	Run: func(cmd *cobra.Command, args []string) {
		publishNatsMessage(args)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}

func publishNatsMessage(args []string) {
	message := strings.Join(args, " ")

	natsURL := fmt.Sprintf("nats://%s:%d", natServer, natsPort)
	nc, err := nats.Connect(natsURL)

	defer nc.Close()

	if err != nil {
		log.Fatalf("Error connecting to nats server %s", err)
	}

	sc, _ := stan.Connect(clusterID, clientID, stan.NatsConn(nc))

	publishERR := sc.Publish(channel, []byte(message))

	if publishERR != nil {
		log.Fatal(publishERR)
	}
	fmt.Printf("Published message with client ID %s", clientID)
}
