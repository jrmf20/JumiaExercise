package main

import (
	"DAO"
	"context"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"phone_add"
	_ "strconv"
	"time"
)

type Phone_List []phone_add.PhoneAdd

func GetCustomersHandler(w http.ResponseWriter, r *http.Request) {
	var out []byte
	p_list := []phone_add.PhoneAdd{}
	if custlist, err := DAO.GetAllCustomers(); err == nil {
		for _, cust := range custlist {
			p_list = append(p_list, phone_add.CreatePhoneAdd(cust.ID, cust.Name, cust.Phone))
		}
		out, err = json.Marshal(Phone_List(p_list))
		if err != nil {
			log.Println(err)
		}
	} else {
		out = []byte("error:\"Unnexpected Error\"")
	}
	w.Write(out)
}

func IndexHander(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("fe/index.html")
	if err != nil {
		log.Fatal("Unnexpected Error Getting File")
	} else {
		w.Write(file)
	}
}

func main() {
	if err := DAO.OpenDBConnection(); err != nil {
		log.Fatal("Could not connect to db \n Error desc: %s", err)
	}

	//initialize routes
	r := mux.NewRouter()

	//main path
	//	staticFileDirectory := http.Dir("fe/")
	//	staticFileHandler := http.StripPrefix("/fe/", http.FileServer(staticFileDirectory))
	//	r.PathPrefix("/fe/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/",IndexHander)
	//scriptpaths
	r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir("fe/scripts")))).Methods("GET")
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("fe/js")))).Methods("GET")
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("fe/css")))).Methods("GET")
	r.Path("/customer").HandlerFunc(GetCustomersHandler).Methods("GET")

	server := &http.Server{Addr: ":80", Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("Unnexpected Error: %s", err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	if err := DAO.CloseDBConnection(); err != nil {
		log.Println("Unnexpected error while closing db: %s", err)
	}

	//Wait
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		// handle err
	}
}
