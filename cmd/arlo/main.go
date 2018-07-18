package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/juju/persistent-cookiejar"
	"github.com/spf13/cobra"

	"github.com/axw/go-arlo"
)

var (
	arloClient *arlo.Client
)

func newCookieJar() *cookiejar.Jar {
	// TODO(axw) adhere to XDG base dir specification.
	homeDir := os.Getenv("HOME")
	arloDir := filepath.Join(homeDir, ".arlo")
	cookieFile := filepath.Join(arloDir, "cookies.json")
	jar, err := cookiejar.New(&cookiejar.Options{
		Filename: cookieFile,
	})
	if err != nil {
		panic(err)
	}
	return jar
}

func Main() error {
	jar := newCookieJar()
	arloClient = &arlo.Client{
		Client: &http.Client{Jar: jar},
	}
	defer jar.Save()

	rootCmd := cobra.Command{Use: "arlo"}
	rootCmd.AddCommand(
		cmdLogin,
		cmdLogout,
		cmdProfile,
		cmdLocations,
		cmdDevices,
	)
	return rootCmd.Execute()
}

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}
