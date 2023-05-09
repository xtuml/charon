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

Typically this will require a `config.json` and job definition json files.

Example curl
```bash
curl --location --request POST '127.0.0.1:9000/upload/aeo-svdc-config' \
--form 'upload=@"/Users/tomgibbs/testFiles/aeo_svdc/config/config.json"' \
--form 'upload=@"/Users/testFiles/aeo_svdc/config/FileRequest.json"' \
--form 'upload=@"/Users/tomgibbs/testFiles/aeo_svdc/config/FileRequest_event_data.json"'
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