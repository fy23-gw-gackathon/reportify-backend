package entity

import (
	"log"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, err error) *Error {
	log.Println(err)
	message, ok := messages[code]
	if !ok {
		return &Error{
			Code:    code,
			Message: messages[http.StatusInternalServerError],
		}
	}
	return &Error{
		Code:    code,
		Message: message,
	}
}

var messages = map[int]string{
	http.StatusBadRequest:          "リクエストの形式が正しくありません。",
	http.StatusUnauthorized:        "認証に失敗しました。ログイン状態を確認してください。",
	http.StatusForbidden:           "管理者権限がありません。",
	http.StatusNotFound:            "お探しのリソースは存在しません。",
	http.StatusConflict:            "リクエストされたコンテンツは既に存在しています。リソースを確認して再度リクエストしてください。",
	http.StatusInternalServerError: "予期せぬエラーが発生しました。",
}
