package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	db := InitDB()
	defer db.Close()

	fmt.Println("====== QUIZ ======")

	// Demander le pseudo
	var pseudo string
	fmt.Println("Quel est votre pseudo ?")
	fmt.Scan(&pseudo)

	// Créer ou récupérer le joueur dans la db
	Users := AjouterUnNouveauJoueur(pseudo)

	fmt.Println("\nQue voulez-vous faire ?")
	
	fmt.Println("1 - Commencer une nouvelle partie")
	fmt.Println("2 - Reprendre une partie en cours")
	fmt.Println("3 - Afficher l'historique des parties")
	var choixAction int
	fmt.Scan(&choixAction)

	var idPartie int
	var scorePartie int

	switch choixAction {
	case 1:
		// Nouvelle partie
		_, _ = db.Exec("UPDATE partie SET etat_partie = 0 WHERE id_joueur = ? AND etat_partie = 1", Users.Id_joueur)
		res, err := db.Exec("INSERT INTO partie (id_joueur, score_final, nombre_vies, etat_partie, date_partie) VALUES (?, ?, ?, ?, ?)",
			Users.Id_joueur, 0, 3, true, time.Now())
		if err != nil {
			log.Fatal("Erreur création partie:", err)
		}
		lastID, _ := res.LastInsertId()
		idPartie = int(lastID)
		scorePartie = 0

	case 2:
		// Reprendre partie
		row := db.QueryRow("SELECT id_partie, score_final FROM partie WHERE id_joueur = ? AND etat_partie = 1", Users.Id_joueur)
		err := row.Scan(&idPartie, &scorePartie)
		if err != nil {
			fmt.Println("Aucune partie en cours trouvée. Une nouvelle partie sera créée.")
		// creation d'une nouvelle partie 
			res, _ := db.Exec("INSERT INTO partie (id_joueur, score_final, nombre_vies, etat_partie, date_partie) VALUES (?, ?, ?, ?, ?)",
				Users.Id_joueur, 0, 3, true, time.Now())
		// Recupere l'id de la nouvelle partie du joueur 
			lastID, _ := res.LastInsertId()
			idPartie = int(lastID)
			scorePartie = 0
		}
	default:
		// Historique
		rows, _ := db.Query("SELECT id_partie, score_final, date_partie FROM partie WHERE id_joueur = ?", Users.Id_joueur)
		defer rows.Close()
		fmt.Println("\n=== Historique des parties ===")
		for rows.Next() {
			var id int
			var score int
			var date time.Time
			rows.Scan(&id, &score, &date)
			fmt.Printf("Partie %d - Score: %d - Date: %s\n", id, score, date.Format("02-01-2006 15:04"))
		}
		return
	}

	// Charger tous les thèmes
	themes := AjoutdesThemes()
	// Map pour stocker les thèmes déjà joués
	themesJouer := make(map[int]bool)

	// Boucle principale : tant que tous les thèmes ne sont pas joués
	for len(themesJouer) < len(themes) {

		// Afficher uniquement les thèmes disponibles
		fmt.Println("\n🎯 Thèmes disponibles :")
		for _, t := range themes {
			if !themesJouer[t.Id_theme] {
				fmt.Printf("%d - %s\n", t.Id_theme, t.Nom_theme)
			}
		}

		// Choix du thème
		var choixTheme int
		fmt.Print("\nChoisissez un thème : ")
		fmt.Scan(&choixTheme)

		// Vérifier si déjà joué
		if themesJouer[choixTheme] {
			fmt.Println("Ce thème a déjà été joué.")
			continue
		}

		// Vérifier que le thème existe
		themeValide := false
		for _, t := range themes {
			if t.Id_theme == choixTheme {
				themeValide = true
				fmt.Printf("\nVous avez choisi le thème %s\n", t.Nom_theme)
				break
			}
		}
		if !themeValide {
			fmt.Println("Thème invalide.")
			continue
		}

		// Jouer le thème
		questions := AjoutdesQuestionsParTheme(choixTheme)
		vies := 3

		for i, q := range questions {
			if vies == 0 {
				fmt.Println("💀 Plus de vies pour ce thème. Passage au thème suivant.")
				break
			}

			fmt.Printf("\nQuestion %d : %s\n", i+1, q.Texte_question)
			var choix string
			fmt.Print("Votre réponse (vrai/faux) : ")
			fmt.Scan(&choix)
			choix = strings.ToLower(strings.TrimSpace(choix))
			// Récupere la bonne réponse
			reponses := AjoutdesReponsesParQuestion(q.Id_question)
			if len(reponses) == 0 {
				fmt.Println("Question sans réponse en base.")
				continue
			}

			bonne := (choix == "vrai" && reponses[0].Est_correct) || (choix == "faux" && !reponses[0].Est_correct)
			if bonne {
				fmt.Println("🟢 Bonne réponse ! +300 points")
				scorePartie += 300
			} else {
				vies--
				fmt.Printf("🔴 Mauvaise réponse ! Vies restantes : %d\n", vies)
			}

			// Sauvegarder score
			_, _ = db.Exec("UPDATE partie SET score_final = ? WHERE id_partie = ?", scorePartie, idPartie)
		}

		// Marquer thème comme joué
		themesJouer[choixTheme] = true
		fmt.Println("✅ Thème terminé !")
	}
	
	// Fin de la partie
	_, _ = db.Exec("UPDATE partie SET etat_partie = 0, score_final = ? WHERE id_partie = ?", scorePartie, idPartie)
	fmt.Println("\n🏁 Tous les thèmes ont été joués !")
	fmt.Printf("🏆 Score final : %d points\n", scorePartie)
	    fmt.Println("*******************************************")
        fmt.Println("*                                         *")
        fmt.Println("*        🌟 BRAVO ! TU AS GAGNÉ 🌟       *")
        fmt.Println("*                                         *")
        fmt.Println("*      🎹 LE CLAVIER D'OR EST À TOI ! 🎹  *")
        fmt.Println("*******************************************")
}