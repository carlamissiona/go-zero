package handler

import (
	"net/http"
     "log"
	"blogsypanther/service/user/api/internal/logic"
	"blogsypanther/service/user/api/internal/svc"
	"blogsypanther/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		log.Println("@logic req");log.Println(req);
		l := logic.NewLoginLogic(r.Context(), svcCtx)

		resp, err := l.Login(&req);log.Println(resp);
		log.Println("@logic resp")
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
