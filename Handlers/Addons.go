package Handlers

import (
	Postgres "../Postgres/PreparedProcedures"
	. "../Utility/HttpHelper"
	log "../Utility/Logrus"
	"net/http"
	"strconv"
)

func HandleAddAddon(w http.ResponseWriter, r *http.Request, guildid int) {

	var Addon struct{
		Name string `json:"name"`
		TwitchLink string `json:"twitch_link"`
	}
	ReadFromRequest(w, r, &Addon)

	if e := Postgres.AddAddon(Addon.Name, Addon.TwitchLink, guildid); e != nil {
		InterErrorHeader(w, e)
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleEditAddon(w http.ResponseWriter, r *http.Request, guildid int) {

	var Addon struct{
		Name string `json:"name"`
		TwitchLink string `json:"twitch_link"`
		Id int `json:"id"`
	}
	ReadFromRequest(w, r, &Addon)

	if e := Postgres.EditAddon(Addon.Name, Addon.TwitchLink, guildid, Addon.Id); e != nil {
		InterErrorHeader(w, e)
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleDeleteAddon(w http.ResponseWriter, r *http.Request, guildid int) {
	raidnightid := r.FormValue("id")
	id, err := strconv.Atoi(raidnightid)
	if err != nil {
		log.WithLocation().WithError(err).Error("Was not an integer?")
	}
	if e := Postgres.DeleteAddon(guildid, id); e != nil {
		InterErrorHeader(w, e)
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleGetAddon(w http.ResponseWriter, r *http.Request, guildid int) {

	addons, e := Postgres.GetAddon(guildid)
	msg, e := json.Marshal(&addons)
	if e != nil {
		InterErrorHeader(w, e)
	} else {
		SuccessHeader(w, msg)
	}

}
