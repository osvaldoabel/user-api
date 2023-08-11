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

func JsonResponse(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}

func GetPaginationInfo(r *http.Request) entity.Pagination {
	offset := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = DEFAULT_PAGE
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = DEFAULT_PER_PAGE
	}

	// debug.DUMP(entity.Pagination{
	// 	Limit:  limitInt,
	// 	Offset: offsetInt,
	// })

	return entity.Pagination{
		Limit:  limitInt,
		Offset: offsetInt,
	}
}
