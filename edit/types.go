/*
Copyright 2016 Google Inc. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/
// Type information for attributes.

package edit

import (
	blaze "github.com/bazelbuild/buildifier/build_proto"
	"github.com/bazelbuild/buildifier/lang"
)

var typeOf = lang.TypeOf

// IsList returns true for all attributes whose type is a list.
func IsList(attr string) bool {
	ty := typeOf[attr]
	return ty == blaze.Attribute_STRING_LIST ||
		ty == blaze.Attribute_LABEL_LIST ||
		ty == blaze.Attribute_OUTPUT_LIST ||
		ty == blaze.Attribute_FILESET_ENTRY_LIST ||
		ty == blaze.Attribute_INTEGER_LIST ||
		ty == blaze.Attribute_LICENSE ||
		ty == blaze.Attribute_DISTRIBUTION_SET
}

// IsString returns true for all attributes whose type is an int list.
func IsIntList(attr string) bool {
	return typeOf[attr] == blaze.Attribute_INTEGER_LIST
}

// IsString returns true for all attributes whose type is a string or a label.
func IsString(attr string) bool {
	ty := typeOf[attr]
	return ty == blaze.Attribute_LABEL ||
		ty == blaze.Attribute_STRING ||
		ty == blaze.Attribute_OUTPUT
}

// IsString returns true for all attributes whose type is a string dictionary.
func IsStringDict(attr string) bool {
	return typeOf[attr] == blaze.Attribute_STRING_DICT
}

// ContainsLabels returns true for all attributes whose type is a label or a label list.
func ContainsLabels(attr string) bool {
	ty := typeOf[attr]
	return ty == blaze.Attribute_LABEL_LIST ||
		ty == blaze.Attribute_LABEL
}
