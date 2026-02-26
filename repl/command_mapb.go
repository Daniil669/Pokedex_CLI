package repl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func commandMapb(c *config) error {
	if c.Previous == "" {
		return fmt.Errorf("No previous page")
	}
	req, err := http.NewRequest("GET", c.Previous, nil)
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

	c.Next = c.Previous

	u, err := url.Parse(c.Next)
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	params := u.Query()
	offsetValStr := params.Get("offset")
	if offsetValStr == "" {
		return fmt.Errorf("No previous page")
	}

	offsetVal, err := strconv.Atoi(offsetValStr)
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	if offsetVal <= 0 {
		offsetVal = 0
	} else {
		offsetVal -= 20
	}

	c.Previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", offsetVal)

	fmt.Printf("Next link: %s\n", c.Next)
	fmt.Printf("Previous link: %s\n", c.Previous)

	return nil
}
