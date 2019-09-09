package cmd

import (
    "os"
	"os/signal"
	"syscall"
	"fmt"

    "github.com/spf13/cobra"

    "github.com/mikaponics/mikaponics-iam/internal"
)

var rootCmd = &cobra.Command{
    Use:   "mikaponics-iam",
    Short: "Mikaponics Identity and Access Management (IAM) Web Service",
    Long: `IAM is a web service that helps you securely control access to Mikaponics resources. You use IAM to control who is authenticated (signed in) and authorized (has permissions) to use resources.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Load up our `environment variables` from our operating system.
        dbHost := os.Getenv("MIKAPONICS_DB_HOST")
        dbPort := os.Getenv("MIKAPONICS_DB_PORT")
        dbUser := os.Getenv("MIKAPONICS_DB_USER")
        dbPassword := os.Getenv("MIKAPONICS_DB_PASSWORD")
        dbName := os.Getenv("MIKAPONICS_DB_NAME")
        webServerAddress := os.Getenv("MIKAPONICS_APP_ADDRESS")

        // Initialize our application.
        app := internal.InitMikaponicsIAM(dbHost, dbPort, dbUser, dbPassword, dbName, webServerAddress)

        // DEVELOPERS CODE:
    	// The following code will create an anonymous goroutine which will have a
    	// blocking chan `sigs`. This blocking chan will only unblock when the
    	// golang app receives a termination command; therfore the anyomous
    	// goroutine will run and terminate our running application.
    	//
    	// Special Thanks:
    	// (1) https://gobyexample.com/signals
    	// (2) https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
    	//
    	sigs := make(chan os.Signal, 1)
    	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    	go func() {
            <-sigs // Block execution until signal from terminal gets triggered here.
            fmt.Println("Starting graceful shut down now.")
    		app.StopMainRuntimeLoop()
        }()

    	app.RunMainRuntimeLoop()
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
