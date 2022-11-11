package workers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

// var (
// 	uploadCountMetrics  *prometheus.CounterVec
// 	uploadTimeMetrics   prometheus.Histogram
// 	uploadStatusMetrics *prometheus.GaugeVec
// )

func UploadTask(task *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResultFromTask(task)

	// Increment upload_task_count
	// uploadCountMetrics.Inc()
	taskCountMetrics.WithLabelValues(task.TaskDefName).Inc()

	log.Infof("received file name: %s", task.InputData["fileName"].(string))
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log:         fmt.Sprintf("received file name: %s", task.InputData["fileName"].(string)),
			TaskId:      task.TaskId,
			CreatedTime: time.Now().Unix(),
		},
	)
	// time.Sleep(1 * time.Minute)
	fileurl := "https://161.185.160.93/storage/" + task.InputData["fileName"].(string)

	min := 30
	max := 60
	uploadtimeInSec := rand.Intn(max-min) + min

	// uploadTimeMetrics.WithLabelValues(
	// 	task.TaskId,
	// 	task.TaskDefName,
	// 	task.WorkflowInstanceId,
	// ).Set(float64(uploadtimeInSec))
	// uploadTimeMetrics.Observe(float64(uploadtimeInSec))

	if uploadtimeInSec > 54 {
		// Failure status 20%
		taskStatusMetrics.WithLabelValues(
			task.TaskId,
			task.TaskDefName,
			task.WorkflowInstanceId,
			"FAILED",
		).Set(1)

	} else {
		// Successful status 80%
		taskStatusMetrics.WithLabelValues(
			task.TaskId,
			task.TaskDefName,
			task.WorkflowInstanceId,
			"SUCCEEDED",
		).Set(1)
	}

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

// func init() {
// 	// uploadCountMetrics = promauto.NewCounter(prometheus.CounterOpts{
// 	// 	Name: "upload_wf_task_count",
// 	// 	Help: "The total number of upload tasks",
// 	// 	ConstLabels: prometheus.Labels{
// 	// 		"task": "upload_task",
// 	// 	},
// 	// })
// 	// if err := prometheus.Register(uploadCountMetrics); err != nil {
// 	// 	fmt.Println("upload_task_count is not registered")
// 	// }

// 	uploadCountMetrics = prometheus.NewCounterVec(
// 		prometheus.CounterOpts{
// 			Name: "upload_wf_task_count",
// 		},
// 		[]string{"taskdef"},
// 	)
// 	prometheus.MustRegister(uploadCountMetrics)

// 	// uploadTimeMetrics = prometheus.NewGaugeVec(
// 	// 	prometheus.GaugeOpts{
// 	// 		Name: "upload_task_time_in_secs",
// 	// 		Help: "The time taken for the upload",
// 	// 	},
// 	// 	[]string{"taskid", "taskdef", "workflowid"},
// 	// )
// 	// prometheus.MustRegister(uploadTimeMetrics)

// 	// Histogram
// 	uploadTimeMetrics = prometheus.NewHistogram(prometheus.HistogramOpts{
// 		Name:    "upload_task_time_in_secs",
// 		Help:    "The time taken for the upload",
// 		Buckets: prometheus.LinearBuckets(25, 5, 9), // 9 buckets, each 5 secs wide. upto 65
// 	})
// 	prometheus.MustRegister(uploadTimeMetrics)

// 	uploadStatusMetrics = prometheus.NewGaugeVec(
// 		prometheus.GaugeOpts{
// 			Name: "upload_wf_task_status",
// 			// Help: "The status of the upload task",
// 			// ConstLabels: prometheus.Labels{
// 			// 	"task": "upload_task",
// 			// },
// 		},
// 		[]string{"taskid", "taskdef", "workflowid", "uploadStatus"},
// 	)
// 	prometheus.MustRegister(uploadStatusMetrics)
// }
