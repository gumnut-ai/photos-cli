// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestAlbumsAssetsAssociationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums:assets-associations", "list",
			"--api-key", "string",
			"--album-id", "album_id",
		)
	})
}

func TestAlbumsAssetsAssociationsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums:assets-associations", "add",
			"--api-key", "string",
			"--album-id", "album_id",
			"--asset-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"asset_ids:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "albums:assets-associations", "add",
			"--api-key", "string",
			"--album-id", "album_id",
		)
	})
}

func TestAlbumsAssetsAssociationsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums:assets-associations", "remove",
			"--api-key", "string",
			"--album-id", "album_id",
			"--asset-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"asset_ids:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "albums:assets-associations", "remove",
			"--api-key", "string",
			"--album-id", "album_id",
		)
	})
}
