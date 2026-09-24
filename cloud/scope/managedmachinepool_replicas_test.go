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

package scope

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/ptr"
	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"sigs.k8s.io/cluster-api-provider-gcp/exp/api/v1beta1"
)

func TestSetReplicasSyncsMachinePool(t *testing.T) {
	tests := []struct {
		name           string
		specReplicas   *int32
		statusReplicas *int32
		observed       int32
		wantSpec       int32
		wantStatus     *int32
	}{
		{name: "GKE scaled up: spec follows", specReplicas: ptr.To(int32(1)), statusReplicas: ptr.To(int32(1)), observed: 3, wantSpec: 3, wantStatus: ptr.To(int32(1))},
		{name: "already aligned: untouched", specReplicas: ptr.To(int32(2)), statusReplicas: ptr.To(int32(2)), observed: 2, wantSpec: 2, wantStatus: ptr.To(int32(2))},
		{name: "scaled to zero with unset status: status set to zero", specReplicas: ptr.To(int32(1)), observed: 0, wantSpec: 0, wantStatus: ptr.To(int32(0))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheme := runtime.NewScheme()
			if err := clusterv1.AddToScheme(scheme); err != nil {
				t.Fatal(err)
			}
			key := client.ObjectKey{Namespace: "ns", Name: "mp"}
			seed := &clusterv1.MachinePool{
				ObjectMeta: metav1.ObjectMeta{Namespace: key.Namespace, Name: key.Name},
				Spec:       clusterv1.MachinePoolSpec{Replicas: tt.specReplicas},
				Status:     clusterv1.MachinePoolStatus{Replicas: tt.statusReplicas},
			}
			c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(seed).WithStatusSubresource(&clusterv1.MachinePool{}).Build()

			machinePool := &clusterv1.MachinePool{}
			if err := c.Get(t.Context(), key, machinePool); err != nil {
				t.Fatal(err)
			}
			s := &ManagedMachinePoolScope{
				client:                c,
				MachinePool:           machinePool,
				GCPManagedMachinePool: &v1beta1.GCPManagedMachinePool{ObjectMeta: metav1.ObjectMeta{Name: key.Name}},
			}

			s.SetReplicas(tt.observed)

			if s.GCPManagedMachinePool.Status.Replicas != tt.observed {
				t.Errorf("GCPManagedMachinePool.Status.Replicas = %d, want %d", s.GCPManagedMachinePool.Status.Replicas, tt.observed)
			}
			got := &clusterv1.MachinePool{}
			if err := c.Get(t.Context(), key, got); err != nil {
				t.Fatal(err)
			}
			if ptr.Deref(got.Spec.Replicas, -1) != tt.wantSpec {
				t.Errorf("spec.replicas = %v, want %d", got.Spec.Replicas, tt.wantSpec)
			}
			if ptr.Deref(got.Status.Replicas, -1) != ptr.Deref(tt.wantStatus, -1) {
				t.Errorf("status.replicas = %v, want %v", got.Status.Replicas, tt.wantStatus)
			}
		})
	}
}
