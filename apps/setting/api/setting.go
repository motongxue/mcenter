package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"

	"github.com/infraboard/mcenter/apps/setting"
)

func (h *handler) UpdateSetting(r *restful.Request, w *restful.Response) {
	req := setting.NewDefaultSetting()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateSetting(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (u *handler) GetSetting(r *restful.Request, w *restful.Response) {
	set, err := h.service.GetSetting(r.Request.Context())
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
