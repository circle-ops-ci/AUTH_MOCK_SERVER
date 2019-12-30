// Copyright (c) 2018-2019 The CYBAVO developers
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

func buildChecksum(params []string, secret string, time int64, r string) string {
	params = append(params, fmt.Sprintf("t=%d", time))
	params = append(params, fmt.Sprintf("r=%s", r))
	sort.Strings(params)
	params = append(params, fmt.Sprintf("secret=%s", secret))
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(params, "&"))))
}

func makeRequest(method string, api string, params []string, postBody []byte) ([]byte, error) {
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
	logs.Debug(fmt.Sprintf("Request URL: %s, params: %s", req.URL.String(), strings.Join(params, "&")))
	logs.Debug("Request URL:", req.URL.String())
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
		showErrorMessage(body)
		return body, errors.New(res.Status)
	}
	return body, nil
}

func showErrorMessage(resp []byte) {
	errResp := &ErrorCodeResponse{}
	err := json.Unmarshal(resp, errResp)
	if err == nil {
		logs.Debug(fmt.Sprintf("Error: %s (code: %d)", errResp.ErrMsg, errResp.ErrCode))
	}
}
