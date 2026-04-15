// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gumnut-ai/photos-cli/internal/apiquery"
	"github.com/gumnut-ai/photos-cli/internal/requestflag"
	"github.com/gumnut-ai/photos-sdk-go"
	"github.com/gumnut-ai/photos-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var albumAssetsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of album-asset links, ordered by creation time,\ndescending. Can be filtered by album_id, asset_id, or specific album-asset IDs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "album-id",
			Usage:     "Filter by album ID",
			QueryPath: "album_id",
		},
		&requestflag.Flag[any]{
			Name:      "asset-id",
			Usage:     "Filter by asset ID",
			QueryPath: "asset_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Filter by specific album-asset IDs (max 100)",
			QueryPath: "ids",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library ID (required if user has multiple libraries)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Max number of results to return (1-200)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "starting-after-id",
			Usage:     "Album-asset ID to start listing after",
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
	Usage:   "Retrieves details for a specific album-asset link.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-asset-id",
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
		return ShowJSON(os.Stdout, os.Stderr, "album-assets list", obj, format, explicitFormat, transform)
	} else {
		iter := client.AlbumAssets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "album-assets list", iter, format, explicitFormat, transform, maxItems)
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
	return ShowJSON(os.Stdout, os.Stderr, "album-assets get", obj, format, explicitFormat, transform)
}
