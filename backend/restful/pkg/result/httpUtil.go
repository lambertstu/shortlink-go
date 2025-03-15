package result

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func Success(w http.ResponseWriter) {
	httpResp := NewResult[any]().SetCode(SUCCESS_CODE)
	httpx.WriteJson(w, http.StatusOK, httpResp)
}

func SuccessWithMsg[T any](w http.ResponseWriter, data T) {
	httpResp := NewResult[T]().SetCode(SUCCESS_CODE).SetData(data)
	httpx.WriteJson(w, http.StatusOK, httpResp)
}

func FailureWithErr[T any](w http.ResponseWriter, data T, err error) {
	httpResp := NewResult[any]().SetCode(FAIL_CODE).SetMessage(err.Error()).SetData(data)
	httpx.WriteJson(w, http.StatusInternalServerError, httpResp)
}

func Response(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		FailureWithErr(w, data, err)
		return
	}

	if data == nil {
		Success(w)
		return
	}

	SuccessWithMsg(w, data)
}
