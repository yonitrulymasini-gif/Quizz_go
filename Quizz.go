package main

import (
	"fmt"
	"strings"
)

func main() {
	/*tableau des questions pour chaque thème*/
	questionsalgorithmes := []string{
		"1️ - Un algorithme doit-il toujours avoir un point de départ et un point d’arrivée ? (oui/non)",
		"2️ - Un algorithme peut-il contenir un nombre infini d’étapes ? (oui/non)",
		"3️ - L’ordre des instructions dans un algorithme n’a pas d’importance ? (oui/non)",
		"4 - Un algorithme peut-il être représenté sous forme de diagramme ? (oui/non)",
		"5 - Un algorithme sert à résoudre un problème étape par étape.",
		`BOSS N°1 : Que fait un algorithme ? 
			A. Il exécute des instructions au hasard 
			B. Il décrit une suite d’étapes pour résoudre un problème 
			C. Il affiche toujours des nombres
			D. Il sert uniquement à trier des listes`,
		"6 - Quel est le but principal d'un algorithme ? (trier des données/résoudre un problème)",
		"7 - Un algorithme peut-il être écrit dans n'importe quel langage de programmation ? (oui/non)",
		"8 - Qu'est-ce qu'une boucle dans un algorithme ? (une répétition d'instructions/une condition)",
		"9 - Un algorithme peut-il être optimisé pour améliorer ses performances ? (oui/non)",
		"10 - Quel est le terme utilisé pour décrire la précision d'un algorithme ? (efficacité/correctitude)",
		`BOSS N°2 : Qu'est-ce que la correctitude d'un algorithme ?
			A. La rapidité d'exécution de l'algorithme
			B. La capacité de l'algorithme à gérer les erreurs
			C. La garantie que l'algorithme produit les résultats attendus
			D. La facilité de compréhension de l'algorithme`,
		"11 - Quel est le terme utilisé pour décrire la rapidité d'un algorithme ? (efficacité/correctitude)",
		"12 - Un algorithme de tri organise les données dans un ordre spécifique. (oui/non)",
		"13 - Quel est le nom de l'algorithme de tri qui divise une liste en sous-listes plus petites ? (tri à bulles/tri rapide)",
		"14 - Un algorithme de recherche binaire fonctionne-t-il sur des listes triées ? (oui/non)",
		"15 - Quel est le nom de l'algorithme qui trouve le plus court chemin dans un graphe pondéré ? (Dijkstra/A*)",
		`BOSS N°3 : Quel algorithme est couramment utilisé pour le chiffrement des données ?
			A. RSA
			B. AES
			C. MD5
			D. SHA-256`,
		"16 - Quel est le nom de l'algorithme qui trie les éléments en les comparant deux par deux ? (tri à bulles/tri rapide)",
		"17 - Un algorithme glouton prend-il toujours la meilleure décision à chaque étape ? (oui/non)",
		"18 - Quel est le nom de l'algorithme qui utilise la technique de diviser pour régner ? (tri fusion/tri par insertion)",
		"19 - Un algorithme récursif s'appelle-t-il lui-même pour résoudre un problème ? (oui/non)",
		"20 - Quel est le nom de l'algorithme qui trouve le plus grand commun diviseur (PGCD) de deux nombres ? (Euclide/Babylonien)",
		`BOSS N°4 : Quel algorithme est utilisé pour compresser les fichiers sans perte de données ?
			A. LZW
			B. Huffman
			C. RLE
			D. JPEG`,

	}
	questionsinformatique := []string{
		"1 - Quel est le composant principal d'un ordinateur ? (processeur/mémoire)",
		"2 - Le HTML est un langage de programmation ? (oui/non)",
		"3 - Qu'est-ce que le 'backend' ? (partie serveur/partie client)",
	}
	questionslogique := []string{
		"1️ - Un algorithme doit-il toujours avoir un point de départ et un point d’arrivée ? (oui/non)",
		"2️ - Un algorithme peut-il contenir un nombre infini d’étapes ? (oui/non)",
		"3️ - L’ordre des instructions dans un algorithme n’a pas d’importance ? (oui/non)",
	}
	questionsculturegenerale := []string{
		"1 - Quelle est la capitale de la France ?",
		"2 - Qui a peint la Joconde ?",
		"3 - Quel est le plus grand océan du monde ?",
	}
	questionsanglais := []string{
		"1 - Quelle est la traduction de 'Bonjour' en anglais ?",
		"2 - Comment dit-on 'Merci' en anglais ?",
		"3 - Quelle est la traduction de 'Au revoir' en anglais ?",
	}

	/*tableau des réponses pour chaque thème*/
	reponsesalgorithmes := []string{"oui", "non", "non", "oui", "oui", "b",
	"résoudre un problème", "oui", "une répétition d'instructions", "oui", "correctitude","c",
	"efficacité","oui", "tri rapide", "oui", "dijkstra", "b",
	"tri à bulles", "oui", "tri fusion", "oui", "euclide", "a"}
	reponsesinformatique := []string{"processeur", "non", "partie serveur"}
	reponseslogique := []string{"oui", "non", "non"}
	reponsesculturegenerale := []string{"Paris", "Leonard de Vinci", "Pacifique"}
	reponsesanglais := []string{"Hello", "Thank you", "Goodbye"}

	var nomdujoueur string
	/*Accueil et choix du thème*/
	fmt.Println("Salut, quel est ton pseudo ?")
	fmt.Scan(&nomdujoueur)
	fmt.Println("\nBienvenue " + nomdujoueur + ", choisis un thème pour ton quizz :\n")

		var choix string
		fmt.Println("--- Thèmes ---")
		fmt.Println("1 - Algorithmes")
		fmt.Println("2 - Métiers de l'informatique")
		fmt.Println("3 - Logique")
		fmt.Println("4 - Culture générale")
		fmt.Println("5 - Anglais")
		fmt.Scan(&choix)
		switch choix {
		/* Boucle pour le thème Algorithmes */
		case "1":
			fmt.Printf("\nTu as choisi le thème Algorithmes.\n\n")
			var reponse_utilisateur string
			for i := 0; i < len(questionsalgorithmes); i++ {
				for {
					fmt.Println(questionsalgorithmes[i])
					fmt.Scan(&reponse_utilisateur)
					reponse_utilisateur = strings.ToLower(reponse_utilisateur)
					
					if reponse_utilisateur == reponsesalgorithmes[i] {
						fmt.Println("Bonne réponse !\n")
						break
					} else {
						fmt.Println("Mauvaise réponse, essaie encore...\n")
					}
				}
			}
		/* Boucle pour le thème Métiers de l'informatique */
		case "2":
			fmt.Println("\nTu as choisi le thème Métiers de l'informatique.\n")
			var reponse_utilisateur string
			for i := 0; i < len(questionsinformatique); i++ {
				for {
					fmt.Println(questionsinformatique[i])
					fmt.Scan(&reponse_utilisateur)
					reponse_utilisateur = strings.ToLower(reponse_utilisateur)

					if reponse_utilisateur == reponsesinformatique[i] {
						fmt.Println("Bonne réponse !\n")
						break
					} else {
						fmt.Println("Mauvaise réponse, essaie encore...\n")
					}
				}
			}
		/* Boucle pour le thème Logique */
		case "3":
			fmt.Println("\nTu as choisi le thème Logique.\n")
			var reponse_utilisateur string
			for i := 0; i < len(questionslogique); i++ {
				for {
					fmt.Println(questionslogique[i])
					fmt.Scan(&reponse_utilisateur)
					reponse_utilisateur = strings.ToLower(reponse_utilisateur)
					
					if reponse_utilisateur == reponseslogique[i] {
						fmt.Println("Bonne réponse !\n")
						break
					} else {
						fmt.Println("Mauvaise réponse, essaie encore...\n")
					}
				}
			}
		/* Boucle pour le thème Culture générale */
		case "4":
			fmt.Println("\nTu as choisi le thème Culture générale.\n")
			var reponse_utilisateur string
			for i := 0; i < len(questionsculturegenerale); i++ {
				for {
					fmt.Println(questionsculturegenerale[i])
					fmt.Scan(&reponse_utilisateur)
					reponse_utilisateur = strings.ToLower(reponse_utilisateur)

					if reponse_utilisateur == reponsesculturegenerale[i] {
						fmt.Println("Bonne réponse !\n")
						break
					} else {
						fmt.Println("Mauvaise réponse, essaie encore...\n")
					}
				}
			}
		/* Boucle pour le thème Anglais */
		case "5":
			fmt.Println("\nTu as choisi le thème Anglais.\n")
			var reponse_utilisateur string
			for i := 0; i < len(questionsanglais); i++ {
				for {
					fmt.Println(questionsanglais[i])
					fmt.Scan(&reponse_utilisateur)
					reponse_utilisateur = strings.ToLower(reponse_utilisateur)

					if reponse_utilisateur == reponsesanglais[i] {
						fmt.Println("Bonne réponse !\n")
						break
					} else {
						fmt.Println("Mauvaise réponse, essaie encore...\n")
					}
				}
			}
		/* Message d'erreur si le thème n'existe pas */
		default:
			fmt.Println("\nCe thème n'existe pas, choisis-en un parmi la liste.\n")

	}

	fmt.Println("Merci d'avoir joué, à bientôt " + nomdujoueur + " !")
}