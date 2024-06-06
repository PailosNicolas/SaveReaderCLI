package menus

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PailosNicolas/GoPkmSaveReader/pokemon"
)

func spaceCalculator(budget int, previosWrods ...string) (string, error) {
	s := ""

	for _, word := range previosWrods {
		if len(word) > budget {
			return "", errors.New("previosWrods longer than budget")
		}

		budget -= len(word)
	}

	s += strings.Repeat(" ", budget)
	return s, nil

}

func pokemonStatView(pkm pokemon.Pokemon) string {
	var s strings.Builder
	budget := 25
	budgetHeader := 10

	s.WriteString("Stats:\n\t")
	stats := pkm.Stats()
	ivs := pkm.IVs()
	evs := pkm.Evs()
	space, _ := spaceCalculator(budget, " Stats:")
	s.WriteString(space)
	s.WriteString("\t Stat")
	space, _ = spaceCalculator(budgetHeader, "Stat")
	s.WriteString(space)
	s.WriteString("IV")
	space, _ = spaceCalculator(budgetHeader, "IV")
	s.WriteString(space)
	s.WriteString("EVs")

	s.WriteString("\n\tHp:")
	space, _ = spaceCalculator(budget, "Hp:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.TotalHP))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.TotalHP))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.Hp))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(ivs.Hp))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(evs.Hp))

	s.WriteString("\n\tAttack:")
	space, _ = spaceCalculator(budget, "Attack:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.Attack))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.Attack))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.Attack))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(ivs.Attack))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(evs.Attack))

	s.WriteString("\n\tDefense:")
	space, _ = spaceCalculator(budget, "Defense:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.Defense))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.Defense))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.Defense))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(ivs.Defense))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(evs.Defense))

	s.WriteString("\n\tSpecial Defense:")
	space, _ = spaceCalculator(budget, "Special Defense:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.SpecialDefense))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.SpecialDefense))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.SpecialDefense))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(ivs.SpecialDefense))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(evs.SpecialDefense))

	s.WriteString("\n\tSpecial Attack:")
	space, _ = spaceCalculator(budget, "Special Attack:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.SpecialAttack))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.SpecialAttack))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.SpecialAttack))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(ivs.SpecialAttack))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(evs.SpecialAttack))

	s.WriteString("\n\tSpeed:")
	space, _ = spaceCalculator(budget, "Speed:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.Speed))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.Speed))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.Speed))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(ivs.Speed))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(evs.Speed))

	s.WriteString("\n")

	return s.String()
}

func moveView(move pokemon.Move) string {
	var s strings.Builder

	s.WriteString(move.Name)
	s.WriteString(":\n\tPP: ")
	s.WriteString(strconv.Itoa(move.PP))

	return s.String()
}
