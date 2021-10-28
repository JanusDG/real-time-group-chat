package front

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"html/template"
	"sync"
)

type WebServer struct {
	Port int
	UserData map[string]string
	Err error
	WgAuth *sync.WaitGroup
	Contacts *[]string
	WgContacts *sync.WaitGroup
	WgSendMessage *sync.WaitGroup
	WgReadMessage *sync.WaitGroup
	MessageData map[string]string
	RecieveData *map[string]string

	// GetDataFormWg sync.WaitGroup
}

func NewWebServer(port int, wgA *sync.WaitGroup, wgC *sync.WaitGroup, wgS *sync.WaitGroup, wgR *sync.WaitGroup) *WebServer{
	return &WebServer{Port: port, UserData: make(map[string]string), MessageData: make(map[string]string), WgAuth: wgA, WgContacts: wgC, WgSendMessage: wgS, WgReadMessage:wgR}
}

func (web *WebServer) MainPage(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
		t, _ := template.ParseFiles("client/front/static/index.html")
		t.Execute(w, nil)
	}
}

func (web *WebServer) GetUserData() map[string]string{
	return web.UserData
}

func (web *WebServer) GetMessageData() map[string]string{
	web.WgSendMessage.Wait()
	return web.MessageData
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
		web.UserData[r.Form["username"][0]] = r.Form["password"][0]
		web.WgAuth.Done()
		fmt.Println("password:", r.Form["password"][0])
		t, _ := template.ParseFiles("client/front/static/user_proceed.html")
		t.Execute(w, nil)
	}
}

type Context struct {
    Title string
    Contacts []string
}

type Message struct {
	From string
	Content string

}

func (web *WebServer) UserChats(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "GET" {
		w.Header().Add("Content Type", " text/html")
		t, _ := template.ParseFiles("client/front/static/user_chats.html")
		web.WgContacts.Wait()
        context := Context{
			Title: "Title",
            Contacts: *web.Contacts,
        }
		t.Execute(w, context)


	} else {
		r.ParseForm()
		fmt.Println("to_user:", r.Form["to_user"][0])
		fmt.Println("message:", r.Form["message"][0])
		web.MessageData[r.Form["to_user"][0]] = r.Form["message"][0]
		web.WgSendMessage.Done()
	}
}

func (web *WebServer) UserMessages(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "GET" {
		w.Header().Add("Content Type", " text/html")
		t, _ := template.ParseFiles("client/front/static/user_messages.html")
		// web.WgContacts.Wait()
		var m map[string]string
		m = *web.RecieveData
        context := Message{
			From: m["from"],
            Content: m["content"],
        }
		t.Execute(w, context)
	}
}

func (web *WebServer) RunWeb() {
	http.Handle("/", http.FileServer(http.Dir("./client/front/static")))
    http.HandleFunc("/user_login", web.LoginUser)
    http.HandleFunc("/user_chats", web.UserChats)
    http.HandleFunc("/user_messages", web.UserMessages)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(web.Port), nil)) // setting listening port
}
