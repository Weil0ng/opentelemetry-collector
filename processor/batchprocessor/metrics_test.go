// Copyright The OpenTelemetry Authors
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

package batchprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/internal/collector/telemetry"
)

func TestBatchProcessorMetrics(t *testing.T) {
	tests := []struct {
		viewNames []string
		level     telemetry.Level
	}{
		{
			viewNames: []string{"batch_size_trigger_send", "timeout_trigger_send", "batch_send_size", "batch_send_size_bytes"},
			level:     telemetry.Detailed,
		},
		{
			viewNames: []string{},
			level:     telemetry.None,
		},
	}
	for _, test := range tests {
		views := MetricViews(test.level)
		if test.viewNames == nil {
			assert.Nil(t, views)
			continue
		}
		for i, viewName := range test.viewNames {
			assert.Equal(t, viewName, views[i].Name)
		}
	}
}
