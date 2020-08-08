package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	help = flag.Bool("h", false, "show this usage information")

	u = flag.Bool("u", false, "output as UTC")
	f = flag.String("f", "Mon Jan 2 15:04 MST", "time format layout")
	l = flag.String("l", "Local", "comma-separated list of time zones to output")

	seconds = flag.Bool("s", false, "output as Unix time in seconds")
	rfc3339 = flag.Bool("rfc-3339", false, "output as RFC 3339 format")
)

func init() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	now := time.Now()
	if *seconds {
		fmt.Println(now.Unix())
		return
	}
	if *u {
		now = now.UTC()
		switch {
		case *rfc3339:
			fmt.Println(now.Format(time.RFC3339))
		case *l == "Local", *l == "UTC", *l == "":
			fmt.Println(now.Format(*f))
		default:
			flag.Usage()
			os.Exit(2)
		}
		return
	}
	if *rfc3339 {
		fmt.Println(now.Format(time.RFC3339))
		return
	}
	timezones := strings.Split(*l, ",")
	for _, tz := range timezones {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(now.In(loc).Format(*f))
	}
}
