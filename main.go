package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){

  r:= chi.NewRouter();
  r.Use(middleware.Logger)
  r.Get("/",func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
     encoder :=	json.NewEncoder(w);
	 encoder.Encode("Rest API");
  })

   http.ListenAndServe(":4000",r);

}



