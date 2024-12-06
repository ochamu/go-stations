package handler

import (
	"context"
	"encoding/json"

	"fmt"
	"strconv"

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

	} else if r.Method == http.MethodPut {
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

		var body model.UpdateTODORequest

		err = json.Unmarshal(rawBody, &body) // 文字列をもとに、構造体にデータを詰め込む関数
		// //fmt.Println("body.subject", body.Subject)'

		if err != nil {
			log.Println("err:", err)
		}

		if body.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if body.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// body にはリクエストボディのデータが入っている状態

		res, err := h.Update(r.Context(), &body)
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
	} else if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var req model.ReadTODORequest
		req.PrevID, _ = strconv.ParseInt(r.URL.Query().Get("prev_id"), 10, 64)
		sizeStr := r.URL.Query().Get("size")
		if sizeStr == "" {
			req.Size = 5
		} else {
			size, err := strconv.ParseInt(sizeStr, 10, 64)
			if err != nil || size <= 0 {
				req.Size = 5
			} else {
				req.Size = size
			}
		}
		res, err := h.Read(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(res)
	} else if r.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var req model.DeleteTODORequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println(err)
		}
		if req.IDs == nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		if len(req.IDs) == 0 {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = h.svc.DeleteTODO(r.Context(), req.IDs)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = json.NewEncoder(w).Encode(&model.DeleteTODOResponse{})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		}
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {

	some, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	// //fmt.Println("err", err)
	// //fmt.Println("some", some)
	// //fmt.Println("create")
	if err != nil {
		return &model.CreateTODOResponse{}, err
	}

	return &model.CreateTODOResponse{TODO: *some}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	todos, err := h.svc.ReadTODO(ctx, req.PrevID, req.Size)

	if err != nil {
		return nil, err
	}

	// todoValues := make([]model.TODO, len(todos))
	// for i, todoPtr := range todos {
	// 	todoValues[i] = *todoPtr
	// }

	return &model.ReadTODOResponse{TODOs: todos}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {

	some, err := h.svc.UpdateTODO(ctx, req.ID, req.Subject, req.Description)
	// //fmt.Println("err", err)
	// //fmt.Println("some", some)
	// //fmt.Println("create")
	if err != nil {
		return &model.UpdateTODOResponse{}, err
	}

	return &model.UpdateTODOResponse{TODO: *some}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

// panic

type DoPanicHandler struct{}

func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

func (h *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("ok")
}
