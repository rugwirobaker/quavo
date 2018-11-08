package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rugwirobaker/quavo/models/cache"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(flushCmd)
}

var flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "flush all the keys from the store",
	Long:  `This command removes all the previously set value from the store`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(port, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("could  not connect to backend: %s\n", err)
			os.Exit(1)
		}
		client := cache.NewCacheClient(conn)

		Flush(context.Background(), client)
	},
}

//Flush calls the Flush method of the CacheService
func Flush(ctx context.Context, client cache.CacheClient) {
	response, err := client.Flush(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Printf("could not flush store error:%v\n", err)
	}
	if response.GetOk() {
		fmt.Println("succesfully flushed store")
	}
}
