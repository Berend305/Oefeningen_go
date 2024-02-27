package main

import (
    "fmt"
    "os"
    "strings"
)

var bekendeKleuren = map[string]bool{
    "blauw": true,
    "rood": true,
    "geel": true,
    "groen": true,
    "paars": true,
    "bruin": true,
    "oranje": true,
}

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
    fmt.Scanf("%s", &gekozenKleur)
    gekozenKleur = strings.ToLower(gekozenKleur)

    // Controleer of de kleur bekend is
    if _, ok := bekendeKleuren[gekozenKleur]; !ok {
        fmt.Printf("Foutmelding: De kleur '%s' is niet bekend.\n", gekozenKleur)
        os.Exit(1)
    }

    // Toon de tekst in de terminal
    fmt.Println(kleurTekst[gekozenKleur])
}
