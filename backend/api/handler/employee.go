package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matheusmosca/nutcache/api/presenter"
	"github.com/matheusmosca/nutcache/service"
)

func getAll(service service.EmployeeService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := context.Background()
		employees, err := service.GetAll(ctx)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := presenter.NewResponse(false, nil, err.Error())
			encodeResponse(w, response)
			return
		}

		encodeResponse(w, presenter.NewResponse(true, employees, ""))
	})
}

func create(service service.EmployeeService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := context.Background()
		var body presenter.EmployeeRequest

		decodeRequestBody(r, &body)
		err := service.Create(ctx, body.Name, body.Gender, body.Email, body.CPF, body.Team, body.StartDate, body.BirthDate)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := presenter.NewResponse(false, nil, err.Error())
			encodeResponse(w, response)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		encodeResponse(w, presenter.NewResponse(true, nil, ""))
	})
}

func update(service service.EmployeeService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := context.Background()
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			response := presenter.NewResponse(false, nil, "")
			encodeResponse(w, response)
			return
		}
		var body presenter.EmployeeUpdate

		decodeRequestBody(r, &body)
		err = service.Update(ctx, id, body.Name, body.Gender, body.Email, body.CPF, body.Team, body.StartDate, body.BirthDate)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := presenter.NewResponse(false, nil, err.Error())
			encodeResponse(w, response)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		encodeResponse(w, presenter.NewResponse(true, nil, ""))
	})
}

func getByID(service service.EmployeeService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := context.Background()
		params := mux.Vars(r)
		ID, err := strconv.Atoi(params["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			response := presenter.NewResponse(false, nil, "")
			encodeResponse(w, response)
			return
		}
		employee, err := service.GetByID(ctx, ID)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := presenter.NewResponse(false, nil, err.Error())
			encodeResponse(w, response)
			return
		}

		response := presenter.EmployeeResponse(*employee)
		encodeResponse(w, presenter.NewResponse(true, response, ""))
	})
}

func delete(service service.EmployeeService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := context.Background()
		params := mux.Vars(r)

		ID, err := strconv.Atoi(params["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			response := presenter.NewResponse(false, nil, "")
			encodeResponse(w, response)
			return
		}
		err = service.Delete(ctx, ID)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			response := presenter.NewResponse(false, nil, err.Error())
			encodeResponse(w, response)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func MakeEmployeeEndpoints(r *mux.Router, service service.EmployeeService) {
	r.Handle("/api/v1/employees", getAll(service)).Methods("GET")
	r.Handle("/api/v1/employees", create(service)).Methods("POST")
	r.Handle("/api/v1/employees/{id}", getByID(service)).Methods("GET")
	r.Handle("/api/v1/employees/{id}", update(service)).Methods("PUT")
	r.Handle("/api/v1/employees/{id}", delete(service)).Methods("DELETE")
}
