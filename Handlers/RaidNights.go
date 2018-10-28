package Handlers

import (
	Postgres "../Postgres/PreparedProcedures"
	. "../Utility/HttpHelper"
	log "../Utility/Logrus"
	"net/http"
	"strconv"
	"time"
)

func HandleAddRaidNight(w http.ResponseWriter, r *http.Request, guildid int) {

	var Raidnight struct {
		Duration time.Duration `json:"duration"`
		Start    time.Duration `json:"start"`
		Day      int           `json:"day"`
	}
	ReadFromRequest(w, r, &Raidnight)

	if e := Postgres.AddRaidNight(Raidnight.Duration, Raidnight.Start, Raidnight.Day, guildid); e != nil {
		InterErrorHeader(w, e)
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleEditRaidNight(w http.ResponseWriter, r *http.Request, guildid int) {

	var Raidnight struct {
		Duration    time.Duration `json:"duration"`
		Start       time.Time     `json:"start"`
		Day         int           `json:"day"`
		RaidnightId int           `json:"raidnight_id"`
	}
	ReadFromRequest(w, r, &Raidnight)

	if e := Postgres.EditRaidNight(Raidnight.Duration, Raidnight.Start, Raidnight.Day, Raidnight.RaidnightId, guildid); e != nil {
		InterErrorHeader(w, e)
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleDeleteRaidNight(w http.ResponseWriter, r *http.Request, guildid int) {
	raidnightid := r.FormValue("id")
	id, err := strconv.Atoi(raidnightid)
	if err != nil {
		log.WithLocation().WithError(err).Error("Was not an integer?")
	}
	if e := Postgres.DeleteRaidNight(id, guildid); e != nil {
		InterErrorHeader(w, e)
		return
	}

	SuccessHeader(w, []byte("Succes"))

}

func HandleGetRaidNights(w http.ResponseWriter, r *http.Request, guildid int) {

	nights, e := Postgres.GetRaidNights(guildid)
	msg, e := json.Marshal(&nights)
	if e != nil {
		InterErrorHeader(w, e)
	} else {
		SuccessHeader(w, msg)
	}

}
