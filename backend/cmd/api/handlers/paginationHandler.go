package handlers

import (
	"backend/pkg/utilities"
	"context"
	"net/http"
)

func (h *handler) Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagQuery := utilities.GetPaginationFromRequest(r)
		if pagQuery == nil {
			http.Error(w, "Could not parse Pagination Query", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), utilities.PaginationQuery{}, pagQuery)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
