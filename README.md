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
go run ./modele/exe/main.go
```

Une fois le serveur démarré, ouvrez votre navigateur et accédez à l'URL suivante :

```arduino
http://localhost:8085
```
## Routes de l'application

### Routes de vues
| Méthode | Route               | Description                               |
|---------|---------------------|-------------------------------------------|
| GET     | `/landingPage`                 | Page d'accueil (Landing page)             |
| GET     | `//hangman/mainGame`          | Jeu du Hangman                            |

### Routes de traitement des données
| Méthode | Route               | Description                               |
|---------|---------------------|-------------------------------------------|
| POST    | `/signup/treatment`        | S'enregistrer         |
| POST     | `/login/treatment`    | Se connecter    |
| POST     | `/landingPage/treatment`    | Permet de choisir un mot aléatoir suivant quel catégorie choisie    |
| POST     | `/hangman/treatment`    | Verifie le mot ou la lettre entrée    |


## Fonctionnalités
- **Landing page** avec une présentation simple de l'application.
- **Jeu du pendu** (Hangman) utilisant des mots tirés aléatoirement d'un fichier.


## Contribution
### Équipe de développement
- **Dimitri Gourrin**
- **Xerly Ji**

