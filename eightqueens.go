package piscine

import (
	"github.com/01-edu/z01"
)

type Point struct {
	x int
	y int
}

// Fonction EightQueens
// Cette fonction à pour but de trouver tout les posibilités pour que
// 8 dames soient positionnés sans se mettre en dangés mutuelement sur
// un plateau de 8x8
//
func EightQueens() {
	for col := 0; col < 8; col++ {
		Recurse(Point{x: col, y: 0}, make([]Point, 0), 8)
	}
	DisplaySolution()
}

var results = make([]rune, 0)

// Fonction récursive
// Cette fonction passe en revue tout les cases et détermine si dans la configuration actuel de placement (current[])
// permet le placement d'une dame sur (point)
// Une fois un placement possible trouvé, elle se rappel en prenant une nouvelle configuration partant du placement trouver
//
func Recurse(point Point, current []Point, n int) {
	if !CanPlace(point, current) { // Si on ne peux pas placer de dame ici
		return // On arrete la
	}
	current = append(current, point) // On ajoute le placement à la configuration actuel
	if len(current) < n {            // Si le nombre de placements dans la configuration est plus petit que le nombre de dames voulus
		for col := 0; col < n; col++ { // On boucle sur toutes les colones du plateau
			for row := point.y; row < n; row++ { // Nouvelle boucle sur toute les ligne du tableau en commencant par le dernière index trouver
				Recurse(Point{x: col, y: row}, current, n) // On rappelle la fonction pour continuer la recherche de placement et de possibilités
			}
		}
		return // Plus aucun placement n'est necessaire dans cette configuration on arrete la
	}
	AddPointToResult(current) // On ajoute la configuration au résultats
}

// Fonction CanPlace
// Cette fonction regarde si un placement est possible en fonction d'une configuration actuelle
//
func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

// Cette fonction regarde si un placement a peut attaquer un placement b et inversement
func CanAttack(a, b Point) bool {
	return a.x == b.x || a.y == b.y || Abs(float64(a.y-b.y)) == Abs(float64(a.x-b.x))
}

// Fonction Abs (venue du package Math)
func Abs(f float64) float64 {
	return map[bool]float64{true: f * float64(-1), false: f}[f < 0]
}

// Ajout d'une position au résultats finaux
func AddPointToResult(current []Point) {
	for _, c := range current {
		results = append(results, rune(c.x+49))
	}
	results = append(results, '\n')
}

// On affiche les résultats dans la console
func DisplaySolution() {
	for _, result := range results {
		z01.PrintRune(result)
	}
}
