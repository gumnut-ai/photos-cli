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

var peopleCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new person record (a named identity for grouping faces). Most people\nare auto-created by face clustering, so this tool is typically used only when\nthe user explicitly wants to introduce a new identity before any faces are\nattached.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "birth-date",
			Usage:    "Optional birth date (ISO 8601 date, YYYY-MM-DD) for this person.",
			BodyPath: "birth_date",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-favorite",
			Usage:    "If true, the person is marked as a favorite. Defaults to false.",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "is_favorite",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-hidden",
			Usage:    "If true, the person is hidden from default listings. Defaults to false.",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "is_hidden",
		},
		&requestflag.Flag[*string]{
			Name:     "library-id",
			Usage:    "Library to create the person in. Optional if the user has a single library; required when they have multiple.",
			BodyPath: "library_id",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Display name for the new person (e.g., 'Alice'). Optional — unnamed people can be named later via `update_person`.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "thumbnail-face-id",
			Usage:    "ID of the face to use as this person's thumbnail (with `face_` prefix). Typically set after the person has at least one associated face — get face IDs from `list_faces`.",
			BodyPath: "thumbnail_face_id",
		},
	},
	Action:          handlePeopleCreate,
	HideHelpCommand: true,
}

var peopleRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetches one person's metadata by ID (name, asset count, thumbnail, etc.). Use\nthis when you already have a `person_id`. To find photos that contain this\nperson, use `search_assets` with `person_ids` or `list_assets` with `person_id`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "person-id",
			Usage:     "Person ID (with `person_` prefix) to fetch. Obtain from `list_people`, `get_face.person_id`, or any response containing a person reference.",
			Required:  true,
			PathParam: "person_id",
		},
		&requestflag.Flag[*string]{
			Name:      "include",
			Usage:     "Comma-separated list of opt-in expansion fields. See `list_people` for supported values.",
			QueryPath: "include",
		},
	},
	Action:          handlePeopleRetrieve,
	HideHelpCommand: true,
}

var peopleUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a person's name, birth date, visibility, or thumbnail. Only the fields\nincluded in the request body are changed. Typical use: assigning a name ('name\nthis face cluster \"Alice\"') or choosing a better thumbnail.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "person-id",
			Usage:     "Person ID (with `person_` prefix) of the person to update.",
			Required:  true,
			PathParam: "person_id",
		},
		&requestflag.Flag[any]{
			Name:     "birth-date",
			Usage:    "New birth date (ISO 8601 date). Omit to leave unchanged.",
			BodyPath: "birth_date",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-favorite",
			Usage:    "Mark or unmark this person as a favorite. Omit to leave unchanged.",
			BodyPath: "is_favorite",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-hidden",
			Usage:    "Hide or unhide this person. Omit to leave unchanged.",
			BodyPath: "is_hidden",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "New display name. Omit to leave unchanged.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "thumbnail-face-id",
			Usage:    "New thumbnail face ID for this person. Omit to leave unchanged. Get face IDs from `list_faces`.",
			BodyPath: "thumbnail_face_id",
		},
	},
	Action:          handlePeopleUpdate,
	HideHelpCommand: true,
}

var peopleList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of people (named identities that group one or more\nfaces), ordered by creation time (newest first), optionally filtered by asset,\nalbum, name, or ID. Use this to enumerate who appears in the library, to resolve\na user-typed name to a `person_id`, or to find who appears in a specific asset\nor album.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "album-id",
			Usage:     "Return only people who appear in at least one asset of this album. Useful for 'who is in this album?'.",
			QueryPath: "album_id",
		},
		&requestflag.Flag[*string]{
			Name:      "asset-id",
			Usage:     "Return only people who have at least one face in this asset. Useful for 'who is in this photo?'.",
			QueryPath: "asset_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Look up specific people by ID (max 100; each ID has the `person_` prefix). When set, `name_filter` defaults to `all` so unnamed clusters are included in the lookup.",
			QueryPath: "ids",
		},
		&requestflag.Flag[*string]{
			Name:      "include",
			Usage:     "Comma-separated list of opt-in expansion fields. Supported values: `cluster_metrics` (adds the nested `cluster_metrics` object — `pairwise_p90`, `pairwise_mean`, `face_count` — for each Person with a populated centroid). Unknown values return 422.",
			QueryPath: "include",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library to list from. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of people to return per page (1–200). Defaults to 20.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[*string]{
			Name:      "name",
			Usage:     "Filter by name using case-insensitive substring matching. Use this to resolve a user-supplied name like 'Alice' into a `person_id`, then pass that ID into `search_assets.person_ids` or `list_assets.person_id`.",
			QueryPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:      "name-filter",
			Usage:     "Filter by name status: `named` returns only people with a name; `unnamed` returns only nameless face clusters awaiting a name; `all` returns both. Defaults to `named` (or `all` when `ids` is provided).",
			QueryPath: "name_filter",
		},
		&requestflag.Flag[*string]{
			Name:      "starting-after-id",
			Usage:     "Cursor for pagination. Pass the `id` of the last person in the previous response's `data` to fetch the next page. Omit for the first page.",
			QueryPath: "starting_after_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handlePeopleList,
	HideHelpCommand: true,
}

var peopleDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the person record; the faces that were attached to this person are not\ndeleted — they become unassigned and will be re-clustered on the next clustering\npass.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "person-id",
			Usage:     "Person ID (with `person_` prefix) of the person to delete.",
			Required:  true,
			PathParam: "person_id",
		},
	},
	Action:          handlePeopleDelete,
	HideHelpCommand: true,
}

var peopleMerge = cli.Command{
	Name:    "merge",
	Usage:   "Merges one or more source people into the primary person identified by the URL.\nAll faces from source people are reassigned to the primary person. Source people\nare permanently deleted (this cannot be undone). The primary person's centroid\nembedding is recalculated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "person-id",
			Required:  true,
			PathParam: "person_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "source-person-id",
			Usage:    "IDs of the people to merge into the primary person. These people will be deleted after their faces are moved.",
			Required: true,
			BodyPath: "source_person_ids",
		},
	},
	Action:          handlePeopleMerge,
	HideHelpCommand: true,
}

func handlePeopleCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := photos.PersonNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.People.New(ctx, params, options...)
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
		Title:          "people create",
		Transform:      transform,
	})
}

func handlePeopleRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
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

	params := photos.PersonGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.People.Get(
		ctx,
		cmd.Value("person-id").(string),
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
		Title:          "people retrieve",
		Transform:      transform,
	})
}

func handlePeopleUpdate(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := photos.PersonUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.People.Update(
		ctx,
		cmd.Value("person-id").(string),
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
		Title:          "people update",
		Transform:      transform,
	})
}

func handlePeopleList(ctx context.Context, cmd *cli.Command) error {
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

	params := photos.PersonListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.People.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "people list",
			Transform:      transform,
		})
	} else {
		iter := client.People.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "people list",
			Transform:      transform,
		})
	}
}

func handlePeopleDelete(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
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

	return client.People.Delete(ctx, cmd.Value("person-id").(string), options...)
}

func handlePeopleMerge(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := photos.PersonMergeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.People.Merge(
		ctx,
		cmd.Value("person-id").(string),
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
		Title:          "people merge",
		Transform:      transform,
	})
}
