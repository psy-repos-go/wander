package command

import (
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"wander/formatter"
	"wander/message"
	"wander/nomad"
)

func FetchJobs(url, token string) tea.Cmd {
	return func() tea.Msg {
		// TODO LEO: error handling
		//body, _ := nomad.GetJobs(url, token)
		body := MockJobsResponse
		var jobResponse []nomad.JobResponseEntry
		if err := json.Unmarshal(body, &jobResponse); err != nil {
			fmt.Println("Failed to decode response")
		}

		return message.NomadJobsMsg{Table: formatter.JobResponseAsTable(jobResponse)}
	}
}
