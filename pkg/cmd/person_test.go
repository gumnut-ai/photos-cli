// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestPeopleCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "people", "create",
			"--api-key", "string",
			"--birth-date", "'2019-12-27'",
			"--is-favorite=true",
			"--is-hidden=true",
			"--library-id", "library_id",
			"--name", "name",
			"--thumbnail-face-id", "thumbnail_face_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"birth_date: '2019-12-27'\n" +
			"is_favorite: true\n" +
			"is_hidden: true\n" +
			"library_id: library_id\n" +
			"name: name\n" +
			"thumbnail_face_id: thumbnail_face_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "people", "create",
			"--api-key", "string",
		)
	})
}

func TestPeopleRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "people", "retrieve",
			"--api-key", "string",
			"--person-id", "person_id",
		)
	})
}

func TestPeopleUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "people", "update",
			"--api-key", "string",
			"--person-id", "person_id",
			"--birth-date", "'2019-12-27'",
			"--is-favorite=true",
			"--is-hidden=true",
			"--name", "name",
			"--thumbnail-face-id", "thumbnail_face_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"birth_date: '2019-12-27'\n" +
			"is_favorite: true\n" +
			"is_hidden: true\n" +
			"name: name\n" +
			"thumbnail_face_id: thumbnail_face_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "people", "update",
			"--api-key", "string",
			"--person-id", "person_id",
		)
	})
}

func TestPeopleList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "people", "list",
			"--api-key", "string",
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

func TestPeopleDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "people", "delete",
			"--api-key", "string",
			"--person-id", "person_id",
		)
	})
}
