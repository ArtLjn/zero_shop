package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"zero_shop/app/user/cmd/rpc/user"
	"zero_shop/pkg"
)

func GetRole(cli user.UserClient, userId string) int64 {
	if len(userId) == 0 {
		return 0
	}
	resp, err := cli.GetUserMsg(context.Background(), &user.GetUserMsgRequest{
		UserId: userId,
	})
	if err != nil {
		return 0
	}
	return resp.GetRoleId()
}

func AuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		result := struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{}
		if len(token) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			result.Code = 401
			result.Msg = "token is empty"
			by, _ := json.Marshal(result)
			_, err := w.Write(by)
			if err != nil {
				return
			}
			return
		}
		data, err := pkg.Verify(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)

			result.Code = 401
			result.Msg = "token is error"
			by, _ := json.Marshal(result)
			_, err := w.Write(by)
			if err != nil {
				return
			}
			return
		}
		ctx := context.WithValue(r.Context(), "userId", data.Username)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

func CheckRoleHandler(next http.HandlerFunc, cli user.UserClient, allowRole int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := GetRole(cli, r.Context().Value("userId").(string))
		if role != allowRole {
			w.WriteHeader(http.StatusUnauthorized)
			result := struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			}{}
			result.Code = 401
			result.Msg = "role is error"
			by, _ := json.Marshal(result)
			_, err := w.Write(by)
			if err != nil {
				return
			}
			return
		}
		next(w, r)
	}
}
