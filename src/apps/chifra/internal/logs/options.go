// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * This file was auto generated with makeClass --gocmds. DO NOT EDIT.
 */

package logsPkg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/globals"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

type LogsOptions struct {
	Transactions []string
	Articulate   bool
	Globals      globals.GlobalOptions
	BadFlag      error
}

func (opts *LogsOptions) TestLog() {
	logger.TestLog(len(opts.Transactions) > 0, "Transactions: ", opts.Transactions)
	logger.TestLog(opts.Articulate, "Articulate: ", opts.Articulate)
	opts.Globals.TestLog()
}

func (opts *LogsOptions) ToCmdLine() string {
	options := ""
	if opts.Articulate {
		options += " --articulate"
	}
	options += " " + strings.Join(opts.Transactions, " ")
	options += fmt.Sprintf("%s", "") // silence go compiler for auto gen
	return options
}

func FromRequest(w http.ResponseWriter, r *http.Request) *LogsOptions {
	opts := &LogsOptions{}
	for key, value := range r.URL.Query() {
		switch key {
		case "transactions":
			for _, val := range value {
				s := strings.Split(val, " ") // may contain space separated items
				opts.Transactions = append(opts.Transactions, s...)
			}
		case "articulate":
			opts.Articulate = true
		default:
			if !globals.IsGlobalOption(key) {
				opts.BadFlag = validate.Usage("Invalid key ({0}) in {1} route.", key, "logs")
				return opts
			}
		}
	}
	opts.Globals = *globals.FromRequest(w, r)

	return opts
}

var Options LogsOptions

func LogsFinishParse(args []string) *LogsOptions {
	// EXISTING_CODE
	Options.Transactions = args
	// EXISTING_CODE
	return &Options
}
