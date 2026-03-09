// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gumnut-ai/photos-cli/internal/apiquery"
	"github.com/gumnut-ai/photos-cli/internal/requestflag"
	"github.com/stainless-sdks/photos-go"
	"github.com/stainless-sdks/photos-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var eventsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieves a list of entity change events for syncing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "after-cursor",
			Usage:     "Cursor from the last event to paginate from. Pass the `cursor` field from the last event to get the next page.",
			QueryPath: "after_cursor",
		},
		&requestflag.Flag[any]{
			Name:      "created-at-gte",
			Usage:     "Only return events created at or after this timestamp (ISO 8601 format)",
			QueryPath: "created_at_gte",
		},
		&requestflag.Flag[any]{
			Name:      "created-at-lt",
			Usage:     "Only return events created before this timestamp (ISO 8601 format). Recommended for bounding sync operations.",
			QueryPath: "created_at_lt",
		},
		&requestflag.Flag[any]{
			Name:      "entity-types",
			Usage:     "Comma-separated list of entity types to include (e.g., 'asset,album'). Valid types: asset, album, person, face, album_asset, exif. Default: all types.",
			QueryPath: "entity_types",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library to list events from. If not provided, uses the user's default library.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of events to return (1-500)",
			Default:   100,
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "events get", obj, format, transform)
}
