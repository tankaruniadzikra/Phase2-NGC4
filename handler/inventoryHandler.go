package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"ngc3/entity"

	"github.com/julienschmidt/httprouter"
)

// Handler untuk mengambil data dari tabel inventories
func GetInventories(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	// Query untuk mengambil data dari database
	rows, err := db.Query("SELECT * FROM inventories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var inventories []entity.Inventory // Slice untuk menyimpan inventories
	for rows.Next() {
		var id int
		var name, itemCode, description, status string
		var stock int
		// Mendekode hasil query ke variabel
		err := rows.Scan(&id, &name, &itemCode, &stock, &description, &status)
		if err != nil {
			log.Fatal(err)
		}
		inventories = append(inventories, entity.Inventory{
			ID:          id,
			Name:        name,
			ItemCode:    itemCode,
			Stock:       stock,
			Description: description,
			Status:      status,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventories)
}

// Handler untuk mengambil data inventory berdasarkan ID
func GetInventoryByID(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	// Mengambil ID dari parameter URL
	paramID := p.ByName("id")

	// Melakukan query untuk mengambil data inventory berdasarkan ID
	var inventory entity.Inventory // Menggunakan struct dari package entity
	err := db.QueryRow("SELECT * FROM inventories WHERE id=?", paramID).Scan(&inventory.ID, &inventory.Name, &inventory.ItemCode, &inventory.Stock, &inventory.Description, &inventory.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound) // Memberi respons status Not Found jika data tidak ditemukan
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Data inventory tidak ditemukan",
			})
			return
		}
		log.Fatal(err)
	}

	// Mengirim data inventory sebagai respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Data inventory ditemukan",
		"inventory": inventory,
	})
}

// Handler untuk menambah data inventory baru
func AddInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	// Menerima data JSON dari body request
	var newInventory entity.Inventory // Menggunakan struct dari package entity
	err := json.NewDecoder(r.Body).Decode(&newInventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Menjalankan query untuk memasukkan data baru ke database
	result, err := db.Exec("INSERT INTO inventories (Name, ItemCode, Stock, Description, Status) VALUES (?, ?, ?, ?, ?)",
		newInventory.Name, newInventory.ItemCode, newInventory.Stock, newInventory.Description, newInventory.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mendapatkan ID dari data yang baru saja dimasukkan
	newID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirim respons yang berisi pesan dan ID dari data baru
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Data inventory berhasil ditambahkan",
		"new_id":  newID,
	})
}

// Handler untuk memperbarui data inventory berdasarkan ID
func UpdateInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	// Mengambil ID dari parameter URL
	paramID := p.ByName("id")

	// Menerima data JSON dari body request
	var updatedInventory entity.Inventory // Slice untuk menyimpan update inventories
	err := json.NewDecoder(r.Body).Decode(&updatedInventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Menjalankan query untuk memperbarui data inventory di database
	_, err = db.Exec("UPDATE inventories SET Name=?, ItemCode=?, Stock=?, Description=?, Status=? WHERE id=?",
		updatedInventory.Name, updatedInventory.ItemCode, updatedInventory.Stock, updatedInventory.Description, updatedInventory.Status, paramID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Data inventory berhasil diperbarui",
	})
}

// Handler untuk menghapus data inventory berdasarkan ID
func DeleteInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	// Mengambil ID dari parameter URL
	paramID := p.ByName("id")

	// Menjalankan query untuk menghapus data inventory dari database
	_, err := db.Exec("DELETE FROM inventories WHERE id=?", paramID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Data inventory berhasil dihapus",
	})
}
