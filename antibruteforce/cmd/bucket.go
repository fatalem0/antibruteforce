package cmd

import (
	"antibruteforce/internal/config"
	"antibruteforce/internal/grpcserver"
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var bucketCmd = &cobra.Command{
	Use: "bucket",
	//Args: cobra.MinimumNArgs(1),
	Short: "Command resets bucket by login or ip",
	Run: func(cmd *cobra.Command, args []string) {
		container := BuildContainer()
		err := container.Invoke(func(conf *config.GrpcConf) {
			conn, err := newGrpcConnection(conf)
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				if err := conn.Close(); err != nil {
					log.Fatal(err)
				}
			}()
			server := grpcserver.NewAntiBruteForceClient(conn)
			ctx := context.Background()
			var typOfBucket grpcserver.BucketKind
			var key string
			if ip != "" {
				typOfBucket = grpcserver.BucketKind_IP
				key = ip
			}
			if login != "" {
				typOfBucket = grpcserver.BucketKind_LOGIN
				key = login
			}
			if ip == "" && login == "" {
				fmt.Println("ip or login required")
				return
			}
			req := &grpcserver.ResetBucketRequest{
				Kind: typOfBucket,
				Key:  key,
			}
			status, err := server.ResetBucket(ctx, req)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s was reset status: %t \n", key, status.Ok)
		})
		if err != nil {
			log.Fatal(err)
		}
		//container := BuildContainer()
	},
}
