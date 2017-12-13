package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	ID int `json: "ID, omitempty"`
	Hari int `json: "Hari, omitempty"`
	Bulan int `json: "Bulan, omitempty"`
	Tahun int `json: "Tahun, omitempty"`
	Jam int `json: "Jam, omitempty"`
	Tempat string `json: "Tempat, omitempty"`
	Kegiatan string `json: "Kegiatan, omitempty"`
	Keterangan string `json: "Keterangan, omitempty"`
}

func main() {
	port := 8080
	
	http.HandleFunc("/lihat/", func(w http.ResponseWriter, r *http.Request) {
		GetAllJadwal(w,r)
	})
	
	http.HandleFunc("/tambah/", func(w http.ResponseWriter, r *http.Request) {
		InsertJadwal(w,r)
		GetAllJadwal(w,r)
	})
	
	http.HandleFunc("/hapus/", func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Path[len("/delete/"):]
		DeleteJadwal(w,r,s)
		GetAllJadwal(w,r)
	})
}

func GetAllJadwal(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/schedule")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	jad := jadwal{}
	
	rows, err := db.Query("SELECT ID, Hari, Bulan, Tahun, Jam, Tempat, Kegiatan, Keterangan FROM jadwal")
	
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&jad.ID, &jad.Hari, &jad.Bulan, &jad.Bulan, &jad.Tahun, &jad.Jam, &jad.Tempat, &jad.Kegiatan, &jad.Keterangan)
		
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
	
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/scheduler")
	if err != nil {
		log.Fatal(err)
	}
	
	stmt, err := db.Prepare("INSERT INTO jadwal (ID, Hari, Bulan, Tahun, Jam, Tempat, Kegiatan, Keterangan VALUES (?,?,?,?,?)")
	if err != inl {
		log.Fatal(err)
	}
	_, err = stmt.Exec(jad.ID, jad.Hari, jad.Bulan, jad.Tahun, jad.Jam, jad.Tempat, jad.Kegiatan, jad.Keterangan)
}

func DeleteJadwal (w http.ResponseWriter, r *http.Request, id string) {
	idjadwal, _ := strconv.Atoi(id)
	
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306/scheduler")
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := db.Query("DELETE FROM jadwal WHERE ID=?", idjadwal)
	defer rows.Close()
}
