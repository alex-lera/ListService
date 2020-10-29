package main

import (
        "fmt"
        "log"
        "os"
        "net/http"
        "github.com/gorilla/mux"
        _"encoding/json"
        _ "io"
        _ "io/ioutil"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "gopkg.in/yaml.v2"
    ) 

type RequestMessage struct {
    ID int
}

type Config struct {
    Server struct {
        Dbname string `yaml:"dbname"`
    } `yaml:"server"`
    Database struct {
        Username string `yaml:"user"`
        Password string `yaml:"pass"`
    } `yaml:"database"`
}

var cfg Config

func main() {

f, err := os.Open("/opt/config.yaml")
        
    if err != nil {
        return
    }
        
    defer f.Close()

    
    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(&cfg)
    
    if err != nil {
        return
}


router := mux.NewRouter().StrictSlash(true)
router.HandleFunc("/service/v1/cars/{id}", getCar)

log.Fatal(http.ListenAndServe(":8081", router))
}

func getCar(w http.ResponseWriter, r *http.Request) {
    
    if(r.Method != "GET"){
        w.WriteHeader(405)
        return
    }

    vars := mux.Vars(r)

    if vars["id"] == "" {
        w.WriteHeader(422);
        return
    } 

       
    conn, err := sql.Open("mysql", cfg.Database.Username+":"+cfg.Database.Password+"@tcp("+cfg.Server.Dbname+")/cars")

    if err != nil {
        w.WriteHeader(500)
        conn.Close()
        return
    }

    statement, err := conn.Prepare("select * from rentals where id = ?")

    if err != nil {
        w.WriteHeader(500)
        conn.Close()
        return
    }

    rows, err := statement.Query(vars["id"])        

    if rows.Next() == false {
        w.WriteHeader(404)
    } else{

        var id int
        var brand string
        var model string
        var horsepow string

        rows.Scan(&id, &brand, &model, &horsepow)

        w.WriteHeader(200)

        fmt.Fprintln(w, "Hola. ID = ", id, ", Brand =", brand, ", Model =", model, ", Horse Power =", horsepow)

    }
    conn.Close()

}
