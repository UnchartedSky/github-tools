// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"

	"github.com/UnchartedSky/github-tools/cmd"
	"github.com/getsentry/raven-go"
)

func main() {
	cmd.Execute()
}

func init() {
	// Remove timestamp from logging records
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Sentry.io
	dsn := os.Getenv("SENTRY_DSN")
	if len(dsn) > 0 {
		err := raven.SetDSN(dsn)
		if err != nil {
			log.Fatal(err)
		}
	}

	raven.CapturePanic(func() {
		// do all of the scary things here
	}, nil)
}
