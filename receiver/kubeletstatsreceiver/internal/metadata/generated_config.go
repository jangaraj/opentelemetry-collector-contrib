// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for kubeletstats metrics.
type MetricsConfig struct {
	ContainerCPUTime                     MetricConfig `mapstructure:"container.cpu.time"`
	ContainerCPUUsage                    MetricConfig `mapstructure:"container.cpu.usage"`
	ContainerCPUUtilization              MetricConfig `mapstructure:"container.cpu.utilization"`
	ContainerFilesystemAvailable         MetricConfig `mapstructure:"container.filesystem.available"`
	ContainerFilesystemCapacity          MetricConfig `mapstructure:"container.filesystem.capacity"`
	ContainerFilesystemUsage             MetricConfig `mapstructure:"container.filesystem.usage"`
	ContainerMemoryAvailable             MetricConfig `mapstructure:"container.memory.available"`
	ContainerMemoryMajorPageFaults       MetricConfig `mapstructure:"container.memory.major_page_faults"`
	ContainerMemoryPageFaults            MetricConfig `mapstructure:"container.memory.page_faults"`
	ContainerMemoryRss                   MetricConfig `mapstructure:"container.memory.rss"`
	ContainerMemoryUsage                 MetricConfig `mapstructure:"container.memory.usage"`
	ContainerMemoryWorkingSet            MetricConfig `mapstructure:"container.memory.working_set"`
	ContainerUptime                      MetricConfig `mapstructure:"container.uptime"`
	K8sContainerCPUNodeUtilization       MetricConfig `mapstructure:"k8s.container.cpu.node.utilization"`
	K8sContainerCPULimitUtilization      MetricConfig `mapstructure:"k8s.container.cpu_limit_utilization"`
	K8sContainerCPURequestUtilization    MetricConfig `mapstructure:"k8s.container.cpu_request_utilization"`
	K8sContainerMemoryNodeUtilization    MetricConfig `mapstructure:"k8s.container.memory.node.utilization"`
	K8sContainerMemoryLimitUtilization   MetricConfig `mapstructure:"k8s.container.memory_limit_utilization"`
	K8sContainerMemoryRequestUtilization MetricConfig `mapstructure:"k8s.container.memory_request_utilization"`
	K8sNodeCPUTime                       MetricConfig `mapstructure:"k8s.node.cpu.time"`
	K8sNodeCPUUsage                      MetricConfig `mapstructure:"k8s.node.cpu.usage"`
	K8sNodeCPUUtilization                MetricConfig `mapstructure:"k8s.node.cpu.utilization"`
	K8sNodeFilesystemAvailable           MetricConfig `mapstructure:"k8s.node.filesystem.available"`
	K8sNodeFilesystemCapacity            MetricConfig `mapstructure:"k8s.node.filesystem.capacity"`
	K8sNodeFilesystemUsage               MetricConfig `mapstructure:"k8s.node.filesystem.usage"`
	K8sNodeMemoryAvailable               MetricConfig `mapstructure:"k8s.node.memory.available"`
	K8sNodeMemoryMajorPageFaults         MetricConfig `mapstructure:"k8s.node.memory.major_page_faults"`
	K8sNodeMemoryPageFaults              MetricConfig `mapstructure:"k8s.node.memory.page_faults"`
	K8sNodeMemoryRss                     MetricConfig `mapstructure:"k8s.node.memory.rss"`
	K8sNodeMemoryUsage                   MetricConfig `mapstructure:"k8s.node.memory.usage"`
	K8sNodeMemoryWorkingSet              MetricConfig `mapstructure:"k8s.node.memory.working_set"`
	K8sNodeNetworkErrors                 MetricConfig `mapstructure:"k8s.node.network.errors"`
	K8sNodeNetworkIo                     MetricConfig `mapstructure:"k8s.node.network.io"`
	K8sNodeUptime                        MetricConfig `mapstructure:"k8s.node.uptime"`
	K8sPodCPUNodeUtilization             MetricConfig `mapstructure:"k8s.pod.cpu.node.utilization"`
	K8sPodCPUTime                        MetricConfig `mapstructure:"k8s.pod.cpu.time"`
	K8sPodCPUUsage                       MetricConfig `mapstructure:"k8s.pod.cpu.usage"`
	K8sPodCPUUtilization                 MetricConfig `mapstructure:"k8s.pod.cpu.utilization"`
	K8sPodCPULimitUtilization            MetricConfig `mapstructure:"k8s.pod.cpu_limit_utilization"`
	K8sPodCPURequestUtilization          MetricConfig `mapstructure:"k8s.pod.cpu_request_utilization"`
	K8sPodFilesystemAvailable            MetricConfig `mapstructure:"k8s.pod.filesystem.available"`
	K8sPodFilesystemCapacity             MetricConfig `mapstructure:"k8s.pod.filesystem.capacity"`
	K8sPodFilesystemUsage                MetricConfig `mapstructure:"k8s.pod.filesystem.usage"`
	K8sPodMemoryAvailable                MetricConfig `mapstructure:"k8s.pod.memory.available"`
	K8sPodMemoryMajorPageFaults          MetricConfig `mapstructure:"k8s.pod.memory.major_page_faults"`
	K8sPodMemoryNodeUtilization          MetricConfig `mapstructure:"k8s.pod.memory.node.utilization"`
	K8sPodMemoryPageFaults               MetricConfig `mapstructure:"k8s.pod.memory.page_faults"`
	K8sPodMemoryRss                      MetricConfig `mapstructure:"k8s.pod.memory.rss"`
	K8sPodMemoryUsage                    MetricConfig `mapstructure:"k8s.pod.memory.usage"`
	K8sPodMemoryWorkingSet               MetricConfig `mapstructure:"k8s.pod.memory.working_set"`
	K8sPodMemoryLimitUtilization         MetricConfig `mapstructure:"k8s.pod.memory_limit_utilization"`
	K8sPodMemoryRequestUtilization       MetricConfig `mapstructure:"k8s.pod.memory_request_utilization"`
	K8sPodNetworkErrors                  MetricConfig `mapstructure:"k8s.pod.network.errors"`
	K8sPodNetworkIo                      MetricConfig `mapstructure:"k8s.pod.network.io"`
	K8sPodUptime                         MetricConfig `mapstructure:"k8s.pod.uptime"`
	K8sVolumeAvailable                   MetricConfig `mapstructure:"k8s.volume.available"`
	K8sVolumeCapacity                    MetricConfig `mapstructure:"k8s.volume.capacity"`
	K8sVolumeInodes                      MetricConfig `mapstructure:"k8s.volume.inodes"`
	K8sVolumeInodesFree                  MetricConfig `mapstructure:"k8s.volume.inodes.free"`
	K8sVolumeInodesUsed                  MetricConfig `mapstructure:"k8s.volume.inodes.used"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		ContainerCPUTime: MetricConfig{
			Enabled: true,
		},
		ContainerCPUUsage: MetricConfig{
			Enabled: true,
		},
		ContainerCPUUtilization: MetricConfig{
			Enabled: false,
		},
		ContainerFilesystemAvailable: MetricConfig{
			Enabled: true,
		},
		ContainerFilesystemCapacity: MetricConfig{
			Enabled: true,
		},
		ContainerFilesystemUsage: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryAvailable: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryMajorPageFaults: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryPageFaults: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryRss: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryUsage: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryWorkingSet: MetricConfig{
			Enabled: true,
		},
		ContainerUptime: MetricConfig{
			Enabled: false,
		},
		K8sContainerCPUNodeUtilization: MetricConfig{
			Enabled: false,
		},
		K8sContainerCPULimitUtilization: MetricConfig{
			Enabled: false,
		},
		K8sContainerCPURequestUtilization: MetricConfig{
			Enabled: false,
		},
		K8sContainerMemoryNodeUtilization: MetricConfig{
			Enabled: false,
		},
		K8sContainerMemoryLimitUtilization: MetricConfig{
			Enabled: false,
		},
		K8sContainerMemoryRequestUtilization: MetricConfig{
			Enabled: false,
		},
		K8sNodeCPUTime: MetricConfig{
			Enabled: true,
		},
		K8sNodeCPUUsage: MetricConfig{
			Enabled: true,
		},
		K8sNodeCPUUtilization: MetricConfig{
			Enabled: false,
		},
		K8sNodeFilesystemAvailable: MetricConfig{
			Enabled: true,
		},
		K8sNodeFilesystemCapacity: MetricConfig{
			Enabled: true,
		},
		K8sNodeFilesystemUsage: MetricConfig{
			Enabled: true,
		},
		K8sNodeMemoryAvailable: MetricConfig{
			Enabled: true,
		},
		K8sNodeMemoryMajorPageFaults: MetricConfig{
			Enabled: true,
		},
		K8sNodeMemoryPageFaults: MetricConfig{
			Enabled: true,
		},
		K8sNodeMemoryRss: MetricConfig{
			Enabled: true,
		},
		K8sNodeMemoryUsage: MetricConfig{
			Enabled: true,
		},
		K8sNodeMemoryWorkingSet: MetricConfig{
			Enabled: true,
		},
		K8sNodeNetworkErrors: MetricConfig{
			Enabled: true,
		},
		K8sNodeNetworkIo: MetricConfig{
			Enabled: true,
		},
		K8sNodeUptime: MetricConfig{
			Enabled: false,
		},
		K8sPodCPUNodeUtilization: MetricConfig{
			Enabled: false,
		},
		K8sPodCPUTime: MetricConfig{
			Enabled: true,
		},
		K8sPodCPUUsage: MetricConfig{
			Enabled: true,
		},
		K8sPodCPUUtilization: MetricConfig{
			Enabled: false,
		},
		K8sPodCPULimitUtilization: MetricConfig{
			Enabled: false,
		},
		K8sPodCPURequestUtilization: MetricConfig{
			Enabled: false,
		},
		K8sPodFilesystemAvailable: MetricConfig{
			Enabled: true,
		},
		K8sPodFilesystemCapacity: MetricConfig{
			Enabled: true,
		},
		K8sPodFilesystemUsage: MetricConfig{
			Enabled: true,
		},
		K8sPodMemoryAvailable: MetricConfig{
			Enabled: true,
		},
		K8sPodMemoryMajorPageFaults: MetricConfig{
			Enabled: true,
		},
		K8sPodMemoryNodeUtilization: MetricConfig{
			Enabled: false,
		},
		K8sPodMemoryPageFaults: MetricConfig{
			Enabled: true,
		},
		K8sPodMemoryRss: MetricConfig{
			Enabled: true,
		},
		K8sPodMemoryUsage: MetricConfig{
			Enabled: true,
		},
		K8sPodMemoryWorkingSet: MetricConfig{
			Enabled: true,
		},
		K8sPodMemoryLimitUtilization: MetricConfig{
			Enabled: false,
		},
		K8sPodMemoryRequestUtilization: MetricConfig{
			Enabled: false,
		},
		K8sPodNetworkErrors: MetricConfig{
			Enabled: true,
		},
		K8sPodNetworkIo: MetricConfig{
			Enabled: true,
		},
		K8sPodUptime: MetricConfig{
			Enabled: false,
		},
		K8sVolumeAvailable: MetricConfig{
			Enabled: true,
		},
		K8sVolumeCapacity: MetricConfig{
			Enabled: true,
		},
		K8sVolumeInodes: MetricConfig{
			Enabled: true,
		},
		K8sVolumeInodesFree: MetricConfig{
			Enabled: true,
		},
		K8sVolumeInodesUsed: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for kubeletstats resource attributes.
type ResourceAttributesConfig struct {
	AwsVolumeID                  ResourceAttributeConfig `mapstructure:"aws.volume.id"`
	ContainerID                  ResourceAttributeConfig `mapstructure:"container.id"`
	FsType                       ResourceAttributeConfig `mapstructure:"fs.type"`
	GcePdName                    ResourceAttributeConfig `mapstructure:"gce.pd.name"`
	GlusterfsEndpointsName       ResourceAttributeConfig `mapstructure:"glusterfs.endpoints.name"`
	GlusterfsPath                ResourceAttributeConfig `mapstructure:"glusterfs.path"`
	K8sContainerName             ResourceAttributeConfig `mapstructure:"k8s.container.name"`
	K8sNamespaceName             ResourceAttributeConfig `mapstructure:"k8s.namespace.name"`
	K8sNodeName                  ResourceAttributeConfig `mapstructure:"k8s.node.name"`
	K8sPersistentvolumeclaimName ResourceAttributeConfig `mapstructure:"k8s.persistentvolumeclaim.name"`
	K8sPodName                   ResourceAttributeConfig `mapstructure:"k8s.pod.name"`
	K8sPodUID                    ResourceAttributeConfig `mapstructure:"k8s.pod.uid"`
	K8sVolumeName                ResourceAttributeConfig `mapstructure:"k8s.volume.name"`
	K8sVolumeType                ResourceAttributeConfig `mapstructure:"k8s.volume.type"`
	Partition                    ResourceAttributeConfig `mapstructure:"partition"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		AwsVolumeID: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerID: ResourceAttributeConfig{
			Enabled: true,
		},
		FsType: ResourceAttributeConfig{
			Enabled: true,
		},
		GcePdName: ResourceAttributeConfig{
			Enabled: true,
		},
		GlusterfsEndpointsName: ResourceAttributeConfig{
			Enabled: true,
		},
		GlusterfsPath: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sContainerName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNamespaceName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNodeName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sPersistentvolumeclaimName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sPodName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sPodUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sVolumeName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sVolumeType: ResourceAttributeConfig{
			Enabled: true,
		},
		Partition: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for kubeletstats metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
