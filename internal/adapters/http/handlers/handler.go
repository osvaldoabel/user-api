package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/osvaldoabel/user-api/internal/entity"
)

const (
	DEFAULT_PER_PAGE = 10
	DEFAULT_PAGE     = 0
)

type AppError struct {
	Message string `json:"message"`
}

func JsonResponse(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}

func GetPaginationInfo(r *http.Request) entity.Pagination {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	offsetInt, err := strconv.Atoi(page)
	if err != nil {
		offsetInt = DEFAULT_PAGE
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = DEFAULT_PER_PAGE
	}

	offset := (offsetInt - 1) * limitInt
	return entity.Pagination{
		Limit: limitInt,
		Page:  offset,
	}
}
