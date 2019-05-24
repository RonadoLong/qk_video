package message

import (
	"qk_video/lib/logger"
	"testing"
)

func TestSendCode(t *testing.T) {
	code := SendCode("13570213647")
	logger.Debug("=======", code)
}

func TestSendRegisterMsg(t *testing.T) {
	_ = SendRegisterMsg("18826073368", "123123", "username")
}
