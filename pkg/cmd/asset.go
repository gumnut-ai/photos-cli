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

var assetsCreate = cli.Command{
	Name:    "create",
	Usage:   "Uploads a new asset file (image or video) along with its metadata to the\nspecified library. If no library_id is provided and the user only has one\nlibrary, uses that library. If the user has multiple libraries, library_id is\nrequired.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "asset-data",
			Usage:     "The asset file to upload",
			Required:  true,
			BodyPath:  "asset_data",
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:     "device-asset-id",
			Required: true,
			BodyPath: "device_asset_id",
		},
		&requestflag.Flag[string]{
			Name:     "device-id",
			Required: true,
			BodyPath: "device_id",
		},
		&requestflag.Flag[any]{
			Name:     "file-created-at",
			Required: true,
			BodyPath: "file_created_at",
		},
		&requestflag.Flag[any]{
			Name:     "file-modified-at",
			Required: true,
			BodyPath: "file_modified_at",
		},
		&requestflag.Flag[any]{
			Name:     "library-id",
			Usage:    "Library to upload asset to (optional)",
			BodyPath: "library_id",
		},
	},
	Action:          handleAssetsCreate,
	HideHelpCommand: true,
}

var assetsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves detailed metadata for a specific asset, including EXIF information,\nasset metrics, faces, and people.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "asset-id",
			Required: true,
		},
	},
	Action:          handleAssetsRetrieve,
	HideHelpCommand: true,
}

var assetsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of assets from the specified library, optionally\nfiltered by album, person, or specific asset IDs. Asset data includes metrics,\nEXIF data, faces, and people. Assets are ordered by local creation time,\ndescending.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "album-id",
			Usage:     "Filter by assets in a specific album",
			QueryPath: "album_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Filter by specific asset IDs (max 100)",
			QueryPath: "ids",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library to list assets from (optional)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Max number of assets to return (1-200)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "local-datetime-after",
			Usage:     "Only include assets with local_datetime after this value (ISO 8601). Naive values compare directly against local_datetime. Timezone-aware values: assets with a known offset are compared in UTC (local_datetime - offset); assets without an offset fall back to wall-clock comparison against local_datetime.",
			QueryPath: "local_datetime_after",
		},
		&requestflag.Flag[any]{
			Name:      "local-datetime-before",
			Usage:     "Only include assets with local_datetime before this value (ISO 8601). Naive values compare directly against local_datetime. Timezone-aware values: assets with a known offset are compared in UTC (local_datetime - offset); assets without an offset fall back to wall-clock comparison against local_datetime.",
			QueryPath: "local_datetime_before",
		},
		&requestflag.Flag[any]{
			Name:      "person-id",
			Usage:     "Filter by assets associated with a specific person ID",
			QueryPath: "person_id",
		},
		&requestflag.Flag[any]{
			Name:      "starting-after-id",
			Usage:     "Cursor for pagination. Pass the `id` of the last asset from the previous page to get the next page.",
			QueryPath: "starting_after_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAssetsList,
	HideHelpCommand: true,
}

var assetsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a specific asset and its associated data (including the file from\nstorage).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "asset-id",
			Required: true,
		},
	},
	Action:          handleAssetsDelete,
	HideHelpCommand: true,
}

var assetsCheckExistence = cli.Command{
	Name:    "check-existence",
	Usage:   "Checks which assets exist in the user's library based on checksums or device\nidentifiers. Provide exactly one of: checksums, checksum_sha1s, or (deviceId AND\ndeviceAssetIds). List parameters are limited to 5000 items.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library to check assets in (optional)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[any]{
			Name:     "checksum-sha1",
			Usage:    "List of base64-encoded SHA-1 checksums to check for existence (for Immich compatibility)",
			BodyPath: "checksum_sha1s",
		},
		&requestflag.Flag[any]{
			Name:     "checksum",
			Usage:    "List of base64-encoded SHA-256 checksums to check for existence",
			BodyPath: "checksums",
		},
		&requestflag.Flag[any]{
			Name:     "device-asset-id",
			Usage:    "List of device asset IDs to check for existence (requires deviceId)",
			BodyPath: "deviceAssetIds",
		},
		&requestflag.Flag[any]{
			Name:     "device-id",
			Usage:    "Device ID to filter assets by (required with deviceAssetIds)",
			BodyPath: "deviceId",
		},
	},
	Action:          handleAssetsCheckExistence,
	HideHelpCommand: true,
}

var assetsCounts = cli.Command{
	Name:    "counts",
	Usage:   "Returns asset counts grouped by time period. Supports optional filtering by\nalbum, person, or date range. Results are ordered by time bucket descending.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "album-id",
			Usage:     "Filter by assets in a specific album",
			QueryPath: "album_id",
		},
		&requestflag.Flag[string]{
			Name:      "group-by",
			Usage:     "Time period to group counts by. Currently only 'month' is supported.",
			Default:   "month",
			QueryPath: "group_by",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library to count assets in (optional)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of time buckets to return (1-200)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "local-datetime-after",
			Usage:     "Only include assets with local_datetime after this value (ISO 8601). Naive values compare directly against local_datetime. Timezone-aware values: assets with a known offset are compared in UTC (local_datetime - offset); assets without an offset fall back to wall-clock comparison against local_datetime.",
			QueryPath: "local_datetime_after",
		},
		&requestflag.Flag[any]{
			Name:      "local-datetime-before",
			Usage:     "Only include assets with local_datetime before this value (ISO 8601). Naive values compare directly against local_datetime. Timezone-aware values: assets with a known offset are compared in UTC (local_datetime - offset); assets without an offset fall back to wall-clock comparison against local_datetime. Use the last time_bucket from a previous response to paginate.",
			QueryPath: "local_datetime_before",
		},
		&requestflag.Flag[any]{
			Name:      "person-id",
			Usage:     "Filter by assets associated with a specific person ID",
			QueryPath: "person_id",
		},
	},
	Action:          handleAssetsCounts,
	HideHelpCommand: true,
}

func handleAssetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AssetNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Assets.New(ctx, params, options...)
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
		Title:          "assets create",
		Transform:      transform,
	})
}

func handleAssetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("asset-id") && len(unusedArgs) > 0 {
		cmd.Set("asset-id", unusedArgs[0])
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
	_, err = client.Assets.Get(ctx, cmd.Value("asset-id").(string), options...)
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
		Title:          "assets retrieve",
		Transform:      transform,
	})
}

func handleAssetsList(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AssetListParams{}

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
		_, err = client.Assets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "assets list",
			Transform:      transform,
		})
	} else {
		iter := client.Assets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "assets list",
			Transform:      transform,
		})
	}
}

func handleAssetsDelete(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("asset-id") && len(unusedArgs) > 0 {
		cmd.Set("asset-id", unusedArgs[0])
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

	return client.Assets.Delete(ctx, cmd.Value("asset-id").(string), options...)
}

func handleAssetsCheckExistence(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AssetCheckExistenceParams{}

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
	_, err = client.Assets.CheckExistence(ctx, params, options...)
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
		Title:          "assets check-existence",
		Transform:      transform,
	})
}

func handleAssetsCounts(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.AssetCountsParams{}

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
	_, err = client.Assets.Counts(ctx, params, options...)
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
		Title:          "assets counts",
		Transform:      transform,
	})
}
