package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type jadwal struct {
	ID int
	Hari int
	Bulan int
	Tahun int
	Jam int
	Tempat string
	Kegiatan string
	Keterangan string
}

type jadwalJS struct {
	ID int `json:"ID, omitempty"`
	Hari int `json:"Hari, omitempty"`
	Bulan int `json:"Bulan, omitempty"`
	Tahun int `json:"Tahun, omitempty"`
	Jam int `json:"Jam, omitempty"`
	Tempat string `json:"Tempat, omitempty"`
	Kegiatan string `json:"Kegiatan, omitempty"`
	Keterangan string `json:"Keterangan, omitempty"`
}

func main() {
	port := 1995
	http.HandleFunc("/jadwal/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				GetAllJadwal(w,r)
				break
			case "POST":
				InsertJadwal(w,r)
				GetAllJadwal(w,r)
				break
			default:
				http.Error(w, "Invalid method.", 405)
				break
		}
	})
	log.Printf("Server starting on port %v\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func GetAllJadwal(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(167.205.67.251:3306)/scheduler")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	jad := jadwal{}
	rows, err := db.Query("SELECT * FROM jadwal")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&jad.ID, &jad.Hari, &jad.Bulan, &jad.Tahun, &jad.Jam, &jad.Tempat, &jad.Kegiatan, &jad.Keterangan)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&jad)
	}
	err = rows.Err()
}

func InsertJadwal (w http.ResponseWriter, r *http.Request) {
	var jad jadwalJS
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&jad)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	db, err := sql.Open("mysql", "root:@tcp(167.205.67.251:3306)/scheduler")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO jadwal (Hari, Bulan, Tahun, Jam, Tempat, Kegiatan, Keterangan) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(jad.Hari, jad.Bulan, jad.Tahun, jad.Jam, jad.Tempat, jad.Kegiatan, jad.Keterangan)
}