package repl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationPokemon struct {
	Name string
}

type ResponsePokemon struct {
	Next    string
	Results []LocationPokemon
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	if c.Next != "" {
		url = c.Next
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
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

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	c.Previous = c.Next
	c.Next = data.Next

	fmt.Printf("Next link: %s\n", c.Next)
	fmt.Printf("Previous link: %s\n", c.Previous)

	return nil
}
