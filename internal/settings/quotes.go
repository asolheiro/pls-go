package settings

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
)


type Quote struct {
    ID          string   `json:"_id"`
    Content     string   `json:"content"`
    Author      string   `json:"author"`
    Tags        []string `json:"tags"`
    AuthorID    string   `json:"authorId"`
    AuthorSlug  string   `json:"authorSlug"`
    Length      int      `json:"length"`
}

func GetRandQuote() (Quote, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return Quote{}, fmt.Errorf("failed to get current file path")
	}
	
	location := filepath.Dir(filename)
	
	quotesFilePath := filepath.Join(location, "quotes.json")
	quotesFile, err := os.Open(quotesFilePath)
	if err != nil {
		return Quote{}, err
	}
	defer quotesFile.Close()
	
	var listOfQuotes []Quote
	if err := json.NewDecoder(quotesFile).Decode(&listOfQuotes); err != nil {
		return Quote{}, err
	}
	
	if len(listOfQuotes) == 0 {
		return Quote{}, fmt.Errorf("no quotes found")
	}
	
	return listOfQuotes[rand.Intn(len(listOfQuotes))], nil
}