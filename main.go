package main

import (
	"fmt"
	"os"

	"github.com/flexphere/discord-gorilla-cli/api"
	inputComponent "github.com/flexphere/discord-gorilla-cli/component/input"
	listComponent "github.com/flexphere/discord-gorilla-cli/component/list"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(inputComponent.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}

	data, err := api.Search(inputComponent.Value)
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// pp.Println(data)

	m := listComponent.Model{List: list.New(data, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Search Results"

	p = tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
