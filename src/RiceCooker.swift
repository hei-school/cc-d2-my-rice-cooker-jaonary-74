import Foundation

class RiceCooker {
    var isPowerOn: Bool
    var contents: [String]
    var hasPower: Bool
    
    init() {
        isPowerOn = false
        contents = []
        hasPower = false
    }
    
    func turnOn() {
        if isPowerOn {
            fatalError("La rice cooker est déjà allumée.")
        }
        
        if !hasPower {
            fatalError("La rice cooker ne peut pas être allumée car il n'y a pas de courant.")
        }
        
        isPowerOn = true
        print("La rice cooker est allumée.")
    }
    
    func turnOff() {
        if !isPowerOn {
            fatalError("La rice cooker est déjà éteinte.")
        }
        
        isPowerOn = false
        print("La rice cooker est éteinte.")
    }
    
    func addContent(content: String) {
        if !isPowerOn {
            fatalError("La rice cooker n'est pas allumée. Vous ne pouvez pas ajouter d'aliments.")
        }
        
        contents.append(content)
        print("Vous avez ajouté \"\(content)\" à la rice cooker.")
    }
    
    func cook() {
        if !isPowerOn {
            fatalError("La rice cooker n'est pas allumée. Vous ne pouvez pas faire cuire les aliments.")
        }
        
        print("La rice cooker est en train de cuire les aliments.")
        // Logique de cuisson des aliments
        print("La cuisson est terminée.")
    }
    
    func setPower(hasPower: Bool) {
        self.hasPower = hasPower
    }
}

func getUserInput(promptMessage: String) -> String? {
    print(promptMessage, terminator: "")
    return readLine()
}

func testRiceCooker() {
    let riceCooker = RiceCooker()
    var menu = 0
    
    while menu != 5 {
        print("Veuillez choisir une action")
        print("1) Allumer la rice cooker")
        print("2) Ajouter un aliment")
        print("3) Cuire les aliments")
        print("4) Éteindre la rice cooker")
        print("5) Quitter")
        
        guard let userInput = getUserInput(promptMessage: "Choisissez une action : ") else {
            print("Une erreur s'est produite lors de la lecture de l'entrée utilisateur.")
            continue
        }
        
        menu = 0
        if let input = Int(userInput) {
            menu = input
        }
        
        switch menu {
        case 1:
            guard let hasPowerInput = getUserInput(promptMessage: "Y a-t-il du courant ? (Oui pour 'o', Non pour 'n') ") else {
                print("Une erreur s'est produite lors de la lecture de l'entrée utilisateur.")
                continue
            }
            
            let hasPower = hasPowerInput.lowercased() == "o"
            riceCooker.setPower(hasPower: hasPower)
            
            do {
                try riceCooker.turnOn()
            } catch {
                print("Une erreur s'est produite :", error)
            }
            
        case 2:
            guard let content = getUserInput(promptMessage: "Veuillez ajouter un aliment : ") else {
                print("Une erreur s'est produite lors de la lecture de l'entrée utilisateur.")
                continue
            }
            
            do {
                try riceCooker.addContent(content: content)
            } catch {
                print("Une erreur s'est produite :", error)
            }
            
        case 3:
            do {
                try riceCooker.cook()
            } catch {
                print("Une erreur s'est produite :", error)
            }
            
        case 4:
            do {
                try riceCooker.turnOff()
            } catch {
                print("Une erreur s'est produite :", error)
            }
            
        case 5:
            print("Vous avez quitté le programme")
            
        default:
            print("Veuillez choisir le bon menu")
        }
    }
}