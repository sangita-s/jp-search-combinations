package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	pipe "github.com/b4b4r07/go-pipe"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/tomoemon/text_normalizer"
)

type SearchResult struct {
	KeywordPattern string `json:"keyword_pattern"`
	Priority       int8   `json:"priority"`
	IsDefault      bool   `json:"is_default"`
	Comment        string `json:"comment"`
}

func main() {
	err := godotenv.Load()
	// don't start the system if environment file is not present
	if err != nil {
		log.Println("Unable to load environment file.")
		panic(err)
	}

	r := chi.NewRouter()

	//CORS Handler
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up and running..."))
	})

	r.Route("/results", func(r chi.Router) {
		r.Get("/", GetResults)
	})

	log.Println("Server started listening on port: " + os.Getenv("SERVER_PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("SERVER_PORT"), r)
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	// search keyword provided by the user
	if r.URL.Query().Get("q") == "" {
		// TODO: if query keyword is empty return error response
	}
	q := r.URL.Query().Get("q")

	n, _ := strconv.Atoi(os.Getenv("NUMBER_OF_KANJI_VARIATION_ATTEMPTS"))
	// number of variations of kanji convertions to attempt against keywords
	if r.URL.Query().Get("n") != "" {
		var err error
		n, err = strconv.Atoi(r.URL.Query().Get("n"))
		// if n is greater than 6, assign default value only to prevent stress on system
		if err != nil || n > 6 {
			n, _ = strconv.Atoi(os.Getenv("NUMBER_OF_KANJI_VARIATION_ATTEMPTS"))
		}
	}

	var results []SearchResult
	results = append(results, SearchResult{KeywordPattern: q, Priority: 1, IsDefault: true, Comment: "Default keywords provided by user."})

	// Convert search katakana keywords to hiragana. Kanjis remain as-is
	normalizer := text_normalizer.NewTextNormalizer(text_normalizer.KatakanaToHiragana)
	hiragana_keywords := normalizer.Replace(q)
	results = append(results, SearchResult{KeywordPattern: hiragana_keywords, Priority: 2, IsDefault: false, Comment: "Katakana keywords converted to hiragana"})

	// Convert user keywords to Kanjis and add top N results to results array
	get_mecab_results(hiragana_keywords, n, &results)

	result, err := json.Marshal(results)
	if err != nil {
		//TODO return error response if unable to marshal the results
		panic(err)
	}

	w.Write([]byte(result))
}

func get_mecab_results(keywords string, number_of_results int, results *[]SearchResult) {
	dictionary_path_cmd := exec.Command("mecab-config", "--dicdir")
	dictionary_path, err := dictionary_path_cmd.Output()
	if err != nil {
		//TODO: Log this error and return error response
		log.Println(err.Error())
		return
	}
	dictionary_path_kkc := strings.TrimSuffix(string(dictionary_path), "\n") + "/mecab-as-kkc"

	var b bytes.Buffer
	pipe.Command(&b,
		exec.Command("echo", keywords),
		exec.Command("mecab", "-d", dictionary_path_kkc, "-N", strconv.Itoa(number_of_results)),
	)

	lines := strings.Split(b.String(), "\n")
	// save results in log also, for validation
	log.Println(lines)
	for _, line := range lines {
		if line != "" {
			*results = append(*results, SearchResult{KeywordPattern: line, Priority: 3, IsDefault: false, Comment: "User keywords converted to Kanjis"})
		}
	}
}
