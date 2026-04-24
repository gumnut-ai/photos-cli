// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/gumnut-ai/photos-cli/internal/apiquery"
	"github.com/gumnut-ai/photos-cli/internal/requestflag"
	"github.com/gumnut-ai/photos-sdk-go"
	"github.com/gumnut-ai/photos-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var eventsGet = cli.Command{
	Name:    "get",
	Usage:   "Returns a paginated stream of change events (create/update/delete) for entities\nin the library. Each event is a lightweight record — `entity_type`, `entity_id`,\n`event_type`, and timestamps — pointing at a concrete entity that has changed.\nFollow up with `get_asset`, `get_album`, `get_person`, or `get_face` to fetch\nfull entity data when needed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "after-cursor",
			Usage:     "Opaque cursor from the last event of the previous page. Pass the `cursor` field from the last event to fetch the next page. Omit for the first page.",
			QueryPath: "after_cursor",
		},
		&requestflag.Flag[any]{
			Name:      "created-at-gte",
			Usage:     "Only return events created at or after this timestamp (ISO 8601). Set this to the previous sync's checkpoint when doing incremental sync.",
			QueryPath: "created_at_gte",
		},
		&requestflag.Flag[any]{
			Name:      "created-at-lt",
			Usage:     "Only return events created strictly before this timestamp (ISO 8601). Recommended for bounding a sync operation — capture `now` once and reuse it as `created_at_lt` across all pages so newly arriving events don't shift the window.",
			QueryPath: "created_at_lt",
		},
		&requestflag.Flag[any]{
			Name:      "entity-types",
			Usage:     "Comma-separated list of entity types to include (e.g., `asset,album`). Valid values: `asset`, `album`, `person`, `face`, `album_asset`, `exif`, `metadata`. Omit to receive events for all types.",
			QueryPath: "entity_types",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library to stream events from. Optional if the user has a single library; required when they have multiple. Use `list_libraries` to enumerate.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of events to return per page (1–200). Defaults to 20.",
			Default:   20,
			QueryPath: "limit",
		},
	},
	Action:          handleEventsGet,
	HideHelpCommand: true,
}

func handleEventsGet(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.EventGetParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Events.Get(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "events get",
		Transform:      transform,
	})
}
