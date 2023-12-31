![French Army Logo](./assets/images/army_logo.png)

# 🪖 ArmyMissionTracker

This project is a simple report generator for the army developed in Go. It lets you create customized reports by filling in a form and saving them in JSON format.
It has been designed mainly for the French army.

## How to Run the Project

1. Make sure you have Go installed on your machine (version 1.16 or higher).
2. Clone this GitHub repository: `git clone <repository URL>`
3. Navigate to the project directory: `cd ArmyMissionTracker`
4. Install project dependencies: `go mod download`
5. Run the application: `go run main.go`
6. Open your browser and access the following URL: `http://localhost:8080`

## Technologies Used

- Go (version 1.16 or higher) - Main programming language.
- HTML/CSS - User interface for the report form.
- HTML Template - Used to generate dynamic HTML pages.
- JSON - Data format used for saving reports.

## Project Directory Structure

```
ArmyMissionTracker/
├── handlers
│ └── report_handler.go # 
├── reports/ # Directory for generated reports (not tracked by Git)
│ └── generated_reports/ # Directory for JSON reports generated by the application
├── main.go # Entry point of the application
├── static/ # Directory for static files (CSS, images)
│ └── style.css # CSS style sheet
│ └── stylereports.css # CSS style sheet
│ └── script.css # JS script
├── templates/ # Directory for HTML templates
│ ├── report_form.html # Report form template
│ └── reports.html # Generated reports list template
└── main.go # Entry point of the application
```

- `main.go` is the main file containing the application logic.
- The `static/` directory contains static files such as CSS style sheets.
- The `templates/` directory contains HTML templates used to generate web pages.
- `report_form.html` is the HTML template for the report form.
- `reports.html` is the HTML template for the generated reports list.
- The `reports/generated_reports/` directory is the location where JSON reports generated by the application are stored. This directory is created automatically when you generate your first report.

Make sure to add the `reports/generated_reports/` directory to your `.gitignore` file to avoid tracking the generated reports in your Git repository.

---
