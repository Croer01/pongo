package main

import (
	"net/http"

	"fmt"
	"golang.org/x/net/websocket"
	"pongo/communication/server"
	"encoding/json"
	"io/ioutil"
)

func newRestServer() *server.RestServer {
	route := server.Route{
		Url:"/login",
		Method:"POST",
		Handler:func(httpResponse http.ResponseWriter, httpRequest *http.Request, session *server.Session) {
			var val map[string]string
			body,_ := ioutil.ReadAll(httpRequest.Body)
			json.Unmarshal(body, &val)
			session.Nick = val["nick"]
			fmt.Printf("you login with %s\n", session.Nick)
		},
	}

	return &server.RestServer{Root:"/api",Routes:[]*server.Route{&route}}
}

// This example demonstrates a trivial echo server.
func main() {
	hub := server.NewHub();
	authentication := server.NewBasicAuthentication()
	//registre handlers
	http.Handle("/", &server.StaticHttpHandler{Root:"webapp"})
	http.Handle("/ws", websocket.Handler(authentication.WrapWsHandler(hub.RegisterConnection)))
	http.HandleFunc("/api/", authentication.WrapHandler(newRestServer().HandleRequest))

	fmt.Println("open server...")
	go hub.Run()
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
