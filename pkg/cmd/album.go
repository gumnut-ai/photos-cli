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
	Usage:   "Creates an album (with optional name and description) and returns it. The album\nstarts empty — follow up with `add_assets_to_album` to populate it. To rename an\nexisting album, use `update_album` instead of creating a new one.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Optional free-form description shown alongside the album name.",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "library-id",
			Usage:    "Library to create the album in. Optional if the user has a single library; required when they have multiple. Use `list_libraries` to enumerate.",
			BodyPath: "library_id",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Display name for the new album. Optional; callers that need to name an album can set it here or via `update_album` after creation.",
			BodyPath: "name",
		},
	},
	Action:          handleAlbumsCreate,
	HideHelpCommand: true,
}

var albumsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetches one album's metadata (name, description, cover, counts). Use when you\nalready have an album ID. Does not include the album's assets — use\n`list_album_assets` or `list_assets` with `album_id` for that.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Usage:    "Album ID (with `album_` prefix) to fetch. Obtain from `list_albums` (optionally filtered by `asset_id` to find albums containing a specific asset), `list_album_assets`, or any response containing an album reference.",
			Required: true,
		},
	},
	Action:          handleAlbumsRetrieve,
	HideHelpCommand: true,
}

var albumsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the `name` and/or `description` of an existing album. Only the fields\nincluded in the request body are changed. To modify the contents of an album,\nuse `add_assets_to_album` / `remove_assets_from_album` instead — this tool only\nchanges album metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Usage:    "Album ID (with `album_` prefix) of the album to rename or re-describe.",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "New free-form description for the album. Omit to leave unchanged.",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "New display name for the album. Omit to leave unchanged.",
			BodyPath: "name",
		},
	},
	Action:          handleAlbumsUpdate,
	HideHelpCommand: true,
}

var albumsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of albums ordered by creation time (newest first). Use\nthis to enumerate a user's albums or to find which albums contain a specific\nasset (via `asset_id`).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "asset-id",
			Usage:     "Return only albums that contain this asset. Useful for answering 'which albums is this photo in?' without calling `list_album_assets`.",
			QueryPath: "asset_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Look up specific albums by ID (max 100; each ID has the `album_` prefix). Use for bulk fetch when IDs are already known.",
			QueryPath: "ids",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library to list albums from. Optional if the user has a single library; required when they have multiple. Use `list_libraries` to enumerate.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of albums to return per page (1–200). Defaults to 20.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "starting-after-id",
			Usage:     "Cursor for pagination. Pass the `id` of the last album in the previous response's `data` to fetch the next page. Omit for the first page.",
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
	Usage:   "Deletes the album itself. Assets that were in the album remain in the library —\nonly the album and its asset-links are removed. Use `delete_asset` to delete the\nunderlying assets, or `remove_assets_from_album` to detach specific assets from\nan album you want to keep.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "album-id",
			Usage:    "Album ID (with `album_` prefix) of the album to delete.",
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
