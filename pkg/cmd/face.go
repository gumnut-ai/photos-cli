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

var facesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetches one face's details by ID (bounding box, assigned person, timestamps,\nthumbnail). Use when you already have a `face_id`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "face-id",
			Usage:     "Face ID (with `face_` prefix) to fetch. Obtain from `list_faces` or from the `faces` array on `get_asset` / `list_assets` responses.",
			Required:  true,
			PathParam: "face_id",
		},
		&requestflag.Flag[any]{
			Name:      "include",
			Usage:     "Opt-in expansion fields. See `list_faces` for supported values. Accepts multiple `include=` query params or a single comma-delimited value.",
			QueryPath: "include",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library the face belongs to. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
	},
	Action:          handleFacesRetrieve,
	HideHelpCommand: true,
}

var facesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Assigns a face to a specific person, or detaches it from its current person (set\n`person_id` to null). This is the right tool for 'this face is Alice' or 'this\nface isn't Bob after all'.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "face-id",
			Usage:     "Face ID (with `face_` prefix) of the face detection to update.",
			Required:  true,
			PathParam: "face_id",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library the face belongs to. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[*string]{
			Name:     "person-id",
			Usage:    "Target person ID (with `person_` prefix) to assign this face to. Pass `null` to detach the face from its current person without deleting either. Get IDs from `list_people`; use `create_person` first if the target identity doesn't exist yet.",
			BodyPath: "person_id",
		},
	},
	Action:          handleFacesUpdate,
	HideHelpCommand: true,
}

var facesList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of individual face detections (with bounding boxes),\nordered by creation time (newest first), optionally filtered by asset, person,\nor ID. Each row is a single face in a single asset — a person with many photos\nwill have many face rows.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "asset-id",
			Usage:     "Return only faces detected in this asset. Useful for 'show me all the faces in this photo'.",
			QueryPath: "asset_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Look up specific faces by ID (max 100). IDs use the `face_` prefix. Accepts multiple `ids=` query params or a single comma-delimited value (e.g., `ids=face_1,face_2`).",
			QueryPath: "ids",
		},
		&requestflag.Flag[any]{
			Name:      "include",
			Usage:     "Opt-in expansion fields. Supported values: `cluster_assignment` (adds the nested `cluster_assignment` object — `distance_to_person` and a top-K `candidates` list of nearby Persons). Accepts multiple `include=` query params or a single comma-delimited value (e.g., `include=cluster_assignment`).",
			QueryPath: "include",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library to list from. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of faces per page (1–200). Defaults to 20.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[*string]{
			Name:      "person-id",
			Usage:     "Return only faces currently assigned to this person. Useful for reviewing or curating a person's face cluster.",
			QueryPath: "person_id",
		},
		&requestflag.Flag[*string]{
			Name:      "starting-after-id",
			Usage:     "Cursor for pagination. Pass the `id` of the last face in the previous response's `data` to fetch the next page. Omit for the first page.",
			QueryPath: "starting_after_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleFacesList,
	HideHelpCommand: true,
}

var facesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Removes one face detection row; the underlying asset and the person this face\nwas assigned to are both preserved.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "face-id",
			Usage:     "Face ID (with `face_` prefix) of the face detection to delete.",
			Required:  true,
			PathParam: "face_id",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library the face belongs to. Optional if the user has a single library; required when they have multiple.",
			QueryPath: "library_id",
		},
	},
	Action:          handleFacesDelete,
	HideHelpCommand: true,
}

func handleFacesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("face-id") && len(unusedArgs) > 0 {
		cmd.Set("face-id", unusedArgs[0])
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

	params := photos.FaceGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Faces.Get(
		ctx,
		cmd.Value("face-id").(string),
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
		Title:          "faces retrieve",
		Transform:      transform,
	})
}

func handleFacesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("face-id") && len(unusedArgs) > 0 {
		cmd.Set("face-id", unusedArgs[0])
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

	params := photos.FaceUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Faces.Update(
		ctx,
		cmd.Value("face-id").(string),
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
		Title:          "faces update",
		Transform:      transform,
	})
}

func handleFacesList(ctx context.Context, cmd *cli.Command) error {
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

	params := photos.FaceListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Faces.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "faces list",
			Transform:      transform,
		})
	} else {
		iter := client.Faces.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "faces list",
			Transform:      transform,
		})
	}
}

func handleFacesDelete(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("face-id") && len(unusedArgs) > 0 {
		cmd.Set("face-id", unusedArgs[0])
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

	params := photos.FaceDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Faces.Delete(
		ctx,
		cmd.Value("face-id").(string),
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
		Title:          "faces delete",
		Transform:      transform,
	})
}
