package workers

import "github.com/prometheus/client_golang/prometheus"

var (
	taskCountMetrics *prometheus.CounterVec
	// uploadTimeMetrics   prometheus.Histogram
	taskStatusMetrics *prometheus.GaugeVec
)

func init() {
	// uploadCountMetrics = promauto.NewCounter(prometheus.CounterOpts{
	// 	Name: "upload_wf_task_count",
	// 	Help: "The total number of upload tasks",
	// 	ConstLabels: prometheus.Labels{
	// 		"task": "upload_task",
	// 	},
	// })
	// if err := prometheus.Register(uploadCountMetrics); err != nil {
	// 	fmt.Println("upload_task_count is not registered")
	// }

	taskCountMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "upload_wf_task_count",
		},
		[]string{"taskdef"},
	)
	prometheus.MustRegister(taskCountMetrics)

	// uploadTimeMetrics = prometheus.NewGaugeVec(
	// 	prometheus.GaugeOpts{
	// 		Name: "upload_task_time_in_secs",
	// 		Help: "The time taken for the upload",
	// 	},
	// 	[]string{"taskid", "taskdef", "workflowid"},
	// )
	// prometheus.MustRegister(uploadTimeMetrics)

	// Histogram
	// taskTimeMetrics = prometheus.NewHistogram(prometheus.HistogramOpts{
	// 	Name:    "upload_task_time_in_secs",
	// 	Help:    "The time taken for the upload",
	// 	Buckets: prometheus.LinearBuckets(25, 5, 9), // 9 buckets, each 5 secs wide. upto 65
	// })
	// prometheus.MustRegister(uploadTimeMetrics)

	taskStatusMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "upload_wf_task_status",
			// Help: "The status of the upload task",
			// ConstLabels: prometheus.Labels{
			// 	"task": "upload_task",
			// },
		},
		[]string{"taskid", "taskdef", "workflowid", "uploadStatus"},
	)
	prometheus.MustRegister(taskStatusMetrics)
}
