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
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&Key, "key", "k", "", "--key [key]")
	setCmd.MarkFlagRequired("key")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "returns value for a given key.",
	Long:  `This command queries the quavo key-store for with a previously set key(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value

		conn, err := grpc.Dial(port, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("could  not connect to backend: %s\n", err)
			os.Exit(1)
		}
		client := cache.NewCacheClient(conn)

		Get(context.Background(), client, key.String())
	},
}

//Get calls the cache service to get a given key's value
func Get(ctx context.Context, client cache.CacheClient, key string) {
	response, err := client.Get(context.Background(), &cache.GetRequest{Key: key})

	if err == nil && response != nil {
		fmt.Printf("Value:\"%s\"\n", response.GetValue())
	} else {
		fmt.Printf("could not get key \"%s\", \"%v\"\n", key, err)
	}
}
