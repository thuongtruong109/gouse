
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Console' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/gouse/console/choice"
	"github.com/thuongtruong109/gouse/console/countdown"
	"github.com/thuongtruong109/gouse/console/dir"
	"github.com/thuongtruong109/gouse/console/glamour"
	"github.com/thuongtruong109/gouse/console/inline"
	"github.com/thuongtruong109/gouse/console/list"
	"github.com/thuongtruong109/gouse/console/parallel"
	"github.com/thuongtruong109/gouse/console/progress"
	"github.com/thuongtruong109/gouse/console/realtime"
	"github.com/thuongtruong109/gouse/console/sequence"
	"github.com/thuongtruong109/gouse/console/spinner"
	"github.com/thuongtruong109/gouse/console/split"
	"github.com/thuongtruong109/gouse/console/stopwatch"
	"github.com/thuongtruong109/gouse/console/table"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. console cmd' />



```go
func ConsoleCmd() {
	gouse.Cmd("echo command is working")

	// first param is default command, second param is windows command (default is empty)
	gouse.Cmd("ls", "clear")
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. console clear' />



```go
func ConsoleClear() {
	println("console will be cleared now")
	gouse.Cls()
	println("console cleared")
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. console with color' />



```go
func ConsoleWithColor() {
	gouse.PrintColor(gouse.DEFAULT_FG, "this is default")
	gouse.PrintColor(gouse.WHITE_FG, "this is white")
	gouse.PrintColor(gouse.RED_FG, "this is red")
	gouse.PrintColor(gouse.GREEN_FG, "this is green")
	gouse.PrintColor(gouse.YELLOW_FG, "this is yellow")
	gouse.PrintColor(gouse.BLUE_FG, "this is blue")
	gouse.PrintColor(gouse.MAGENTA_FG, "this is magenta")
	gouse.PrintColor(gouse.CYAN_FG, "this is cyan")
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='4. console banner' />



```go
func ConsoleBanner() {
	// param1: font name, param2: your input string
	gouse.Banner(gouse.DOUBLE_ALPHA, "gouse - type double")
	gouse.Banner(gouse.DOUBLE_ALPHA, "gouse - type single")
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='5. console help' />



```go
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
	// 	fmt.Printf("Option %s: %v\n", option.Opt, option.Val)
	// }
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='6. console select' />



```go
func ConsoleSelect() {
	optconsolens := []string{"a", "b", "c"}
	selected, err := gouse.Select("Select an optconsolen:", optconsolens)
	if err != nil {
		panic(err)
	}
	gouse.Confirm("Are you sure?")
	println("You selected: ", selected)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='7. console list' />



```go
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
```

### <Badge style='font-size: 1.1rem;' type='tip' text='8. console progress' />



```go
func ConsoleProgress() {
	// first param is fail message
	// second param is done message
	// third param is increment percent (default 0.25)
	progress.Run("^_^ Oh no, something went wrong", "✔️ Done!", 0.5)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='9. console realtime' />



```go
func ConsoleRealtime() {
	realtime.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='10. console choice' />



```go
func ConsoleChoice() {
	question := "What's your favorite flavor?"
	options := []string{"Taro", "Coffee", "Lychee"}

	update := &choice.Model{}
	choice.Run(question, options, update)

	fmt.Printf("\n---\nYou chose %s!\n", update.Choice)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='11. console spinner' />



```go
func ConsoleSpinner() {
	spinner.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='12. console split' />



```go
func ConsoleSplit() {
	split.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='13. console stopwatch' />



```go
func ConsoleStopwatch() {
	stopwatch.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='14. console table' />



```go
func ConsoleTable() {
	table.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='15. console countdown' />



```go
func ConsoleCountdown() {
	countdown.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='16. console sequence' />



```go
func ConsoleSequence() {
	sequence.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='17. console inline' />



```go
func ConsoleInline() {
	customMode := &inline.Mode{
		AltscreenMode: " altscreen mode ", // string will be displayed when switch to altscreen mode
		InlineMode:    " inline mode ",    // string will be displayed when switch to inline mode
		Key:           " ",                // space key to switch mode
	}
	inline.Run(customMode)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='18. console parallel' />



```go
func ConsoleParallel() {
	parallel.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='19. console dir' />



```go
func ConsoleDir() {
	dir.Run()
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='20. console glamour' />



```go
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
```
