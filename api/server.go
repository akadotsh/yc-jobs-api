package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gocolly/colly"
)

type Server struct {
	listenAddr string
}


type Job struct {
	Name string `json:"name"`;
	Logo string `json:"logo"`;
	Role string `json:"role"`
    Location string `json:"location"`
   }
   
   

var ycJobs []Job

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	router:= chi.NewRouter();
	router.Use(middleware.Logger);

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}));

	router.Get("/latest",s.handleGetLatestJobs)

    return http.ListenAndServe(s.listenAddr,router)
}


func (s *Server) handleGetLatestJobs(w http.ResponseWriter, r * http.Request){
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ycJobs)
}