package tools

import (
	"fmt"
	"os"
	"os/signal" // Import the signal package
	"syscall"   // Import syscall for specific signal types
)

// ListenForGracefulShutdown sets up a signal listener to catch
// Ctrl+C (SIGINT) and termination (SIGTERM) signals.
// When a signal is received, it prints "Program Closed!" and exits.
func ListenForGracefulShutdown() {
	signs := make(chan os.Signal, 1)

	// Register to receive notification of interrupt (Ctrl+C) and terminate signals
	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to handle the signals
	go func() {
		<-signs                          // This blocks until a signal is received
		fmt.Println("\nProgram Closed!") // Print the message when a signal is caught
		os.Exit(0)                       // Exit the program cleanly
	}()
}
