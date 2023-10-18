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

// Handler untuk mendapatkan laporan kejahatan berdasarkan ID
func GetCriminalReportByID(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	paramID := p.ByName("id") // Mendapatkan ID dari parameter URL

	var report entity.CriminalReport
	// Query database untuk mendapatkan laporan kejahatan berdasarkan ID
	err := db.QueryRow("SELECT * FROM criminalreports WHERE id=?", paramID).
		Scan(&report.ID, &report.HeroID, &report.VillainID, &report.Description, &report.Time)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Data laporan kejahatan tidak ditemukan",
			})
			return
		}
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Data laporan kejahatan ditemukan",
		"report":  report,
	})
}
