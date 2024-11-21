package handler

import (
	"context"
	"encoding/json"

	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// ServeHTTP implements http.Handler interface.
func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		// w, r
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// 1. リクエストボディの文字列を変数に取り出す
		rawBody := make([]byte, r.ContentLength)
		_, err := r.Body.Read(rawBody)
		// //fmt.Println(string(rawBody), "頼む")
		if err != nil {
			log.Println(err)
		}
		// rawBody にリクエストボディの文字列が入っている状態

		// //fmt.Println(rawBody)

		// 2. リクエストボディを構造体にする

		var body model.CreateTODORequest

		err = json.Unmarshal(rawBody, &body) // 文字列をもとに、構造体にデータを詰め込む関数
		// //fmt.Println("body.subject", body.Subject)'

		if err != nil {
			log.Println("err:", err)
		}

		if body.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// body にはリクエストボディのデータが入っている状態

		res, err := h.Create(r.Context(), &body)
		if err != nil {
			//fmt.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ここでレスポンスを整形する
		//fmt.Println(&res)
		responseJSON, err := json.Marshal(&res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(responseJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {

	some, thing := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	// //fmt.Println("err", err)
	// //fmt.Println("some", some)
	// //fmt.Println("create")
	if thing != nil {
		return &model.CreateTODOResponse{}, thing
	}

	return &model.CreateTODOResponse{TODO: *some}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
