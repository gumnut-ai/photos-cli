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

var oauthAuthURL = cli.Command{
	Name:    "auth-url",
	Usage:   "Generate OAuth authorization URL with state and nonce for CSRF and replay attack\nprotection. State is stored with TTL for validation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "redirect-uri",
			Usage:     "The URI to redirect to after OAuth consent. Must match the registered redirect URI in OAuth client configuration.",
			Required:  true,
			QueryPath: "redirect_uri",
		},
		&requestflag.Flag[any]{
			Name:      "code-challenge",
			Usage:     "PKCE code challenge derived from code_verifier. Required for public clients to prevent authorization code interception attacks.",
			QueryPath: "code_challenge",
		},
		&requestflag.Flag[any]{
			Name:      "code-challenge-method",
			Usage:     "PKCE code challenge method, typically 'S256' (SHA-256 hash). Must be provided if code_challenge is specified.",
			QueryPath: "code_challenge_method",
		},
	},
	Action:          handleOAuthAuthURL,
	HideHelpCommand: true,
}

var oauthExchange = cli.Command{
	Name:    "exchange",
	Usage:   "Exchange OAuth authorization code for application JWT after validating state,\nnonce, and ID token signature. User is retrieved from or created in the database\nand details added to the JWT.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "code",
			Usage:    "Authorization code returned by the OAuth provider after user consent",
			BodyPath: "code",
		},
		&requestflag.Flag[any]{
			Name:     "code-verifier",
			Usage:    "PKCE code verifier that corresponds to the code_challenge sent in the authorization request",
			BodyPath: "code_verifier",
		},
		&requestflag.Flag[any]{
			Name:     "error",
			Usage:    "Error code if OAuth provider returned an error instead of authorization code",
			BodyPath: "error",
		},
		&requestflag.Flag[any]{
			Name:     "state",
			Usage:    "State token from the initial auth request, used for CSRF protection",
			BodyPath: "state",
		},
	},
	Action:          handleOAuthExchange,
	HideHelpCommand: true,
}

var oauthLogoutEndpoint = cli.Command{
	Name:            "logout-endpoint",
	Usage:           "Returns the OAuth provider's logout endpoint URL from OIDC discovery. This can\nbe used to redirect users to logout from the OAuth provider after logging out\nlocally.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleOAuthLogoutEndpoint,
	HideHelpCommand: true,
}

func handleOAuthAuthURL(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.OAuthAuthURLParams{}

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
	_, err = client.OAuth.AuthURL(ctx, params, options...)
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
		Title:          "oauth auth-url",
		Transform:      transform,
	})
}

func handleOAuthExchange(ctx context.Context, cmd *cli.Command) error {
	client := photos.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := photos.OAuthExchangeParams{}

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
	_, err = client.OAuth.Exchange(ctx, params, options...)
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
		Title:          "oauth exchange",
		Transform:      transform,
	})
}

func handleOAuthLogoutEndpoint(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.OAuth.LogoutEndpoint(ctx, options...)
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
		Title:          "oauth logout-endpoint",
		Transform:      transform,
	})
}
