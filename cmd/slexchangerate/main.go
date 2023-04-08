package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"plugin"
	"text/tabwriter"

	"github.com/DiLRandI/sl-bank-exchangerate/app/result"
	"github.com/DiLRandI/sl-bank-exchangerate/common"
	"github.com/DiLRandI/sl-bank-exchangerate/config"
	"github.com/DiLRandI/sl-bank-exchangerate/contract"
	"github.com/rs/zerolog"
)

var Version = "development"

const (
	appName = "SL Bank Exchange Rate Monitor"
)

var currency = flag.String("currency", "USD", "currency to search ex USD, EUR")

func main() {
	logger := zerolog.New(os.Stdout).With().
		Str("app", appName).
		Str("version", Version).
		Logger()

	logger.Info().Msg("starting the app")

	flag.Parse()

	validateCurrency(currency, logger)

	config, err := config.ParseConfig("config.json")
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to parse config.json")
	}

	plugins := make(map[string]contract.Convert)

	for _, config := range config.Plugins {
		plugin, err := plugin.Open(config.File)
		if err != nil {
			logger.Fatal().Err(err).Msgf("unable to open the plugin %s", config.Name)
		}

		InitializeLookup, err := plugin.Lookup("Initialize")
		if err != nil {
			logger.Fatal().Err(err).Msg("Initialize function not found in plugin")
		}

		fn, ok := InitializeLookup.(contract.Initialize)
		if !ok {
			logger.Fatal().Msg("function not satisfy the contract")
		}

		plugins[config.Name] = fn(config.Endpoint)
	}

	output := make(map[string][]result.Result)

	for bankName, convert := range plugins {
		logger.Info().Msgf("executing %s", bankName)

		results := []result.Result{}

		// for _, code := range common.LookupKeys {
		value, err := convert(*currency)
		if err != nil {
			logger.Fatal().Err(err)
		}

		results = append(results, result.Result{
			BankName: bankName,
			Currency: *currency,
			Value:    value,
		})
		// }

		output[bankName] = results
	}

	logger.Info().Msg("printing output")

	printOutput(output)
}

func validateCurrency(currency *string, logger zerolog.Logger) {
	for _, key := range common.LookupKeys {
		if key == *currency {
			return
		}
	}

	logger.Fatal().Msgf("currency %q is not supported", *currency)
}

func printOutput(output map[string][]result.Result) {
	writer := tabwriter.NewWriter(os.Stdout, 20, 8, 1, ' ',
		tabwriter.Debug|tabwriter.AlignRight)

	printHeaders(writer)
	for bankName, rates := range output {
		for _, r := range rates {
			fmt.Fprintf(writer, "| %s \t| %s \t| %.2f \t|\n", bankName, r.Currency, float64(r.Value)/common.CentsFactor)
		}
	}
}

func printHeaders(write io.Writer) {
	printVerticalLine(write)
	fmt.Fprint(write, "| Bank Name \t| Currency \t| Exchange Rate (LKR) \t|\n")
	printVerticalLine(write)
}

func printVerticalLine(write io.Writer) {
	fmt.Fprint(write, "| ------------------ \t"+
		"| ------------------ \t"+
		"| ------------------ \t"+
		"|\n")
}
