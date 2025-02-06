package filter

import (
	"fmt"
	"github.com/phrp720/aw-sync-agent-plugins/models"
	"github.com/phrp720/aw-sync-agent-plugins/util"
	"log"
	"strconv"
	"strings"
)

// ValidateFilters validates the filters in the List
func ValidateFilters(filters []models.Filter) ([]models.Filter, int, int, int) {
	validFilters := []models.Filter{}
	invalid := 0
	disabled := 0
	total := len(filters)
	for _, filter := range filters {
		if filter.Enable {
			var targetList []models.Target
			for _, target := range filter.Target {
				if strings.TrimSpace(target.Key) != "" && target.Value != nil {
					targetList = append(targetList, target)
				}
			}

			if len(targetList) != 0 { // Check if the filter has at least one valid target
				if filter.Category != "" && (filter.Drop || (filter.RegexReplace != nil && len(filter.RegexReplace) > 0) || (filter.PlainReplace != nil && len(filter.PlainReplace) > 0)) {
					log.Print("Warning: Filter `", filter.FilterName, "` has a category assigned, but it must not have any additional filtering. This filter will be ignored.")
					invalid++
					continue
				}
				validFilters = append(validFilters, filter)
			} else {
				invalid++
			}
		} else {
			disabled++
		}
	}
	return validFilters, total, invalid, disabled
}

func GetCategories(filters []models.Filter) []string {
	var categories []string
	for _, filter := range filters {
		if filter.Category != "" && !util.Contains(categories, filter.Category) {
			categories = append(categories, filter.Category)
		}
	}
	return categories
}

// Apply applies the filters to the data
func Apply(data map[string]interface{}, filters []models.Filter) (map[string]interface{}, bool) {
	for _, filter := range filters {
		allMatch := true
		for _, target := range filter.Target {
			// Check if the data contains the key to be filtered
			if value, ok := data[target.Key]; ok {
				// Check if the value matches the target's regex
				if !target.Value.MatchString(fmt.Sprintf("%v", value)) {
					allMatch = false
					break
				}
			} else {
				allMatch = false
				break
			}
		}

		if allMatch {
			if filter.Category != "" {
				// Add the category to the data
				data["category"] = filter.Category
				continue
			}
			// Drop the event if the filter matches
			if filter.Drop {
				return nil, true
			}

			// Apply replacements
			data = Replace(data, filter.PlainReplace, filter.RegexReplace)
		}
	}
	return data, false
}

// Replace replaces the values in the data
func Replace(data map[string]interface{}, plain []models.PlainReplace, regex []models.RegexReplace) map[string]interface{} {

	// Apply replacements

	// Plain replacements
	for _, replace := range plain {
		if _, exists := data[replace.Key]; exists {
			data[replace.Key] = replace.Value
		}
	}

	// Regex replacements
	for _, replace := range regex {
		if value, exists := data[replace.Key]; exists {
			// Check if the value matches the target's regex
			if replace.Expression.MatchString(fmt.Sprintf("%v", value)) {
				// Replace the value with the formatted string
				data[replace.Key] = replace.Expression.ReplaceAllString(fmt.Sprintf("%v", value), replace.Value)
			}
		}
	}
	return data
}

// GetMatchingFilters returns filters that match the given watcher or have an empty watcher list
func GetMatchingFilters(filters []models.Filter, watcher string) []models.Filter {
	var matchingFilters []models.Filter
	for _, filter := range filters {
		if len(filter.Watchers) == 0 || util.Contains(filter.Watchers, watcher) {
			matchingFilters = append(matchingFilters, filter)
		}
	}
	if len(matchingFilters) > 0 {
		log.Print("Filters applied for [", watcher, "]: ", len(matchingFilters))
	}

	return matchingFilters
}

// PrintFiltersDebug prints the filters in the List
func PrintFiltersDebug(filters []models.Filter) {
	for i, filter := range filters {
		fmt.Printf("Filter %d:\n", i+1)
		fmt.Printf("  Name: %s\n", filter.FilterName)
		fmt.Printf("  Watchers: %v\n", filter.Watchers)
		for k, target := range filter.Target {
			fmt.Printf("  Target %d:\n", k+1)
			fmt.Printf("    Key: %s\n", target.Key)
			fmt.Printf("    Value: %s\n", target.Value)
		}
		for j, replace := range filter.PlainReplace {
			fmt.Printf("  Plain String Replace %d:\n", j+1)
			fmt.Printf("    Key: %s\n", replace.Key)
			fmt.Printf("    Value: %s\n", replace.Value)
		}
		for m, replace := range filter.RegexReplace {
			fmt.Printf("  Regex Value Replace %d:\n", m+1)
			fmt.Printf("    Key: %s\n", replace.Key)
			fmt.Printf("    Value: %s\n", replace.Value)
			fmt.Printf("    Expression: %s\n", replace.Expression)
		}
		fmt.Printf("  Enabled: %t\n", filter.Enable)
		fmt.Printf("  Drop: %t\n", filter.Drop)
		fmt.Printf("  Category: %s\n", filter.Category)
	}
}

// PrintFilters prints the filters in the List in a dashboard format
func PrintFilters(totalFilters, invalidFilters, disabledFilters int) {
	log.Print("Filters :")
	// Create a map of settings for easier iteration
	filtersMap := map[string]int{
		"Total filters":    totalFilters,
		"Valid filters":    totalFilters - invalidFilters,
		"Invalid filters":  invalidFilters,
		"Disabled filters": disabledFilters,
	}
	maxKeyLength := 0
	maxValueLength := 0
	for key, value := range filtersMap {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
		if len(fmt.Sprintf("%d", value)) > maxValueLength {
			maxValueLength = len(fmt.Sprintf("%d", value))
		}
	}

	borderLength := maxKeyLength + maxValueLength + 7
	border := strings.Repeat("-", borderLength)
	fmt.Println(border)
	for key := range filtersMap {
		value := filtersMap[key]
		fmt.Printf("| %-*s | %-*d |\n", maxKeyLength, key, maxValueLength, value)
	}
	fmt.Println(border)

}

// PrintCategories prints the categories in the List in a dashboard format
func PrintCategories(categories []string) {
	log.Print("Categories:")

	filtersCategoriesMap := map[string]string{
		"Categories found": strconv.Itoa(len(categories)),
		"Categories":       strings.Join(categories, ", "),
	}

	maxKeyLength := 0
	maxValueLength := 0
	for key, value := range filtersCategoriesMap {
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
	for key, value := range filtersCategoriesMap {
		fmt.Printf("| %-*s | %-*s |\n", maxKeyLength, key, maxValueLength, value)
	}
	fmt.Println(border)
}
