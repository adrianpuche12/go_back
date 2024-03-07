package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go_fundaments/internal/user"
	"github.com/go_fundaments/pkg/transport"
	"log"
	"net/http"
	"strconv"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, enpoints user.Endpoint) {
	router.HandleFunc("/users/", UserServer(ctx, enpoints))
}

func UserServer(ctx context.Context, enpoints user.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.Path
		log.Println(r.Method, " : ", url)
		transport.Clean(url)
		path, pathSize := transport.Clean(url)

		params := make(map[string]string)
		if pathSize == 4 && path[2] != "" {
			params["userID"] = path[2]
		}

		tran := transport.New(w, r, context.WithValue(ctx, "params", params))

		var end user.Controller
		var deco func(ctx context.Context, r *http.Request) (interface{}, error)

		switch r.Method {
		case http.MethodGet:
			switch pathSize {
			case 3:
				end = enpoints.GetAll
				deco = decodeGetAllUser
			case 4:
				end = enpoints.Get
				deco = decodeGetUser
			}
		case http.MethodPost:
			switch pathSize {
			case 3:
				end = enpoints.Create
				deco = decodeCreateUser
			}
		case http.MethodPatch:
			switch pathSize {
			case 4:
				end = enpoints.Update
				deco = decoUpdateUser
			}
		}

		if end != nil && deco != nil {
			tran.Server(
				transport.Endpoint(end),
				deco,
				encodeResponse,
				encodeError)

		} else {
			InvalidMethod(w)
		}
	}
}

func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)
	fmt.Println(params)

	id, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		return nil, err
	}

	return user.GetReq{
		ID: id,
	}, nil

}

func decoUpdateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req user.UpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}
	params := ctx.Value("params").(map[string]string)

	id, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		return nil, err
	}
	req.ID = id
	return req, nil
}

func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {

	// la funcion Marshal transforma una entidad en un Json
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	status := http.StatusOK
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, data)
	return nil
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	status := http.StatusInternalServerError
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "message": %s}`, status, err.Error())
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "method doesnt exist"}`, status)
}
