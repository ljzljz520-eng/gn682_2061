package main

import (
	"inspectionbase/internal/api"
	"inspectionbase/internal/audit"
	"inspectionbase/internal/config"
	"inspectionbase/internal/inspection"
	"inspectionbase/internal/store"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	st, e := store.Open(cfg.DSN())
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	au := audit.New(st)
	svc := inspection.New(st, au.Record)
	log.Printf("inspection service listening on %s", cfg.Addr)
	log.Fatal(http.ListenAndServe(cfg.Addr, api.New(svc).Routes()))
}
