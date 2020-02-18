package main

import (
	"DAO"
	"context"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"phone_add"
	_ "strconv"
	"time"
)

type Phone_List []phone_add.PhoneAdd

//como fazer uma melhor
func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	var out []byte
	switch r.Method {
	case "GET":
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
		break
		//	case "POST":
		//		break
	default:
		log.Println("Call to CustomerHandler not handled due to unrecognized method")
	}
}

func main() {
	if err := DAO.OpenDBConnection(); err != nil {
		log.Fatal("Could not connect to db \n Error desc: %s", err)
	}

	//initialize routes
	r := mux.NewRouter()

	//main path
	staticFileDirectory := http.Dir("fe/")
	staticFileHandler := http.StripPrefix("/fe/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/fe/").Handler(staticFileHandler).Methods("GET")

	//scriptpaths
	r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir("fe/scripts")))).Methods("GET")
	r.Path("/customer").HandlerFunc(CustomerHandler)

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
