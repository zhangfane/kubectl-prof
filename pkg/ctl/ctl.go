package ctl

import (
	"time"

	"github.com/josepdcs/kubectl-prof/internal/cli/config"
	"github.com/josepdcs/kubectl-prof/internal/cli/kubernetes"
	"github.com/josepdcs/kubectl-prof/internal/cli/profiler"
	apiprof "github.com/josepdcs/kubectl-prof/internal/cli/profiler/api"
	clientgo "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewProfiler(connectionInfo kubernetes.ConnectionInfo) *profiler.Profiler {
	jobProfiler := profiler.NewJobProfiler(
		apiprof.NewPodApi(connectionInfo),
		apiprof.NewProfilingJobApi(connectionInfo),
		apiprof.NewProfilingContainerApi(connectionInfo),
	)
	return jobProfiler
}
func NewConnectionInfo(ns string, restCfg *rest.Config) kubernetes.ConnectionInfo {

	return kubernetes.ConnectionInfo{
		ClientSet:  clientgo.NewForConfigOrDie(restCfg),
		RestConfig: restCfg,
		Namespace:  ns,
	}
}

func NewConfig(podName string, target map[any]any) (*config.ProfilerConfig, error) {
	return config.NewProfilerConfig(&config.TargetConfig{
		Namespace:            target["namespace"].(string),
		PodName:              podName,
		ContainerName:        "",
		LabelSelector:        "",
		ContainerID:          "",
		Event:                "itimer",
		Duration:             time.Minute,
		Interval:             0,
		Id:                   "",
		LocalPath:            "./",
		Alpine:               false,
		DryRun:               false,
		Image:                "registry.cn-hangzhou.aliyuncs.com/zfane/kubectl-prof:v1.7.0-dev-python",
		ContainerRuntime:     "containerd",
		ContainerRuntimePath: "/run/containerd",
		Language:             "python",
		Compressor:           "gzip",
		ImagePullSecret:      "",
		ServiceAccountName:   "",
		ProfilingTool:        "pyspy",
		OutputType:           "flamegrpah",
		ImagePullPolicy:      "IfNotPresent",
		ExtraTargetOptions: config.ExtraTargetOptions{
			PoolSizeLaunchProfilingJobs: 0,
			PrintLogs:                   true,
			GracePeriodEnding:           time.Minute * 5,
			HeapDumpSplitInChunkSize:    "50M",
			PoolSizeRetrieveChunks:      5,
			RetrieveFileRetries:         3,
			PID:                         "",
			Pgrep:                       "",
			NodeHeapSnapshotSignal:      12,
		},
	}, config.WithJob(&config.JobConfig{
		ContainerConfig: config.ContainerConfig{
			RequestConfig: config.ResourceConfig{},
			LimitConfig:   config.ResourceConfig{},
			Privileged:    false,
			Capabilities:  nil,
		},
		Namespace:      "",
		Tolerations:    nil,
		TolerationsRaw: nil,
	}), config.WithLogLevel("info"))
}
