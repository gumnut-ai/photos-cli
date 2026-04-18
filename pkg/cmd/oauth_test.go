// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestOAuthAuthURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth", "auth-url",
			"--redirect-uri", "redirect_uri",
			"--code-challenge", "code_challenge",
			"--code-challenge-method", "code_challenge_method",
		)
	})
}

func TestOAuthExchange(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth", "exchange",
			"--code", "code",
			"--code-verifier", "code_verifier",
			"--error", "error",
			"--state", "state",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"code: code\n" +
			"code_verifier: code_verifier\n" +
			"error: error\n" +
			"state: state\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"oauth", "exchange",
		)
	})
}

func TestOAuthLogoutEndpoint(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth", "logout-endpoint",
		)
	})
}
