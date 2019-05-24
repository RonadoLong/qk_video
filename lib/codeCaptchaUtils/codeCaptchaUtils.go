package codeCaptchaUtils

import (
	"github.com/mojocn/base64Captcha"
	"log"
)
//数字验证码配置
var configD = base64Captcha.ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 5,
}
//config struct for audio
//声音验证码配置
var configA = base64Captcha.ConfigAudio{
	CaptchaLen: 6,
	Language:   "zh",
}
//字符,公式,验证码配置
var configC = base64Captcha.ConfigCharacter{
	Height:             40,
	Width:              200,
	//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
	Mode:               base64Captcha.CaptchaModeArithmetic,
	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         6,
}
func CodeCaptchaCreate() (string, string) {
	////create a audio captcha.
	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	////以base64编码
	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//create a characters captcha.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	////create a digits captcha.
	//idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	////以base64编码
	//base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	//fmt.Println(idKeyC, base64stringC)
	return idKeyC, base64stringC
}

func VerifyCaptcha(idkey, verifyValue string) bool{
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		return true
	} else {
		log.Println("e64Captcha.VerifyCaptcha fail")
		return false
	}
}
