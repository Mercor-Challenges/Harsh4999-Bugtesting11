/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package util

import (
	"fmt"

	g "github.com/onsi/ginkgo"
)

//DumpAndReturnTagging takes and array of tags and obtains the hex image IDs, dumps them to ginkgo for printing, and then returns them
func DumpAndReturnTagging(tags []string) ([]string, error) {
	hexIDs, err := GetImageIDForTags(tags)
	if err != nil {
		return nil, err
	}
	for i, hexID := range hexIDs {
		fmt.Fprintf(g.GinkgoWriter, "tag %s hex id %s ", tags[i], hexID)
	}
	return hexIDs, nil
}

//CreateResource creates the resources from the supplied json file (not a template); ginkgo error checking included
func CreateResource(jsonFilePath string, oc *CLI) error {
	err := oc.Run("create").Args("-f", jsonFilePath).Execute()
	return err
}
