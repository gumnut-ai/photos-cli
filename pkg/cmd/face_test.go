// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestFacesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"faces", "retrieve",
			"--face-id", "face_id",
			"--library-id", "library_id",
		)
	})
}

func TestFacesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"faces", "update",
			"--face-id", "face_id",
			"--library-id", "library_id",
			"--person-id", "person_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("person_id: person_id")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"faces", "update",
			"--face-id", "face_id",
			"--library-id", "library_id",
		)
	})
}

func TestFacesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"faces", "list",
			"--max-items", "10",
			"--asset-id", "asset_id",
			"--id", "[string, string]",
			"--library-id", "library_id",
			"--limit", "1",
			"--person-id", "person_id",
			"--starting-after-id", "starting_after_id",
		)
	})
}

func TestFacesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"faces", "delete",
			"--face-id", "face_id",
			"--library-id", "library_id",
		)
	})
}

func TestFacesDownloadThumbnail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"faces", "download-thumbnail",
			"--face-id", "face_id",
			"--output", "/dev/null",
		)
	})
}
