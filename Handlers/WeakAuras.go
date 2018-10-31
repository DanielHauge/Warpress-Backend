package Handlers

import (
	Postgres "../Postgres/PreparedProcedures"
	. "../Utility/HttpHelper"
	log "../Utility/Logrus"
	"github.com/json-iterator/go"
	"net/http"
	"strconv"
)

var json = jsoniter.ConfigFastest

func HandleAddWeakaura(w http.ResponseWriter, r *http.Request, guildid int) {

	var Weak struct {
		Name   string `json:"name"`
		Link   string `json:"link"`
		Import string `json:"import"`
	}
	ReadFromRequest(w, r, &Weak)

	if e := Postgres.AddWeakaura(guildid, Weak.Name, Weak.Link, Weak.Import); e != nil {
		InterErrorHeader(w, e, "Weakaura could not be added", GetStatusCodeByError(e))
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleEditWeakaura(w http.ResponseWriter, r *http.Request, guildid int) {

	var Weak struct {
		Name   string `json:"name"`
		Link   string `json:"link"`
		Import string `json:"import"`
		Id     int    `json:"id"`
	}
	ReadFromRequest(w, r, &Weak)

	if e := Postgres.EditWeakaura(guildid, Weak.Name, Weak.Link, Weak.Import, Weak.Id); e != nil {
		InterErrorHeader(w, e, "Weakaura could not be edited", GetStatusCodeByError(e))
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleDeleteWeakaura(w http.ResponseWriter, r *http.Request, guildid int) {
	raidnightid := r.FormValue("id")
	id, err := strconv.Atoi(raidnightid)
	if err != nil {
		log.WithLocation().WithError(err).Error("Was not an integer?")
	}
	if e := Postgres.DeleteWeakaura(guildid, id); e != nil {
		InterErrorHeader(w, e, "Weakaura could not be deleted", GetStatusCodeByError(e))
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleGetWeakauras(w http.ResponseWriter, r *http.Request, guildid int) {

	weakauras, e := Postgres.GetWeakaura(guildid)
	if e != nil {
		InterErrorHeader(w, e, "Cannot get weakauras", GetStatusCodeByError(e))
		return
	}

	msg, e := json.Marshal(&weakauras)
	if e != nil {
		InterErrorHeader(w, e, "Cannot marshal", GetStatusCodeByError(e))
	} else {
		SuccessHeader(w, msg)
	}

}
