package middleware

import (
	"net/http"

	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/claims"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/recipe/userroles/userrolesclaims"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func VerifySession(handler http.Handler) http.Handler {
	return session.VerifySession(nil, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionContainer := session.GetSessionFromRequestContext(r.Context())
		userID := sessionContainer.GetUserID()
		r.Header.Add("X-User", userID)
		handler.ServeHTTP(w, r)
	}))
}
func VerifyAdmin(handler http.Handler) http.Handler {
	return session.VerifySession(&sessmodels.VerifySessionOptions{
		OverrideGlobalClaimValidators: func(globalClaimValidators []claims.SessionClaimValidator, sessionContainer sessmodels.SessionContainer, userContext supertokens.UserContext) ([]claims.SessionClaimValidator, error) {
			request := supertokens.GetRequestFromUserContext(userContext)
			teamID := request.Header.Get("X-Team")
			globalClaimValidators = append(globalClaimValidators, userrolesclaims.UserRoleClaimValidators.Includes(teamID+"_owner", nil, nil))
			return globalClaimValidators, nil
		},
	}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionContainer := session.GetSessionFromRequestContext(r.Context())
		userID := sessionContainer.GetUserID()
		r.Header.Add("X-User", userID)
		handler.ServeHTTP(w, r)
	}))
}
