package middleware

import (
	"context"
	"net/http"

	"github.com/karankap00r/employee_portal/storage/repository"
	"github.com/karankap00r/employee_portal/util"
)

type key int

const (
	OrgIDKey       key = iota
	ClientIDHeader     = "X-Client-ID"
)

// OrgResolver is a middleware to resolve the org ID from the client ID
func OrgResolver(repo repository.OrgRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientID := r.Header.Get(ClientIDHeader)
			if clientID == "" {
				util.WriteErrorResponse(w, http.StatusBadRequest, "Client ID is missing")
				return
			}

			org, err := repo.GetOrgByClientID(clientID)
			if err != nil {
				util.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid Client ID")
				return
			}

			ctx := context.WithValue(r.Context(), OrgIDKey, org.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetOrgIDFromContext returns the org ID from the context
func GetOrgIDFromContext(ctx context.Context) (int, bool) {
	orgID, ok := ctx.Value(OrgIDKey).(int)
	return orgID, ok
}
