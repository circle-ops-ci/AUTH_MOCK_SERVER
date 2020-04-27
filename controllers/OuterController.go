// Copyright (c) 2018-2019 The CYBAVO developers
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
	"errors"
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

	var request api.RegisterUserRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	resp, err := api.RegisterUser(&request)
	if err != nil {
		logs.Error("RegisterUser failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /devices [post]
func (c *OuterController) PairDevice() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.PairDevice(qs)
	if err != nil {
		logs.Error("PairDevice failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/pin [post]
func (c *OuterController) SetupPIN() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.SetupPIN(qs)
	if err != nil {
		logs.Error("SetupPIN failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /devices [get]
func (c *OuterController) GetDevices() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetDevices(qs)
	if err != nil {
		logs.Error("GetDevices failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /devices [delete]
func (c *OuterController) UnpairDevices() {
	defer c.ServeJSON()

	var request api.UnpairDevicesRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.UnpairDevices(&request, qs)
	if err != nil {
		logs.Error("UnpairDevices failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /devices/2fa [post]
func (c *OuterController) SendPushToDevices() {
	defer c.ServeJSON()

	var request api.SendPushToDevicesRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.SendPushToDevices(&request, qs)
	if err != nil {
		logs.Error("SendPushToDevices failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/2fa [get]
func (c *OuterController) GetDevice2FA() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetDevice2FA(qs)
	if err != nil {
		logs.Error("GetDevice2FA failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/2fa/:token [delete]
func (c *OuterController) CancelDevice2FA() {
	defer c.ServeJSON()

	token := c.Ctx.Input.Param(":token")
	if token == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.CancelDevice2FA(token, qs)
	if err != nil {
		logs.Error("CancelDevice2FA failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/me [get]
func (c *OuterController) GetUserInfo() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetUserInfo(qs)
	if err != nil {
		logs.Error("GetUserInfo failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /order/status [post]
func (c *OuterController) GetCallbackStatus() {
	defer c.ServeJSON()

	var request api.GetCallbackStatusRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetCallbackStatus(&request, qs)
	if err != nil {
		logs.Error("GetCallbackStatus failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/emailotp [post]
func (c *OuterController) SendEmailOTP() {

	defer c.ServeJSON()

	var request api.SendEmailOTPRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.SendEmailOTP(&request, qs)
	if err != nil {
		logs.Error("SendEmailOTP => ", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/totpverify [get]
func (c *OuterController) UserTotpVerify() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.UserTotpVerify(qs)
	if err != nil {
		logs.Error("UserTotpVerify failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/emailotp/verify [get]
func (c *OuterController) VerifyEmailOTP() {

	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.VerifyEmailOTP(qs)
	if err != nil {
		logs.Error("VerifyEmailOTP => ", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}
