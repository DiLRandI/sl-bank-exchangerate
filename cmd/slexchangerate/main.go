package main

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"plugin"
	"text/tabwriter"

	"github.com/DiLRandI/sl-bank-exchangerate/app/result"
	"github.com/DiLRandI/sl-bank-exchangerate/common"
	"github.com/DiLRandI/sl-bank-exchangerate/config"
	"github.com/DiLRandI/sl-bank-exchangerate/contract"
)

var Version = "development"

const (
	appName = "SL Bank Exchange Rate Monitor"
)

var currency = flag.String("currency", "USD", "currency to search ex USD, EUR")

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).
		With(appName, Version)

	logger.Info("starting the app")

	flag.Parse()

	validateCurrency(currency, logger)

	config, err := config.ParseConfig("config.json")
	if err != nil {
		logger.Error("unable to parse config.json", err)
	}

	plugins := make(map[string]contract.Convert)

	for _, config := range config.Plugins {
		plugin, err := plugin.Open(config.File)
		if err != nil {
			logger.Error("unable to open the plugin", config.Name)
			os.Exit(1)
		}

		InitializeLookup, err := plugin.Lookup("Initialize")
		if err != nil {
			logger.Error("Initialize function not found in plugin")
			os.Exit(1)
		}

		fn, ok := InitializeLookup.(contract.Initialize)
		if !ok {
			logger.Error("function not satisfy the contract")
			os.Exit(1)
		}

		plugins[config.Name] = fn(config.Endpoint)
	}

	output := make(map[string][]result.Result)

	for bankName, convert := range plugins {
		logger.Info("executing", bankName)

		results := []result.Result{}

		// for _, code := range common.LookupKeys {
		value, err := convert(*currency)
		if err != nil {
			logger.Error("unable to get the value", err)
		}

		results = append(results, result.Result{
			BankName: bankName,
			Currency: *currency,
			Value:    value,
		})
		// }

		output[bankName] = results
	}

	logger.Info("execution completed")

	printOutput(output)
}

func validateCurrency(currency *string, logger *slog.Logger) {
	for _, key := range common.LookupKeys {
		if key == *currency {
			return
		}
	}

	logger.Error("currency %q is not supported", *currency)
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
