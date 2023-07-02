package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// Structure de données pour stocker les informations du rapport
type Report struct {
	Date        string
	Time        string
	Location    string
	Units       string
	Description string
}

func main() {
	// Gestion des fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route pour le formulaire
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/generate", generateHandler)
	http.HandleFunc("/reports", reportsHandler)
	http.HandleFunc("/download", downloadHandler)

	// Démarrer le serveur HTTP
	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Handler pour la page d'accueil
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Chargement du template HTML
	tmpl, err := template.ParseFiles("templates/report_form.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	// Affichage du formulaire de rapport
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
		return
	}
}

// Handler pour la génération de rapport
func generateHandler(w http.ResponseWriter, r *http.Request) {
	// Vérification de la méthode HTTP
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
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

	// Génération du rapport (ici, nous affichons simplement les données dans la console)
	fmt.Println("Rapport généré :")
	fmt.Println("Date :", report.Date)
	fmt.Println("Heure :", report.Time)
	fmt.Println("Lieu :", report.Location)
	fmt.Println("Unités :", report.Units)
	fmt.Println("Description :", report.Description)

	// Sauvegarde du rapport au format JSON avec le nom spécifié
	err := saveReport(report, report.Units)
	if err != nil {
		http.Error(w, "Erreur lors de la sauvegarde du rapport", http.StatusInternalServerError)
		return
	}

	// Redirection vers la page d'accueil ou une page de confirmation
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Handler pour la liste des rapports
func reportsHandler(w http.ResponseWriter, r *http.Request) {
	// Liste des fichiers JSON
	fileList, err := getFileList("reports/generated_reports/")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de la liste des fichiers", http.StatusInternalServerError)
		return
	}

	// Chargement du template HTML
	tmpl, err := template.ParseFiles("templates/reports.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	// Affichage de la liste des fichiers
	err = tmpl.Execute(w, fileList)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
		return
	}
}

// Handler pour le téléchargement d'un rapport
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération du nom du fichier à télécharger
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Nom de fichier manquant", http.StatusBadRequest)
		return
	}

	// Chemin du fichier
	filePath := filepath.Join("reports/generated_reports", filename)

	// Lecture du fichier
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier", http.StatusInternalServerError)
		return
	}

	// Envoi du fichier en tant que téléchargement
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/json")
	w.Write(fileContent)
}

// Fonction pour sauvegarder le rapport au format JSON avec un nom de fichier spécifié
func saveReport(report Report, filename string) error {
	// Conversion du rapport en JSON
	jsonData, err := json.Marshal(report)
	if err != nil {
		return err
	}

	// Écriture des données JSON dans un fichier avec le nom spécifié
	err = ioutil.WriteFile(filepath.Join("reports/generated_reports", filename+".json"), jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Fonction pour générer un nom de fichier unique
func generateFileName() string {
	return "report.json"
}

// Fonction utilitaire pour récupérer la liste des fichiers JSON dans un répertoire
func getFileList(dirPath string) ([]string, error) {
	var fileList []string

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList, nil
}
