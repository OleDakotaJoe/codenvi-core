package terminal

import (
	"../utils"
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"os/exec"
	"strings"
)

func EnviTerminal() {
	utils.InitEnvi()
	enviRoot, _ := os.Getwd()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("envi >%s$ ", enviRoot)
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")

func execInput(input string, writer io.Writer) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home with empty path not yet supported.
		if len(args) < 2 {
			return ErrNoPath
		}
		// Change the directory and return the error.
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = writer

	// Execute the command and return the error.
	return cmd.Run()
}

func EnviTermialV2() error {
	if !terminal.IsTerminal(0) || !terminal.IsTerminal(1) {
		return fmt.Errorf("stdin/stdout should be terminal")
	}
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return err
	}
	defer terminal.Restore(0, oldState)
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	term := terminal.NewTerminal(screen, "")
	enviColor := string(term.Escape.Green)
	reset := string(term.Escape.Reset)
	wdColor := string(term.Escape.Blue)
	wd, _ := os.Getwd()
	homeDir, _ := utils.GetHomeDir()
	wd = strings.Replace(wd, homeDir, "~", 1)
	colonColor := string(term.Escape.Reset)

	prompt := fmt.Sprintf("(%s)%s envi%s:%s%s%s$", "test", enviColor, colonColor, wdColor, wd, reset)
	term.SetPrompt(prompt)

	for {
		line, err := term.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if line == "" {
			continue
		}
		if line == "exit" {
			return nil
		}
		execInput(line, term)
		fmt.Fprint(term)
	}
}
