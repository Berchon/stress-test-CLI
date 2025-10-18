package cli

import (
	"fmt"
	"log"
	"os"
	"time"

	httpService "github.com/Berchon/stress-test-cli/internal/business/service"
	"github.com/Berchon/stress-test-cli/internal/business/usecase"
	"github.com/Berchon/stress-test-cli/internal/infra/statistics"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var url string
	var requests int
	var concurrency int

	rootCmd := &cobra.Command{
		Use:   "stress-test <url> <requests> <concurrency>",
		Short: "CLI tool to perform stress testing on a URL",
		Long:  `A CLI tool that executes a stress test on a given URL with specified total requests and concurrency.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(os.Args) == 1 {
				PrintHelp()
				return
			}

			if err := ValidateParams(url, requests, concurrency); err != nil {
				log.Fatalf("Invalid parameters: %v", err)
			}

			fmt.Printf("Executing stress test for URL: %s\n", url)
			fmt.Println()
			fmt.Println()

			httpSvc := httpService.NewHTTPService(10 * time.Second)
			statsSvc := statistics.NewStatisticsService()

			uc := usecase.NewLoadTestUseCase(httpSvc, statsSvc)

			err := uc.StartLoadTest(url, requests, concurrency)
			if err != nil {
				log.Fatalf("Error running load test: %v", err)
			}
		},
	}

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL to perform the stress test (required)")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 0, "Total number of requests (required)")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 0, "Number of concurrent requests (required)")

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		PrintHelp()
	})

	rootCmd.AddCommand(NewHelpCmd())

	return rootCmd
}
