package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Shopify/go-lua"
)

type Group struct {
	monkeys []*Monkey
}

type Monkey struct {
	items     []int
	operation string
	divider   int
	targetYes int
	targetNo  int
	activity  int
}

func main() {
	inputFileName := "input.txt"
	data, _ := os.ReadFile(inputFileName)

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	group := Group{monkeys: []*Monkey{}}

	for _, line := range lines {
		line = strings.Trim(line, " ")

		if strings.HasPrefix(line, "Monkey") {
			group.addMonkey()
		}
		if strings.HasPrefix(line, "Starting items") {
			itemStringFull := line[len("Starting items:")+1:]
			itemStringList := strings.Split(itemStringFull, ", ")
			for _, itemString := range itemStringList {
				item, _ := strconv.Atoi(itemString)
				group.addItem(item)
			}
		}
		if strings.HasPrefix(line, "Operation") {
			operationString := line[len("Operation:")+1:]
			group.setOperation(operationString)
		}
		if strings.HasPrefix(line, "Test") {
			dividerString := line[len("Test: divisible by")+1:]
			divider, _ := strconv.Atoi(dividerString)
			group.setDivider(divider)
		}
		if strings.HasPrefix(line, "If true") {
			targetYesString := line[len("If true: throw to monkey")+1:]
			targetYes, _ := strconv.Atoi(targetYesString)
			group.setTargetYes(targetYes)
		}
		if strings.HasPrefix(line, "If false") {
			targetNoString := line[len("If false: throw to monkey")+1:]
			targetNo, _ := strconv.Atoi(targetNoString)
			group.setTargetNo(targetNo)
		}
	}

	group.dump()

	for i := 0; i < 20; i++ {
		group.round()
		group.dump()
	}

	fmt.Println("answer1", group.business()) //66124
}

func (g *Group) addMonkey() {
	g.monkeys = append(g.monkeys, &Monkey{activity: 0})
}

func (g *Group) lastMonkey() *Monkey {
	return g.monkeys[len(g.monkeys)-1]
}

func (g *Group) addItem(item int) {
	monkey := g.lastMonkey()
	monkey.addItem(item)
}

func (g *Group) setOperation(operation string) {
	g.lastMonkey().operation = operation
}

func (g *Group) setDivider(divider int) {
	g.lastMonkey().divider = divider
}

func (g *Group) setTargetYes(targetYes int) {
	g.lastMonkey().targetYes = targetYes
}

func (g *Group) setTargetNo(targetNo int) {
	g.lastMonkey().targetNo = targetNo
}

func (g *Group) round() {
	for _, monkey := range g.monkeys {
		for _, item := range monkey.items {
			target, worry := monkey.process(item)
			g.monkeys[target].addItem(worry)
			monkey.activity++
		}
		monkey.items = []int{}
	}
}

func (m *Monkey) addItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) process(item int) (int, int) {
	vm := lua.NewStateEx()
	vm.PushNumber(float64(item))
	vm.SetGlobal("old")
	err := lua.DoString(vm, m.operation)
	if err != nil {
		panic(err)
	}
	vm.Global("new")
	worry, _ := vm.ToInteger(0)

	worry = worry / 3

	if worry%m.divider == 0 {
		return m.targetYes, worry
	} else {
		return m.targetNo, worry
	}
}

func (g *Group) business() int {
	activities := []int{}
	for _, monkey := range g.monkeys {
		activities = append(activities, monkey.activity)
	}
	sort.Ints(activities)
	return activities[len(activities)-1] * activities[len(activities)-2]
}

func (g *Group) dump() {
	fmt.Println("Group")
	for index, monkey := range g.monkeys {
		fmt.Println(index, monkey.activity, monkey.items)
	}
}
