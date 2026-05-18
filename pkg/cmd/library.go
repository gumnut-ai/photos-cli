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
	Usage:   "Fetches one library's metadata by ID. Returns the library regardless of trash\nstate.",
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
	Name:    "list",
	Usage:   "Returns libraries owned by the authenticated user (no pagination — users\ntypically have one or a handful). Call this when another tool's `library_id`\nparameter is required but you don't yet know which libraries exist. A\nsingle-library user can usually omit `library_id` on other tools entirely.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Which set of libraries to return: `live` (default — excludes trashed), `trashed` (only trashed, ordered by most recently trashed), or `all` (both).",
			Default:   "live",
			QueryPath: "state",
		},
	},
	Action:          handleLibrariesList,
	HideHelpCommand: true,
}

var librariesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Expedites the background purge on a **trashed** library: the 90-day undo window\nis waived and the drain begins claiming this library on the next scheduled tick.\nReturns 204 immediately; the drain proceeds asynchronously in bounded batches\nand does not block on completion. Restore still works until the drain finishes\npurging all assets, but past this point it will recover only the assets the\ndrain hasn't gotten to yet. Returns 409 if the library has not been trashed yet;\ntrash it first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "library-id",
			Usage:     "Library ID (with `lib_` prefix) of the trashed library to expedite.",
			Required:  true,
			PathParam: "library_id",
		},
	},
	Action:          handleLibrariesDelete,
	HideHelpCommand: true,
}

var librariesRestore = cli.Command{
	Name:    "restore",
	Usage:   "Restores a previously-trashed library so it reappears in default list/search\nresults. Works as long as the library row still exists — once `get_library`\nreturns 404 the row is gone and restore is no longer possible. If the background\ndrain has already started purging assets, restore succeeds but recovers only the\nassets the drain hasn't gotten to yet.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "library-id",
			Usage:     "Library ID (with `lib_` prefix) of the trashed library to restore.",
			Required:  true,
			PathParam: "library_id",
		},
	},
	Action:          handleLibrariesRestore,
	HideHelpCommand: true,
}

var librariesTrash = cli.Command{
	Name:    "trash",
	Usage:   "Moves the library and all its contents into the trash. The library becomes\ninaccessible by default and can be fully restored within 90 days by calling\n`restore_library`. After 90 days the library's assets are gradually purged in\nthe background; until the library row itself is removed, restore still works but\nrecovers only the assets not yet purged.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "library-id",
			Usage:     "Library ID (with `lib_` prefix) of the library to trash.",
			Required:  true,
			PathParam: "library_id",
		},
	},
	Action:          handleLibrariesTrash,
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

	params := photos.LibraryListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Libraries.List(ctx, params, options...)
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

func handleLibrariesRestore(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Libraries.Restore(ctx, cmd.Value("library-id").(string), options...)
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
		Title:          "libraries restore",
		Transform:      transform,
	})
}

func handleLibrariesTrash(ctx context.Context, cmd *cli.Command) error {
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

	return client.Libraries.Trash(ctx, cmd.Value("library-id").(string), options...)
}
