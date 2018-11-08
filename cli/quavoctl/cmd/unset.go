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
	rootCmd.AddCommand(unsetCmd)
	unsetCmd.Flags().StringVarP(&Key, "key", "k", "", "--key=[key]")
	unsetCmd.MarkFlagRequired("key")
}

//unsetCmd removes a key from the cache
var unsetCmd = &cobra.Command{
	Use:   "unset --key=[key], --key [key]",
	Short: "removes a given key-value pair",
	Long:  `This command removes a key from cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()

		conn, err := grpc.Dial(port, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("could  not connect to backend: %s\n", err)
			os.Exit(1)
		}
		client := cache.NewCacheClient(conn)

		Unset(context.Background(), client, key)
	},
}

//Unset calls the cacheService to remove a given key
func Unset(cxt context.Context, client cache.CacheClient, key string) {
	request := &cache.UnsetRequest{
		Key: key,
	}
	response, err := client.Unset(context.Background(), request)

	if err == nil && response.Ok {
		fmt.Printf("key \"%s\"removed from the cache\n", request.GetKey())
	} else {
		fmt.Printf("could not remove key \"%s\", \"%v\"\n", key, err)
	}
}
