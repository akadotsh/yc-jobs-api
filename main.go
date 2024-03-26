package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Job struct {
 name string;
 logo string;
 role string
//  location string
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

		post.name = e.ChildText("span.block.font-bold");
        post.logo = e.ChildAttr("img", "src");
		post.role = e.ChildText("div>a.font-semibold.text-linkColor")
		
		// fmt.Println("location",e.ChildText("div.flex.flex-row.flex-wrap.justify-center"))

        ycJobs = append(ycJobs, post)
	})
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})
	
	if err := c.Visit("https://www.ycombinator.com/jobs"); err != nil {
		log.Fatal("Error visiting initial URL:", err)
	}

	fmt.Println(ycJobs)
}


