package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

// Structure de données pour stocker les informations du rapport
type Report struct {
	Date        string
	Time        string
	Location    string
	Units       string
	Description string
}

// Handler pour la page de création de rapport
func ReportFormHandler(w http.ResponseWriter, r *http.Request) {
	// Vérification de la méthode HTTP
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Chargement du template HTML
	tmpl, err := template.ParseFiles("templates/report_form.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Affichage du formulaire de rapport
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Handler pour la soumission du rapport
func GenerateReportHandler(w http.ResponseWriter, r *http.Request) {
	// Vérification de la méthode HTTP
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupération des données du formulaire
	report := Report{
		Date:        r.FormValue("date"),
		Time:        r.FormValue("time"),
		Location:    r.FormValue("location"),
		Units:       r.FormValue("units"),
		Description: r.FormValue("description"),
	}

	// Génération du rapport
	generateReportPDF(report)

	// Redirection vers la page de confirmation
	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

// Fonction pour générer le rapport au format PDF
func generateReportPDF(report Report) {
	// Création du nom de fichier unique basé sur la date et l'heure actuelles
	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("reports/generated_reports/report_%s.pdf", timestamp)

	// Ouverture du fichier en écriture
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}
	defer file.Close()

	// Affichage d'un message de confirmation
	fmt.Println("Rapport généré :", filename)
}
