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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
)

// MultiTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/aggregations/bucket.ts#L265-L274
type MultiTermsAggregation struct {
	CollectMode           *termsaggregationcollectmode.TermsAggregationCollectMode `json:"collect_mode,omitempty"`
	Meta                  map[string]interface{}                                   `json:"meta,omitempty"`
	MinDocCount           *int64                                                   `json:"min_doc_count,omitempty"`
	Name                  *string                                                  `json:"name,omitempty"`
	Order                 *AggregateOrder                                          `json:"order,omitempty"`
	ShardMinDocCount      *int64                                                   `json:"shard_min_doc_count,omitempty"`
	ShardSize             *int                                                     `json:"shard_size,omitempty"`
	ShowTermDocCountError *bool                                                    `json:"show_term_doc_count_error,omitempty"`
	Size                  *int                                                     `json:"size,omitempty"`
	Terms                 []MultiTermLookup                                        `json:"terms"`
}

// NewMultiTermsAggregation returns a MultiTermsAggregation.
func NewMultiTermsAggregation() *MultiTermsAggregation {
	r := &MultiTermsAggregation{}

	return r
}