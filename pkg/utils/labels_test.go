/*
Copyright 2021 The Clusternet Authors.

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

package utils

import (
	"encoding/json"
	"testing"

	utilpointer "k8s.io/utils/pointer"
)

func TestMarshalLabelOption(t *testing.T) {
	tests := []struct {
		name        string
		labelOption LabelOption
		want        string
	}{
		{
			name: "add/update label",
			labelOption: LabelOption{Meta: Meta{Labels: map[string]*string{
				"key": utilpointer.StringPtr("val"),
			}}},
			want: `{"metadata":{"labels":{"key":"val"}}}`,
		},
		{
			name: "remove label",
			labelOption: LabelOption{Meta: Meta{Labels: map[string]*string{
				"key": nil,
			}}},
			want: `{"metadata":{"labels":{"key":null}}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.labelOption)
			if err != nil {
				t.Errorf("failed to marshal LabelOption: error = %v", err)
				return
			}
			if string(got) != tt.want {
				t.Errorf("got: %s, want: %s", string(got), tt.want)
			}
		})
	}
}
