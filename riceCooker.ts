class RiceCooker {
    private isPowerOn: boolean;
    private contents: string[];
    private hasPower: boolean;
  
    constructor() {
      this.isPowerOn = false;
      this.contents = [];
      this.hasPower = false;
    }
  
    turnOn(): void {
      if (this.isPowerOn) {
        throw new Error("La rice cooker est déjà allumée.");
      }
  
      if (!this.hasPower) {
        throw new Error("La rice cooker ne peut pas être allumée car il n'y a pas de courant.");
      }
  
      this.isPowerOn = true;
      console.log("La rice cooker est allumée.");
    }
  
    turnOff(): void {
      if (!this.isPowerOn) {
        throw new Error("La rice cooker est déjà éteinte.");
      }
      this.isPowerOn = false;
      console.log("La rice cooker est éteinte.");
    }
  
    addContent(content: string): void {
      if (!this.isPowerOn) {
        throw new Error(
          "La rice cooker n'est pas allumée. Vous ne pouvez pas ajouter d'aliments."
        );
      }
      this.contents.push(content);
      console.log(`Vous avez ajouté "${content}" à la rice cooker.`);
    }
  
    cook(): void {
      if (!this.isPowerOn) {
        throw new Error(
          "La rice cooker n'est pas allumée. Vous ne pouvez pas faire cuire les aliments."
        );
      }
      console.log("La rice cooker est en train de cuire les aliments.");
      // Logique de cuisson des aliments
      console.log("La cuisson est terminée.");
    }
  
    setPower(hasPower: boolean): void {
      this.hasPower = hasPower;
    }
  }
  
  async function getUserInput(promptMessage: string): Promise<string | null> {
    const readline = require('readline');
    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout
    });
  
    return new Promise((resolve) => {
      rl.question(promptMessage, (answer) => {
        rl.close();
        resolve(answer);
      });
    });
  }
  
  async function testRiceCooker(): Promise<void> {
    const riceCooker = new RiceCooker();
    let menu = 0;
  
    while (menu !== 5) {
      console.log("Veuillez choisir une action");
      console.log("1) Allumer la rice cooker");
      console.log("2) Ajouter un aliment");
      console.log("3) Cuire les aliments");
      console.log("4) Éteindre la rice cooker");
      console.log("5) Quitter");
  
      const userInput = await getUserInput("Choisissez une action : ");
      if (userInput !== null) {
        menu = parseInt(userInput, 10);
  
        if (menu === 1) {
          const hasPowerInput = await getUserInput("Y a-t-il du courant ? (Oui pour 'o', Non pour 'n') ");
          const hasPower = hasPowerInput === 'o';
          riceCooker.setPower(hasPower);
  
          try {
            riceCooker.turnOn();
          } catch (error) {
            console.log("Une erreur s'est produite : " + error.message);
          }
        } else if (menu === 2) {
          const content = await getUserInput("Veuillez ajouter un aliment : ");
          if (content !== null && content.length !== 0) {
            try {
              riceCooker.addContent(content);
            } catch (error) {
              console.log("Une erreur s'est produite : " + error.message);
            }
          }
        } else if (menu === 3) {
          try {
            riceCooker.cook();
          } catch (error) {
            console.log("Une erreur s'est produite : " + error.message);
          }
        } else if (menu === 4) {
          try {
            riceCooker.turnOff();
          } catch (error) {
            console.log("Une erreur s'est produite : " + error.message);
          }
        } else if (menu === 5) {
          console.log("Vous avez quitté le programme");
        } else {
          console.log("Veuillez choisir le bon menu");
        }
      } else {
        console.log("La valeur de l'entrée utilisateur est null.");
      }
    }
  }
  
  testRiceCooker();