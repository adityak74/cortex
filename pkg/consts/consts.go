/*
Copyright 2021 Cortex Labs, Inc.

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

package consts

import (
	"os"
)

var (
	CortexVersion      = "master" // CORTEX_VERSION
	CortexVersionMinor = "master" // CORTEX_VERSION_MINOR

	DefaultMaxQueueLength = int64(100)
	DefaultMaxConcurrency = int64(1)

	DefaultUserPodPortStr   = "8080"
	DefaultUserPodPortInt32 = int32(8080)

	ProxyListeningPortStr   = "8888"
	ProxyListeningPortInt32 = int32(8888)

	AdminPortStr   = "15000"
	AdminPortInt32 = int32(15000)

	StatsDPortStr   = "9125"
	StatsDPortInt32 = int32(9125)

	AuthHeader = "X-Cortex-Authorization"

	DefaultInClusterConfigPath   = "/configs/cluster/cluster.yaml"
	MaxBucketLifecycleRules      = 100
	AsyncWorkloadsExpirationDays = int64(7)

	ReservedContainerPorts = []int32{
		ProxyListeningPortInt32,
		AdminPortInt32,
	}
	ReservedContainerNames = []string{
		"dequeuer",
		"proxy",
	}
)

func DefaultRegistry() string {
	if registryOverride := os.Getenv("CORTEX_DEV_DEFAULT_IMAGE_REGISTRY"); registryOverride != "" {
		return registryOverride
	}
	return "quay.io/cortexlabs"
}
