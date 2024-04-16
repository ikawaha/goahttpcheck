// Code generated by goa v3.16.1, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen calc/design

package cli

import (
	calcc "calc/gen/http/calc/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `calc (multiply|divide|redirect)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` calc multiply --a 3793862871819669726 --b 8399553735696626949` + "\n" +
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
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcMultiplyFlags = flag.NewFlagSet("multiply", flag.ExitOnError)
		calcMultiplyAFlag = calcMultiplyFlags.String("a", "REQUIRED", "Left operand")
		calcMultiplyBFlag = calcMultiplyFlags.String("b", "REQUIRED", "Right operand")

		calcDivideFlags = flag.NewFlagSet("divide", flag.ExitOnError)
		calcDivideAFlag = calcDivideFlags.String("a", "REQUIRED", "Left operand")
		calcDivideBFlag = calcDivideFlags.String("b", "REQUIRED", "Right operand")

		calcRedirectFlags = flag.NewFlagSet("redirect", flag.ExitOnError)
	)
	calcFlags.Usage = calcUsage
	calcMultiplyFlags.Usage = calcMultiplyUsage
	calcDivideFlags.Usage = calcDivideUsage
	calcRedirectFlags.Usage = calcRedirectUsage

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
		case "calc":
			svcf = calcFlags
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
		case "calc":
			switch epn {
			case "multiply":
				epf = calcMultiplyFlags

			case "divide":
				epf = calcDivideFlags

			case "redirect":
				epf = calcRedirectFlags

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
		case "calc":
			c := calcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "multiply":
				endpoint = c.Multiply()
				data, err = calcc.BuildMultiplyPayload(*calcMultiplyAFlag, *calcMultiplyBFlag)
			case "divide":
				endpoint = c.Divide()
				data, err = calcc.BuildDividePayload(*calcDivideAFlag, *calcDivideBFlag)
			case "redirect":
				endpoint = c.Redirect()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers.
Usage:
    %[1]s [globalflags] calc COMMAND [flags]

COMMAND:
    multiply: Multiply implements multiply.
    divide: Divide implements divide.
    redirect: Redirect implements redirect.

Additional help:
    %[1]s calc COMMAND --help
`, os.Args[0])
}
func calcMultiplyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc multiply -a INT -b INT

Multiply implements multiply.
    -a INT: Left operand
    -b INT: Right operand

Example:
    %[1]s calc multiply --a 3793862871819669726 --b 8399553735696626949
`, os.Args[0])
}

func calcDivideUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc divide -a INT -b INT

Divide implements divide.
    -a INT: Left operand
    -b INT: Right operand

Example:
    %[1]s calc divide --a 5401762099778430809 --b 1918630006328122782
`, os.Args[0])
}

func calcRedirectUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc redirect

Redirect implements redirect.

Example:
    %[1]s calc redirect
`, os.Args[0])
}
