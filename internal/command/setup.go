package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Setup struct {
	options Options
}

func (s *Setup) Run(args []string) error {

	if len(args) == 0 {
		return s.CreateSetup(args)
	}

	if len(args) == 1 && args[0] == "get" {
		return s.GetSetup(args)
	}

	return errors.New("bad parameters")
}

func (s *Setup) CreateSetup(args []string) error {

	if !existSetupFile(s.GetCurrentProfile()) {

		if err := createSetupFile(s.GetCurrentProfile()); err != nil {
			return err
		}
	}

	data := getSetupDataFromUser()

	file, _ := os.OpenFile(getSetupFilePath(s.GetCurrentProfile()), os.O_WRONLY, os.ModePerm)

	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(data); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *Setup) GetSetup(args []string) error {

	setup, err := getSetup(s.GetCurrentProfile())

	if err != nil {
		log.Fatal(err)
	}

	keys := []string{
		"host",
		"port",
		"user",
		"password",
		"from",
	}

	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, setup[key])
	}

	return nil
}

func getSetupDataFromUser() map[string]interface{} {

	var host string
	var port int
	var user string
	var password string
	var from string

	fmt.Print("Inserisci l'host: ")
	fmt.Scan(&host)

	fmt.Print("Inserisci la porta: ")
	fmt.Scan(&port)

	fmt.Print("Inserisci l'utente: ")
	fmt.Scan(&user)

	fmt.Print("Inserisci la password: ")
	fmt.Scan(&password)

	fmt.Print("Inserisci il mittente: ")
	fmt.Scan(&from)

	return map[string]interface{}{
		"host":     host,
		"port":     port,
		"user":     user,
		"password": password,
		"from":     from,
	}
}

func (s *Setup) GetCurrentProfile() string {
	return s.options["profile"].(string)
}
