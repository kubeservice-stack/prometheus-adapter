/*
Copyright 2023 The Kubernetes Authors.
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

package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

var _ cache.TransformFunc = partialMetadataRemoveAll

// partialMetadataRemoveAll implements a cache.TransformFunc that removes
// annotations and managed
// fields from PartialObjectMetadata.
func partialMetadataRemoveAll(obj interface{}) (interface{}, error) {
	partialMeta, ok := obj.(*metav1.PartialObjectMetadata)
	if !ok {
		return nil, fmt.Errorf("internal error: cannot cast object %#+v to PartialObjectMetadata", obj)
	}
	partialMeta.Annotations = nil
	partialMeta.ManagedFields = nil
	return partialMeta, nil
}
