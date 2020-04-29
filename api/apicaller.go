// Copyright (c) 2018-2020 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var baseURL = beego.AppConfig.DefaultString("api_server_url", "")
var APICode = beego.AppConfig.DefaultString("api_code", "")
var APISecret = beego.AppConfig.DefaultString("api_secret", "")

type ErrorCodeResponse struct {
	ErrMsg          string `json:"error,omitempty"`
	ErrCode         int    `json:"error_code,omitempty"`
	Message         string `json:"message,omitempty"`
	ServerTimestamp int64  `json:"server_timestamp,omitempty"`
}

func (m *ErrorCodeResponse) String() string {
	var msg, time string
	if m.Message != "" {
		msg = fmt.Sprintf(" (msg:%s)", m.Message)
	}
	if m.ServerTimestamp != 0 {
		time = fmt.Sprintf(" (timestamp:%d)", m.ServerTimestamp)
	}
	return fmt.Sprintf("%s%s%s (code:%d)", m.ErrMsg, msg, time, m.ErrCode)
}

func buildChecksum(params []string, secret string, time int64, r string) string {
	params = append(params, fmt.Sprintf("t=%d", time))
	params = append(params, fmt.Sprintf("r=%s", r))
	sort.Strings(params)
	params = append(params, fmt.Sprintf("secret=%s", secret))
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(params, "&"))))
}

func MakeRequest(method string, api string, params []string, postBody []byte) ([]byte, error) {
	if method == "" || api == "" {
		return nil, errors.New("invalid parameters")
	}

	r := RandomString(8)
	if r == "" {
		return nil, errors.New("can't generate random byte string")
	}
	t := time.Now().Unix()

	url := fmt.Sprintf("%s%s?t=%d&r=%s", baseURL, api, t, r)
	if len(params) > 0 {
		url += fmt.Sprintf("&%s", strings.Join(params, "&"))
	}

	var req *http.Request
	var err error
	if postBody == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewReader(postBody))
		params = append(params, string(postBody))
	}
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-CODE", APICode)
	req.Header.Set("X-CHECKSUM", buildChecksum(params, APISecret, t, r))
	if postBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	logs.Debug(fmt.Sprintf("Request URL: %s", url))
	logs.Debug("\tX-CHECKSUM:\t", req.Header.Get("X-CHECKSUM"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		result := &ErrorCodeResponse{}
		_ = json.Unmarshal(body, result)
		msg := fmt.Sprintf("%s, Error: %s", res.Status, result.String())
		return body, errors.New(msg)
	}
	return body, nil
}
