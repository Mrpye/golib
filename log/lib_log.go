// Package log provides a set of functions for logging to the console and storing the output in a buffer
package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/drewstinnett/gout/v2"
	gout_json "github.com/drewstinnett/gout/v2/formats/json"
	"github.com/drewstinnett/gout/v2/formats/plain"
	"github.com/drewstinnett/gout/v2/formats/toml"
	"github.com/drewstinnett/gout/v2/formats/xml"
	"github.com/drewstinnett/gout/v2/formats/yaml"
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

// actionLog is a function that takes a string and a rune as arguments and prints a header with the string in the middle, and the
// rune as the border
// It also prints the date and time in the middle of the header
func actionLog(Header string, char rune, show_date bool, border_color int) {

	//conmpare rune
	if char == 0 {
		char = '*'
	}

	//****************************
	//cal the length of the header
	//****************************
	count := strings.Count(Header, "\x1b")
	str_len := len(Header) + 6
	if count > 0 {
		str_len = str_len - (count * 4) - 1
	}

	//*****************************************
	//See if the date is longer than the header
	//If it is, then use the date length
	//*****************************************
	dt_string := ""
	if show_date {
		dt := time.Now()
		dt_string = dt.Format("01-02-2006 15:04:05")
		if len(dt_string)+6 > str_len {
			str_len = len(dt_string) + 6
		}
	}

	color_blue := color.FgBlue.Render
	color_red := color.FgRed.Render
	color_green := color.FgGreen.Render
	//colour_blue := color.FgBlue.Render
	str_len_as_str := fmt.Sprintf("%d", str_len-6)
	border := ""

	// convert the str_len to a string minus the 6 characters for the "** " and " **"
	// create a string of the length of the border
	if border_color == 0 {
		border = color_blue(strings.Repeat(string(char), str_len))
	} else if border_color == 1 {
		border = color_red(strings.Repeat(string(char), str_len))
	} else if border_color == 2 {
		border = color_green(strings.Repeat(string(char), str_len))
	}

	if border_color == 0 {
		Header = fmt.Sprintf("%s %-"+str_len_as_str+"s %s", color_blue("**"), Header, color_blue("**")) // create a string for the header
	} else if border_color == 1 {
		Header = fmt.Sprintf("%s %-"+str_len_as_str+"s %s", color_red("**"), Header, color_red("**")) // create a string for the header
	} else if border_color == 2 {
		Header = fmt.Sprintf("%s %-"+str_len_as_str+"s %s", color_green("**"), Header, color_green("**")) // create a string for the header
	}

	if show_date {
		if border_color == 0 {
			dt_string = fmt.Sprintf("%s %-"+str_len_as_str+"s %s", color_blue("**"), dt_string, color_blue("**")) // create a string for the date
		} else if border_color == 1 {
			dt_string = fmt.Sprintf("%s %-"+str_len_as_str+"s %s", color_red("**"), dt_string, color_red("**")) // create a string for the date
		} else if border_color == 2 {
			dt_string = fmt.Sprintf("%s %-"+str_len_as_str+"s %s", color_green("**"), dt_string, color_green("**")) // create a string for the date
		}
		//********************
		//Print the action log
		//********************
		fmt.Printf("%s\n%s\n%s\n%s\n", border, Header, dt_string, border)
	} else {
		//********************
		//Print the action log
		//********************
		fmt.Printf("%s\n%s\n%s\n", border, Header, border)
	}
}

//ActionLogDT is a function that takes a string and a rune as arguments and prints a header with the string , and the
// rune as the border
// It also prints the date and time in the middle of the header
func ActionLogDT(Header string, char rune) {
	actionLog(Header, char, true, 0)
}

// It takes a string and a rune as input, and prints a header with the string in the middle, and the
// rune as the border
func ActionLog(Header string, char rune) {
	actionLog(Header, char, true, 0)
}

// ActionLogOK() is a function that takes a string and a rune as arguments and returns nothing
func ActionLogGreen(Header string, char rune) {
	color_green := color.FgGreen.Render
	actionLog(fmt.Sprintf("%s: %s", Header, color_green("OK")), char, true, 2)
}

// ActionLogFail() is a function that takes a string and a rune as arguments and prints a red "Fail"
// message to the console
func ActionLogRed(Header string, char rune) {
	color_red := color.FgRed.Render
	actionLog(fmt.Sprintf("%s: %s", Header, color_red("Fail")), char, true, 1)
}

// ActionLogOK() is a function that takes a string and a rune as arguments and returns nothing
func ActionLogOK(Header string, char rune) {
	color_green := color.FgGreen.Render
	actionLog(fmt.Sprintf("%s: %s", Header, color_green("OK")), char, false, 0)
}

// ActionLogFail() is a function that takes a string and a rune as arguments and prints a red "Fail"
// message to the console
func ActionLogFail(Header string, char rune) {
	color_red := color.FgRed.Render
	actionLog(fmt.Sprintf("%s: %s", Header, color_red("Fail")), char, false, 0)
}

// ActionLogOK() is a function that takes a string and a rune as arguments and returns nothing
func ActionLogDateOK(Header string, char rune) {
	color_green := color.FgGreen.Render
	actionLog(fmt.Sprintf("%s: %s", Header, color_green("OK")), char, true, 0)
}

// ActionLogFail() is a function that takes a string and a rune as arguments and prints a red "Fail"
// message to the console
func ActionLogDateFail(Header string, char rune) {
	color_red := color.FgRed.Render
	actionLog(fmt.Sprintf("%s: %s", Header, color_red("Fail")), char, true, 0)
}

// `PrintlnOK` prints a message in green color
func PrintlnOK(msg string) {
	color_green := color.FgGreen.Render
	log.Printf("%s: %s\n", msg, color_green("OK"))
}

// `PrintlnFail` is a function that takes a string as an argument and prints it to the console in red
func PrintlnFail(msg string) {
	color_red := color.FgRed.Render
	log.Printf("%s: %s\n", msg, color_red("Fail"))
}

// `PrintOK` prints a message in green color
func PrintOK(msg string) {
	color_green := color.FgGreen.Render
	log.Printf("%s: %s", msg, color_green("OK"))
}

// `PrintFail` is a function that takes a string as an argument and prints it in red
func PrintFail(msg string) {
	color_red := color.FgRed.Render
	log.Printf("%s: %s", msg, color_red("Fail"))
}

// `LogVerbose` is a function that takes a string and prints it to the console in magenta
func LogVerbose(msg string) {
	color_magenta := color.FgMagenta.Render
	fmt.Print(color_magenta(msg))
}

func LogInterfaceVerbose(data interface{}) {
	color_magenta := color.FgMagenta.Render
	msg, err := json.Marshal(data)
	if err != nil {
		color_red := color.FgRed.Render
		fmt.Println(color_red(err))
	}
	fmt.Print(color_magenta(msg))
}

// It takes a title, data, and format, and prints the data in the specified format
func FormatResults(title string, data interface{}, format string) error {
	color_magenta := color.FgMagenta.Render
	str_buf := fmt.Sprintf("%s\n", color_magenta(title))
	w, err := gout.New()
	if err != nil {
		panic(err)
	}

	switch strings.ToLower(format) {
	case "json":
		w.SetFormatter(gout_json.Formatter{})
	case "yaml":
		w.SetFormatter(yaml.Formatter{})
	case "toml":
		w.SetFormatter(toml.Formatter{})
	case "xml":
		w.SetFormatter(xml.Formatter{})
	case "plain":
		w.SetFormatter(plain.Formatter{})
	default:
		w.SetFormatter(gout_json.Formatter{})
	}

	b := new(strings.Builder)
	w.SetWriter(b)
	err = w.Print(data)
	if err != nil {
		return err
	}
	data = b.String()

	// Print it on out!
	w.Print(fmt.Sprintf("%s%s\n", str_buf, data))
	return nil
}
