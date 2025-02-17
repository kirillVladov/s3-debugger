# s3-debugger

## Description

S3 Debug Tool is a utility designed to simplify debugging and testing interactions with S3-compatible storage, such as Amazon S3 and MinIO. This tool allows you to perform various operations on objects

## Installation

### Requirements

- Go 1.23 or higher
- Access to an S3-compatible storage Amazon s3

### Installation Steps

1. Clone the repository:

   ```bash
   git clone git@github.com:kirillVladov/s3-debugger.git
   cd s3-debug-tool
2. Setup .env file

      ```bash
      ENDPOINT=
      REGION=
      ACCESS_KEY=
      SECRET_KEY=
      BUCKET_NAME=
      GETTING_COUNT=
3. Build
    ```bash
    make build
3. Run
    ```bash
    make check-latency