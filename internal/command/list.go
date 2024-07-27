package command

import "fmt"

type ListCommand struct{}

func (l *ListCommand) Run(args []string) error {

	fmt.Println("Listing commands")
	fmt.Println("list -- Lista comandi")
	fmt.Println("setup -- Configura il sistema")
	fmt.Println("setup get -- Recupera configurazione del sistema")
	fmt.Println("sendmail -- Invia una mail")

	return nil
}
