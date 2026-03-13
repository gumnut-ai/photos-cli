// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestLibrariesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "libraries", "create",
			"--api-key", "string",
			"--name", "name",
			"--description", "description",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "libraries", "create",
			"--api-key", "string",
		)
	})
}

func TestLibrariesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "libraries", "retrieve",
			"--api-key", "string",
			"--library-id", "library_id",
		)
	})
}

func TestLibrariesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "libraries", "update",
			"--api-key", "string",
			"--library-id", "library_id",
			"--description", "description",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "libraries", "update",
			"--api-key", "string",
			"--library-id", "library_id",
		)
	})
}

func TestLibrariesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "libraries", "list",
			"--api-key", "string",
		)
	})
}

func TestLibrariesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "libraries", "delete",
			"--api-key", "string",
			"--library-id", "library_id",
		)
	})
}
