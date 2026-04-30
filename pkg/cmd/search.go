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

var searchSearch = cli.Command{
	Name:    "search",
	Usage:   "Searches for assets using semantic (CLIP-based) image-content matching and/or\nstructured filters. Use this tool when the user describes _what's in_ the photos\nthey want — subjects, scenes, places, activities, moods, objects — as opposed to\nbrowsing by album membership or exact ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "captured-after",
			Usage:     "Only include assets captured strictly after this instant (ISO 8601; exclusive). Equivalent in purpose to `local_datetime_after` on `list_assets` (naming inconsistency is tracked as a follow-up).",
			QueryPath: "captured_after",
		},
		&requestflag.Flag[any]{
			Name:      "captured-before",
			Usage:     "Only include assets captured strictly before this instant (ISO 8601; exclusive). Equivalent in purpose to `local_datetime_before` on `list_assets` (naming inconsistency is tracked as a follow-up).",
			QueryPath: "captured_before",
		},
		&requestflag.Flag[*string]{
			Name:      "library-id",
			Usage:     "Library to search. Optional if the user has a single library; required when they have multiple. Use `list_libraries` to enumerate available libraries.",
			QueryPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results per page (1–200). Defaults to 20.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "1-indexed page number. `search_assets` uses page-number pagination; the sibling `list_assets` uses cursor pagination via `starting_after_id`. Increment `page` to fetch subsequent pages.",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[any]{
			Name:      "person-id",
			Usage:     "Filter to assets containing ALL of these person IDs (intersection, not union). Accepts multiple `person_ids=` query params or a single comma-delimited value (e.g., `person_123,person_abc`). Get person IDs from `list_people`. Plural on this tool; the sibling `list_assets` uses `person_id` (singular).",
			QueryPath: "person_ids",
		},
		&requestflag.Flag[*string]{
			Name:      "query",
			Usage:     "Natural-language description of the image content to search for. Matched against CLIP image embeddings, so it works best with concrete visual concepts: subjects, scenes, objects, settings ('beach sunset', 'birthday cake', 'mountain hike').\n\nPrefer structured params when available: use `person_ids` for people (not names in `query`) and `captured_before`/`captured_after` for dates (not phrases like 'in 2023' in `query`).",
			QueryPath: "query",
		},
		&requestflag.Flag[float64]{
			Name:      "threshold",
			Usage:     "Maximum semantic distance for a result to be included (0.0 = identical, 1.0 = unrelated). Lower values return fewer, more confident matches; higher values return more results with looser matching. Default 0.8 is moderate — try 0.6 for high-precision queries, 0.9 for exploratory searches. **Note:** this is inverted from the usual 'similarity score' convention where higher means more similar.",
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
		&requestflag.Flag[*string]{
			Name:      "image",
			Usage:     "Image file to search for similar assets. Can be combined with text query.",
			BodyPath:  "image",
			FileInput: true,
		},
		&requestflag.Flag[*string]{
			Name:     "library-id",
			Usage:    "Library to search assets from (optional)",
			BodyPath: "library_id",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Number of results per page (1-200)",
			Default:  20,
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
		&requestflag.Flag[*string]{
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "search search",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "search search-assets",
		Transform:      transform,
	})
}
