/*
Copyright 2026 The Kubernetes Authors.

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

package nodepools

import (
	"testing"

	"cloud.google.com/go/container/apiv1/containerpb"
	"k8s.io/utils/ptr"
	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"sigs.k8s.io/cluster-api-provider-gcp/cloud/scope"
	infrav1exp "sigs.k8s.io/cluster-api-provider-gcp/exp/api/v1beta1"
)

func TestCheckDiffAndPrepareUpdateSizeWithAutoscaling(t *testing.T) {
	s := &Service{scope: &scope.ManagedMachinePoolScope{
		GCPManagedControlPlane: &infrav1exp.GCPManagedControlPlane{
			Spec: infrav1exp.GCPManagedControlPlaneSpec{
				GCPManagedControlPlaneClassSpec: infrav1exp.GCPManagedControlPlaneClassSpec{
					Project:  "test-project",
					Location: "us-central1-a",
				},
				ClusterName: "test-cluster",
			},
		},
		GCPManagedMachinePool: &infrav1exp.GCPManagedMachinePool{
			Spec: infrav1exp.GCPManagedMachinePoolSpec{
				GCPManagedMachinePoolClassSpec: infrav1exp.GCPManagedMachinePoolClassSpec{
					NodePoolName: "pool",
					Scaling: &infrav1exp.NodePoolAutoScaling{
						EnableAutoscaling: ptr.To(true),
						MinCount:          ptr.To(int32(1)),
						MaxCount:          ptr.To(int32(5)),
					},
				},
			},
		},
		MachinePool: &clusterv1.MachinePool{Spec: clusterv1.MachinePoolSpec{Replicas: ptr.To(int32(3))}},
	}}

	needUpdate, req := s.checkDiffAndPrepareUpdateSize(&containerpb.NodePool{InitialNodeCount: 1, Locations: []string{"us-central1-a"}})
	if !needUpdate {
		t.Fatalf("expected a size update with autoscaling enabled")
	}
	if req.GetNodeCount() != 3 {
		t.Errorf("NodeCount = %d, want 3", req.GetNodeCount())
	}

	needUpdate, _ = s.checkDiffAndPrepareUpdateSize(&containerpb.NodePool{InitialNodeCount: 3, Locations: []string{"us-central1-a"}})
	if needUpdate {
		t.Errorf("expected no size update when the node pool already has the desired size")
	}
}
