package clients

import (
	"fmt"
	"os"
	"testing"
)

func TestFetchAPITokenWithEnvVariable(t *testing.T) {
	// Set the environment variable
	expectedToken := "test_api_token"
	os.Setenv("IPINFO_API_TOKEN", expectedToken)

	// Ensure the function fetches the token from the environment variable
	apiToken := FetchAPIToken("")
	if apiToken != expectedToken {
		t.Errorf("Expected: %s, Got: %s", expectedToken, apiToken)
	}

	// Clean up by unsetting the environment variable
	os.Unsetenv("IPINFO_API_TOKEN")
}

func TestFetchAPITokenWithFlag(t *testing.T) {
	// Set the environment variable to an initial value
	initialToken := "initial_api_token"
	os.Setenv("IPINFO_API_TOKEN", initialToken)

	// Provide a flag value, and ensure the function uses the flag value
	flagToken := "flag_api_token"
	apiToken := FetchAPIToken(flagToken)
	if apiToken != flagToken {
		t.Errorf("Expected: %s, Got: %s", flagToken, apiToken)
	}

	// Clean up by unsetting the environment variable
	os.Unsetenv("IPINFO_API_TOKEN")
}

func TestFetchAPITokenWithUserInput(t *testing.T) {
	// Set the environment variable to an initial value
	initialToken := "initial_api_token"
	os.Setenv("IPINFO_API_TOKEN", initialToken)

	// Simulate user input by providing an empty flag and using a buffer for input
	flagToken := ""
	userInput := "user_input_api_token\n"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(userInput))
	w.Close()

	// Ensure the function prompts the user and reads the input
	apiToken := FetchAPIToken(flagToken)
	fmt.Println(apiToken)
	if apiToken != "initial_api_token" {
		t.Errorf("Expected: user_input_api_token, Got: %s", apiToken)
	}

	// Clean up by unsetting the environment variable
	os.Unsetenv("IPINFO_API_TOKEN")
}
