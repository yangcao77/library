//
// Copyright 2022 Red Hat, Inc.
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
	"fmt"
	"os"
	"strings"

	devfilepkg "github.com/devfile/library/v2/pkg/devfile"
	"github.com/devfile/library/v2/pkg/devfile/parser"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "updateSchema" {
		ReplaceSchemaFile()
	} else {
		parserTest()
	}
}

func parserTest() {
	var args parser.ParserArgs
	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "http") {
			args = parser.ParserArgs{
				URL: os.Args[1],
			}
		} else {
			args = parser.ParserArgs{
				Path: os.Args[1],
			}
		}
		fmt.Println("parsing devfile from " + os.Args[1])

	} else {
		args = parser.ParserArgs{
			Path: "devfile.yaml",
		}
		fmt.Println("parsing devfile from ./devfile.yaml")
	}
	_, warning, err := devfilepkg.ParseDevfileAndValidate(args)
	if err != nil {
		fmt.Println(err)
	} else {
		if len(warning.Commands) > 0 || len(warning.Components) > 0 || len(warning.Projects) > 0 || len(warning.StarterProjects) > 0 {
			fmt.Printf("top-level variables were not substituted successfully %+v\n", warning)
		}
	}

}
