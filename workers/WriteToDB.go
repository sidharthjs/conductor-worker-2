package workers

import (
	"math/rand"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

// var (
// 	dbWriteCountMetrics  *prometheus.CounterVec
// 	dbWriteTimeMetrics   *prometheus.GaugeVec
// 	dbWriteStatusMetrics *prometheus.GaugeVec
// )

func WriteToDBTask(task *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResultFromTask(task)

	// Increment db_write_task_count
	// dbWriteCountMetrics.Inc()
	taskCountMetrics.WithLabelValues(task.TaskDefName).Inc()

	log.Infof("writing to db")
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log:         "writing to db",
			TaskId:      task.TaskId,
			CreatedTime: time.Now().Unix(),
		},
	)

	// time.Sleep(1 * time.Minute)

	min := 1
	max := 6
	dbWriteTimeInSec := rand.Intn(max-min) + min
	// dbWriteTimeMetrics.WithLabelValues(
	// 	task.TaskId,
	// 	task.TaskDefName,
	// 	task.WorkflowInstanceId,
	// ).Set(float64(dbWriteTimeInSec))

	if dbWriteTimeInSec > 4 {
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

// func init() {
// 	// dbWriteCountMetrics = promauto.NewCounter(prometheus.CounterOpts{
// 	// 	Name: "upload_wf_task_count",
// 	// 	Help: "The total number of db write tasks",
// 	// 	ConstLabels: prometheus.Labels{
// 	// 		"task": "db_write_task",
// 	// 	},
// 	// })
// 	// if err := prometheus.Register(dbWriteCountMetrics); err != nil {
// 	// 	fmt.Println("db_write_task_count is not registered")
// 	// }

// 	dbWriteCountMetrics = prometheus.NewCounterVec(
// 		prometheus.CounterOpts{
// 			Name: "upload_wf_task_count",
// 		},
// 		[]string{"taskdef"},
// 	)
// 	prometheus.Register(dbWriteCountMetrics)

// 	dbWriteTimeMetrics = prometheus.NewGaugeVec(
// 		prometheus.GaugeOpts{
// 			Name: "db_write_time_in_secs",
// 			Help: "The time taken for the db write",
// 		},
// 		[]string{"taskid", "taskdef", "workflowid"},
// 	)
// 	prometheus.MustRegister(dbWriteTimeMetrics)

// 	dbWriteStatusMetrics = prometheus.NewGaugeVec(
// 		prometheus.GaugeOpts{
// 			Name: "upload_wf_task_status",
// 			// Help: "The status of the db write task",
// 			// ConstLabels: prometheus.Labels{
// 			// 	"task": "db_write_task",
// 			// },
// 		},
// 		[]string{"taskid", "taskdef", "workflowid"},
// 	)
// 	prometheus.Register(dbWriteStatusMetrics)
// }
