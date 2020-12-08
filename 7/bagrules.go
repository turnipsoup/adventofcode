package main

import (
    "fmt"
	"io/ioutil"
	"strings"
)

// Rule structure for when we parse the list of rules
type Rule struct {
	bagColor string
	ruleOne string
	ruleTwo string
	ruleThree string
	ruleFour string
}

// Open the rules file, return a slice of the file seprarated by newlines
func loadRules(filename string) []string {
	rulesArr, err := ioutil.ReadFile("./rules.txt")
	
	if err != nil {
		fmt.Println(err)
	}
	rules := strings.Split(string(rulesArr), "\n")

	return rules
}

// Returns a Rules struct containing the bagColor and rules
func parseRule(ruleString string) Rule {
	currentRule := Rule{}

	ruleSplit := strings.Split(ruleString, " contain ")

	// Get bag color of rule
	bagColor := ruleSplit[0]
	currentRule.bagColor = strings.Split(bagColor, " bags")[0]

	// Get rules that are present
	bagRules := strings.Split(ruleSplit[1], ",")

	for i := 0; i < len(bagRules); i++ {
		switch i {
			case 0:
			currentRule.ruleOne = bagRules[i]
			case 1:
			currentRule.ruleTwo = bagRules[i]
			case 2:
			currentRule.ruleThree = bagRules[i]
			case 3:
			currentRule.ruleFour = bagRules[i]
		}
	}

	return currentRule
}

// Checks all of the bag rules to see if the targets color is holdable
func checkBagRules(bagColor string, rule Rule) bool {

	var rulesArr = []string{rule.ruleOne, rule.ruleTwo, rule.ruleThree, rule.ruleFour}
	colorPresent := false

	for i := 0; i < len(rulesArr); i++ {
		if strings.Contains(rulesArr[i], bagColor) {
			colorPresent = true
		}
	}

	return colorPresent
}

func sliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func checkAllRulesForSupporters(rulesArr []Rule, supportsShinyGold []string) []string {
	supports := supportsShinyGold

	for i := 0; i < len(supports); i++ {
		for t := 0; t < len(rulesArr); t++ {
			otherSupported := checkBagRules(supports[i], rulesArr[t])

			if otherSupported {
				supports = append(supports, rulesArr[t].bagColor)
			}
		}
		
	}

	return supports
}

// https://www.golangprograms.com/remove-duplicate-values-from-slice.html
// https://stackoverflow.com/questions/33207197/how-can-i-create-an-array-that-contains-unique-strings
func uniqueValues(stringSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{} 
    for _, entry := range stringSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

// Part 1
func partOne() {
	// Load rules slice
	rules := loadRules("./rules.txt")

	// Define globals
	parsedRulesArr := []Rule{}
	supportsShinyGold := []string{}
	
	// Iterate over the rules, parse them, and add parsed Rule structs to parsedRulesArr
	for i := 0; i < len(rules); i++ {
		parsedRule := parseRule(rules[i])
		parsedRulesArr = append(parsedRulesArr, parsedRule)
	}

	// Check if shiny gold bags are in these bag rules
	for i := 0; i < len(parsedRulesArr); i++ {
		goldSupported := checkBagRules("shiny gold", parsedRulesArr[i])
		if goldSupported {
			supportsShinyGold = append(supportsShinyGold, parsedRulesArr[i].bagColor)
		}
		
	}

	// Check which bag colors support supporters of shiny gold
	// Basically iterate over an ever growing like like 4 times, then get unique
	// values from that list
	for i := 0; i < 4; i++ {
		supportsShinyGold = checkAllRulesForSupporters(parsedRulesArr, supportsShinyGold)
	}

	//totalValidBags := len(supportsShinyGold) + len(otherGoldSupporters)
	fmt.Printf("There are a total of %d bags that support shiny gold bags.", len(uniqueValues(supportsShinyGold)))

}

func partTwo() {
	// Load rules slice
	rules := loadRules("./rules.txt")

	// Define globals
	parsedRulesArr := []Rule{}
	
	// Iterate over the rules, parse them, and add parsed Rule structs to parsedRulesArr
	for i := 0; i < len(rules); i++ {
		parsedRule := parseRule(rules[i])
		parsedRulesArr = append(parsedRulesArr, parsedRule)
	}
}

func main() {
	partOne()
	
}
