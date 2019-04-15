package exe7_11

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type database map[string]dollars
type dollars float64

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range *db {
		fmt.Fprintf(w, "%s: %f\n", item, price)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	val := req.URL.Query().Get("val")
	_, present := (*db)[item]
	if present {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item already exist: %s\n", item)
		return
	}
	parsedVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid parameter for price: %f\n", parsedVal)
		return
	}
	(*db)[item] = dollars(parsedVal)
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, present := (*db)[item]
	if !present {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	delete(*db, item)
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	val := req.URL.Query().Get("val")
	_, present := (*db)[item]
	if !present {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	parsedVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid parameter for price: %f\n", parsedVal)
		return
	}
	(*db)[item] = dollars(parsedVal)
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	conversion := *db
	price, ok := conversion[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	fmt.Fprintf(w, "%f\n", price)
}

func startServer() {
	db := database{"mouse": 100, "laptop": 2000, "headset": 50}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
