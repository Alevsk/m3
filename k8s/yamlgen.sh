#!/bin/bash
# Get's the latest deployment file from MinIO Operator
curl https://raw.githubusercontent.com/minio/minio-operator/master/minio-operator.yaml > k8s/base/minio-operator.yaml
