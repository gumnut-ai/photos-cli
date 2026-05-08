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

var librariesCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new, empty photo library for the authenticated user. A library is the\ntop-level container for assets, albums, people, and faces — most users have\nexactly one. Only create a new library when the user explicitly asks for a\nseparate container.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Display name for the new library. Required.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Optional free-form description shown alongside the library name.",
			BodyPath: "description",
		},
	},
	Action:          handleLibrariesCreate,
	HideHelpCommand: true,
}

var librariesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetches one library's metadata by ID (name, description, asset count). Use when\nyou already have a specific `library_id`; for enumerating a user's libraries\nprefer `list_libraries`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "library-id",
			Usage:     "Library ID (with `lib_` prefix) to fetch. Obtain from `list_libraries` or any response containing a library reference.",
			Required:  true,
			PathParam: "library_id",
		},
	},
	Action:          handleLibrariesRetrieve,
	HideHelpCommand: true,
}

var librariesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Renames a library or changes its description. Only the fields included in the\nrequest body are changed. Library contents (assets, albums, people, faces) are\nnot affected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "library-id",
			Usage:     "Library ID (with `lib_` prefix) of the library to update.",
			Required:  true,
			PathParam: "library_id",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "New free-form description for the library. Omit to leave unchanged.",
			BodyPath: "description",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "New display name for the library. Omit to leave unchanged.",
			BodyPath: "name",
		},
	},
	Action:          handleLibrariesUpdate,
	HideHelpCommand: true,
}

var librariesList = cli.Command{
	Name:            "list",
	Usage:           "Returns every library owned by the authenticated user (no pagination — users\ntypically have one or a handful). Call this when another tool's `library_id`\nparameter is required but you don't yet know which libraries exist. A\nsingle-library user can usually omit `library_id` on other tools entirely.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleLibrariesList,
	HideHelpCommand: true,
}

var librariesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the library and all its contents — assets (including their stored\nfiles), albums, people, and faces. **Destructive and irreversible** — should be\nused only when the user explicitly confirms they want to destroy an entire\nlibrary.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "library-id",
			Usage:     "Library ID (with `lib_` prefix) of the library to delete.",
			Required:  true,
			PathParam: "library_id",
		},
	},
	Action:          handleLibrariesDelete,
	HideHelpCommand: true,
}

func handleLibrariesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := photos.LibraryNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Libraries.New(ctx, params, options...)
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
		Title:          "libraries create",
		Transform:      transform,
	})
}

func handleLibrariesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("library-id") && len(unusedArgs) > 0 {
		cmd.Set("library-id", unusedArgs[0])
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
	_, err = client.Libraries.Get(ctx, cmd.Value("library-id").(string), options...)
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
		Title:          "libraries retrieve",
		Transform:      transform,
	})
}

func handleLibrariesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("library-id") && len(unusedArgs) > 0 {
		cmd.Set("library-id", unusedArgs[0])
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

	params := photos.LibraryUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Libraries.Update(
		ctx,
		cmd.Value("library-id").(string),
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
		Title:          "libraries update",
		Transform:      transform,
	})
}

func handleLibrariesList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Libraries.List(ctx, options...)
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
		Title:          "libraries list",
		Transform:      transform,
	})
}

func handleLibrariesDelete(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("library-id") && len(unusedArgs) > 0 {
		cmd.Set("library-id", unusedArgs[0])
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

	return client.Libraries.Delete(ctx, cmd.Value("library-id").(string), options...)
}
