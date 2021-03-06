/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package types

import (
	"net/http"
	"time"
)

// DockerID is an ID of docker container. It is a type to make it clear when we're working with docker container Ids
type DockerID string

type HttpGetter interface {
	Get(url string) (*http.Response, error)
}

// Timestamp wraps around time.Time and offers utilities to format and parse
// the time using RFC3339Nano
type Timestamp struct {
	time time.Time
}

// NewTimestamp returns a Timestamp object using the current time.
func NewTimestamp() *Timestamp {
	return &Timestamp{time.Now()}
}

// ConvertToTimestamp takes a string, parses it using the RFC3339Nano layout,
// and converts it to a Timestamp object.
func ConvertToTimestamp(timeString string) *Timestamp {
	parsed, _ := time.Parse(time.RFC3339Nano, timeString)
	return &Timestamp{parsed}
}

// Get returns the time as time.Time.
func (t *Timestamp) Get() time.Time {
	return t.time
}

// GetString returns the time in the string format using the RFC3339Nano
// layout.
func (t *Timestamp) GetString() string {
	return t.time.Format(time.RFC3339Nano)
}
