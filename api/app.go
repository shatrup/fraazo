package api

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type API struct {
	Cache map[string]string
}
func (a API) LoadCSVFile()  error {
	file, err := ioutil.ReadFile("test.csv")
	if err != nil {
		log.Println("Error while reading file...", err.Error())
		return  err
	}

	reader := csv.NewReader(bytes.NewReader(file))
	all, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading records from csv...", err.Error())
		return  err
	}
	for _, rec := range all {
		a.Cache[rec[0]] = rec[1]
	}
	return nil
}

func (a API) Start() {
	r := mux.NewRouter()

	//defining routes
	r.HandleFunc("/{key}", a.GetKeyAndValue).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8082",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("Listen to sport ",srv)
	log.Fatal(srv.ListenAndServe())
}

func (a API) GetKeyAndValue(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		value := key
		if val, ok := a.Cache[key]; ok {
			value = val
		} else {
			value = "This key is not present in our system."
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
	}
}