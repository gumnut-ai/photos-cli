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

var albumsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new, empty album with optional name and description in the specified\nlibrary.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "library-id",
			BodyPath: "library_id",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleAlbumsCreate,
	HideHelpCommand: true,
}

var albumsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves details for a specific album.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Required: true,
		},
	},
	Action:          handleAlbumsRetrieve,
	HideHelpCommand: true,
}

var albumsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the name and/or description of a specific album.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleAlbumsUpdate,
	HideHelpCommand: true,
}

var albumsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of albums from the specified library, ordered by\ncreation time, descending. Can be filtered by asset_id or specific album IDs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "asset-id",
			Usage:     "Filter albums containing this asset ID (optional)",
			QueryPath: "asset_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Filter by specific album IDs (max 100)",
			QueryPath: "ids",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library to list albums from (optional)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Max number of albums to return (1-200)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "starting-after-id",
			Usage:     "Cursor for pagination. Pass the `id` of the last album from the previous page to get the next page.",
			QueryPath: "starting_after_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAlbumsList,
	HideHelpCommand: true,
}

var albumsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a specific album. Note: This does not delete the assets within the\nalbum.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Required: true,
		},
	},
	Action:          handleAlbumsDelete,
	HideHelpCommand: true,
}

func handleAlbumsCreate(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AlbumNewParams{}

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
	_, err = client.Albums.New(ctx, params, options...)
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
		Title:          "albums create",
		Transform:      transform,
	})
}

func handleAlbumsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("album-id") && len(unusedArgs) > 0 {
		cmd.Set("album-id", unusedArgs[0])
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
	_, err = client.Albums.Get(ctx, cmd.Value("album-id").(string), options...)
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
		Title:          "albums retrieve",
		Transform:      transform,
	})
}

func handleAlbumsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("album-id") && len(unusedArgs) > 0 {
		cmd.Set("album-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AlbumUpdateParams{}

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
	_, err = client.Albums.Update(
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
		Title:          "albums update",
		Transform:      transform,
	})
}

func handleAlbumsList(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AlbumListParams{}

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
		_, err = client.Albums.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "albums list",
			Transform:      transform,
		})
	} else {
		iter := client.Albums.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "albums list",
			Transform:      transform,
		})
	}
}

func handleAlbumsDelete(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("album-id") && len(unusedArgs) > 0 {
		cmd.Set("album-id", unusedArgs[0])
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

	return client.Albums.Delete(ctx, cmd.Value("album-id").(string), options...)
}
