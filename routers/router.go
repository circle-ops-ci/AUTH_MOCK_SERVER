// Copyright (c) 2018-2020 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package routers

import (
	"github.com/astaxie/beego"
	"github.com/cybavo/AUTH_MOCK_SERVER/controllers"
)

func init() {
	beego.Router("/v1/mock/users", &controllers.OuterController{}, "POST:RegisterUser")
	beego.Router("/v1/mock/devices", &controllers.OuterController{}, "POST:PairDevice")
	beego.Router("/v1/mock/users/pin", &controllers.OuterController{}, "POST:SetupPIN")
	beego.Router("/v1/mock/devices", &controllers.OuterController{}, "GET:GetDevices")
	beego.Router("/v1/mock/devices", &controllers.OuterController{}, "DELETE:UnpairDevices")
	beego.Router("/v1/mock/devices/2fa", &controllers.OuterController{}, "POST:SendPushToDevices")
	beego.Router("/v1/mock/users/2fa", &controllers.OuterController{}, "GET:GetDevice2FA")
	beego.Router("/v1/mock/users/2fa/:token", &controllers.OuterController{}, "DELETE:CancelDevice2FA")
	beego.Router("/v1/mock/users/me", &controllers.OuterController{}, "GET:GetUserInfo")
	beego.Router("/v1/mock/order/status", &controllers.OuterController{}, "POST:GetCallbackStatus")
	beego.Router("/v1/mock/users/totpverify", &controllers.OuterController{}, "GET:UserTotpVerify")
	beego.Router("/v1/mock/users/emailotp", &controllers.OuterController{}, "POST:SendEmailOTP")
	beego.Router("/v1/mock/users/emailotp/verify", &controllers.OuterController{}, "GET:VerifyEmailOTP")
	beego.Router("/v1/mock/callback", &controllers.CallbackController{}, "POST:Callback")
}
