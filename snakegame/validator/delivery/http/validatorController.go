package delivery

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
)

type ValidatorController struct {
	ValidatorSvc domain.Validator
}

func NewValidatorController(r *mux.Router, svc domain.Validator) {
	controller := &ValidatorController{
		ValidatorSvc: svc,
	}
	Route(r, controller)
}

func (con *ValidatorController) GetNewGame(w http.ResponseWriter, r *http.Request) {
	wi := r.URL.Query().Get("w")
	h := r.URL.Query().Get("h")
	if (wi == "") && (h == "") {
		con.HandleGETError(400, errors.New("missing width and height"), w)
		return
	}
	width, e := strconv.Atoi(wi)
	if e != nil {
		con.HandleGETError(400, errors.New("invalid width"), w)
		return
	}
	height, e := strconv.Atoi(h)
	if e != nil {
		con.HandleGETError(400, errors.New("invalid height"), w)
		return
	}
	state, err := con.ValidatorSvc.NewGame(width, height)
	if err != nil {
		con.HandleGETError(400, err, w)
		return
	}
	json.NewEncoder(w).Encode(&state)
}
func (con *ValidatorController) Validate(w http.ResponseWriter, r *http.Request) {
	req, e := ioutil.ReadAll(r.Body)
	var resp domain.State
	if len(req) == 0 {
		con.HandlePOSTError(400, resp, w)
		return
	}
	if e != nil {
		con.HandlePOSTError(400, resp, w)
		return
	}
	var validateReq domain.ValidateRequest
	json.Unmarshal(req, &validateReq)

	resp, valerr := con.ValidatorSvc.Validate(&validateReq)
	if valerr != nil && valerr.(*domain.ValidationError).ErrCode >= 400 {
		con.HandlePOSTError(valerr.(*domain.ValidationError).ErrCode, resp, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

func (con *ValidatorController) HandleGETError(status int, e error, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(e.Error()))
}

func (con *ValidatorController) HandlePOSTError(status int, resp domain.State, w http.ResponseWriter) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&resp)
}

// 2021/11/09 00:33:46 {1 5 5 1 {2 3} {4 4 1 0}}

// expected: main.state{GameID:"1", Width:5, Height:5, Score:1, Fruit:main.fruit{X:2, Y:3}, Snake:main.snake{X:4, Y:4, VelX:1, VelY:0}}
// actual  : main.state{GameID:"1", Width:5, Height:5, Score:1, Fruit:main.fruit{X:2, Y:3}, Snake:main.snake{X:4, Y:4, VelX:0, VelY:1}}
