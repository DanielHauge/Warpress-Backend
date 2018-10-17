package Filters

import (
	"../Integrations/BlizzardOauthAPI"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RequireAuthentication(HandleFunction func(w http.ResponseWriter, r *http.Request, id int, region string)) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		acces, id, region := BlizzardOauthAPI.DoesUserHaveAccess(w, r)
		if acces{

			HandleFunction(w, r, id, region)


		} else {
			log.Info("Unautherized user tried to access")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unfortunately it seemed like you didn't have access, try login with blizzard again"))
		}



	})
}
