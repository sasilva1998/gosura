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
	SET_TABLE_IS_ENUM_TYPE string = `set_table_is_enum`
)

type SetTableIsEnum struct {
	Arguments SetTableIsEnumArgs `json:"args"`
	Ver       int                `json:"version"`
	QueryType string             `json:"type"`
}

type SetTableIsEnumArgs struct {
	Table  TableArgs `json:"table"`
	IsEnum bool      `json:"is_enum"`
}

type SetTableIsEnumResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *SetTableIsEnum) SetArgs(args interface{}) error {
	switch args.(type) {
	case SetTableIsEnumArgs:
		t.Arguments = args.(SetTableIsEnumArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *SetTableIsEnum) SetVersion(version int) {
	t.Ver = version
}

func (t *SetTableIsEnum) SetType(name string) {
	t.QueryType = name
}

func (t *SetTableIsEnum) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *SetTableIsEnum) Method() string {
	return http.MethodPost
}

func (t *SetTableIsEnum) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var trackTableResponse SetTableIsEnumResponse
	if err := json.Unmarshal(body, &trackTableResponse); err != nil {
		return nil, err
	}
	return trackTableResponse, nil
}

func NewSetTableIsEnumQuery() Query {
	query := SetTableIsEnum{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: "set_table_is_enum",
	}

	return Query(&query)
}
