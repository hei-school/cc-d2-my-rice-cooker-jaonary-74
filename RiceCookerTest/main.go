package TestRiceCooker

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type RiceCooker struct {
	isPowerOn bool
	contents  []string
	hasPower  bool
}

func NewRiceCooker() *RiceCooker {
	return &RiceCooker{
		isPowerOn: false,
		contents:  make([]string, 0),
		hasPower:  false,
	}
}

func (rc *RiceCooker) TurnOn() {
	if rc.isPowerOn {
		panic("La rice cooker est déjà allumée.")
	}

	if !rc.hasPower {
		panic("La rice cooker ne peut pas être allumée car il n'y a pas de courant.")
	}

	rc.isPowerOn = true
	fmt.Println("La rice cooker est allumée.")
}

func (rc *RiceCooker) TurnOff() {
	if !rc.isPowerOn {
		panic("La rice cooker est déjà éteinte.")
	}

	rc.isPowerOn = false
	fmt.Println("La rice cooker est éteinte.")
}

func (rc *RiceCooker) AddContent(content string) {
	if !rc.isPowerOn {
		panic("La rice cooker n'est pas allumée. Vous ne pouvez pas ajouter d'aliments.")
	}

	rc.contents = append(rc.contents, content)
	fmt.Printf("Vous avez ajouté \"%s\" à la rice cooker.\n", content)
}

func (rc *RiceCooker) Cook() {
	if !rc.isPowerOn {
		panic("La rice cooker n'est pas allumée. Vous ne pouvez pas faire cuire les aliments.")
	}

	fmt.Println("La rice cooker est en train de cuire les aliments.")
	// Logique de cuisson des aliments
	fmt.Println("La cuisson est terminée.")
}

func (rc *RiceCooker) SetPower(hasPower bool) {
	rc.hasPower = hasPower
}

func GetUserInput(promptMessage string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promptMessage)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSuffix(input, "\n")
	return input, nil
}

func TestRiceCooker() {
	riceCooker := NewRiceCooker()
	menu := 0

	for menu != 5 {
		fmt.Println("Veuillez choisir une action")
		fmt.Println("1) Allumer la rice cooker")
		fmt.Println("2) Ajouter un aliment")
		fmt.Println("3) Cuire les aliments")
		fmt.Println("4) Éteindre la rice cooker")
		fmt.Println("5) Quitter")

		userInput, err := GetUserInput("Choisissez une action : ")
		if err != nil {
			fmt.Println("Une erreur s'est produite lors de la lecture de l'entrée utilisateur :", err)
			continue
		}

		menu = 0
		fmt.Sscan(userInput, &menu)

		switch menu {
		case 1:
			hasPowerInput, err := GetUserInput("Y a-t-il du courant ? (Oui pour 'o', Non pour 'n') ")
			if err != nil {
				fmt.Println("Une erreur s'est produite lors de la lecture de l'entrée utilisateur :", err)
				continue
			}

			hasPower := hasPowerInput == "o"
			riceCooker.SetPower(hasPower)

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Une erreur s'est produite :", r)
				}
			}()

			riceCooker.TurnOn()

		case 2:
			content, err := GetUserInput("Veuillez ajouter un aliment : ")
			if err != nil {
				fmt.Println("Une erreur s'est produite lors de la lecture de l'entrée utilisateur :", err)
				continue
			}

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Une erreur s'est produite :", r)
				}
			}()

			riceCooker.AddContent(content)

		case 3:
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Une erreur s'est produite :", r)
				}
			}()

			riceCooker.Cook()

		case 4:
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Une erreur s'est produite :", r)
				}
			}()

			riceCooker.TurnOff()

		case 5:
			fmt.Println("Vous avez quitté le programme")

		default:
			fmt.Println("Veuillez choisir le bon menu")
		}
	}
}