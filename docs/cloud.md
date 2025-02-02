## cloud.go Docs

### Connect()

Connect() is a function that connects to the Cloud Instance

- Parameters:
  - endpoint: string (ip and port of the instance)
  - accessKey: string (Access Key of the instance)
  - secretAccessKey: string (Secret access key of the instance)
- Returns:
  - Returns error (if occurred)

### UploadFile()

UploadFile() is a function that uploads file into Cloud

- Parameters:
  - bucketName: string (Name of the bucket where file will be uploaded)
  - objectName: string (Server side name of the uploaded file)
  - filePath: string (File location which will be uploade)
  - contentType: string (MIME type of the file)
- Returns:
  - Returns error (if occurred)

### DownloadFile()

DownloadFile() is a function that download file from Cloud

- Parameters:
  - bucketName: string (Name of the bucket where the remote file is stored)
  - objectName: string (Server side name of the file)
  - filePath: string (File name which will be stored)
- Returns:
  - Returns error (if occurred)
