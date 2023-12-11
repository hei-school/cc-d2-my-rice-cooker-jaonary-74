#include <iostream>
#include <vector>
#include <string>

class RiceCooker {
private:
    bool isPowerOn;
    std::vector<std::string> contents;
    bool hasPower;

public:
    RiceCooker() : isPowerOn(false), hasPower(false) {}

    void turnOn() {
        if (isPowerOn) {
            throw std::runtime_error("La rice cooker est déjà allumée.");
        }

        if (!hasPower) {
            throw std::runtime_error("La rice cooker ne peut pas être allumée car il n'y a pas de courant.");
        }

        isPowerOn = true;
        std::cout << "La rice cooker est allumée." << std::endl;
    }

    void turnOff() {
        if (!isPowerOn) {
            throw std::runtime_error("La rice cooker est déjà éteinte.");
        }

        isPowerOn = false;
        std::cout << "La rice cooker est éteinte." << std::endl;
    }

    void addContent(const std::string& content) {
        if (!isPowerOn) {
            throw std::runtime_error("La rice cooker n'est pas allumée. Vous ne pouvez pas ajouter d'aliments.");
        }

        contents.push_back(content);
        std::cout << "Vous avez ajouté \"" << content << "\" à la rice cooker." << std::endl;
    }

    void cook() {
        if (!isPowerOn) {
            throw std::runtime_error("La rice cooker n'est pas allumée. Vous ne pouvez pas faire cuire les aliments.");
        }

        std::cout << "La rice cooker est en train de cuire les aliments." << std::endl;
        // Logique de cuisson des aliments
        std::cout << "La cuisson est terminée." << std::endl;
    }

    void setPower(bool hasPower) {
        this->hasPower = hasPower;
    }
};

std::string getUserInput(const std::string& promptMessage) {
    std::string input;
    std::cout << promptMessage;
    std::getline(std::cin, input);
    return input;
}

void testRiceCooker() {
    RiceCooker riceCooker;
    int menu = 0;

    while (menu != 5) {
        std::cout << "Veuillez choisir une action" << std::endl;
        std::cout << "1) Allumer la rice cooker" << std::endl;
        std::cout << "2) Ajouter un aliment" << std::endl;
        std::cout << "3) Cuire les aliments" << std::endl;
        std::cout << "4) Éteindre la rice cooker" << std::endl;
        std::cout << "5) Quitter" << std::endl;

        std::string userInput = getUserInput("Choisissez une action : ");

        menu = 0;
        try {
            menu = std::stoi(userInput);
        } catch (const std::invalid_argument&) {
            std::cout << "Veuillez entrer un numéro valide." << std::endl;
            continue;
        }

        switch (menu) {
            case 1: {
                std::string hasPowerInput = getUserInput("Y a-t-il du courant ? (Oui pour 'o', Non pour 'n') ");
                bool hasPower = (hasPowerInput == "o");
                riceCooker.setPower(hasPower);

                try {
                    riceCooker.turnOn();
                } catch (const std::runtime_error& e) {
                    std::cout << "Une erreur s'est produite : " << e.what() << std::endl;
                }
                break;
            }

            case 2: {
                std::string content = getUserInput("Veuillez ajouter un aliment : ");

                try {
                    riceCooker.addContent(content);
                } catch (const std::runtime_error& e) {
                    std::cout << "Une erreur s'est produite : " << e.what() << std::endl;
                }
                break;
            }

            case 3: {
                try {
                    riceCooker.cook();
                } catch (const std::runtime_error& e) {
                    std::cout << "Une erreur s'est produite : " << e.what() << std::endl;
                }
                break;
            }

            case 4: {
                try {
                    riceCooker.turnOff();
                } catch (const std::runtime_error& e) {
                    std::cout << "Une erreur s'est produite : " << e.what() << std::endl;
                }
                break;
            }

            case 5:
                std::cout << "Vousavez quitté le programme" << std::endl;
                break;

            default:
                std::cout << "Veuillez choisir le bon menu" << std::endl;
                break;
        }
    }
}