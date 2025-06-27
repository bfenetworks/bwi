// Copyright (c) 2019 The BFE Authors.
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

package bwi

import (
	"net"
	"net/http"
)

const (
	WAF_RESULT_PASS  = 0
	WAF_RESULT_BLOCK = 1
)

type WafResult interface {
	//get result flag (WAF_RESULT_PASS or WAF_RESULT_BLOCK)
	GetResultFlag() int

	//get attack event id
	GetEventId() string
}

// WAF server agent in client side
type WafServer interface {
	DetectRequest(req *http.Request, logId string) (WafResult, error)
	UpdateSockFactory(socketFactory func() (net.Conn, error))
	Close()
}

// require waf sdk implement this function
// func NewWafServerWithPoolSize(socketFactory func() (net.Conn, error), poolSize int) (IWafServer, error)

// require waf sdk implement this function
// func HealthCheck(conn net.Conn) error
