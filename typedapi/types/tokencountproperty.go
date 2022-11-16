// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4ca0cc05d3ae3fa06c2cd7be91905b656a474334


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// TokenCountProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/_types/mapping/specialized.ts#L78-L85
type TokenCountProperty struct {
	Analyzer                 *string                        `json:"analyzer,omitempty"`
	Boost                    *float64                       `json:"boost,omitempty"`
	CopyTo                   []string                       `json:"copy_to,omitempty"`
	DocValues                *bool                          `json:"doc_values,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EnablePositionIncrements *bool                          `json:"enable_position_increments,omitempty"`
	Fields                   map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`
	// Meta Metadata about the field.
	Meta       map[string]string   `json:"meta,omitempty"`
	NullValue  *float64            `json:"null_value,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
	Similarity *string             `json:"similarity,omitempty"`
	Store      *bool               `json:"store,omitempty"`
	Type       string              `json:"type,omitempty"`
}

// NewTokenCountProperty returns a TokenCountProperty.
func NewTokenCountProperty() *TokenCountProperty {
	r := &TokenCountProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	r.Type = "token_count"

	return r
}
