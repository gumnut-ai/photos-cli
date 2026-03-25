// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestAssetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "create",
			"--asset-data", "asset_data",
			"--device-asset-id", "device_asset_id",
			"--device-id", "device_id",
			"--file-created-at", "'2019-12-27T18:11:19.117Z'",
			"--file-modified-at", "'2019-12-27T18:11:19.117Z'",
			"--library-id", "library_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"asset_data: asset_data\n" +
			"device_asset_id: device_asset_id\n" +
			"device_id: device_id\n" +
			"file_created_at: '2019-12-27T18:11:19.117Z'\n" +
			"file_modified_at: '2019-12-27T18:11:19.117Z'\n" +
			"library_id: library_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"assets", "create",
		)
	})
}

func TestAssetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "retrieve",
			"--asset-id", "asset_id",
		)
	})
}

func TestAssetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "list",
			"--max-items", "10",
			"--album-id", "album_id",
			"--id", "[string, string]",
			"--library-id", "library_id",
			"--limit", "1",
			"--local-datetime-after", "'2019-12-27T18:11:19.117Z'",
			"--local-datetime-before", "'2019-12-27T18:11:19.117Z'",
			"--person-id", "person_id",
			"--starting-after-id", "starting_after_id",
		)
	})
}

func TestAssetsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "delete",
			"--asset-id", "asset_id",
		)
	})
}

func TestAssetsCheckExistence(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "check-existence",
			"--library-id", "library_id",
			"--checksum-sha1", "[string]",
			"--checksum", "[string]",
			"--device-asset-id", "[string]",
			"--device-id", "deviceId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"checksum_sha1s:\n" +
			"  - string\n" +
			"checksums:\n" +
			"  - string\n" +
			"deviceAssetIds:\n" +
			"  - string\n" +
			"deviceId: deviceId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"assets", "check-existence",
			"--library-id", "library_id",
		)
	})
}

func TestAssetsCounts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "counts",
			"--album-id", "album_id",
			"--group-by", "group_by",
			"--library-id", "library_id",
			"--limit", "1",
			"--local-datetime-after", "'2019-12-27T18:11:19.117Z'",
			"--local-datetime-before", "'2019-12-27T18:11:19.117Z'",
			"--person-id", "person_id",
		)
	})
}

func TestAssetsDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "download",
			"--asset-id", "asset_id",
			"--output", "/dev/null",
		)
	})
}

func TestAssetsDownloadThumbnail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "download-thumbnail",
			"--asset-id", "asset_id",
			"--size", "size",
			"--output", "/dev/null",
		)
	})
}
