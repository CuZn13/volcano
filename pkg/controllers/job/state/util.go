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

package state

import (
	vkv1 "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
)

//DefaultMaxRetry is the default number of retries.
const DefaultMaxRetry int32 = 3

//TotalTasks returns number of tasks in a given volcano job
func TotalTasks(job *vkv1.Job) int32 {
	var rep int32

	for _, task := range job.Spec.Tasks {
		rep += task.Replicas
	}

	return rep
}
