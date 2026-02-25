package repl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationPokemon struct {
	Id   int
	Name string
}

type ResponsePokemon struct {
	Next    string
	Results []LocationPokemon
}

func commandMap() error {
	const mapUrl string = "https://pokeapi.co/api/v2/location-area/"
	req, err := http.NewRequest("GET", mapUrl, nil)
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	req.Header.Set("content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	defer res.Body.Close()

	var data ResponsePokemon

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	fmt.Printf("Next link: %s", data.Next)

	return nil
}
