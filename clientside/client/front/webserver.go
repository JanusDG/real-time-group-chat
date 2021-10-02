package front

import (
	"fmt"
	// "flag"
	// "log"
	// "net/url"
	// // "os"
	// // "os/signal"
	// // "time"
	// "bufio"
	// "strings"
	"log"
	"net/http"
	// "strconv"
	"html/template"
	// "sync"
	// "github.com/gorilla/mux"


	// "github.com/gorilla/websocket"
	// "github.com/JanusDG/real-time-group-chat/odt"
	// "github.com/satori/go.uuid"
)

type WebServer struct {
	Port int
	UserData map[string]string
	// GetDataFormWg sync.WaitGroup
}

func NewWebServer(port int) *WebServer{
	return &WebServer{Port: port, UserData: make(map[string]string)}
}

func (web *WebServer) MainPage(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
		t, _ := template.ParseFiles("client/front/static/index.html")
		t.Execute(w, nil)
	}
}

func (web *WebServer) GetUserData() map[string]string{
	// web.GetDataFormWg.Add(1)

	// web.GetDataFormWg.Wait()
	return web.UserData
}

func (web *WebServer) LoginUser(w http.ResponseWriter, r *http.Request) {
	
    fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("client/front/static/user_login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		if (r.Form["username"][0] == "" || r.Form["password"][0] == ""){
			//TODO make error here
		}
		fmt.Println("username:", r.Form["username"][0])
		web.UserData["username:"] = r.Form["username"][0]
		fmt.Println("password:", r.Form["password"][0])
		web.UserData["password:"] = r.Form["password"][0]
		t, _ := template.ParseFiles("client/front/static/user_proceed.html")
		t.Execute(w, nil)
	}
	// web.GetDataFormWg.Done()
}



func (web *WebServer) RunWeb() {
	
	http.HandleFunc("/", web.MainPage) // setting router rule
	
    http.HandleFunc("/user_login", web.LoginUser)
	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "hello") })


    err := http.ListenAndServe(":8081", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
