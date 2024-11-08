/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package support

import (
	"os"
	"strings"
)

const (
	// The environment variables hereafter can be used to change the components
	// used for testing.

	CodeFlareTestRayVersion   = "CODEFLARE_TEST_RAY_VERSION"
	CodeFlareTestRayImage     = "CODEFLARE_TEST_RAY_IMAGE"
	CodeFlareTestPyTorchImage = "CODEFLARE_TEST_PYTORCH_IMAGE"

	// The testing output directory, to write output files into.
	CodeFlareTestOutputDir = "CODEFLARE_TEST_OUTPUT_DIR"

	// The namespace where a secret containing InstaScale OCM token is stored and the secret name.
	InstaScaleOcmSecret = "INSTASCALE_OCM_SECRET"

	// Cluster ID for OSD cluster used in tests, used for testing InstaScale
	ClusterID = "CLUSTERID"

	// Type of cluster test is run on
	ClusterTypeEnvVar = "CLUSTER_TYPE"

	// Hostname of the Kubernetes cluster
	ClusterHostname = "CLUSTER_HOSTNAME"

	// URL for downloading MNIST dataset
	mnistDatasetURL = "MNIST_DATASET_URL"

	// URL for PiPI index containing all the required test Python packages
	pipIndexURL    = "PIP_INDEX_URL"
	pipTrustedHost = "PIP_TRUSTED_HOST"

	// Storage bucket credentials
	storageDefaultEndpoint = "AWS_DEFAULT_ENDPOINT"
	storageDefaultRegion   = "AWS_DEFAULT_REGION"
	storageAccessKeyId     = "AWS_ACCESS_KEY_ID"
	storageSecretKey       = "AWS_SECRET_ACCESS_KEY"
	storageBucketName      = "AWS_STORAGE_BUCKET"
	storageBucketMnistDir  = "AWS_STORAGE_BUCKET_MNIST_DIR"
)

type ClusterType string

const (
	OsdCluster        ClusterType = "OSD"
	OcpCluster        ClusterType = "OCP"
	HypershiftCluster ClusterType = "HYPERSHIFT"
	KindCluster       ClusterType = "KIND"
	UndefinedCluster  ClusterType = "UNDEFINED"
)

func GetRayVersion() string {
	return lookupEnvOrDefault(CodeFlareTestRayVersion, RayVersion)
}

func GetRayImage() string {
	return lookupEnvOrDefault(CodeFlareTestRayImage, RayImage)
}

func GetRayROCmImage() string {
	return lookupEnvOrDefault(CodeFlareTestRayImage, RayROCmImage)
}

func GetRayTorchCudaImage() string {
	return lookupEnvOrDefault(CodeFlareTestRayImage, RayTorchCudaImage)
}

func GetRayTorchROCmImage() string {
	return lookupEnvOrDefault(CodeFlareTestRayImage, RayTorchROCmImage)
}

func GetPyTorchImage() string {
	return lookupEnvOrDefault(CodeFlareTestPyTorchImage, "pytorch/pytorch:1.11.0-cuda11.3-cudnn8-runtime")
}

func GetInstascaleOcmSecret() (string, string) {
	res := strings.SplitN(lookupEnvOrDefault(InstaScaleOcmSecret, "default/instascale-ocm-secret"), "/", 2)
	return res[0], res[1]
}

func GetClusterId() (string, bool) {
	return os.LookupEnv(ClusterID)
}

func GetClusterType(t Test) ClusterType {
	clusterType, ok := os.LookupEnv(ClusterTypeEnvVar)
	if !ok {
		t.T().Logf("Environment variable %s is unset.", ClusterTypeEnvVar)
		return UndefinedCluster
	}
	switch clusterType {
	case "OSD":
		return OsdCluster
	case "OCP":
		return OcpCluster
	case "HYPERSHIFT":
		return HypershiftCluster
	case "KIND":
		return KindCluster
	default:
		t.T().Logf("Environment variable %s is unset or contains an incorrect value: '%s'", ClusterTypeEnvVar, clusterType)
		return UndefinedCluster
	}
}

func GetClusterHostname(t Test) string {
	hostname, ok := os.LookupEnv(ClusterHostname)
	if !ok {
		t.T().Fatalf("Expected environment variable %s not found, please define cluster hostname.", ClusterHostname)
	}
	return hostname
}

func GetMnistDatasetURL() string {
	return lookupEnvOrDefault(mnistDatasetURL, "https://ossci-datasets.s3.amazonaws.com/mnist/")
}

func GetStorageBucketDefaultEndpoint() (string, bool) {
	storage_endpoint, exists := os.LookupEnv(storageDefaultEndpoint)
	return storage_endpoint, exists
}

func GetStorageBucketDefaultRegion() (string, bool) {
	storage_default_region, exists := os.LookupEnv(storageDefaultRegion)
	return storage_default_region, exists
}

func GetStorageBucketAccessKeyId() (string, bool) {
	storage_access_key_id, exists := os.LookupEnv(storageAccessKeyId)
	return storage_access_key_id, exists
}

func GetStorageBucketSecretKey() (string, bool) {
	storage_secret_key, exists := os.LookupEnv(storageSecretKey)
	return storage_secret_key, exists
}

func GetStorageBucketName() (string, bool) {
	storage_bucket_name, exists := os.LookupEnv(storageBucketName)
	return storage_bucket_name, exists
}

func GetStorageBucketMnistDir() (string, bool) {
	storage_bucket_mnist_dir, exists := os.LookupEnv(storageBucketMnistDir)
	return storage_bucket_mnist_dir, exists
}

func GetPipIndexURL() string {
	return lookupEnvOrDefault(pipIndexURL, "https://pypi.python.org/simple")
}

func GetPipTrustedHost() string {
	return lookupEnvOrDefault(pipTrustedHost, "")
}

func lookupEnvOrDefault(key, value string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return value
}
