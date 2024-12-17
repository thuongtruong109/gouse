package samples

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/gouse/cookbook/console/choice"
	"github.com/thuongtruong109/gouse/cookbook/console/countdown"
	"github.com/thuongtruong109/gouse/cookbook/console/dir"
	"github.com/thuongtruong109/gouse/cookbook/console/glamour"
	"github.com/thuongtruong109/gouse/cookbook/console/inline"
	"github.com/thuongtruong109/gouse/cookbook/console/list"

	// "github.com/thuongtruong109/gouse/console/paper"
	"github.com/thuongtruong109/gouse/cookbook/console/parallel"
	"github.com/thuongtruong109/gouse/cookbook/console/progress"
	"github.com/thuongtruong109/gouse/cookbook/console/realtime"
	"github.com/thuongtruong109/gouse/cookbook/console/sequence"
	"github.com/thuongtruong109/gouse/cookbook/console/spinner"
	"github.com/thuongtruong109/gouse/cookbook/console/split"
	"github.com/thuongtruong109/gouse/cookbook/console/stopwatch"

	// "github.com/thuongtruong109/gouse/console/tab"
	"github.com/thuongtruong109/gouse/cookbook/console/table"
)

/*
Description: Run a command in the console
Input params: (...command string)
*/
func ConsoleCmd() {
	gouse.Cmd("echo command is working")

	// first param is default command, second param is windows command (default is empty)
	gouse.Cmd("ls", "clear")
}

/*
Description: Clear the console
*/
func ConsoleClear() {
	println("console will be cleared now")
	gouse.Cls()
	println("console cleared")
}

/*
Description: Print with color in the console
Input params: (color string, text string)
*/
func ConsoleWithColor() {
	gouse.OutputColor(gouse.DEFAULT_CONSOLE, "this is default")
	gouse.OutputColor(gouse.WHITE_CONSOLE, "this is white")
	gouse.OutputColor(gouse.RED_CONSOLE, "this is red")
	gouse.OutputColor(gouse.GREEN_CONSOLE, "this is green")
	gouse.OutputColor(gouse.YELLOW_CONSOLE, "this is yellow")
	gouse.OutputColor(gouse.BLUE_CONSOLE, "this is blue")
	gouse.OutputColor(gouse.MAGENTA_CONSOLE, "this is magenta")
	gouse.OutputColor(gouse.CYAN_CONSOLE, "this is cyan")
}

/*
Description: Display a banner in the console
Input params: (font string, text string)
*/
func ConsoleBanner() {
	gouse.Banner(gouse.DOUBLE_ALPHA, "gouse - type double")
	gouse.Banner(gouse.DOUBLE_ALPHA, "gouse - type single")
}

/*
Description: Display a help menu in the console
Input params: (name string, options []*gouse.IHelpOptions)
*/
func ConsoleHelp() {
	name := "myprogram"
	options := []*gouse.IHelpOptions{
		{
			Opt:  "name",
			Desc: "Enter your name",
			Val:  "Example name",
			Action: func(name string) {
				println("Your name is: ", name)
			},
		},
		{
			Opt:  "age",
			Desc: "Enter your age",
			Val:  18,
			Action: func(age int) {
				println("this is age: ", age)
			},
		},
		{
			Opt:  "learning",
			Desc: "Enter your confirm",
			Val:  true,
			Action: func(confirm bool) {
				println("this is your confirm", confirm)
			},
		},
	}
	gouse.Help(name, options)
	// for _, option := range options {
	// 	gouse.Printf("Option %s: %v\n", option.Opt, option.Val)
	// }
}

/*
Description: Display a select menu in the console
*/
func ConsoleSelect() {
	optconsolens := []string{"a", "b", "c"}
	selected, err := gouse.Select("Select an optconsolen:", optconsolens)
	if err != nil {
		panic(err)
	}
	gouse.Confirm("Are you sure?")
	println("You selected: ", selected)
}

func ConsoleList() {
	title := "My Fave Things"
	items := []list.Item{
		{Label: "Raspberry Pi’s", Desc: "I have ’em all over my house"},
		{Label: "Nutella", Desc: "It's good on toast"},
		{Label: "Bitter melon", Desc: "It cools you down"},
		{Label: "Nice socks", Desc: "And by that I mean socks without holes"},
		{Label: "Eight hours of sleep", Desc: "I had this once"},
		{Label: "Cats", Desc: "Usually"},
		{Label: "Plantasia, the album", Desc: "My plants love it too"},
		{Label: "Pour over coffee", Desc: "It takes forever to make though"},
		{Label: "VR", Desc: "Virtual reality...what is there to say?"},
		{Label: "Noguchi Lamps", Desc: "Such pleasing organic forms"},
		{Label: "Linux", Desc: "Pretty much the best OS"},
		{Label: "Business school", Desc: "Just kidding"},
		{Label: "Pottery", Desc: "Wet clay is a great feeling"},
		{Label: "Shampoo", Desc: "Nothing like clean hair"},
		{Label: "Table tennis", Desc: "It’s surprisingly exhausting"},
		{Label: "Milk crates", Desc: "Great for packing in your extra stuff"},
		{Label: "Afternoon tea", Desc: "Especially the tea sandwich part"},
		{Label: "Stickers", Desc: "The thicker the vinyl the better"},
		{Label: "20° Weather", Desc: "Celsius, not Fahrenheit"},
		{Label: "Warm light", Desc: "Like around 2700 Kelvin"},
		{Label: "The vernal equinox", Desc: "The autumnal equinox is pretty good too"},
		{Label: "Gaffer’s tape", Desc: "Basically sticky fabric"},
		{Label: "Terrycloth", Desc: "In other words, towel fabric"},
	}

	list.Default(title, items)
}

// func ConsolePaper() {
// 	paper.Run("main.go")
// }

func ConsoleProgress() {
	// first param is fail message
	// second param is done message
	// third param is increment percent (default 0.25)
	progress.Run("^_^ Oh no, something went wrong", "✔️ Done!", 0.5)
}

func ConsoleRealtime() {
	realtime.Run()
}

func ConsoleChoice() {
	question := "What's your favorite flavor?"
	options := []string{"Taro", "Coffee", "Lychee"}

	update := &choice.Model{}
	choice.Run(question, options, update)

	gouse.Printf("\n---\nYou chose %s!\n", update.Choice)
}

func ConsoleSpinner() {
	spinner.Run()
}

func ConsoleSplit() {
	split.Run()
}

func ConsoleStopwatch() {
	stopwatch.Run()
}

func ConsoleTable() {
	table.Run()
}

// func ConsoleTab() {
// 	// tabContent items must be same length with tabs items
// 	// and according to similar order in tabs
// 	tabs := []string{"Lip Gloss", "Blush", "Eye Shadow", "Mascara", "Foundation"}
// 	tabContent := []string{"Lip Gloss Tab", "Blush Tab", "Eye Shadow Tab", "Mascara Tab", "Foundation Tab"}

// 	println("Use arrow keys to switch tabs. Press q to quit.")

// 	tab.Run(tabs, tabContent)
// }

func ConsoleCountdown() {
	countdown.Run()
}

func ConsoleSequence() {
	sequence.Run()
}

func ConsoleInline() {
	customMode := &inline.Mode{
		AltscreenMode: " altscreen mode ", // string will be displayed when switch to altscreen mode
		InlineMode:    " inline mode ",    // string will be displayed when switch to inline mode
		Key:           " ",                // space key to switch mode
	}
	inline.Run(customMode)
}

func ConsoleParallel() {
	parallel.Run()
}

func ConsoleDir() {
	dir.Run()
}

func ConsoleGlamour() {
	const content = `
# Today’s Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] René Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appétit!
`
	glamour.Run(content)
}
