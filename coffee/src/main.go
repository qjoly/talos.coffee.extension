package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

type PageData struct {
	Hostname string
	IP       string
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("Error getting hostname: %v", err)
			hostname = "Unknown"
		}

		addrs, err := net.InterfaceAddrs()
		var ip string
		if err == nil {
			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						ip = ipNet.IP.String()
						break
					}
				}
			}
		}
		if ip == "" {
			ip = "Unknown"
		}

		tmpl, err := template.ParseFiles(filepath.Join("tmpl", "index.tmpl"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}

		data := PageData{
			Hostname: hostname,
			IP:       ip,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
		}
	})

	log.Println("Server started on :1337")
	log.Fatal(http.ListenAndServe(":1337", nil))
}
