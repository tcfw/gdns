package main

// BEFORE RUNNING:
// ---------------
// 1. If not already done, enable the Google Cloud DNS API
//    and check the quota for your project at
//    https://console.developers.google.com/apis/api/dns
// 2. This sample uses Application Default Credentials for authentication.
//    If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run
//    `gcloud beta auth application-default login`.
//    For more information, see
//    https://developers.google.com/identity/protocols/application-default-credentials
// 3. Install and update the Go dependencies by running `go get -u` in the
//    project directory.

import (
	"log"
	"time"

	"github.com/spf13/cobra"
)

var (
	// Identifies the GCP project
	project string

	// Identifies the managed zone in the GCP. Can be the managed zone name or id.
	managedZone string

	// Identifies the record name in the zone
	recordName string

	// Logs when not updating and other useless messages
	verbose bool

	// Ticker duration used for how often to check for IP changes
	tickerDuration time.Duration
)

var rootCmd = &cobra.Command{
	Use:   "gdns",
	Short: "GNS Updates records in GCP Cloud DNS similar to dynamic DNS",
	Run: func(cmd *cobra.Command, args []string) {
		project, _ = cmd.Flags().GetString("project")
		managedZone, _ = cmd.Flags().GetString("zone")
		recordName, _ = cmd.Flags().GetString("record")
		verbose, _ = cmd.Flags().GetBool("verbose")
		tickerDuration, _ = cmd.Flags().GetDuration("duration")
		startWatch()
	},
}

func main() {
	rootCmd.Flags().StringP("project", "p", "", "GCP Project")
	rootCmd.Flags().StringP("zone", "m", "", "GCP Managed Zone")
	rootCmd.Flags().StringP("record", "r", "", "GCP Managed Zone")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose")
	rootCmd.Flags().DurationP("duration", "d", 30*time.Second, "Check duration")
	rootCmd.Execute()
}

func startWatch() {
	curIP, err := getCurrentIP()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Last recorded:", curIP)

	t := time.NewTicker(tickerDuration)
	for {
		select {
		case <-t.C:
			changed, nip, err := checkIP(curIP)
			if err != nil {
				log.Fatal(err)
			}
			if changed {
				log.Println("Updating:", nip)
				updateGCP(curIP, nip)
				curIP = nip
			} else if verbose {
				log.Println("No update. Next check in", tickerDuration.String())
			}
		}
	}
}

func checkIP(curIP string) (bool, string, error) {
	nIP, err := getPublicIP()
	if err != nil {
		return false, "", err
	}

	return nIP != curIP, nIP, nil
}
