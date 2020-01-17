/**
 * Copyright 2019-2020 Wargaming Group Limited
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
**/

package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ADD_REMOTE_SCHEMA_TYPE string = `add_remote_schema`
)

type AddRemoteSchema struct {
	Arguments AddRemoteSchemaArgs `json:"args"`
	Ver       int                 `json:"version"`
	QueryType string              `json:"type"`
}

type AddRemoteSchemaArgs struct {
	Name       string                        `json:"name"`
	Definition AddRemoteSchemaArgsDefinition `json:"definition"`
	Comment    string                        `json:"comment,omitempty"`
}

type AddRemoteSchemaArgsDefinition struct {
	URL                  string              `json:"url"`
	Headers              []map[string]string `json:"headers"`
	ForwardClientHeaders bool                `json:"forward_client_headers"`
	TimeoutSecs          int                 `json:"timeout_seconds"`
}

type AddRemoteSchemaResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *AddRemoteSchema) SetArgs(args interface{}) error {
	switch args.(type) {
	case AddRemoteSchemaArgs:
		t.Arguments = args.(AddRemoteSchemaArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *AddRemoteSchema) SetVersion(version int) {
	t.Ver = version
}

func (t *AddRemoteSchema) SetType(name string) {
	t.QueryType = name
}

func (t *AddRemoteSchema) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *AddRemoteSchema) Method() string {
	return http.MethodPost
}

func (t *AddRemoteSchema) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var addRemoteSchemaResponse AddRemoteSchemaResponse
	if err := json.Unmarshal(body, &addRemoteSchemaResponse); err != nil {
		return nil, err
	}
	return addRemoteSchemaResponse, nil
}

func NewAddRemoteSchemaQuery() Query {
	query := AddRemoteSchema{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: ADD_REMOTE_SCHEMA_TYPE,
	}

	return Query(&query)
}
