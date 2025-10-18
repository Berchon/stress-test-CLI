package cli

import "log"

func Run() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute CLI command: %v", err)
	}
}
