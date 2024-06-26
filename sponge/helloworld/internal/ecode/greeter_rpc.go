// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// greeter business-level rpc error codes.
// the _greeterNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_greeterNO       = 29
	_greeterName     = "greeter"
	_greeterBaseCode = errcode.RCode(_greeterNO)

	StatusSayHelloGreeter = errcode.NewRPCStatus(_greeterBaseCode+1, "failed to SayHello "+_greeterName)
	// error codes are globally unique, adding 1 to the previous error code
)
