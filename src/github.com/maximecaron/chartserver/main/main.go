package main

import (
        _ "github.com/lib/pq"
        "database/sql"
        "github.com/ant0ine/go-json-rest"
        "net/http"
        "log"
)

func GetVersion(w *rest.ResponseWriter, req *rest.Request) {
        // Open does not directly open a database connection: this is deferred until a query is made
        db, err := sql.Open("postgres", "sslmode=disable")
        defer db.Close()
        if err != nil {
            log.Fatal(err)
        }
    
        row := db.QueryRow("SELECT version();")
        var value string
        err2 := row.Scan(&value)
        if err2 != nil {
            log.Fatal(err)
        }
        w.WriteJson(&value)
}

type User struct {
        Id   string
        Name string
}

func GetUser(w *rest.ResponseWriter, req *rest.Request) {
        user := User{
                Id:   req.PathParam("id"),
                Name: "Antoine",
        }
        w.WriteJson(&user)
}

func main() {
        handler := rest.ResourceHandler{}
        handler.SetRoutes(
                rest.Route{"GET", "/users/:id", GetUser},
                rest.Route{"GET", "/version/", GetVersion},
        )
        http.ListenAndServe(":8080", &handler)
}