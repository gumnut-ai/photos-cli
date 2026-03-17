// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestAlbumsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"albums", "create",
			"--description", "description",
			"--library-id", "library_id",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"library_id: library_id\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"albums", "create",
		)
	})
}

func TestAlbumsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"albums", "retrieve",
			"--album-id", "album_id",
		)
	})
}

func TestAlbumsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"albums", "update",
			"--album-id", "album_id",
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
			t, pipeData,
			"--api-key", "string",
			"albums", "update",
			"--album-id", "album_id",
		)
	})
}

func TestAlbumsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"albums", "list",
			"--max-items", "10",
			"--asset-id", "asset_id",
			"--id", "[string, string]",
			"--library-id", "library_id",
			"--limit", "1",
			"--starting-after-id", "starting_after_id",
		)
	})
}

func TestAlbumsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"albums", "delete",
			"--album-id", "album_id",
		)
	})
}
