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
	Usage:   "Adds one or more existing assets to a specific album. Assets must be in the same\nlibrary as the album. Duplicate assets are ignored.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "asset-id",
			Required: true,
			BodyPath: "asset_ids",
		},
	},
	Action:          handleAlbumsAssetsAssociationsAdd,
	HideHelpCommand: true,
}

var albumsAssetsAssociationsRemove = cli.Command{
	Name:    "remove",
	Usage:   "Removes one or more assets from a specific album. Note: This does not delete the\nassets themselves.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "asset-id",
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
