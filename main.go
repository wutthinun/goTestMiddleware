package main

import (
	"fmt"
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
)

type LoginMiddleware struct {
}

func (ip *LoginMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	fmt.Println("LoginMiddleware")
	return func(w rest.ResponseWriter, r *rest.Request) {
		fmt.Println("Befor LoginMiddleware")
		handler(w, r)
		fmt.Println("After LoginMiddleware")
	}
}

func main() {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    api.Use(&LoginMiddleware{})

    rount, err := rest.MakeRouter(
    	rest.Post("/",helloApi),
    )
    if err != nil {
    	log.Fatal(err);
    }
    api.SetApp(rount)
    log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func helloApi(w rest.ResponseWriter, r *rest.Request) {
	fmt.Println("MakeHandler")
    w.WriteJson(map[string]string{"Body": "Hello World!"})
}