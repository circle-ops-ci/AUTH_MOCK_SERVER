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
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

type Platform int

const (
	Android Platform = 0x1
	iOS     Platform = 0x2
	Browser Platform = 0x4
)

type RuleType int

const (
	PolicyTypeNone      RuleType = iota
	PolicyTypeGlobal    RuleType = 1
	PolicyTypeCustomize RuleType = 2
	PolicyTypeRiskRule  RuleType = 3
)

type Action int

// Action
const (
	NoAction    Action = iota
	AlwaysAllow Action = 1
	Require2FA  Action = 2
	DenyAccess  Action = 3
)

type UserAction int

const (
	UserActionNone   UserAction = iota
	UserActionAccept UserAction = 1
	UserActionReject UserAction = 2
)

type State int

const (
	Pending2FACreated  State = iota
	Pending2FADone     State = 1
	Pending2FAPassed   State = 2
	Pending2FACanceled State = 3
	Pending2FAFailed   State = 4
)

type ErrorCodeResponse struct {
	ErrMsg  string `json:"error,omitempty"`
	ErrCode int    `json:"error_code,omitempty"`
}

type RegisterUserRequest struct {
	Account  string `json:"account"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Language string `json:"locale"`
}

type RegisterUserResponse struct {
	Account string `json:"account"`
	Email   string `json:"email"`
}

type PairDeviceResponse struct {
	OrderID int64  `json:"order_id"`
	URL     string `json:"url"`
}

type SetupPINResponse struct {
	OrderID int64 `json:"order_id"`
}

type DeviceInfo struct {
	Name           string `json:"name"`
	Platform       string `json:"platform"`
	DeviceID       string `json:"device_id"`
	ServiceID      int64  `json:"service_id"`
	LastActiveTime int64  `json:"last_active_time"`
	CreateTime     int64  `json:"create_time"`
}

type GetDevicesResponse struct {
	Devices []*DeviceInfo `json:"devices"`
}

type UnpairDevicesRequest struct {
	DeviceIDs []string `json:"devices"`
}

type UnpairDevicesResponse struct {
	RemovedDeviceIDs []string `json:"removed_devices"`
}

type SendPushToDevicesRequest struct {
	Type     int64                  `json:"type"`
	Title    string                 `json:"title"`
	Body     string                 `json:"body"`
	Data     map[string]interface{} `json:"data"`
	IP       string                 `json:"client_ip"`
	Platform Platform               `json:"client_platform"`
}

type MatchedPolicyInfo struct {
	PolicyID   int64    `json:"policy_id"`
	RuleType   RuleType `json:"rule_type"`
	PolicyName string   `json:"policy_name"`
}

type SendPushToDevicesResponse struct {
	SuccessDevices    []string          `json:"success_devices,omitempty"`
	Action            Action            `json:"action"`
	OrderId           int64             `json:"order_id"`
	MatchedPolicyInfo MatchedPolicyInfo `json:"matched_policy"`
}

type GetUserInfoResponse struct {
	Account     string `json:"account"`
	UserEmail   string `json:"user_email"`
	ServiceID   int64  `json:"service_id"`
	DeviceCount int    `json:"device_count"`
	IsSetupPin  bool   `json:"is_setup_pin"`
}

type CancelDevice2faResponse struct {
	CanceledDevices []string `json:"canceled_devices,omitempty"`
}

type Device2faResponse struct {
	OrderId      int64  `json:"order_id"`
	Type         int    `json:"type"`
	UserAction   int    `json:"user_action"`
	State        int    `json:"state"`
	UpdatedTime  int64  `json:"updated_time"`
	MessageType  int64  `json:"message_type,omitempty"`
	MessageTitle string `json:"message_title,omitempty"`
	MessageBody  string `json:"message_body,omitempty"`
	DeviceSent   int    `json:"device_sent"`
}

type GetDevice2FAResponse struct {
	Items []*Device2faResponse `json:"items"`
}

type CommonResponse struct {
	OrderID int64 `json:"order_id"`
}

type GetCallbackStatusRequest struct {
	OrderIDs []int64 `json:"order_ids"`
}

type OrderStatus struct {
	IsExist        bool                   `json:"is_exist"`
	OrderID        int64                  `json:"order_id"`
	BehaviorType   int                    `json:"behavior_type"`
	BehaviorResult int                    `json:"behavior_result"`
	Addon          map[string]interface{} `json:"addon"`
}

type GetCallbackStatusResponse struct {
	OrderStatus []OrderStatus `json:"order_status"`
}

type SendEmailOTPRequest struct {
	Url      string `json:"url"`
	Duration int    `json:"duration"`
}

type SendEmailOTPResponse struct {
}

type UserTotpVerifyResponse struct {
	Result bool `json:"result"`
}

const (
	BehaviorResultPending = 0
	BehaviorResultReject  = 1
	BehaviorResultAccept  = 2
	BehaviorResultExpired = 3
	BehaviorResultFailed  = 4
)

const (
	BehaviorTypePairedDevice  = 1
	BehaviorTypeSetupPIN      = 2
	BehaviorTypeCustomMessage = 9
)

type CallbackStruct struct {
	ServiceID      int64 `json:"service_id"`
	OrderID        int64 `json:"order_id"`
	BehaviorType   int   `json:"behavior_type"`
	BehaviorResult int   `json:"behavior_result"`
}

func RegisterUser(request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest("POST", "/v1/api/users", nil, jsonRequest)
	if err != nil {
		return nil, err
	}

	response = &RegisterUserResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("RegisterUser() => ", response)
	return
}

func PairDevice(qs []string) (response *PairDeviceResponse, err error) {
	resp, err := makeRequest("POST", "/v1/api/devices", qs, nil)
	if err != nil {
		return nil, err
	}

	response = &PairDeviceResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("PairDevice() => ", response)
	return
}

func SetupPIN(qs []string) (response *SetupPINResponse, err error) {
	resp, err := makeRequest("POST", "/v1/api/users/pin", qs, nil)
	if err != nil {
		return nil, err
	}

	response = &SetupPINResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("SetupPIN() => ", response)
	return
}

func GetDevices(qs []string) (response *GetDevicesResponse, err error) {
	resp, err := makeRequest("GET", "/v1/api/devices", qs, nil)
	if err != nil {
		return nil, err
	}

	response = &GetDevicesResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("GetDevices() => ", response)
	return
}

func UnpairDevices(request *UnpairDevicesRequest, qs []string) (response *UnpairDevicesResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest("DELETE", "/v1/api/devices", qs, jsonRequest)
	if err != nil {
		return nil, err
	}

	response = &UnpairDevicesResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("UnpairDevices() => ", response)
	return
}

func SendPushToDevices(request *SendPushToDevicesRequest, qs []string) (response *SendPushToDevicesResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest("POST", "/v1/api/devices/2fa", qs, jsonRequest)
	if err != nil {
		return nil, err
	}

	response = &SendPushToDevicesResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("SendPushToDevices() => ", response)
	return
}

func GetDevice2FA(qs []string) (response *GetDevice2FAResponse, err error) {
	resp, err := makeRequest("GET", "/v1/api/users/2fa", qs, nil)
	if err != nil {
		return nil, err
	}

	response = &GetDevice2FAResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("GetDevice2FA() => ", response)
	return
}

func CancelDevice2FA(token string, qs []string) (response *CancelDevice2faResponse, err error) {
	resp, err := makeRequest("DELETE", fmt.Sprintf("/v1/api/users/2fa/%s", token), qs, nil)
	if err != nil {
		return nil, err
	}

	response = &CancelDevice2faResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("CancelDevice2FA() => ", response)
	return
}

func GetUserInfo(qs []string) (response *GetUserInfoResponse, err error) {
	resp, err := makeRequest("GET", "/v1/api/users/me", qs, nil)
	if err != nil {
		return nil, err
	}

	response = &GetUserInfoResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("GetUserInfo() => ", response)
	return
}

func GetCallbackStatus(request *GetCallbackStatusRequest, qs []string) (response *GetCallbackStatusResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := makeRequest("POST", "/v1/api/order/status", qs, jsonRequest)
	if err != nil {
		return nil, err
	}

	response = &GetCallbackStatusResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("GetCallbackStatus() => ", response)
	return
}

func SendEmailOTP(request *SendEmailOTPRequest, qs []string) (response *SendEmailOTPResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest("POST", "/v1/api/users/emailotp", qs, jsonRequest)
	if err != nil {
		return nil, err
	}

	response = &SendEmailOTPResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("SendEmailOTP() => ", response)
	return
}

func UserTotpVerify(qs []string) (response *UserTotpVerifyResponse, err error) {
	resp, err := makeRequest("GET", "/v1/api/users/totpverify", qs, nil)
	if err != nil {
		return nil, err
	}

	response = &UserTotpVerifyResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("UserTotpVerify() => ", response)
	return
}
