package main

import (
	"DAO"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"phone_add"
	_ "strconv"
)
type Phone_List []phone_add.PhoneAdd

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
			if err != nil{
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
	//initialize routes
	
	r := mux.NewRouter()
	
	//main path
	staticFileDirectory := http.Dir("fe/")
	staticFileHandler := http.StripPrefix("/fe/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/fe/").Handler(staticFileHandler).Methods("GET")

	//scriptpaths
	r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir("fe/scripts")))).Methods("GET")
	r.Path("/customer").HandlerFunc(CustomerHandler)

	log.Fatal(http.ListenAndServe(":80", r))
} //*/
