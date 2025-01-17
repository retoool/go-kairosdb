// Copyright 2016 Ajit Yagaty
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

package aggregator

import (
	"testing"

	"github.com/retoool/go-kairosdb/builder/utils"
	"github.com/stretchr/testify/assert"
)

// Success test.
func TestRateAggregator(t *testing.T) {
	ra := NewRateAggregator(utils.MINUTES)
	assert.Equal(t, "rate", ra.Name(), "Rate aggregator name field must be set to 'rate'")
	assert.EqualValues(t, utils.MINUTES, ra.Unit(), "Rate aggregator unit must be set minutes")
}
