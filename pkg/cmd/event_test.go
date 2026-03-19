// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/gumnut-ai/photos-cli/internal/mocktest"
)

func TestEventsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"events", "get",
			"--after-cursor", "after_cursor",
			"--created-at-gte", "'2019-12-27T18:11:19.117Z'",
			"--created-at-lt", "'2019-12-27T18:11:19.117Z'",
			"--entity-types", "entity_types",
			"--library-id", "library_id",
			"--limit", "1",
		)
	})
}
