package env

import (
	"github.com/opencost/opencost/core/pkg/env"
)

// FilePaths
const (
	ClusterInfoFile = "cluster-info.json"
	ClusterCacheFile
	GCPAuthSecretFile        = "key.json"
	MetricConfigFile         = "metrics.json"
	DefaultLocalCollectorDir = "collector"
)

// Env Variables
const (
	AWSECSPricingURL = "AWS_ECS_PRICING_URL"

	// Currently being used for OCI and DigitalOcean
	ProviderPricingURL = "PROVIDER_PRICING_URL"

	CollectorDataSourceEnabledEnvVar = "COLLECTOR_DATA_SOURCE_ENABLED"
	LocalCollectorDirectoryEnvVar    = "LOCAL_COLLECTOR_DIRECTORY"

	KubernetesResourceAccessEnvVar = "KUBERNETES_RESOURCE_ACCESS"
	UseCacheV1                     = "USE_CACHE_V1"

	// Cloud provider override
	CloudProviderVar = "CLOUD_PROVIDER"
)

func GetGCPAuthSecretFilePath() string {
	return GetPathFromConfig(GCPAuthSecretFile)
}

func GetClusterInfoFilePath() string {
	return GetPathFromConfig(ClusterInfoFile)
}

func GetClusterCacheFilePath() string {
	return GetPathFromConfig(ClusterCacheFile)
}

// GetAWSECSPricingURL returns an optional alternative URL to fetch AmazonECS pricing data from; for use in airgapped environments
func GetAWSECSPricingURL() string {
	return env.Get(AWSECSPricingURL, "")
}

func IsCollectorDataSourceEnabled() bool {
	return env.GetBool(CollectorDataSourceEnabledEnvVar, false)
}

func IsAllocationNodeLabelsEnabled() bool {
	return env.GetBool(AllocationNodeLabelsEnabled, true)
}

func IsAssetIncludeLocalDiskCost() bool {
	return env.GetBool(AssetIncludeLocalDiskCostEnvVar, true)
}

// HasKubernetesResourceAccess can be set to false if Opencost is run without access to the kubernetes resources
func HasKubernetesResourceAccess() bool { return env.GetBool(KubernetesResourceAccessEnvVar, true) }

// GetUseCacheV1 is a temporary flag to allow users to opt-in to using the old cache
// Mainly for comparison purposes
func GetUseCacheV1() bool {
	return env.GetBool(UseCacheV1, false)
}

// GetCloudProvider returns the explicitly set cloud provider from environment variable
func GetCloudProvider() string {
	return env.Get(CloudProviderVar, "")
}

func GetMetricConfigFile() string {
	return GetPathFromConfig(MetricConfigFile)
}

func GetLocalCollectorDirectory() string {
	dir := env.Get(LocalCollectorDirectoryEnvVar, DefaultLocalCollectorDir)
	return GetPathFromConfig(dir)

}

func GetDOKSPricingURL() string {
	return env.Get(ProviderPricingURL, "https://api.digitalocean.com/v2/billing/pricing")
}
