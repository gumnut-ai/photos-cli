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

var searchSearch = cli.Command{
	Name:    "search",
	Usage:   "Searches for assets using semantic similarity and/or metadata filters. Results\ninclude asset metadata, faces, and people. At least one search criterion must be\nprovided.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "captured-after",
			Usage:     "Filter to only include assets captured after this date (ISO format).",
			QueryPath: "captured_after",
		},
		&requestflag.Flag[any]{
			Name:      "captured-before",
			Usage:     "Filter to only include assets captured before this date (ISO format).",
			QueryPath: "captured_before",
		},
		&requestflag.Flag[any]{
			Name:      "library-id",
			Usage:     "Library to search assets from (optional)",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Number of results per page (1-200)",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[any]{
			Name:      "person-id",
			Usage:     "Filter to only include assets containing ALL of these person IDs. Can be comma-delimited string (e.g. 'person_123,person_abc') or multiple query parameters.",
			QueryPath: "person_ids",
		},
		&requestflag.Flag[any]{
			Name:      "query",
			Usage:     "The text query to search for. If you want to search for a specific person or set of people, use the person_ids parameter instead.If you want to search for a photos taken during a specific date range, use the captured_before and captured_after parameters instead.",
			QueryPath: "query",
		},
		&requestflag.Flag[float64]{
			Name:      "threshold",
			Usage:     "Similarity threshold (lower means more similar)",
			Default:   0.8,
			QueryPath: "threshold",
		},
	},
	Action:          handleSearchSearch,
	HideHelpCommand: true,
}

var searchSearchAssets = cli.Command{
	Name:    "search-assets",
	Usage:   "Searches for assets using semantic similarity and/or metadata filters. Results\ninclude asset metadata, faces, and people. At least one search criterion must be\nprovided. Can search by text query, uploaded image, or both combined.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "captured-after",
			Usage:    "Filter to only include assets captured after this date (ISO format).",
			BodyPath: "captured_after",
		},
		&requestflag.Flag[any]{
			Name:     "captured-before",
			Usage:    "Filter to only include assets captured before this date (ISO format).",
			BodyPath: "captured_before",
		},
		&requestflag.Flag[any]{
			Name:     "image",
			Usage:    "Image file to search for similar assets. Can be combined with text query.",
			BodyPath: "image",
		},
		&requestflag.Flag[any]{
			Name:     "library-id",
			Usage:    "Library to search assets from (optional)",
			BodyPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Number of results per page (1-200)",
			Default:  50,
			BodyPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:     "page",
			Usage:    "Page number",
			Default:  1,
			BodyPath: "page",
		},
		&requestflag.Flag[any]{
			Name:     "person-id",
			Usage:    "Filter to only include assets containing ALL of these person IDs. Can be comma-delimited string (e.g. 'person_123,person_abc') or multiple query parameters.",
			BodyPath: "person_ids",
		},
		&requestflag.Flag[any]{
			Name:     "query",
			Usage:    "The text query to search for. If you want to search for a specific person or set of people, use the person_ids parameter instead.If you want to search for a photos taken during a specific date range, use the captured_before and captured_after parameters instead.",
			BodyPath: "query",
		},
		&requestflag.Flag[float64]{
			Name:     "threshold",
			Usage:    "Similarity threshold (lower means more similar)",
			Default:  0.8,
			BodyPath: "threshold",
		},
	},
	Action:          handleSearchSearchAssets,
	HideHelpCommand: true,
}

func handleSearchSearch(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.SearchSearchParams{}

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
	_, err = client.Search.Search(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search search", obj, format, transform)
}

func handleSearchSearchAssets(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.SearchSearchAssetsParams{}

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
	_, err = client.Search.SearchAssets(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search search-assets", obj, format, transform)
}
