package main

import (
	"bytes"
	"os/exec"
	"testing"
)

// The function `TestMainOutput` tests the output of a Go program to ensure it matches the expected
// "Hello, World!" message.
func TestMainOutput(t *testing.T) {
	var out bytes.Buffer
	cmd := exec.Command("go", "run", "../cmd/api/main.go")
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {

		t.Fatalf("Erro ao executar o comando: %v", err)
	}

	expected := "Hello, World!\n"
	if out.String() != expected {
		t.Errorf("Saída inesperada: got %q, want %q", out.String(), expected)
	}
}
