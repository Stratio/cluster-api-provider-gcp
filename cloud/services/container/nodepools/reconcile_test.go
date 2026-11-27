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
	clusterv1exp "sigs.k8s.io/cluster-api/exp/api/v1beta1"

	"sigs.k8s.io/cluster-api-provider-gcp/cloud/scope"
	infrav1exp "sigs.k8s.io/cluster-api-provider-gcp/exp/api/v1beta1"
)

func newSizeTestService(scaling *infrav1exp.NodePoolAutoScaling) *Service {
	return &Service{scope: &scope.ManagedMachinePoolScope{
		GCPManagedControlPlane: &infrav1exp.GCPManagedControlPlane{
			Spec: infrav1exp.GCPManagedControlPlaneSpec{
				Project:     "test-project",
				Location:    "us-central1-a",
				ClusterName: "test-cluster",
			},
		},
		GCPManagedMachinePool: &infrav1exp.GCPManagedMachinePool{
			Spec: infrav1exp.GCPManagedMachinePoolSpec{
				NodePoolName: "pool",
				Scaling:      scaling,
			},
		},
		MachinePool: &clusterv1exp.MachinePool{Spec: clusterv1exp.MachinePoolSpec{Replicas: ptr.To(int32(3))}},
	}}
}

func TestCheckDiffAndPrepareUpdateSizeWithAutoscaling(t *testing.T) {
	s := newSizeTestService(&infrav1exp.NodePoolAutoScaling{
		EnableAutoscaling: ptr.To(true),
		MinCount:          ptr.To(int32(1)),
		MaxCount:          ptr.To(int32(5)),
	})

	needUpdate, _ := s.checkDiffAndPrepareUpdateSize(&containerpb.NodePool{InitialNodeCount: 1, Locations: []string{"us-central1-a"}})
	if needUpdate {
		t.Errorf("expected no size update with autoscaling enabled: the GKE autoscaler owns the size")
	}
}

func TestCheckDiffAndPrepareUpdateSizeWithoutAutoscaling(t *testing.T) {
	for name, scaling := range map[string]*infrav1exp.NodePoolAutoScaling{
		"autoscaling disabled": {EnableAutoscaling: ptr.To(false)},
		"no scaling block":     nil,
	} {
		t.Run(name, func(t *testing.T) {
			s := newSizeTestService(scaling)

			needUpdate, req := s.checkDiffAndPrepareUpdateSize(&containerpb.NodePool{InitialNodeCount: 1, Locations: []string{"us-central1-a"}})
			if !needUpdate {
				t.Fatalf("expected a size update")
			}
			if req.GetNodeCount() != 3 {
				t.Errorf("NodeCount = %d, want 3", req.GetNodeCount())
			}

			needUpdate, _ = s.checkDiffAndPrepareUpdateSize(&containerpb.NodePool{InitialNodeCount: 3, Locations: []string{"us-central1-a"}})
			if needUpdate {
				t.Errorf("expected no size update when the node pool already has the desired size")
			}
		})
	}
}
