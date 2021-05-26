package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const adminID = "vfYZ35mry2cVqOPNo1xnL1HE0VW5tp7oMX"
const adminNOTE = "inctf{theres_a_part_2_6552637428346}"

var Notes = make(map[string]string)

// Set these common headers for all responses
var headerz = map[string]string{
	"x-permitted-cross-domain-policies": "none",
	"x-xss-protection":                  "1; mode=block",
	"Cross-Origin-Opener-Policy":        "same-origin",
	"x-content-type-options":            "nosniff",
	"X-Frame-Options":                   "DENY",
}

// Prevent XSS on api-endpoints ¬‿¬
var cType = map[string]string{
	// "content-security-policy": "script-src 'strict-dynamic';object-src 'none';base-uri 'none';require-trusted-types-for 'script';default-src 'self';frame-ancestors 'self'",
	"Content-Type": "text/plain",
}

func cookGenerator() string {
	hash := md5.Sum([]byte(string(rand.Intn(30))))
	return hex.EncodeToString((hash)[:])
}

func headerSetter(w http.ResponseWriter, header map[string]string) {
	for k, v := range header {
		w.Header().Set(k, v)
	}
}

func getIDFromCooke(r *http.Request, w http.ResponseWriter) string {
	var cooke, err = r.Cookie("id")
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	var cookeval string
	if err == nil && re.MatchString(cooke.Value) && len(cooke.Value) <= 35 && len(cooke.Value) >= 30 {
		cookeval = cooke.Value
	} else {
		cookeval = cookGenerator()
		c := http.Cookie{
			Name:     "id",
			Value:    cookeval,
			SameSite: 2,
			HttpOnly: true,
			// Secure:   true,
		}
		http.SetCookie(w, &c)
	}
	return cookeval
}

func add(w http.ResponseWriter, r *http.Request) {
	headerSetter(w, headerz)

	id := getIDFromCooke(r, w)
	if id != adminID {
		r.ParseForm()
		noteConte := r.Form.Get("content")
		if len(noteConte) < 75 {
			Notes[id] = noteConte
		}
	}
	fmt.Fprintf(w, "OK")
}

func get(w http.ResponseWriter, r *http.Request) {
	headerSetter(w, headerz)

	id := getIDFromCooke(r, w)
	x := Notes[id]
	headerSetter(w, cType)
	if x == "" {
		fmt.Fprintf(w, "404 No Note Found")
	} else {
		fmt.Fprintf(w, x)
	}
}

func find(w http.ResponseWriter, r *http.Request) {
	headerSetter(w, headerz)

	id := getIDFromCooke(r, w)

	param := r.URL.Query()
	x := Notes[id]

	var which string
	str, err := param["condition"]
	if !err {
		which = "any"
	} else {
		which = str[0]
	}

	var start bool
	str, err = param["startsWith"]
	if !err {
		start = strings.HasPrefix(x, "legen")
	} else {
		start = strings.HasPrefix(x, str[0])
	}
	var responseee string
	var end bool
	str, err = param["endsWith"]
	if !err {
		end = strings.HasSuffix(x, "dary")
	} else {
		end = strings.HasSuffix(x, str[0])
	}

	if which == "starts" && start {
		responseee = x
	} else if which == "ends" && end {
		responseee = x
	} else if which == "both" && (start && end) {
		responseee = x
	} else if which == "any" && (start || end) {
		responseee = x
	} else {
		_, present := param["debug"]

		if present {
			for k, v := range param {
				fmt.Println(k)
				fmt.Println(v)
				for _, d := range v {
					w.Header().Set(k, d)
				}
			}
		}
		fmt.Println("++++++++++++++")
		responseee = "404 No Note Found"
	}
	headerSetter(w, cType)
	fmt.Fprintf(w, responseee)
}

// Reset notes every 30 mins
func resetNotes() {
	Notes[adminID] = adminNOTE
	for range time.Tick(time.Second * 1 * 60 * 30) {
		Notes = make(map[string]string)
		Notes[adminID] = adminNOTE
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var dir string
	flag.StringVar(&dir, "dir", "./public", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	go resetNotes()
	r := mux.NewRouter()
	r.HandleFunc("/add", add).Methods("POST")
	r.HandleFunc("/get", get).Methods("GET")
	r.HandleFunc("/find", find).Methods("GET")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	fmt.Println("Server started at http://0.0.0.0:3000")
	srv := &http.Server{
		Addr: "0.0.0.0:3000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
