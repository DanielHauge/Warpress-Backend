package Filters

import (
	"../../Integrations/BlizzardOauthAPI"
	. "../HttpHelper"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RequireAuthentication(HandleFunction func(w http.ResponseWriter, r *http.Request, id int, region string)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acces, id, region := BlizzardOauthAPI.DoesUserHaveAccess(w, r)
		if acces {
			HandleFunction(w, r, id, region)
		} else {
			log.Info("Unautherized user tried to access")
			e := errors.New("Authentication with blizzard failed, try login with blizzard again")
			InterErrorHeader(w, e, "Cannot procede without authentication", GetStatusCodeByError(e))
		}
	})
}
