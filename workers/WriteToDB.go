package workers

import (
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

func WriteToDBTask(task *model.Task) (interface{}, error) {
	// fmt.Println("Executing write_to_db task")
	// fmt.Println("File url: ", task.InputData["file_url"])
	// fmt.Println("File location: ", task.InputData["country"])
	taskResult := model.NewTaskResultFromTask(task)

	log.Infof("writing to db")
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log:         "writing to db",
			TaskId:      task.TaskId,
			CreatedTime: time.Now().Unix(),
		},
	)

	time.Sleep(1 * time.Minute)

	log.Infof("write db success")
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log:         "write db success",
			TaskId:      task.TaskId,
			CreatedTime: time.Now().Unix(),
		},
	)

	taskResult.OutputData = map[string]interface{}{
		"db_write_status": "SUCCESS",
	}
	taskResult.Status = model.CompletedTask

	return taskResult, nil
}
