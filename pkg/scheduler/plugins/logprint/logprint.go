/*
Copyright 2019 The Volcano Authors.

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

package logprint

import (
	"k8s.io/klog/v2"

	"volcano.sh/volcano/pkg/scheduler/api"
	"volcano.sh/volcano/pkg/scheduler/framework"
)

const (
	// PluginName indicates name of volcano scheduler plugin.
	PluginName = "logprint"
)

type logPrint struct {
	// Arguments given for the plugin
	pluginArguments framework.Arguments
}

// New function returns prioritizePlugin object
func New(aruguments framework.Arguments) framework.Plugin {
	return &logPrint{aruguments}
}

func (logPrint *logPrint) Name() string {
	return PluginName
}

func (logPrint *logPrint) OnSessionOpen(ssn *framework.Session) {
	klog.V(5).Infof("Enter logPrint plugin ...")
	defer func() {
		klog.V(5).Infof("Leaving logPrint plugin. %s ...", logPrint.pluginArguments)
	}()
	loggingFn := func(jobs []*api.JobInfo) {
		for _, job := range jobs {
			if job.IsPending() {
				klog.V(1).Infof("Job %s/%s is pending, try to schedule it on node %s", job.Namespace, job.Name)
			}
			if job.IsReady() {
				klog.V(1).Infof("Job %s/%s is ready, detail info : %s", job.Namespace, job.Name, job.String())
			}
		}
	}

	ssn.AddPrintFn(PluginName, loggingFn)
}

func (logPrint *logPrint) OnSessionClose(ssn *framework.Session) {
}
