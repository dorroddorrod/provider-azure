/*
Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	"context"

	networkv1alpha3 "github.com/crossplane-contrib/provider-azure/apis/network/v1alpha3"
	"github.com/crossplane-contrib/provider-azure/apis/v1alpha3"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Openshift) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.resourceGroupName
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.ResourceGroupName,
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To:           reference.To{Managed: &v1alpha3.ResourceGroup{}, List: &v1alpha3.ResourceGroupList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.resourceGroupName")
	}
	
	mg.Spec.ForProvider.ResourceGroupName = rsp.ResolvedValue
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.MasterProfile.SubnetID,
		Reference:    mg.Spec.ForProvider.MasterProfile.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.MasterProfile.SubnetIDSelector,
		To:           reference.To{Managed: &networkv1alpha3.Subnet{}, List: &networkv1alpha3.SubnetList{}},
		Extract:      networkv1alpha3.SubnetID(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.subnetID")
	}
	mg.Spec.ForProvider.MasterProfile.SubnetID = rsp.ResolvedValue
	mg.Spec.ForProvider.MasterProfile.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.WorkerProfile.SubnetID,
		Reference:    mg.Spec.ForProvider.WorkerProfile.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.WorkerProfile.SubnetIDSelector,
		To:           reference.To{Managed: &networkv1alpha3.Subnet{}, List: &networkv1alpha3.SubnetList{}},
		Extract:      networkv1alpha3.SubnetID(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.subnetID")
	}
	mg.Spec.ForProvider.WorkerProfile.SubnetID = rsp.ResolvedValue
	mg.Spec.ForProvider.WorkerProfile.SubnetIDRef = rsp.ResolvedReference

	return nil
}
