/*******************************************************************************
 * Copyright 2017 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	"encoding/json"
)

type Service struct {
	DescribedObject
	Id             string         `json:"id"`
	Name           string         `json:"name"`           // time in milliseconds that the device last provided any feedback or responded to any request
	LastConnected  int64          `json:"lastConnected"`  // time in milliseconds that the device last reported data to the core
	LastReported   int64          `json:"lastReported"`   // operational state - either enabled or disabled
	OperatingState OperatingState `json:"operatingState"` // operational state - ether enabled or disableddc
	Labels         []string       `json:"labels"`         // tags or other labels applied to the device service for search or other identification needs
	Addressable    Addressable    `json:"addressable"`    // address (MQTT topic, HTTP address, serial bus, etc.) for reaching the service
}

// Custom Marshaling to make empty strings null
func (s Service) MarshalJSON() ([]byte, error) {
	test := struct {
		DescribedObject
		Id             *string        `json:"id"`
		Name           *string        `json:"name"`           // time in milliseconds that the device last provided any feedback or responded to any request
		LastConnected  int64          `json:"lastConnected"`  // time in milliseconds that the device last reported data to the core
		LastReported   int64          `json:"lastReported"`   // operational state - either enabled or disabled
		OperatingState OperatingState `json:"operatingState"` // operational state - ether enabled or disableddc
		Labels         []string       `json:"labels"`         // tags or other labels applied to the device service for search or other identification needs
		Addressable    Addressable    `json:"addressable"`    // address (MQTT topic, HTTP address, serial bus, etc.) for reaching the service
	}{
		DescribedObject: s.DescribedObject,
		LastConnected:   s.LastConnected,
		LastReported:    s.LastReported,
		OperatingState:  s.OperatingState,
		Labels:          s.Labels,
		Addressable:     s.Addressable,
	}

	// Empty strings are null
	if s.Name != "" {
		test.Name = &s.Name
	}
	if s.Id != "" {
		test.Id = &s.Id
	}

	return json.Marshal(test)
}

/*
 * To String function for Service
 */
func (dp Service) String() string {
	out, err := json.Marshal(dp)
	if err != nil {
		return err.Error()
	}

	return string(out)
}
