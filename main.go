package main 

import (
	"fmt"
)

func main() {
	// --- Connexion DB ---
	db := InitDB()
	defer db.Close()

	fmt.Println("=== QUIZ ===")

	// --- Demander pseudo ---
	var pseudo string
	fmt.Print("Quel est ton pseudo ? ")
	fmt.Scan(&pseudo)

	Users := AjouterOuModifierUnNouveauJoueur(pseudo)

	fmt.Printf("Bienvenue %s\n", Users.Username)
}
	/*themes := LoadThemes(db)

	fmt.Printf("Bienvenue, %s !\n", player.Name)

	fmt.Printf("Thèmes disponibles :\n")
	for _, theme := range themes {
		fmt.Printf("- %s\n", theme.Nom)
	}

	questions := LoadQuestionsForTheme(db, themes[0].ID, 5)

	fmt.Printf("Questions chargées pour le thème %s :\n", themes[0].Nom)
	for _, q := range questions {
		fmt.Printf("- %s\n", q.Text)
	}
}*/
