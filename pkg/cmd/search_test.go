// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestSearchSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "search", "search",
			"--api-key", "string",
			"--captured-after", "'2019-12-27T18:11:19.117Z'",
			"--captured-before", "'2019-12-27T18:11:19.117Z'",
			"--library-id", "library_id",
			"--limit", "1",
			"--page", "1",
			"--person-id", "[string, string]",
			"--query", "query",
			"--threshold", "0",
		)
	})
}

func TestSearchSearchAssets(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "search", "search-assets",
			"--api-key", "string",
			"--captured-after", "'2019-12-27T18:11:19.117Z'",
			"--captured-before", "'2019-12-27T18:11:19.117Z'",
			"--image", "Example data",
			"--library-id", "library_id",
			"--limit", "1",
			"--page", "1",
			"--person-id", "[string]",
			"--query", "query",
			"--threshold", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"captured_after: '2019-12-27T18:11:19.117Z'\n" +
			"captured_before: '2019-12-27T18:11:19.117Z'\n" +
			"image: Example data\n" +
			"library_id: library_id\n" +
			"limit: 1\n" +
			"page: 1\n" +
			"person_ids:\n" +
			"  - string\n" +
			"query: query\n" +
			"threshold: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "search", "search-assets",
			"--api-key", "string",
		)
	})
}
