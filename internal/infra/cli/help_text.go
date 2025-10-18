package cli

import "fmt"

func PrintHelp() {
	fmt.Println("Usage: stress-test --url <url> --requests <number> --concurrency <number>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  help          Display help information")
	fmt.Println()
	fmt.Println("Flags for 'stress-test' command:")
	fmt.Println("  Flags          Shorthand  Description")
	fmt.Println(" =====================================================================================")
	fmt.Println("  --url          -u         URL to perform the stress test (required)")
	fmt.Println("  --requests     -r         Total number of requests (required, positive integer)")
	fmt.Println("  --concurrency  -c         Number of concurrent requests (required, positive integer)")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  go run cmd/cli/main.go --url https://www.google.com --requests 1000 --concurrency 2")
	fmt.Println()
}
