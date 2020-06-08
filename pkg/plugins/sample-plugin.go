package plugins

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

// Name 插件名称
const Name = "sample-plugin"

//Args ...
type Args struct {
	FavoriteColor  string `json:"favorite_color,omitempty"`
	FavoriteNumber int    `json:"favorite_number,omitempty"`
	ThanksTo       string `json:"thanks_to,omitempty"`
}

//Sample ...
type Sample struct {
	args   *Args
	handle framework.FrameworkHandle
}

//Name ...
func (s *Sample) Name() string {
	return Name
}

//PreFilter ...
func (s *Sample) PreFilter(pc *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}

//Filter ...
func (s *Sample) Filter(pc *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

//Score ...
func (s *Sample) Score(pc *framework.CycleState, pod *v1.Pod, nodeName string) (int, *framework.Status) {
	klog.V(3).Infof("Score pod: %v, node: %v", pod.Name, nodeName)
	score := rand.Intn(100)
	return score, framework.NewStatus(framework.Success, "")
}

//NormalizeScore ...
func (s *Sample) NormalizeScore(pc *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	klog.V(3).Infof("NormalizeScore pod: %v, score: %s", pod.Name, scores)
	return framework.NewStatus(framework.Success, "")
}

//PreBind ...
func (s *Sample) PreBind(pc *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	nodeInfo, err := s.handle.SnapshotSharedLister().NodeInfos().Get(nodeName)
	if err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("prebind get node info error: %+v", nodeName))
	}

	klog.V(3).Infof("prebind node info: %+v", nodeInfo.Node())
	return framework.NewStatus(framework.Success, "")

}

//New ...
//type PluginFactory = func(configuration *runtime.Unknown, f FrameworkHandle) (Plugin, error)
func New(configuration *runtime.Unknown, f framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(configuration, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("get plugin config args: %+v", args)
	return &Sample{
		args:   args,
		handle: f,
	}, nil
}
