package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Users struct {
	Id_joueur int
	Username string
	Date_inscription time.Time
}

type Partie struct {
	Id_partie int
	Id_joueur int
	Score_final int
	Date_partie time.Time
	Nombre_vies int
	Etat_partie bool
}

type Theme struct {
	Id_theme int
	Nom_theme string
}

type Question struct {
	Id_question int
	Texte_question string
	Niveau_difficulte string
	Id_theme int
}

type Reponse_joueur struct {
	Id_reponsejoueur int
	Id_question int
	Texte_reponsejoueur string
	Id_joueur int
	Est_correct bool
}

type Reponse_question struct {
	Id_reponsequestion int
	Id_question int
	Texte_reponsequestion string
	Est_correct bool
}

type Historique struct {
	Id_historique int
	Score_enregistre int
	Id_partie int
	Date_sauvegarde time.Time
}

// ----- InitDB (exportée) -----
func InitDB() *sql.DB {

	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "root"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306" // Adresse à changer avec l'adresse de votre serveur MariaDB
	cfg.DBName = "clavier_dor"              // Nom à changer avec le nom de votre BDD

	// Vérifier la connexion
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}
	fmt.Println("Connexion à la base de données réussie.")
	return db
}

// ----- Créer ou charger un joueur (exportée) -----
func AjouterUnNouveauJoueur(name string) Users {
	var joueur Users
	err := db.QueryRow("SELECT id_joueur, username FROM users WHERE username = ?", name).Scan(&joueur.Id_joueur, &joueur.Username)
	if err == sql.ErrNoRows {
		// Créer un nouveau joueur
		res, err := db.Exec("INSERT INTO users (username, date_inscription) VALUES (?, ?)", name, time.Now())
		if err != nil {
			log.Fatal("Erreur lors de la création du joueur:", err)
		}
		newID, _ := res.LastInsertId()
		joueur = Users{Id_joueur: int(newID), Username: name}
	} else if err != nil {
		log.Fatal("Erreur lors du chargement du joueur:", err)
	}
	return joueur
}


// ----- Charger les thèmes disponibles (exportée) -----
func AjoutdesThemes() []Theme {
	rows, err := db.Query("SELECT id_theme, nom_theme FROM theme")
	if err != nil {
		log.Fatal("Erreur lors du chargement des thèmes:", err)
	}
	defer rows.Close()

	var themes []Theme
	for rows.Next() {
		var theme Theme
		err := rows.Scan(&theme.Id_theme, &theme.Nom_theme)
		if err != nil {
			log.Fatal("Erreur lors du scan des thèmes:", err)
		}
		themes = append(themes, theme)
	}
	return themes
}
// ----- Charger les questions par thème (exportée) -----
func AjoutdesQuestionsParTheme(idTheme int) []Question {
	rows, err := db.Query("SELECT id_question, texte_question, niveau_difficulte, id_theme FROM question WHERE id_theme = ?", idTheme)
	if err != nil {
		log.Fatal("Erreur lors du chargement des questions:", err)
	}	
	defer rows.Close()
	var questions []Question
	for rows.Next() {
		var question Question
		err := rows.Scan(&question.Id_question, &question.Texte_question, &question.Niveau_difficulte, &question.Id_theme)
		if err != nil {
			log.Fatal("Erreur lors du scan des questions:", err)
		}	
		questions = append(questions, question)
	}
	return questions
}
// ----- Charger les réponses par question (exportée) -----
func AjoutdesReponsesParQuestion(idQuestion int) []Reponse_question {
	rows, err := db.Query("SELECT id_reponsequestion, id_question, texte_reponsequestion, est_correct FROM reponse_question WHERE id_question = ?", idQuestion)
	if err != nil {
		log.Fatal("Erreur lors du chargement des réponses:", err)
	}
	defer rows.Close()
	var reponses []Reponse_question
	for rows.Next() {
		var reponse Reponse_question
		err := rows.Scan(&reponse.Id_reponsequestion, &reponse.Id_question, &reponse.Texte_reponsequestion, &reponse.Est_correct)
		if err != nil {
			log.Fatal("Erreur lors du scan des réponses:", err)
		}
		reponses = append(reponses, reponse)
	}
	return reponses
}
// ----- Histoire des parties (exportée) -----
func AjoutdesHistoriquesParPartie(idPartie int) []Historique {
	rows, err := db.Query("SELECT id_historique, score_enregistre, id_partie, date_sauvegarde FROM historique WHERE id_partie = ?", idPartie)
	if err != nil {
		log.Fatal("Erreur lors du chargement des historiques:", err)
	}
	defer rows.Close()

	var historiques []Historique
	for rows.Next() {
		var h Historique
		var dateStr []byte // lire d'abord comme byte slice
		err := rows.Scan(&h.Id_historique, &h.Score_enregistre, &h.Id_partie, &dateStr)
		if err != nil {
			log.Fatal("Erreur lors du scan des historiques:", err)
		}
		historiques = append(historiques, h)
	}
	return historiques
}
