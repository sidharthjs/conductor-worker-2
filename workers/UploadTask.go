package workers

import (
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

func UploadTask(task *model.Task) (interface{}, error) {
	// fmt.Println("Executing upload task")
	// fmt.Println("Uploading the file: ", task.InputData["fileName"])
	// fmt.Println(task)
	taskResult := model.NewTaskResultFromTask(task)

	log.Infof("received file name: %s", task.InputData["fileName"].(string))
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log:         fmt.Sprintf("received file name: %s", task.InputData["fileName"].(string)),
			TaskId:      task.TaskId,
			CreatedTime: time.Now().Unix(),
		},
	)
	time.Sleep(1 * time.Minute)
	fileurl := "https://161.185.160.93/storage/" + task.InputData["fileName"].(string)

	log.Infof("successfully uploaded file. url: %s", fileurl)
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log:         fmt.Sprintf("successfully uploaded file. url: %s", fileurl),
			TaskId:      task.TaskId,
			CreatedTime: time.Now().Unix(),
		},
	)

	taskResult.OutputData = map[string]interface{}{
		"file_url":   fileurl,
		"ip_address": "161.185.160.93",
	}
	taskResult.Status = model.CompletedTask

	return taskResult, nil
}
