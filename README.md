# Jeu du pendu en Go

Projet fait en binôme (Alexis et Leandro) pour le TP de Go à Ynov.
C'est un jeu du pendu qui se joue dans le terminal.

## Lancer le jeu

Il faut avoir Go installé, puis dans le dossier du projet :

```
go run . jouer       -> nouvelle partie
go run . reprendre   -> reprendre la partie sauvegardée
go run . score       -> voir l'historique des parties
```

Si on lance sans commande ou avec une commande qui existe pas, ça affiche l'usage.

## Comment on joue

- On choisit un pseudo puis une catégorie de mots
- Une lettre du mot est donnée au début
- On propose une lettre à chaque tour, ou directement le mot entier
- On a 6 essais : une mauvaise lettre = -1, un mauvais mot = -2
- `!save` pour sauvegarder et quitter
- `!menu` pour revenir au menu

## Les fichiers

- `main.go` : les commandes
- `newgame.go` : le pseudo, les règles et le choix de la catégorie
- `selection.go` : lit les fichiers de mots et en tire un au hasard
- `indicateletter.go` : la partie en elle-même
- `save.go` : la sauvegarde dans save.json
- `score.go` : l'historique dans scores.csv
- `mot-*.txt` : les listes de mots par catégorie

## Ce qu'on a ajouté en plus

- Des catégories (animaux, métiers, pays, nourriture, simple, hard)
- Le dessin du pendu qui avance à chaque erreur
- Pouvoir rejouer le même mot à la fin

## Aide

On s'est servi de la doc de Go et on s'est aidé d'une IA sur certaines parties
(le dessin du pendu, la sauvegarde et les scores).