package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"ngc3/entity"

	"github.com/julienschmidt/httprouter"
)

// Handler untuk mendapatkan semua laporan kejahatan
func GetCriminalReports(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	// Query database untuk mendapatkan semua laporan kejahatan
	rows, err := db.Query("SELECT * FROM criminalreports")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var reports []entity.CriminalReport // Slice untuk menyimpan laporan kejahatan
	for rows.Next() {
		var id, heroID, villainID int
		var description string
		var time string
		err := rows.Scan(&id, &heroID, &villainID, &description, &time)
		if err != nil {
			log.Fatal(err)
		}
		// Menambahkan laporan kejahatan ke dalam slice
		reports = append(reports, entity.CriminalReport{
			ID:          id,
			HeroID:      heroID,
			VillainID:   villainID,
			Description: description,
			Time:        time,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports) // Mengirim respons JSON dengan laporan kejahatan
}
