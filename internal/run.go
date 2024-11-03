package internal;

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"github.com/go-chi/chi/v5/middleware"
	"flag"
);

func Run() {
	initLogger := flag.Bool("run_logger", false, "Run logger - yes or no");
	flag.Parse();

	r := chi.NewRouter();
	if *initLogger == true {
		r.Use(middleware.Logger);
	};
	
	fileServer := http.FileServer(http.Dir("../static/"))
  r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	
	r.Get("/", Index);

	http.ListenAndServe(":8080", r);
};
