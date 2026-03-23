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

var peopleCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new person entry.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "birth-date",
			BodyPath: "birth_date",
		},
		&requestflag.Flag[any]{
			Name:     "is-favorite",
			Default:  false,
			BodyPath: "is_favorite",
		},
		&requestflag.Flag[any]{
			Name:     "is-hidden",
			Default:  false,
			BodyPath: "is_hidden",
		},
		&requestflag.Flag[any]{
			Name:     "library-id",
			BodyPath: "library_id",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "thumbnail-face-id",
			BodyPath: "thumbnail_face_id",
		},
	},
	Action:          handlePeopleCreate,
	HideHelpCommand: true,
}

var peopleRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves details for a specific person.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "person-id",
			Required: true,
		},
	},
	Action:          handlePeopleRetrieve,
	HideHelpCommand: true,
}

var peopleUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the details of a specific person.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "person-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "birth-date",
			BodyPath: "birth_date",
		},
		&requestflag.Flag[any]{
			Name:     "is-favorite",
			BodyPath: "is_favorite",
		},
		&requestflag.Flag[any]{
			Name:     "is-hidden",
			BodyPath: "is_hidden",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "thumbnail-face-id",
			BodyPath: "thumbnail_face_id",
		},
	},
	Action:          handlePeopleUpdate,
	HideHelpCommand: true,
}

var peopleList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of people, ordered by creation time, descending. Can\nbe filtered by specific person IDs, name, or whether the person has been named.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "album-id",
			Usage:     "Include only people associated with this album ID",
			QueryPath: "album_id",
		},
		&requestflag.Flag[any]{
			Name:      "asset-id",
			Usage:     "Include only people associated with this asset ID",
			QueryPath: "asset_id",
		},
		&requestflag.Flag[any]{
			Name:      "has-name",
			Usage:     "Filter by whether the person has a name assigned (true = named only, false = unnamed only)",
			QueryPath: "has_name",
		},
		&requestflag.Flag[any]{
			Name:      "id",
			Usage:     "Filter by specific person IDs (max 100)",
			QueryPath: "ids",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library ID (required if user has multiple libraries)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Max number of people to return (1-200)",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "name",
			Usage:     "Filter by name using case-insensitive substring matching",
			QueryPath: "name",
		},
		&requestflag.Flag[any]{
			Name:      "starting-after-id",
			Usage:     "Person ID to start listing people after",
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
	Usage:   "Deletes a specific person. Associated faces will have their person_id set to the\nclosest matching person, or null if no one matches.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "person-id",
			Required: true,
		},
	},
	Action:          handlePeopleDelete,
	HideHelpCommand: true,
}

func handlePeopleCreate(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.PersonNewParams{}

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
	_, err = client.People.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "people create", obj, format, transform)
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.People.Get(ctx, cmd.Value("person-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "people retrieve", obj, format, transform)
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

	params := photos.PersonUpdateParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "people update", obj, format, transform)
}

func handlePeopleList(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.PersonListParams{}

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
		_, err = client.People.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "people list", obj, format, transform)
	} else {
		iter := client.People.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "people list", iter, format, transform, maxItems)
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
