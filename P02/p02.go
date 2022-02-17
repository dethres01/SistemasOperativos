package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type Colony struct {
	Name         string `json:"name"`
	Municipality string `json:"municipality"`
}
type PostalCode struct {
	Code     string   `json:"code"`
	Colonies []Colony `json:"colonies"`
}
type State struct {
	Name        string                 `json:"name"`
	PostalCodes map[string]*PostalCode `json:"postal_codes"`
}

var states = make(map[string]State)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		word, err := charmap.Windows1252.NewDecoder().String(scanner.Text())
		if err != nil {
			panic(err)
		}
		words := strings.Split(word, "|")

		state := words[4]

		postalCode := words[0]

		municipality := fmt.Sprintf("%s, %s", words[3], words[4])

		colony := words[1]
		isColony := words[2] == "Colonia"

		colonyObject := Colony{
			Name:         colony,
			Municipality: municipality,
		}
		if isColony {
			if _, ok := states[state]; ok {

				if _, ok := states[state].PostalCodes[postalCode]; ok {

					states[state].PostalCodes[postalCode].Colonies = append(states[state].PostalCodes[postalCode].Colonies, colonyObject)
				} else {

					states[state].PostalCodes[postalCode] = &PostalCode{
						Code:     postalCode,
						Colonies: []Colony{colonyObject},
					}
				}

			} else {

				states[state] = State{
					Name:        state,
					PostalCodes: make(map[string]*PostalCode),
				}

				states[state].PostalCodes[postalCode] = &PostalCode{
					Code:     postalCode,
					Colonies: []Colony{},
				}

				states[state].PostalCodes[postalCode].Colonies = append(states[state].PostalCodes[postalCode].Colonies, colonyObject)
			}
		}
	}

	if _, err := os.Stat("json_files"); os.IsNotExist(err) {
		os.Mkdir("json_files", 0755)
	}
	json_file, err := os.Create("json_files/data.json")
	if err != nil {
		panic(err)
	}
	defer json_file.Close()

	json_data, err := json.MarshalIndent(states, "", "    ")
	if err != nil {
		panic(err)
	}

	json_file.Write(json_data)
	json_file.Close()

	for _, state := range states {
		json_file, err := os.Create("json_files/" + state.Name + ".json")
		if err != nil {
			panic(err)
		}
		defer json_file.Close()

		json_data, err := json.MarshalIndent(state.PostalCodes, "", "    ")
		if err != nil {
			panic(err)
		}

		json_file.Write(json_data)
		json_file.Close()
	}
}
