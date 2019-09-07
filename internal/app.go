package internal // github.com/mikaponics/mikapod-iam/internal

import (
	// "context"
	"log"
	// "os"
	// "time"
	// "fmt"

	// "google.golang.org/grpc"
	// "github.com/golang/protobuf/ptypes/timestamp"

	"github.com/mikaponics/mikaponics-iam/internal/models"
)

type MikapodIAM struct {
	db *models.MikapodDB
	done chan bool
}

// Function will construct the Mikapod IAM application.
func InitMikapodIAM(dbHost, dbPort, dbUser, dbPassword, dbName, appAddress string) (*MikapodIAM) {

	// Initialize and connect our database layer for the entire application.
    dbInstance := models.InitDB(dbHost, dbPort, dbUser, dbPassword, dbName)

    // Create our app's models if they haven't been created previously.
    dbInstance.CreateUserTable(false)

	// Create our application instance.
 	return &MikapodIAM{
		db: dbInstance,
		done: make(chan bool, 1), // Create a execution blocking channel.
	}
}

// Function will consume the main runtime loop and run the business logic
// of the Mikapod IAM application.
func (app *MikapodIAM) RunMainRuntimeLoop() {
	defer app.shutdown()

	go func() {
        for {
            select {
				case <- app.done:
					log.Printf("Interrupted")
					return
			}
		}
	}()
	<-app.done
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (app *MikapodIAM) StopMainRuntimeLoop() {
	app.done <- true
}

func (app *MikapodIAM) shutdown()  {
	app.db.Shutdown()
    // app.storageCon.Close()
	// app.remoteCon.Close()                                  //TODO: Uncommment
}
