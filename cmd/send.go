/*
Copyright © 2024 Sebastian Meyer sebastian.meyer1@gmail.com
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to an entity",
	Long: `Use this to send a message to an Service Bus Entity.
	Only queues are supported right now.`,
	Run: func(cmd *cobra.Command, args []string) {
		queue, _ := cmd.Flags().GetString("queue")
		connection, _ := cmd.Flags().GetString("connection")
		data, _ := cmd.Flags().GetString("data")
		contentType, _ := cmd.Flags().GetString("content-type")

		sendDataToQueue(queue, connection, data, contentType)

		fmt.Println("Send Message to queue successfully.")
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("queue", "q", "", "Name of the Entity (Queue)")
	sendCmd.MarkFlagRequired("queue")
	sendCmd.Flags().StringP("connection", "c", "", "Service Bus Connection String")
	sendCmd.MarkFlagRequired("connection")
	sendCmd.Flags().StringP("data", "d", "", "Data to send")
	sendCmd.Flags().StringP("content-type", "t", "application/xml", "Content Type of Data. [application/json, text/xml, application/xml, text/plain]")
}

func sendDataToQueue(queue string, connection string, data string, contentType string) {
	context := context.Background()
	client, err := azservicebus.NewClientFromConnectionString(connection, nil)
	if err != nil {
		panic("Service Bus Connection failed.")
	}

	sender, err := client.NewSender(queue, nil)
	if err != nil {
		panic("Queue not found")
	}

	defer sender.Close(context)

	sbMessage := &azservicebus.Message{
		Body:        []byte(data),
		ContentType: &contentType,
	}
	err = sender.SendMessage(context, sbMessage, nil)
	if err != nil {
		panic(err)
	}
}
