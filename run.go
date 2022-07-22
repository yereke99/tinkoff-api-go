package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func notifiHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "POST":
		w.WriteHeader(http.StatusOk)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)

		resp["message"] = "OK"

		jsonResp, err := json.marshal(resp)
		if err != nil{
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonResp)
		return
	}
}

func main(){
	handler := http.HandlerFunc(notifiHandler)
	http.Handle("/ok", handler)
	http.ListenAndServe(":8080", nil)
}

/*
func handle_notification(w http.ResponseWriter, r *http.Request){
  switch r.Method{
  case "POST":
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    resp := make(map[string]string)

    resp["message"] = "OK"

    jsonResp, err := json.Marshal(resp)
    if err != nil{
      log.Fatalf("Error happened in JSON marshal. Err: %s", err)
    }

    w.Write(jsonResp)
    return
  }

}


func main(){
  handler := http.HandlerFunc(handle_notification)
  http.Handle("/ok", handler)
  http.ListenAndServe(":8080", nil)
}
*/
