package script

import (
	"fmt"
	"github.com/phrp720/aw-sync-agent-plugins/models"
	"github.com/phrp720/aw-sync-agent-plugins/util"
	"log"
	"strconv"
	"strings"
)

// PrintScripts prints the array of the registered scripts with borders
func PrintScripts(scripts []string) {
	log.Print("Scripts Registered:")

	scriptsMap := map[string]string{
		"Scripts found": strconv.Itoa(len(scripts)),
		"Scripts":       strings.Join(scripts, ", "),
	}

	maxKeyLength := 0
	maxValueLength := 0
	for key, value := range scriptsMap {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
		if len(value) > maxValueLength {
			maxValueLength = len(value)
		}
	}

	borderLength := maxKeyLength + maxValueLength + 7
	border := strings.Repeat("-", borderLength)
	fmt.Println(border)
	for key, value := range scriptsMap {
		fmt.Printf("| %-*s | %-*s |\n", maxKeyLength, key, maxValueLength, value)
	}
	fmt.Println(border)
}

// GetScriptNames returns the names of the registered scripts
func GetScriptNames(scripts []models.Script) []string {
	var scriptNames []string
	for _, script := range scripts {
		if script.Name != "" && !util.Contains(scriptNames, script.Name) {
			scriptNames = append(scriptNames, script.Name)
		}
	}
	return scriptNames
}
