package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/rugwirobaker/quavo/models/cache"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&Key, "key", "k", "", "--key=[key]")
	setCmd.Flags().StringVarP(&Value, "value", "v", "", "--value [value]")
	setCmd.MarkFlagRequired("key")
	setCmd.MarkFlagRequired("value")

}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "sets an new key-value pair",
	Long:  `This command adds a new key-value pair to the cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()
		value := cmd.Flag("value").Value.String()

		conn, err := grpc.Dial(port, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("could  not connect to backend: %s\n", err)
			os.Exit(1)
		}
		client := cache.NewCacheClient(conn)

		Set(context.Background(), client, key, value)
	},
}

//Set calls the cache service to set a new key.
func Set(ctx context.Context, client cache.CacheClient, key, value string) {
	request := &cache.SetRequest{
		Key:   key,
		Value: []byte(value),
	}
	response, err := client.Set(context.Background(), request)

	if err == nil && response.Ok {
		fmt.Printf("key \"%s\" added to the cache\n", request.GetKey())
	} else {
		fmt.Printf("could not set key \"%s\", \"%v\"\n", key, err)
	}
}
