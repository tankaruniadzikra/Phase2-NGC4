package main

import (
	"log"
	"net/http"

	"ngc3/config"
	"ngc3/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Membuat koneksi ke database MySQL
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Tutup koneksi database setelah selesai digunakan

	router := httprouter.New()

	// Handler untuk character
	router.GET("/heroes", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.GetHeroes(w, r, p, db)
	})

	router.GET("/villains", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.GetVillains(w, r, p, db)
	})

	router.GET("/characters", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.GetCharacters(w, r, p, db)
	})

	// Handler untuk inventory
	router.GET("/inventories", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.GetInventories(w, r, p, db)
	})

	router.GET("/inventories/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.GetInventoryByID(w, r, p, db)
	})

	router.POST("/inventories", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.AddInventory(w, r, p, db)
	})

	router.PUT("/inventories/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.UpdateInventory(w, r, p, db)
	})

	router.DELETE("/inventories/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.DeleteInventory(w, r, p, db)
	})

	// Handler untuk criminal
	router.GET("/criminal_reports", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.GetCriminalReports(w, r, p, db)
	})

	router.GET("/criminal_reports/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler.GetCriminalReportByID(w, r, p, db)
	})

	// Membuat server HTTP
	app := http.Server{
		Addr:    "localhost:8001",
		Handler: router,
	}

	// Menjalankan server
	err = app.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to run app: %s\n", err.Error())
	}
}
