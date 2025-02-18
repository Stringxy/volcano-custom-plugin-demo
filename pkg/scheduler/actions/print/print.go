/*
 Copyright 2021 The Volcano Authors.

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

package print

import (
	"k8s.io/klog/v2"

	"volcano.sh/volcano/pkg/scheduler/api"
	"volcano.sh/volcano/pkg/scheduler/framework"
)

type Action struct {
	session *framework.Session
}

func New() *Action {
	return &Action{}
}

func (print *Action) Name() string {
	return "print"
}

func (print *Action) Initialize() {}

func (print *Action) Execute(ssn *framework.Session) {
	klog.V(5).Infof("Enter Print Action ...")
	defer klog.V(5).Infof("Leaving Action ...")
	jobs := make([]*api.JobInfo, 0)
	for _, job := range ssn.Jobs {
		jobs = append(jobs, job)
	}
	ssn.PrintJobInfo(jobs)
}

func (alloc *Action) UnInitialize() {}
