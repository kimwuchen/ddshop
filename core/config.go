// Copyright © 2022 zc2638 <zc2638@qq.com>.
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

package core

type Config struct {
	Cookie   string       `json:"cookie"`
	Interval int64        `json:"interval"`
	PayType  string       `json:"payType"`
	BarkKey  string       `json:"barkKey"`
	Periods  []TimePeriod `json:"periods"`
}

type TimePeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`

	startHour   int
	startMinute int
	endHour     int
	endMinute   int
}
