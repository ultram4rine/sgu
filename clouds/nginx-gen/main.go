package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

const nginxTmpl = `user www-data;
worker_processes auto;
pid /run/nginx.pid;

events {
        worker_connections 768;
        # multi_accept on;
}

http {
	{{- range .Hosts }}
	server {
		server_name {{ .Name }};
		listen      80;
		listen      [::]:80;
		charset     utf-8;
		location / {
			proxy_pass              {{ .Gateway }};
			proxy_set_header        Host $http_host;
			proxy_set_header        X-Real-IP $remote_addr;
			proxy_set_header        X-Forwarder-For $proxy_add_x_forwarded_for;
			proxy_set_header        Upgrade $http_upgrade;
			proxy_set_header        Connection "upgrade";
			proxy_buffers           4 32k;
			client_max_body_size    500m;
			client_body_buffer_size 128k;
		}
	}
	{{- end }}
}`

type NginxConf struct {
	Hosts []NginxServer
}

type NginxServer struct {
	Name    string `json:"name"`
	Gateway string `json:"gateway"`
}

func main() {
	file, err := os.Open("hosts.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var servers []NginxServer
	if err := json.Unmarshal(jsonBytes, &servers); err != nil {
		log.Fatal(err)
	}
	var nc = &NginxConf{
		Hosts: servers,
	}

	tmpl := template.New("nginx")
	tmpl, err = tmpl.Parse(nginxTmpl)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(os.Stdout, nc)
}
