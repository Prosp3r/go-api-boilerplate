/*
Package firewall allow to guard handlers
*/
package firewall

import (
	"net/http"

	"github.com/vardius/go-api-boilerplate/pkg/http/response"
	"github.com/vardius/go-api-boilerplate/pkg/identity"
)

// GrantHTTPAccessFor returns Status Unauthorized if
// Identity not set within request's context
// or user does not have required role
func GrantHTTPAccessFor(role string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			i, ok := identity.FromContext(r.Context())
			if ok {
				for _, userRole := range i.Roles {
					if userRole == role {
						next.ServeHTTP(w, r)
						return
					}
				}
			}

			response.WithError(r.Context(), response.NewErrorFromHTTPStatus(http.StatusUnauthorized))
		}

		return http.HandlerFunc(fn)
	}
}
