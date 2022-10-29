# conductor-worker-2

## Steps to run a workflow:

### Create the tasks definitions:

1. Upload file task:
```
{
  "createTime": 1666094354937,
  "updateTime": 1666094435110,
  "createdBy": "",
  "updatedBy": "",
  "accessPolicy": {},
  "name": "upload_file",
  "description": "Upload file task",
  "retryCount": 3,
  "timeoutSeconds": 3600,
  "inputKeys": [],
  "outputKeys": [],
  "timeoutPolicy": "TIME_OUT_WF",
  "retryLogic": "FIXED",
  "retryDelaySeconds": 60,
  "responseTimeoutSeconds": 600,
  "inputTemplate": {},
  "rateLimitPerFrequency": 0,
  "rateLimitFrequencyInSeconds": 1,
  "ownerEmail": "orkes@orkes.io",
  "backoffScaleFactor": 1
}
```

2. Write-to-db task:
```
{
  "createTime": 1666097404782,
  "createdBy": "",
  "accessPolicy": {},
  "name": "write_to_db",
  "description": "Edit or extend this sample task. Set the task name to get started",
  "retryCount": 3,
  "timeoutSeconds": 3600,
  "inputKeys": [],
  "outputKeys": [],
  "timeoutPolicy": "TIME_OUT_WF",
  "retryLogic": "FIXED",
  "retryDelaySeconds": 60,
  "responseTimeoutSeconds": 600,
  "inputTemplate": {},
  "rateLimitPerFrequency": 0,
  "rateLimitFrequencyInSeconds": 1,
  "ownerEmail": "orkes@orkes.com",
  "backoffScaleFactor": 1
}
```

### Create the workflow definition:

1. Upload workflow:
```
{
  "createTime": 1666094624930,
  "updateTime": 1666097559034,
  "accessPolicy": {},
  "name": "Upload_Cycle",
  "description": "Edit or extend this sample workflow. Set the workflow name to get started",
  "version": 3,
  "tasks": [
    {
      "name": "upload_file",
      "taskReferenceName": "upload_file",
      "inputParameters": {
        "fileName": "${workflow.input.file_name}"
      },
      "type": "SIMPLE",
      "startDelay": 0,
      "optional": false,
      "asyncComplete": false
    },
    {
      "name": "Get_IP",
      "taskReferenceName": "get_IP",
      "inputParameters": {
        "http_request": {
          "uri": "http://ip-api.com/json/${upload_file.output.ip_address}?fields=status,message,country,countryCode,region,regionName,city,zip,lat,lon,timezone,offset,isp,org,as,query",
          "method": "GET",
          "connectionTimeOut": "0",
          "readTimeOut": "0"
        }
      },
      "type": "HTTP",
      "startDelay": 0,
      "optional": false,
      "asyncComplete": false
    },
    {
      "name": "write_to_db",
      "taskReferenceName": "write_to_db",
      "inputParameters": {
        "fileName": "${workflow.input.file_name}",
        "country": "${get_IP.output.response.body.country}",
        "file_url": "${upload_file.output.file_url}"
      },
      "type": "SIMPLE",
      "startDelay": 0,
      "optional": false,
      "asyncComplete": false
    }
  ],
  "inputParameters": [],
  "outputParameters": {
    "file_url": "${upload_file.output.file_url}",
    "ip_address": "${upload_file.output.ip_address}",
    "country": "${get_IP.output.response.body.country}",
    "db_write_status": "${write_to_db.output.db_write_status}"
  },
  "schemaVersion": 2,
  "restartable": true,
  "workflowStatusListenerEnabled": false,
  "ownerEmail": "example@email.com",
  "timeoutPolicy": "ALERT_ONLY",
  "timeoutSeconds": 0,
  "variables": {},
  "inputTemplate": {}
}
```

### Run the workflow:

A workflow is triggered in the workbench page in conductor UI.

1. Workflow name: `Upload_Cycle`
2. Workflow version: `3`
3. Input json:
```
{
  "file_name": "shipments.db"
}
```