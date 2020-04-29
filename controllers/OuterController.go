// Copyright (c) 2018-2020 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/cybavo/AUTH_MOCK_SERVER/api"
)

type OuterController struct {
	beego.Controller
}

func getQueryString(ctx *context.Context) []string {
	var qs []string
	tokens := strings.Split(ctx.Request.URL.RawQuery, "&")
	for _, token := range tokens {
		qs = append(qs, token)
	}
	return qs
}

var debugPrint = func(ctx *context.Context) {
	var params string
	qs := getQueryString(ctx)
	if qs != nil {
		params = strings.Join(qs, "&")
	}
	logs.Debug(fmt.Sprintf("Recv requst => %s, params: %s, body: %s", ctx.Input.URL(), params, ctx.Input.RequestBody))
}

func init() {
	beego.InsertFilter("/v1/mock/*", beego.BeforeExec, debugPrint)
}

func (c *OuterController) AbortWithError(status int, err error) {
	resp := api.ErrorCodeResponse{
		ErrMsg:  err.Error(),
		ErrCode: status,
	}
	c.Data["json"] = resp
	c.Abort(strconv.Itoa(status))
}

// @router /users [post]
func (c *OuterController) RegisterUser() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("POST", "/v1/api/users", nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("RegisterUser failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /devices [post]
func (c *OuterController) PairDevice() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("POST", "/v1/api/devices", getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("PairDevice failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/pin [post]
func (c *OuterController) SetupPIN() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("POST", "/v1/api/users/pin", getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("SetupPIN failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /devices [get]
func (c *OuterController) GetDevices() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("GET", "/v1/api/devices", getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("GetDevices failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /devices [delete]
func (c *OuterController) UnpairDevices() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("DELETE", "/v1/api/devices", getQueryString(c.Ctx), c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("UnpairDevices failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /devices/2fa [post]
func (c *OuterController) SendPushToDevices() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("POST", "/v1/api/devices/2fa", getQueryString(c.Ctx), c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("SendPushToDevices failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/2fa [get]
func (c *OuterController) GetDevice2FA() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("GET", "/v1/api/users/2fa", getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("GetDevice2FA failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/2fa/:token [delete]
func (c *OuterController) CancelDevice2FA() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("DELETE",
		fmt.Sprintf("/v1/api/users/2fa/%s", c.Ctx.Input.Param(":token")), getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("CancelDevice2FA failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/me [get]
func (c *OuterController) GetUserInfo() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("GET", "/v1/api/users/me", getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("GetUserInfo failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /order/status [post]
func (c *OuterController) GetCallbackStatus() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("POST", "/v1/api/order/status", getQueryString(c.Ctx), c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("GetCallbackStatus failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/emailotp [post]
func (c *OuterController) SendEmailOTP() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("POST", "/v1/api/users/emailotp", getQueryString(c.Ctx), c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("SendEmailOTP => ", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/totpverify [get]
func (c *OuterController) UserTotpVerify() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("GET", "/v1/api/users/totpverify", getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("UserTotpVerify failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/emailotp/verify [get]
func (c *OuterController) VerifyEmailOTP() {
	defer c.ServeJSON()

	resp, err := api.MakeRequest("GET", "/v1/api/users/emailotp/verify", getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("VerifyEmailOTP => ", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/info/email[post]
func (c *OuterController) CheckUserInfoEmail() {

	defer c.ServeJSON()

	resp, err := api.MakeRequest("POST", "/v1/api/users/info/email", getQueryString(c.Ctx), c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("CheckUserInfoEmail => ", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @router /users/info/verify[get]
func (c *OuterController) VerifyUserOTP() {

	defer c.ServeJSON()

	resp, err := api.MakeRequest("GET", "/v1/api/users/info/verify", getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("CheckUserInfoEmail => ", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}
