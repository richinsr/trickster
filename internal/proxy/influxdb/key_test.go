/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package influxdb

import (
	"net/url"
	"testing"
	"time"

	"github.com/Comcast/trickster/internal/proxy"
	"github.com/Comcast/trickster/internal/timeseries"
)

func TestDeriveCacheKey(t *testing.T) {

	client := &Client{}
	tu := &url.URL{Path: "/", RawQuery: "db=test&u=test&p=test&q=select * where <$TIME_TOKEN$> group by time(1m)"}
	r := &proxy.Request{TemplateURL: tu, TimeRangeQuery: &timeseries.TimeRangeQuery{Step: time.Duration(60) * time.Second}}
	key := client.DeriveCacheKey(r, "extra")

	if key != "147fe576515e5ac107195a3eb69c8cbe" {
		t.Errorf("expected %s got %s", "147fe576515e5ac107195a3eb69c8cbe", key)
	}

}
