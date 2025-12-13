package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dbPath = "data.db" 

var QuestionsParTheme = 20

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
	Text_question string
	Niveau_difficulte string
	Id_theme int
}

type Reponse_joueur struct {
	Id_reponsejoueur int
	Id_question int
	Text_reponsejoueur string
	Id_joueur int
	Est_correcte bool
}

type Reponse_question struct {
	Id_reponsequestion int
	Id_question int
	Text_reponsequestion string
	Est_correcte bool
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
func AjouterOuModifierUnNouveauJoueur(name string) Users {
	var joueur Users
	err := db.QueryRow("SELECT username FROM users WHERE username = ?", name).Scan(&joueur.Username)
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
/* 
// ----- Charger les thèmes disponibles (exportée) -----
func LoadThemes() []Theme {
	rows, err := db.Query("SELECT id, nom FROM themes")
	if err != nil {
		log.Fatal("Erreur lors du chargement des thèmes:", err)
	}
	defer rows.Close()

	var themes []Theme
	for rows.Next() {
		var theme Theme
		err := rows.Scan(&theme.ID, &theme.Nom)
		if err != nil {
			log.Fatal("Erreur lors du scan des thèmes:", err)
		}
		themes = append(themes, theme)
	}
	return themes
}

// ----- Charger les questions pour un thème donné (exportée) -----
func LoadQuestionsForTheme(themeID int, limit int) []Question {
	rows, err := db.Query("SELECT id, category, text, options, correct_idx FROM questions WHERE category_id = ? LIMIT ?", themeID, limit)
	if err != nil {
		log.Fatal("Erreur lors du chargement des questions:", err)
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var q Question
		var optionsStr string
		err := rows.Scan(&q.ID, &q.Category, &q.Text, &optionsStr, &q.CorrectIdx)
		if err != nil {
			log.Fatal("Erreur lors du scan des questions:", err)
		}
		q.Options = strings.Split(optionsStr, ";") // Supposant que les options sont séparées par des points-virgules
		questions = append(questions, q)
	}	
	return questions
}*/