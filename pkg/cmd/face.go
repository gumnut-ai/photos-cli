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

var facesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves details for a specific face.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "face-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library ID (required if user has multiple libraries)",
			QueryPath: "library_id",
		},
	},
	Action:          handleFacesRetrieve,
	HideHelpCommand: true,
}

var facesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the details of a specific face, currently only supporting\nassociating/disassociating with a person.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "face-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library ID (required if user has multiple libraries)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[any]{
			Name:     "person-id",
			BodyPath: "person_id",
		},
	},
	Action:          handleFacesUpdate,
	HideHelpCommand: true,
}

var facesList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of faces, optionally filtered by asset, person, or\nspecific face IDs, ordered by creation time, descending.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "asset-id",
			Usage:     "Filter by faces in a specific asset",
			QueryPath: "asset_id",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Filter by specific face IDs (max 100)",
			QueryPath: "ids",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library ID (required if user has multiple libraries)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Max number of faces to return (1-200)",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "person-id",
			Usage:     "Filter by faces associated with a specific person",
			QueryPath: "person_id",
		},
		&requestflag.Flag[any]{
			Name:      "starting-after-id",
			Usage:     "Face ID to start listing faces after",
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
	Usage:   "Deletes a specific face entry. This does not delete the associated asset or\nperson.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "face-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library ID (required if user has multiple libraries)",
			QueryPath: "library_id",
		},
	},
	Action:          handleFacesDelete,
	HideHelpCommand: true,
}

var facesDownloadThumbnail = cli.Command{
	Name:    "download-thumbnail",
	Usage:   "Retrieves a thumbnail for a specific face.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "face-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleFacesDownloadThumbnail,
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

	params := photos.FaceGetParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "faces retrieve", obj, format, transform)
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

	params := photos.FaceUpdateParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "faces update", obj, format, transform)
}

func handleFacesList(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.FaceListParams{}

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
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Faces.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "faces list", obj, format, transform)
	} else {
		iter := client.Faces.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "faces list", iter, format, transform, maxItems)
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

	params := photos.FaceDeleteParams{}

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

	return client.Faces.Delete(
		ctx,
		cmd.Value("face-id").(string),
		params,
		options...,
	)
}

func handleFacesDownloadThumbnail(ctx context.Context, cmd *cli.Command) error {
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

	response, err := client.Faces.DownloadThumbnail(ctx, cmd.Value("face-id").(string), options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}
