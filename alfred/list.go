package alfred

import (
	"fmt"
	"sort"
)

func list(tasks map[string]Task) {

	max := 0
	labels := make([]string, 0)
	for label := range tasks {
		// lets add the label to the list(so we an alphabatize the list)
		labels = append(labels, label)

		// figure out max label size
		if len(label) > max {
			max = len(label)
		}
	}

	// alphabatize the list
	sort.Strings(labels)

	// insignifigant tasks
	for _, label := range labels {
		task := tasks[label]
		if task.Summary == "" {
			fmt.Print(translate("{{ .Text.Grey }}"+padRight(label, max, " ")+"{{ .Text.Reset }}", emptyContext()), "\t")
		}
	}

	fmt.Println()
	fmt.Println()

	// signifigant tasks
	for _, label := range labels {
		task := tasks[label]
		if task.Summary != "" {
			fmt.Println(translate("{{ .Text.Task }}"+padRight(label, max, " ")+"{{ .Text.Grey }}| {{ .Text.Reset }}"+task.Summary, emptyContext()))
		}
	}
}