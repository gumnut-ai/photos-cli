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
		&requestflag.Flag[*string]{
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
	Usage:   "Fetches one asset and its associated metadata. Use this when you already have a\nspecific asset ID (e.g., from `list_assets`, `search_assets`, or\n`list_album_assets`) and need its full details. For bulk fetch of multiple known\nIDs, prefer `list_assets` with the `ids` parameter to avoid N round trips.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "asset-id",
			Usage:     "Asset ID (with `asset_` prefix) to fetch. Obtain from `list_assets`, `search_assets`, or `list_album_assets`.",
			Required:  true,
			PathParam: "asset_id",
		},
	},
	Action:          handleAssetsRetrieve,
	HideHelpCommand: true,
}

var assetsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of assets ordered by local capture time (newest first).\nUse this tool for structured browsing and filtering — when the request can be\nexpressed as exact filters on album membership, people, date range, or specific\nasset IDs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "album-id",
			Usage:     "Return only assets that are in the album with this ID. Equivalent to calling `list_album_assets` with `album_id` and then fetching each asset — prefer this param when you need the full asset metadata in one call.",
			QueryPath: "album_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Look up specific assets by ID (max 100; each ID has the `asset_` prefix). Use this for bulk fetch when you already have asset IDs. Combines with other filters (album_id, person_id, datetime range) using AND logic — the result is the intersection.",
			QueryPath: "ids",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library to list assets from. Optional if the user has a single library; required when they have multiple. Use `list_libraries` to enumerate available libraries.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of assets to return per page (1–200). Defaults to 20.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "local-datetime-after",
			Usage:     "Only include assets captured strictly after this instant (ISO 8601; exclusive). `local_datetime` is the photo's wall-clock time in the device's own timezone. Naive values compare directly against `local_datetime`. Timezone-aware values: assets with a known offset are compared in UTC (`local_datetime - offset`); assets without an offset fall back to wall-clock comparison against `local_datetime`. Equivalent in purpose to `captured_after` on `search_assets` (naming inconsistency is tracked as a follow-up).",
			QueryPath: "local_datetime_after",
		},
		&requestflag.Flag[any]{
			Name:      "local-datetime-before",
			Usage:     "Only include assets captured strictly before this instant (ISO 8601; exclusive). Same awareness/offset semantics as `local_datetime_after`. Equivalent in purpose to `captured_before` on `search_assets` (naming inconsistency is tracked as a follow-up).",
			QueryPath: "local_datetime_before",
		},
		&requestflag.Flag[*string]{
			Name:      "person-id",
			Usage:     "Return only assets containing a face belonging to this person. Singular on this tool; the sibling `search_assets` uses `person_ids` (plural, ALL-of).",
			QueryPath: "person_id",
		},
		&requestflag.Flag[*string]{
			Name:      "starting-after-id",
			Usage:     "Cursor for pagination. Pass the `id` of the last asset in the previous response's `data` to fetch the next page. Omit for the first page. `list_assets` uses cursor pagination; the sibling `search_assets` uses 1-indexed `page` numbers (naming inconsistency is tracked as a follow-up).",
			QueryPath: "starting_after_id",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Which set of assets to read from: `live` (default — only assets that are not trashed), `trashed` (only trashed assets, ordered by most recently trashed), or `all` (both live and trashed, ordered by capture time like `live`).",
			Default:   "live",
			QueryPath: "state",
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
	Usage:   "Deletes the asset entirely — the database record, the stored file, and all\nassociated data (faces, album links, etc.). **Irreversible.** Prefer\n`trash_assets` for the user's standard delete action so accidents can be\nrecovered.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "asset-id",
			Usage:     "Asset ID (with `asset_` prefix) of the asset to permanently delete.",
			Required:  true,
			PathParam: "asset_id",
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
		&requestflag.Flag[*string]{
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
		&requestflag.Flag[*string]{
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
		&requestflag.Flag[*string]{
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
		&requestflag.Flag[*string]{
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
		&requestflag.Flag[*string]{
			Name:      "person-id",
			Usage:     "Filter by assets associated with a specific person ID",
			QueryPath: "person_id",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Which set of assets to count: `live` (default — excludes trashed assets), `trashed` (only trashed assets), or `all` (both live and trashed).",
			Default:   "live",
			QueryPath: "state",
		},
	},
	Action:          handleAssetsCounts,
	HideHelpCommand: true,
}

var assetsDeleteList = cli.Command{
	Name:    "delete-list",
	Usage:   "Hard-deletes each specified asset — the database record, the stored file, and\nall associated data (faces, album links, etc.). **Irreversible.** Prefer\n`trash_assets` for the user's standard delete action so accidents can be\nrecovered.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "id",
			Usage:    "Asset IDs (each with the `asset_` prefix) to operate on. Up to 100 ids per request.",
			Required: true,
			BodyPath: "ids",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library that owns the assets. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
	},
	Action:          handleAssetsDeleteList,
	HideHelpCommand: true,
}

var assetsEmptyTrash = cli.Command{
	Name:    "empty-trash",
	Usage:   "Hard-deletes every trashed asset in the caller's library in one shot — storage\nand CDN are cleaned up via the same outbox path as the scheduled purge task.\n**Irreversible**. Deliberately not exposed as an MCP tool.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library whose trashed assets to permanently delete. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
	},
	Action:          handleAssetsEmptyTrash,
	HideHelpCommand: true,
}

var assetsRestore = cli.Command{
	Name:    "restore",
	Usage:   "Restores trashed assets so they reappear in default list/search results.\nIdempotent — assets that are already live are silently skipped.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "id",
			Usage:    "Asset IDs (each with the `asset_` prefix) to operate on. Up to 100 ids per request.",
			Required: true,
			BodyPath: "ids",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library that owns the assets. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
	},
	Action:          handleAssetsRestore,
	HideHelpCommand: true,
}

var assetsTrash = cli.Command{
	Name:    "trash",
	Usage:   "Soft-deletes the given assets. Trashed assets are excluded from default\nlist/search results and are purged after the configured retention window.\n**Reversible** via `restore_assets` until purge.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "id",
			Usage:    "Asset IDs (each with the `asset_` prefix) to operate on. Up to 100 ids per request.",
			Required: true,
			BodyPath: "ids",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library that owns the assets. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
	},
	Action:          handleAssetsTrash,
	HideHelpCommand: true,
}

func handleAssetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := photos.AssetNewParams{}

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

	params := photos.AssetListParams{}

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

	params := photos.AssetCheckExistenceParams{}

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

	params := photos.AssetCountsParams{}

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

func handleAssetsDeleteList(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := photos.AssetDeleteListParams{}

	return client.Assets.DeleteList(ctx, params, options...)
}

func handleAssetsEmptyTrash(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := photos.AssetEmptyTrashParams{}

	return client.Assets.EmptyTrash(ctx, params, options...)
}

func handleAssetsRestore(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := photos.AssetRestoreParams{}

	return client.Assets.Restore(ctx, params, options...)
}

func handleAssetsTrash(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := photos.AssetTrashParams{}

	return client.Assets.Trash(ctx, params, options...)
}
