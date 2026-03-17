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
			t,
			"--api-key", "string",
			"people", "create",
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
			t, pipeData,
			"--api-key", "string",
			"people", "create",
		)
	})
}

func TestPeopleRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"people", "retrieve",
			"--person-id", "person_id",
		)
	})
}

func TestPeopleUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"people", "update",
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
			t, pipeData,
			"--api-key", "string",
			"people", "update",
			"--person-id", "person_id",
		)
	})
}

func TestPeopleList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"people", "list",
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
			t,
			"--api-key", "string",
			"people", "delete",
			"--person-id", "person_id",
		)
	})
}
