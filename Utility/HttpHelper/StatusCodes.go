package HttpHelper

import (
	"net/http"
	"strings"
)

func GetStatusCodeByError(e error) int {
	if strings.Contains(e.Error(), "Guild is allready registrered") {
		return http.StatusConflict
	} else if strings.Contains(e.Error(), "No such guild is registered") {
		return http.StatusNotFound
	} else if strings.Contains(e.Error(), "Insufficient rank in guild") {
		return http.StatusForbidden
	} else if strings.Contains(e.Error(), "No main registered") {
		return http.StatusForbidden
	} else if strings.Contains(e.Error(), "Authentication with blizzard failed") {
		return http.StatusUnauthorized
	} else if strings.Contains(e.Error(), "No Addons have been added for this guild") {
		return http.StatusNotFound
	} else if strings.Contains(e.Error(), "No Weakaura have been added for this guild") {
		return http.StatusNotFound
	} else if strings.Contains(e.Error(), "No raid nights have been added for this guild") {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
