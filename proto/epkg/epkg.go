package epkg

import (
	"encoding/json"
	"fmt"
	"github.com/dangersalad/httperror"
	"net/http"
	"strconv"
	"strings"
)

const (
	BraveTroopsCodeSeparatorBegin = "-----BraveTroops CODE SEPARATOR BEGIN-----"
	BraveTroopsCodeSeparatorEnd   = "-----BraveTroops CODE SEPARATOR END-----"
)

func WrapFail(code int, message string) string {
	text := fmt.Sprintf("%s%v%s%s",
		BraveTroopsCodeSeparatorBegin,
		code,
		BraveTroopsCodeSeparatorEnd,
		message,
	)
	return text
}

func WrapFailV2(err error) string {
	if httpError, ok := err.(httperror.HTTPError); ok {
		return WrapFail(httpError.Code, httpError.Message)
	}
	return WrapFail(http.StatusInternalServerError, err.Error())
}

func UnwrapFail(err error) (code int, message string) {
	text := err.Error()
	code, message = http.StatusInternalServerError, text
	tokenList := strings.Split(text, BraveTroopsCodeSeparatorBegin)
	if len(tokenList) < 2 {
		return
	}
	headPart, restString := tokenList[0], tokenList[1]
	tokenList = strings.Split(restString, BraveTroopsCodeSeparatorEnd)
	if len(tokenList) < 2 {
		return
	}
	codeString, tailPart := tokenList[0], tokenList[1]
	codeInt, err := strconv.Atoi(codeString)
	if err != nil {
		return
	}
	code, message = codeInt, headPart+tailPart
	return
}

func WrapSucc(result interface{}) ([]byte, error) {
	return json.Marshal(result)
}

func UnwrapSucc(payload []byte, result interface{}) error {
	return json.Unmarshal(payload, result)
}

func Wrapf(errCode int, format string, args ...interface{}) error {
	sprint := fmt.Sprint(format, args)
	return httperror.New(errCode, sprint)
}
