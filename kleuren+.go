package main

import (
    "fmt"
    "os"
    "strings"
)

// De kleuren die de applicatie herkent
var bekendeKleuren = map[string]bool{
    "blauw": true,
    "rood": true,
    "geel": true,
    "groen": true,
    "paars": true,
    "bruin": true,
    "oranje": true,
}

// De tekst gekoppeld aan de kleuren
var kleurTekst = map[string]string{
    "blauw": "Blauw zoals de lucht.",
    "rood": "Rood met passie.",
    "geel": "Geel als de stralen van de zon.",
    "groen": "Groen van de natuur.",
    "paars": "Paars als de avondhemel.",
    "bruin": "Bruin als de aarde.",
    "oranje": "Oranje als een zonsondergang.",
}

func main() {
    // Gebruikersinvoer
    fmt.Print("Voer je lievelingskleur in: ")
    var gekozenKleur string
    fmt.Scanln(&gekozenKleur)
    gekozenKleur = strings.ToLower(gekozenKleur)

    // Controleer of de kleur bekend is
    if _, ok := bekendeKleuren[gekozenKleur]; !ok {
        fmt.Printf("Foutmelding: De kleur '%s' is niet bekend.\n", gekozenKleur)
        os.Exit(1)
    }

    // Schrijf de tekst naar een bestand
    bestand, err := os.OpenFile("favoriet_kleur.txt", os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Foutmelding:", err)
        os.Exit(1)
    }
    defer bestand.Close()

    _, err = bestand.WriteString(kleurTekst[gekozenKleur])
    if err != nil {
        fmt.Println("Foutmelding:", err)
        os.Exit(1)
    }

    // Bevestiging
    fmt.Println("De tekst is geschreven naar het bestand 'favoriet_kleur.txt'.")
}
