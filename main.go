package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akadotsh/yc-jobs-api/api"
	"github.com/joho/godotenv"
)



func main() {
   
   godotenv.Load()
   port:=os.Getenv("PORT")
   
   if port == "" {
    port = "8080" 
   }

   server:= api.NewServer(port)
   fmt.Println("server running on port",port)
   log.Fatal(server.Start())

}


