package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Job struct {
	Name string `json:"name"`;
	Logo string `json:"logo"`;
	Role string `json:"role"`
    Location string `json:"location"`
   }
   

var ycJobs []Job

func main() {

	// r:= chi.NewRouter();
	// r.Use(middleware.Logger)
	// r.Get("/",func(w http.ResponseWriter, r *http.Request){
	//   w.Header().Set("Content-Type","application/json");
	//    encoder :=	json.NewEncoder(w);
	//    encoder.Encode("Rest API");
	// })
  
	//  http.ListenAndServe(":4000",r);

	c := colly.NewCollector(
		colly.AllowedDomains("www.ycombinator.com","www.ycombinator.com/jobs"),
	)
	
	

	c.OnHTML("ul.space-y-2.overflow-hidden > li", func(e *colly.HTMLElement) {
		post := Job{};

		post.Name = e.ChildText("span.block.font-bold");
        post.Logo = e.ChildAttr("img", "src");
		post.Role = e.ChildText("div>a.font-semibold.text-linkColor")
		post.Location = e.ChildText("div.flex.flex-row.flex-wrap.justify-center > :nth-child(2)")

        ycJobs = append(ycJobs, post)
	})
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})
	
	if err := c.Visit("https://www.ycombinator.com/jobs"); err != nil {
		log.Fatal("Error visiting initial URL:", err)
	}

	_j,_:= json.MarshalIndent(ycJobs, "", "  ")

     fmt.Println(string(_j))  
}


