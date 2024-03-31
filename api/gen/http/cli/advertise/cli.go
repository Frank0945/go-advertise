// Code generated by goa v3.15.2, DO NOT EDIT.
//
// advertise HTTP client CLI support package
//
// Command:
// $ goa gen github.com/Frank0945/go-advertise/api/design -o api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	advertisec "github.com/Frank0945/go-advertise/api/gen/http/advertise/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `advertise (create|list)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` advertise create --body '{
      "Country": "TW",
      "age_end": 60,
      "age_start": 18,
      "end_at": "2024-10-01 00:00:00",
      "gender": "M",
      "platform": "ios",
      "start_at": "2024-01-01 00:00:00",
      "title": "AD 1"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		advertiseFlags = flag.NewFlagSet("advertise", flag.ContinueOnError)

		advertiseCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		advertiseCreateBodyFlag = advertiseCreateFlags.String("body", "REQUIRED", "")

		advertiseListFlags    = flag.NewFlagSet("list", flag.ExitOnError)
		advertiseListBodyFlag = advertiseListFlags.String("body", "REQUIRED", "")
	)
	advertiseFlags.Usage = advertiseUsage
	advertiseCreateFlags.Usage = advertiseCreateUsage
	advertiseListFlags.Usage = advertiseListUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "advertise":
			svcf = advertiseFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "advertise":
			switch epn {
			case "create":
				epf = advertiseCreateFlags

			case "list":
				epf = advertiseListFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "advertise":
			c := advertisec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = advertisec.BuildCreatePayload(*advertiseCreateBodyFlag)
			case "list":
				endpoint = c.List()
				data, err = advertisec.BuildListPayload(*advertiseListBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// advertiseUsage displays the usage of the advertise command and its
// subcommands.
func advertiseUsage() {
	fmt.Fprintf(os.Stderr, `Service is the advertise service interface.
Usage:
    %[1]s [globalflags] advertise COMMAND [flags]

COMMAND:
    create: Create a new edge
    list: List all ADs by filter

Additional help:
    %[1]s advertise COMMAND --help
`, os.Args[0])
}
func advertiseCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] advertise create -body JSON

Create a new edge
    -body JSON: 

Example:
    %[1]s advertise create --body '{
      "Country": "TW",
      "age_end": 60,
      "age_start": 18,
      "end_at": "2024-10-01 00:00:00",
      "gender": "M",
      "platform": "ios",
      "start_at": "2024-01-01 00:00:00",
      "title": "AD 1"
   }'
`, os.Args[0])
}

func advertiseListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] advertise list -body JSON

List all ADs by filter
    -body JSON: 

Example:
    %[1]s advertise list --body '{
      "Country": "TW",
      "age_end": 60,
      "age_start": 18,
      "gender": "M",
      "limit": 10,
      "offset": 0,
      "platform": "ios"
   }'
`, os.Args[0])
}