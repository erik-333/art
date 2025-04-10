package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type PageData struct {
	Result string
	Status string
}

func PressEnterTwice() {
	fmt.Fprintln(os.Stdout, "")
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(os.Stdout, "")
}

func executeDecoderCLI(mode string, input string) (string, error) {
	var modeFlag string
	switch mode {
	case "encode":
		modeFlag = "-en"
	case "decode":
		modeFlag = "-de"
	default:
		return "", fmt.Errorf("invalid mode: %s", mode)
	}

	// Use proper flags: -en/-de for mode and -m for multiline
	cmd := exec.Command("../decoder/diamond-decoder", "-m", modeFlag)

	// Create pipes for stdin, stdout, and stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stdin pipe: %v", err)
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("error starting command: %v", err)
	}

	// Write the input to stdin
	_, err = stdin.Write([]byte(input + "\n\n"))
	if err != nil {
		return "", fmt.Errorf("error writing to stdin: %v", err)
	}

	// Close stdin to signal we're done writing
	stdin.Close()

	// Wait for the command to complete
	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("error waiting for command: %v: %s", err, stderr.String())
	}

	// get raw output
	rawOutput := out.String()

	// remove prompt message
	cleanOutput := strings.TrimPrefix(rawOutput, "Enter your text, press Enter twice to finish")

	return cleanOutput, nil
}

// The decoder handler
func decoder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get form values
	input := r.FormValue("text")
	mode := r.FormValue("mode") // "decode" or "encode"

	if input == "" {
		http.Error(w, "Bad Request: No input provided", http.StatusBadRequest)
		return
	}

	// Validate mode
	if mode != "decode" && mode != "encode" {
		mode = "decode" // default to decode if invalid
	}

	// Execute CLI command
	result, err := executeDecoderCLI(mode, input)
	if err != nil {
		http.Error(w, "Processing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create page data with result
	data := PageData{
		Result: result,
		Status: "202 Accepted",
	}

	w.WriteHeader(http.StatusAccepted)
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, data)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup routes
	http.HandleFunc("/", home)
	http.HandleFunc("/decoder", decoder)

	// Start server
	fmt.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
