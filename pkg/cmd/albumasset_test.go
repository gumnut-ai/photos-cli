// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestAlbumAssetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"album-assets", "list",
			"--max-items", "10",
			"--album-id", "album_id",
			"--asset-id", "asset_id",
			"--id", "[string, string]",
			"--library-id", "library_id",
			"--limit", "1",
			"--starting-after-id", "starting_after_id",
		)
	})
}

func TestAlbumAssetsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"album-assets", "get",
			"--album-asset-id", "album_asset_id",
		)
	})
}
