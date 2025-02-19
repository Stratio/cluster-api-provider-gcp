/*
Copyright 2022 The Kubernetes Authors.

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

package v1beta1

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/pointer"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api-provider-gcp/util/hash"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const (
	maxClusterNameLength = 40
	resourcePrefix       = "capg-"
)

// log is for logging in this package.
var gcpmanagedcontrolplanelog = logf.Log.WithName("gcpmanagedcontrolplane-resource")

func (r *GCPManagedControlPlane) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-infrastructure-cluster-x-k8s-io-v1beta1-gcpmanagedcontrolplane,mutating=true,failurePolicy=fail,sideEffects=None,groups=infrastructure.cluster.x-k8s.io,resources=gcpmanagedcontrolplanes,verbs=create;update,versions=v1beta1,name=mgcpmanagedcontrolplane.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &GCPManagedControlPlane{}

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (r *GCPManagedControlPlane) Default() {
	gcpmanagedcontrolplanelog.Info("default", "name", r.Name)

	if r.Spec.ClusterName == "" {
		gcpmanagedcontrolplanelog.Info("ClusterName is empty, generating name")
		name, err := generateGKEName(r.Name, r.Namespace, maxClusterNameLength)
		if err != nil {
			gcpmanagedcontrolplanelog.Error(err, "failed to create GKE cluster name")
			return
		}

		gcpmanagedcontrolplanelog.Info("defaulting GKE cluster name", "cluster-name", name)
		r.Spec.ClusterName = name
	}
}

//+kubebuilder:webhook:path=/validate-infrastructure-cluster-x-k8s-io-v1beta1-gcpmanagedcontrolplane,mutating=false,failurePolicy=fail,sideEffects=None,groups=infrastructure.cluster.x-k8s.io,resources=gcpmanagedcontrolplanes,verbs=create;update,versions=v1beta1,name=vgcpmanagedcontrolplane.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &GCPManagedControlPlane{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPManagedControlPlane) ValidateCreate() (admission.Warnings, error) {
	gcpmanagedcontrolplanelog.Info("validate create", "name", r.Name)
	var allErrs field.ErrorList

	if len(r.Spec.ClusterName) > maxClusterNameLength {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "ClusterName"),
				r.Spec.ClusterName, fmt.Sprintf("cluster name cannot have more than %d characters", maxClusterNameLength)),
		)
	}

	if r.Spec.EnableAutopilot && r.Spec.ReleaseChannel == nil {
		allErrs = append(allErrs, field.Required(field.NewPath("spec", "ReleaseChannel"), "Release channel is required for an autopilot enabled cluster"))
	}

	if len(allErrs) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(GroupVersion.WithKind("GCPManagedControlPlane").GroupKind(), r.Name, allErrs)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPManagedControlPlane) ValidateUpdate(oldRaw runtime.Object) (admission.Warnings, error) {
	gcpmanagedcontrolplanelog.Info("validate update", "name", r.Name)
	var allErrs field.ErrorList
	old := oldRaw.(*GCPManagedControlPlane)

	if !cmp.Equal(r.Spec.ClusterName, old.Spec.ClusterName) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "ClusterName"),
				r.Spec.ClusterName, "field is immutable"),
		)
	}

	if !cmp.Equal(r.Spec.Project, old.Spec.Project) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Project"),
				r.Spec.Project, "field is immutable"),
		)
	}

	if !cmp.Equal(r.Spec.Location, old.Spec.Location) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Location"),
				r.Spec.Location, "field is immutable"),
		)
	}

	// Add IPAllocationPolicy for CIDR support (PLT-1246)

	if !cmp.Equal(r.Spec.ClusterIpv4Cidr, old.Spec.ClusterIpv4Cidr) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "ClusterIpv4Cidr"),
				pointer.StringDeref(r.Spec.ClusterIpv4Cidr, ""), "field is immutable"),
		)
	}

	if !cmp.Equal(r.Spec.EnableAutopilot, old.Spec.EnableAutopilot) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "EnableAutopilot"),
				r.Spec.EnableAutopilot, "field is immutable"),
		)
	}

	if !cmp.Equal(r.Spec.NetworkPolicy, old.Spec.NetworkPolicy) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "NetworkPolicy"),
				r.Spec.NetworkPolicy, "field is immutable"),
		)
	}

	if len(allErrs) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(GroupVersion.WithKind("GCPManagedControlPlane").GroupKind(), r.Name, allErrs)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPManagedControlPlane) ValidateDelete() (admission.Warnings, error) {
	gcpmanagedcontrolplanelog.Info("validate delete", "name", r.Name)

	return nil, nil
}

func generateGKEName(resourceName, namespace string, maxLength int) (string, error) {
	escapedName := strings.ReplaceAll(resourceName, ".", "-")
	gkeName := fmt.Sprintf("%s-%s", namespace, escapedName)

	if len(gkeName) < maxLength {
		return gkeName, nil
	}

	hashLength := 32 - len(resourcePrefix)
	hashedName, err := hash.Base36TruncatedHash(gkeName, hashLength)
	if err != nil {
		return "", errors.Wrap(err, "creating hash from name")
	}

	return fmt.Sprintf("%s%s", resourcePrefix, hashedName), nil
}

// Add IPAllocationPolicy for CIDR support (PLT-1246)

func validateIPAllocationPolicy(spec GCPManagedControlPlaneSpec) field.ErrorList {
	var allErrs field.ErrorList

	if spec.IPAllocationPolicy == nil {
		return allErrs
	}

	path := field.NewPath("spec", "IPAllocationPolicy")

	isUseIPAliases := pointer.BoolDeref(spec.IPAllocationPolicy.UseIPAliases, false)
	if spec.IPAllocationPolicy.ClusterSecondaryRangeName != nil && !isUseIPAliases {
		allErrs = append(allErrs,
			field.Invalid(path.Child("ClusterSecondaryRangeName"),
				spec.IPAllocationPolicy.ClusterSecondaryRangeName,
				"field cannot be set unless UseIPAliases is set to true"),
		)
	}
	if spec.IPAllocationPolicy.ServicesSecondaryRangeName != nil && !isUseIPAliases {
		allErrs = append(allErrs,
			field.Invalid(path.Child("ServicesSecondaryRangeName"),
				spec.IPAllocationPolicy.ServicesSecondaryRangeName,
				"field cannot be set unless UseIPAliases is set to true"),
		)
	}
	if spec.IPAllocationPolicy.ServicesIpv4CidrBlock != nil && !isUseIPAliases {
		allErrs = append(allErrs,
			field.Invalid(path.Child("ServicesIpv4CidrBlock"),
				spec.IPAllocationPolicy.ServicesIpv4CidrBlock,
				"field cannot be set unless UseIPAliases is set to true"),
		)
	}
	if spec.IPAllocationPolicy.ClusterIpv4CidrBlock != nil && !isUseIPAliases {
		allErrs = append(allErrs,
			field.Invalid(path.Child("ClusterIpv4CidrBlock"),
				spec.IPAllocationPolicy.ClusterIpv4CidrBlock,
				"field cannot be set unless UseIPAliases is set to true"),
		)
	}
	if spec.IPAllocationPolicy.ClusterIpv4CidrBlock != nil && spec.ClusterIpv4Cidr != nil {
		allErrs = append(allErrs,
			field.Invalid(path.Child("ClusterIpv4CidrBlock"),
				spec.IPAllocationPolicy.ClusterIpv4CidrBlock,
				"only one of spec.ClusterIpv4Cidr and spec.IPAllocationPolicy.ClusterIpv4CidrBlock can be set"),
		)
	}

	return allErrs
}
