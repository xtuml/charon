# protocol verifier http server

## Paths

### /

This will be for a webpage front end which is still in development.

### /health

The health endpoint will return a `200 OK` when the application is running. 

Example curl
```bash
curl --location --request GET '127.0.0.1:9000/health'
```

### /upload

#### /upload/aer-config

- Single file upload only

This path allows you to upload a single `config.json` to define the configuration for the AEReception application

Example curl
```bash
curl --location --request POST '127.0.0.1:9000/upload/aer-config' \
--form 'upload=@"/Users/tomgibbs/repo/smartdcs/cdsdt/protocol-verifier-http-server/testFiles/aereception/config.json"'
```

#### /upload/aeo-svdc-config

- Multi file upload

This path allows you to upload multiple files following the Protocol Verifiers guide to setting AEO_SVDC configuration.

Typically this will require a `config.json`.

Example curl
```bash
curl --location --request POST '127.0.0.1:9000/upload/aeo-svdc-config' \
--form 'upload=@"/Users/tomgibbs/testFiles/aeo_svdc/config/config.json"'
```

#### /upload/job-definitions

- Multi file upload

This path allows you to upload multiple files following the Protocol Verifiers guide to setting job definitions.

Typically this will require job definition json files.

Example curl
```bash
curl --location --request POST '127.0.0.1:9000/upload/aeo-svdc-config' \
--form 'upload=@"./FileRequest.json"' \
--form 'upload=@"./ANDFork_ANDFork_a.json"'
```

#### /upload/events

- Multi file upload

Allows posting events to AERecption incoming directory for consumption by the protocol verifier.

Example curl
```bash
curl --location --request POST '127.0.0.1:9000/upload/events' \
--form 'upload=@"/Users/tomgibbs/repo/smartdcs/cdsdt/protocol-verifier-http-server/testFiles/events/FileRequest_HappyPath.json"' \
--form 'upload=@"/Users/tomgibbs/repo/smartdcs/cdsdt/protocol-verifier-http-server/testFiles/events/FileRequest_LoopConstraintViolation.json"'
```

#### /download/verifierlog

This POST request path allows the download of a single log from the verifier domain file (zipped or unzipped) given the files name as a JSON request with mime type application/json with JSON of the form 
```
{
    "fileName":"<given filename>"
}
```

Example curl
```bash
curl --location --request POST \
-d '{"fileName":"Verifier.log"}' \
-H 'Content-Type: application/json' \ 
'127.0.0.1:9000/download/verifierlog'
```

#### /download/aerlog

This POST request path allows the download of a single log from the AER domain file (zipped or unzipped) given the files name as a JSON request with mime type application/json with JSON of the form 
```
{
    "fileName":"<given filename>"
}
```

Example curl
```bash
curl --location --request POST \
-d '{"fileName":"Reception.log"}' \
-H 'Content-Type: application/json' \ 
'127.0.0.1:9000/download/aerlog'
```

#### /download/verifier-log-file-names

This GET request path responds with a JSON response that contains a list of log files that are in the verifier domain that start with the string "Verifier". The JSON is of the form
```
{
    "fileName": [
        "Verifier.log",
        "Verifier.log.gz"
    ]
}
```

Example curl
```bash
curl --location --request GET \
'127.0.0.1:9000/download/verifier-log-file-names'
```

#### /download/verifier-log-file-names

This GET request path responds with a JSON response that contains a list of log files that are in the AER domain that start with the string "Reception". The JSON is of the form
```
{
    "fileName": [
        "Reception.log",
        "Reception.log.gz"
    ]
}
```

Example curl
```bash
curl --location --request GET \
'127.0.0.1:9000/download/aer-log-file-names'
```

#### /ioTracking/aer-incoming

GET request path that provides a JSON response with the number of files in the AER incoming folder and a timestamp in UNIX format (i.e. the number of nanoseconds since 1970) the reading was taken at. The JSON response has the following form
```
{
    "num_files": 1,
    "t": 123456
}
```

Example curl
```bash
curl --location --request GET \
'127.0.0.1:9000/ioTracking/aer-incoming'
```

#### /ioTracking/verifier-processed

GET request path that provides a JSON response with the number of files in the verifier processed folder and a timestamp in UNIX format (i.e. the number of nanoseconds since 1970) the reading was taken at. The JSON response has the following form
```
{
    "num_files": 1,
    "t": 123456
}
```

Example curl
```bash
curl --location --request GET \
'127.0.0.1:9000/ioTracking/verifier-processed'
```

#### /io/cleanup-test

GET request path that cleans up the folders of the protocol verifier. WARNING should only be used once a test is completely finished.


Example curl
```bash
curl --location --request GET \
'127.0.0.1:9000/io/cleanup-test'
```