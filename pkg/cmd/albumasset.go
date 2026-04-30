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

var albumAssetsList = cli.Command{
	Name:    "list",
	Usage:   "Returns paginated _link_ records describing which assets are in which albums —\neach row contains `album_id` + `asset_id` + link timestamps, not the full asset\nor album metadata. Use this when you specifically need the junction records (for\nsync or change tracking).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "album-id",
			Usage:     "Return only link records for this album ID. Equivalent to 'list the assets in this album' — in most cases prefer `list_assets` with `album_id` to get the asset metadata directly instead of the lightweight link records.",
			QueryPath: "album_id",
		},
		&requestflag.Flag[*string]{
			Name:      "asset-id",
			Usage:     "Return only link records for this asset ID. Equivalent to 'which albums contain this asset' — in most cases prefer `list_albums` with `asset_id` to get the album metadata directly.",
			QueryPath: "asset_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Look up specific album-asset link records by ID (max 100). The ID has the `album_asset_` prefix.",
			QueryPath: "ids",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library to list from. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of link records per page (1–200). Defaults to 20.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[*string]{
			Name:      "starting-after-id",
			Usage:     "Cursor for pagination. Pass the `id` of the last album-asset in the previous response's `data` to fetch the next page. Omit for the first page.",
			QueryPath: "starting_after_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAlbumAssetsList,
	HideHelpCommand: true,
}

var albumAssetsGet = cli.Command{
	Name:    "get",
	Usage:   "Fetches one album-asset link record (the junction row between an album and an\nasset). Rarely needed directly; most callers want `get_asset` or `get_album`\ninstead.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-asset-id",
			Usage:    "Album-asset junction row ID (with `album_asset_` prefix). Obtain from `list_album_assets`. Not the same as `asset_id` or `album_id`.",
			Required: true,
		},
	},
	Action:          handleAlbumAssetsGet,
	HideHelpCommand: true,
}

func handleAlbumAssetsList(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AlbumAssetListParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AlbumAssets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "album-assets list",
			Transform:      transform,
		})
	} else {
		iter := client.AlbumAssets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "album-assets list",
			Transform:      transform,
		})
	}
}

func handleAlbumAssetsGet(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("album-asset-id") && len(unusedArgs) > 0 {
		cmd.Set("album-asset-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.AlbumAssets.Get(ctx, cmd.Value("album-asset-id").(string), options...)
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
		Title:          "album-assets get",
		Transform:      transform,
	})
}
