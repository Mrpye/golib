package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/drewstinnett/gout/v2"
	"github.com/gookit/color"
)

// LogStreamer is a io.Writer that logs to the console and stores the output in a buffer
type LogStreamer struct {
	b bytes.Buffer
}

// Gets the log streamer output as a string
func (l *LogStreamer) String() string {
	return l.b.String()
}

// Write implements the io.Writer interface
func (l *LogStreamer) Write(p []byte) (n int, err error) {
	a := strings.TrimSpace(string(p))
	l.b.WriteString(a + "\n")
	log.Println(a)
	return len(p), nil
}

// ActionLog is a function to log a header with a line of characters
func ActionLog(Header string, char rune) {
	//conmpare rune
	if char == 0 {
		char = '*'
	}
	count := strings.Count(Header, "\x1b")
	str_len := len(Header) + 6
	if count > 0 {
		str_len = str_len - (count * 4) - 1
	}
	b := make([]rune, str_len)
	for i := range b {
		b[i] = char
	}
	colour := color.FgBlue.Render

	log.Println(colour(string(b)))
	log.Println(colour("** "+Header) + colour(" **"))
	log.Println(colour(string(b)))
	//if char == '-' || char == '<' {
	//fmt.Println("")
	//fmt.Println("")
	//}
}

// ActionLogOK is a function to log a header with a line of characters and a OK
func ActionLogOK(Header string, char rune) {
	color := color.FgGreen.Render
	ActionLog(fmt.Sprintf("%s: %s", Header, color("OK")), char)
}

func ActionLogFail(Header string, char rune) {
	color := color.FgRed.Render
	ActionLog(fmt.Sprintf("%s: %s", Header, color("Fail")), char)
}

// PrintlnOK is a function to log a message with a OK with a new line
func PrintlnOK(msg string) {
	color := color.FgGreen.Render
	log.Printf("%s: %s\n", msg, color("OK"))
}

// PrintOK is a function to log a message with a OK
func PrintOK(msg string) {
	color := color.FgGreen.Render
	log.Printf("%s: %s", msg, color("OK"))
}

// PrintlnFail is a function to log a message with a Fail with a new line
func PrintlnFail(msg string) {
	color := color.FgRed.Render
	log.Printf("%s: %s\n", msg, color("Fail"))
}

// PrintFail is a function to log a message with a Fail
func PrintFail(msg string) {
	color := color.FgRed.Render
	log.Printf("%s: %s", msg, color("Fail"))
}

// LogVerbose prints a message to the console in magenta
func LogVerbose(msg string) {
	color := color.FgMagenta.Render
	fmt.Print(color(msg))

}

// FormatResults prints a message to the console in magenta and then prints the data in a table
func FormatResults(title string, data interface{}) error {
	color := color.FgMagenta.Render
	fmt.Printf("%s\n", color(title))
	w, err := gout.New()
	if err != nil {
		panic(err)
	}
	// By Default, print to stdout. Let's change it to stderr though
	w.SetWriter(os.Stderr)

	// Print it on out!
	w.Print(data)
	return nil
}
