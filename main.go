package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type JsonUserid struct {
	Username string `json:"username"`
	Follower int    `json:"followers"`
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome")
}

func GetByUserid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonUserUserid := vars["userid"]

	jsonUser := map[string]JsonUserid{}

	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonUser)

	result := jsonUser[jsonUserUserid]
	defer jsonFile.Close()
	json.NewEncoder(w).Encode(result)
}

func GetByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonUserUsername := vars["username"]

	jsonUsers := map[string]JsonUserid{}

	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonUsers)

	var result JsonUserid

	for _, v := range jsonUsers {
		if v.Username == jsonUserUsername {
			result = v
		}
	}

	str := []string{"followers: ", strconv.Itoa(result.Follower)}
	res := strings.Join(str, " ")

	defer jsonFile.Close()
	json.NewEncoder(w).Encode(res)
}

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	r.HandleFunc("/", home)
	r.HandleFunc("/{userid}", GetByUserid)
	r.HandleFunc("/follower/{username}", GetByUsername)
	http.Handle("/", r)
	log.Print("Listening on:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
