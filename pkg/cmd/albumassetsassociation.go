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

var albumsAssetsAssociationsAdd = cli.Command{
	Name:    "add",
	Usage:   "Adds one or more existing assets to the specified album. Assets must already be\nin the same library as the album (this tool does not upload new assets). Assets\nalready in the album are silently skipped and returned separately as\n`duplicate_assets`. Idempotent: calling with the same IDs twice leaves the album\nin the same state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Usage:    "Album ID (with `album_` prefix) of the album to add the assets to.",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "asset-id",
			Usage:    "Asset IDs (with `asset_` prefix) to associate with the album. Get IDs from `list_assets`, `search_assets`, or `list_album_assets`.",
			Required: true,
			BodyPath: "asset_ids",
		},
	},
	Action:          handleAlbumsAssetsAssociationsAdd,
	HideHelpCommand: true,
}

var albumsAssetsAssociationsRemove = cli.Command{
	Name:    "remove",
	Usage:   "Detaches one or more assets from the given album. The assets remain in the\nlibrary and in any other albums they belong to. Use `delete_asset` to delete the\nasset entirely. To empty an album completely, call `list_album_assets` to get\nthe links and then remove them, or delete the album itself with `delete_album`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Usage:    "Album ID (with `album_` prefix) of the album to detach assets from.",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "asset-id",
			Usage:    "Asset IDs (with `asset_` prefix) to associate with the album. Get IDs from `list_assets`, `search_assets`, or `list_album_assets`.",
			Required: true,
			BodyPath: "asset_ids",
		},
	},
	Action:          handleAlbumsAssetsAssociationsRemove,
	HideHelpCommand: true,
}

func handleAlbumsAssetsAssociationsAdd(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("album-id") && len(unusedArgs) > 0 {
		cmd.Set("album-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AlbumAssetsAssociationAddParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Albums.AssetsAssociations.Add(
		ctx,
		cmd.Value("album-id").(string),
		params,
		options...,
	)
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
		Title:          "albums:assets-associations add",
		Transform:      transform,
	})
}

func handleAlbumsAssetsAssociationsRemove(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("album-id") && len(unusedArgs) > 0 {
		cmd.Set("album-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AlbumAssetsAssociationRemoveParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Albums.AssetsAssociations.Remove(
		ctx,
		cmd.Value("album-id").(string),
		params,
		options...,
	)
}
