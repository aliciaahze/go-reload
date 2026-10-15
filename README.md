# Go Reloaded

## Description

Go Reloaded est un programme écrit en Go qui transforme un fichier texte à partir de marqueurs et corrige certains éléments de mise en forme.

## Fonctionnalités

- Conversion de nombres hexadécimaux avec `(hex)`.
- Conversion de nombres binaires avec `(bin)`.
- Mise en majuscules avec `(up)`.
- Mise en minuscules avec `(low)`.
- Mise en forme avec une majuscule initiale grâce à `(cap)`.
- Transformation de plusieurs mots avec `(up, n)`, `(low, n)` et `(cap, n)`.
- Correction de la ponctuation et des apostrophes.
- Remplacement de `a` par `an` devant une voyelle ou un `h`.
- Préservation des sauts de ligne.

## Prérequis

Go doit être installé sur votre ordinateur.

## Utilisation

Depuis le dossier du projet, exécutez :

```bash
go run . sample.txt result.txt
```

Le programme lit le fichier `sample.txt` et écrit le résultat transformé dans `result.txt`.

## Technologies

- Go
- Bibliothèques standard de Go
