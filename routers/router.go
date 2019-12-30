// Copyright (c) 2018-2019 The CYBAVO developers
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
	InitUpdateSRVNameSpace()
}

func InitUpdateSRVNameSpace() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/mock",
				beego.NSInclude(
					&controllers.OuterController{},
					&controllers.CallbackController{},
				),
			),
		)
	beego.AddNamespace(ns)
}
