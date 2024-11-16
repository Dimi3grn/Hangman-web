# Projet Web avec Go, HTML et CSS

## Description
Ce projet est une application web développée en **Go**, utilisant des **templates HTML** et du **CSS** pour la présentation. L'application permet de gérer des pages comme **Landing**, **Hangman**, ainsi qu'une section pour afficher des informations sur les étudiants. Aucune bibliothèque JavaScript n'est utilisée, en restant fidèle à une approche basée uniquement sur Go, HTML, et CSS.

## Pré-requis
Avant de lancer le projet, assurez-vous d'avoir installé :
- **Go** (version 1.20 ou supérieure)
- Un éditeur de texte ou un IDE (comme Visual Studio Code)
- Un navigateur web moderne (comme Google Chrome ou Firefox)

## Installation
1. Clonez ce dépôt sur votre machine locale :
    ```bash
    git clone <url_du_dépôt>
    cd <nom_du_dépôt>
    ```

2. Installez les dépendances Go :
    ```bash
    go mod tidy
    ```

3. Assurez-vous que la structure des dossiers est bien en place :
    ```
    .
    ├── go.mod
    ├── main.go
    ├── template/
    │   ├── index.html
    │   ├── hangman.html
    │   └── etudiants.html
    ├── images/
    │   ├── homme.png
    │   └── femme.png
    ├── static/
    │   └── styles.css
    └── README.md
    ```

## Lancement de l'application
Pour lancer le serveur web en local, exécutez la commande suivante :
```bash
go run main.go