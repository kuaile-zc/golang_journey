package inf

import (
	"flag"
	"fmt"
	"time"
)

type celsiusFlag struct{ celsius float64 }

func (f *celsiusFlag) String() string {
	return fmt.Sprintf("%g°C", f.celsius)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.celsius = value
		return nil
	case "F", "°F":
		f.celsius = (value - 32) * 5 / 9
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value float64, usage string) string {
	f := celsiusFlag{celsius: value}
	flag.CommandLine.Var(&f, name, usage)
	return f.String()
}

var period = flag.Duration("period", 1*time.Second, "sleep period")

func FlagTest() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
