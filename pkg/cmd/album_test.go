// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/photos-cli/internal/mocktest"
)

func TestAlbumsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums", "create",
			"--api-key", "string",
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
			t, pipeData, "albums", "create",
			"--api-key", "string",
		)
	})
}

func TestAlbumsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums", "retrieve",
			"--api-key", "string",
			"--album-id", "album_id",
		)
	})
}

func TestAlbumsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums", "update",
			"--api-key", "string",
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
			t, pipeData, "albums", "update",
			"--api-key", "string",
			"--album-id", "album_id",
		)
	})
}

func TestAlbumsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums", "list",
			"--api-key", "string",
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
			t, "albums", "delete",
			"--api-key", "string",
			"--album-id", "album_id",
		)
	})
}
