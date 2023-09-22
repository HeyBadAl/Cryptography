package main 

import (
  "os/exec"
  "strings"
  "testing"
)

func TestHello(t *testing.T) {
  command := exec.Command("go", "run", "main.go")

  output, err := command.CombinedOutput()
  if err != nil {
    t.Fatalf("Main Program failed: %v", err)
  }
  
  outputStr := string(output)

  if !strings.Contains(outputStr, "Hello"){
    t.Errorf("Expected 'Hello' in output, got: %s", outputStr)
  }
}
