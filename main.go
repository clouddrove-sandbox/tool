package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <target>", os.Args[0])
	}

	target := os.Args[1]
	if err := runMake(target); err != nil {
		log.Fatalf("Error running make target '%s': %v", target, err)
	}
	fmt.Printf("Make target '%s' completed successfully.\n", target)
}

func runMake(target string) error {
	cmd := exec.Command("make", target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("make command failed: %v", err)
	}
	return nil
}
