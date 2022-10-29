package main

import (
	"os"
	"strconv"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/sidharthjs/conductor-worker-2/workers"
	log "github.com/sirupsen/logrus"
)

func main() {
	conductorServer := os.Getenv("CONDUCTOR_SERVER")
	batchSizeEnv := os.Getenv("BATCH_SIZE")

	if conductorServer == "" {
		log.Fatal("conductor server url is invalid")
		// fmt.Println("conductor server url is invalid")
	}
	batchSize, err := strconv.Atoi(batchSizeEnv)
	if err != nil {
		log.Fatalf("batch size value invalid: %s", err)
		// fmt.Printf("batch size value invalid: %s\n", err)
	}

	apiClient := client.NewAPIClient(
		nil,
		settings.NewHttpSettings(
			conductorServer,
		),
	)
	// fmt.Println("sleeping")
	// time.Sleep(1 * time.Minute)

	tr := worker.NewTaskRunnerWithApiClient(apiClient)
	// tr.StartWorker("upload_file", workers.UploadTask, batchSize, 5*time.Second)
	tr.StartWorker("write_to_db", workers.WriteToDBTask, batchSize, 5*time.Second)
	log.Info("started workers")
	tr.WaitWorkers()
}
