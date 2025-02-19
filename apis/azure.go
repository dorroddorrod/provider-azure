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

// Package apis contains Kubernetes API groups for Azure cloud provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	cachev1beta1 "github.com/crossplane-contrib/provider-azure/apis/cache/v1beta1"
	computev1alpha3 "github.com/crossplane-contrib/provider-azure/apis/compute/v1alpha3"
	databasev1alpha3 "github.com/crossplane-contrib/provider-azure/apis/database/v1alpha3"
	databasev1beta1 "github.com/crossplane-contrib/provider-azure/apis/database/v1beta1"
	dnsv1alpha1 "github.com/crossplane-contrib/provider-azure/apis/dns/v1alpha1"
	keyvaultv1alpha1 "github.com/crossplane-contrib/provider-azure/apis/keyvault/v1alpha1"
	networkv1alpha3 "github.com/crossplane-contrib/provider-azure/apis/network/v1alpha3"
	storagev1alpha3 "github.com/crossplane-contrib/provider-azure/apis/storage/v1alpha3"
	azurev1alpha1 "github.com/crossplane-contrib/provider-azure/apis/v1alpha1"
	azurev1alpha3 "github.com/crossplane-contrib/provider-azure/apis/v1alpha3"
	azurev1beta1 "github.com/crossplane-contrib/provider-azure/apis/v1beta1"
	openshiftv1alpha1 "github.com/crossplane-contrib/provider-azure/apis/containers/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		azurev1alpha1.SchemeBuilder.AddToScheme,
		azurev1alpha3.SchemeBuilder.AddToScheme,
		azurev1beta1.SchemeBuilder.AddToScheme,
		cachev1beta1.SchemeBuilder.AddToScheme,
		computev1alpha3.SchemeBuilder.AddToScheme,
		databasev1alpha3.SchemeBuilder.AddToScheme,
		databasev1beta1.SchemeBuilder.AddToScheme,
		keyvaultv1alpha1.SchemeBuilder.AddToScheme,
		networkv1alpha3.SchemeBuilder.AddToScheme,
		storagev1alpha3.SchemeBuilder.AddToScheme,
		dnsv1alpha1.SchemeBuilder.AddToScheme,
		openshiftv1alpha1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
