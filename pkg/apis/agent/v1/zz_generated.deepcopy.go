// +build !ignore_autogenerated

// (c) Copyright IBM Corporation 2019, 2020. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// U.S. Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.
// Licensed Materials - Property of IBM
//
// Copyright (c) 2020 Red Hat, Inc.

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationManager) DeepCopyInto(out *ApplicationManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationManager.
func (in *ApplicationManager) DeepCopy() *ApplicationManager {
	if in == nil {
		return nil
	}
	out := new(ApplicationManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationManagerList) DeepCopyInto(out *ApplicationManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationManagerList.
func (in *ApplicationManagerList) DeepCopy() *ApplicationManagerList {
	if in == nil {
		return nil
	}
	out := new(ApplicationManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationManagerSpec) DeepCopyInto(out *ApplicationManagerSpec) {
	*out = *in
	in.GlobalValues.DeepCopyInto(&out.GlobalValues)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationManagerSpec.
func (in *ApplicationManagerSpec) DeepCopy() *ApplicationManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationManagerStatus) DeepCopyInto(out *ApplicationManagerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationManagerStatus.
func (in *ApplicationManagerStatus) DeepCopy() *ApplicationManagerStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertPolicyController) DeepCopyInto(out *CertPolicyController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertPolicyController.
func (in *CertPolicyController) DeepCopy() *CertPolicyController {
	if in == nil {
		return nil
	}
	out := new(CertPolicyController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertPolicyController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertPolicyControllerList) DeepCopyInto(out *CertPolicyControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertPolicyController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertPolicyControllerList.
func (in *CertPolicyControllerList) DeepCopy() *CertPolicyControllerList {
	if in == nil {
		return nil
	}
	out := new(CertPolicyControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertPolicyControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertPolicyControllerSpec) DeepCopyInto(out *CertPolicyControllerSpec) {
	*out = *in
	in.GlobalValues.DeepCopyInto(&out.GlobalValues)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertPolicyControllerSpec.
func (in *CertPolicyControllerSpec) DeepCopy() *CertPolicyControllerSpec {
	if in == nil {
		return nil
	}
	out := new(CertPolicyControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertPolicyControllerStatus) DeepCopyInto(out *CertPolicyControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertPolicyControllerStatus.
func (in *CertPolicyControllerStatus) DeepCopy() *CertPolicyControllerStatus {
	if in == nil {
		return nil
	}
	out := new(CertPolicyControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalValues) DeepCopyInto(out *GlobalValues) {
	*out = *in
	if in.ImageOverrides != nil {
		in, out := &in.ImageOverrides, &out.ImageOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalValues.
func (in *GlobalValues) DeepCopy() *GlobalValues {
	if in == nil {
		return nil
	}
	out := new(GlobalValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyController) DeepCopyInto(out *IAMPolicyController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyController.
func (in *IAMPolicyController) DeepCopy() *IAMPolicyController {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPolicyController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyControllerList) DeepCopyInto(out *IAMPolicyControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMPolicyController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyControllerList.
func (in *IAMPolicyControllerList) DeepCopy() *IAMPolicyControllerList {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPolicyControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyControllerSpec) DeepCopyInto(out *IAMPolicyControllerSpec) {
	*out = *in
	in.GlobalValues.DeepCopyInto(&out.GlobalValues)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyControllerSpec.
func (in *IAMPolicyControllerSpec) DeepCopy() *IAMPolicyControllerSpec {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyControllerStatus) DeepCopyInto(out *IAMPolicyControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyControllerStatus.
func (in *IAMPolicyControllerStatus) DeepCopy() *IAMPolicyControllerStatus {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfig) DeepCopyInto(out *KlusterletAddonConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfig.
func (in *KlusterletAddonConfig) DeepCopy() *KlusterletAddonConfig {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KlusterletAddonConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigApplicationManagerSpec) DeepCopyInto(out *KlusterletAddonConfigApplicationManagerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigApplicationManagerSpec.
func (in *KlusterletAddonConfigApplicationManagerSpec) DeepCopy() *KlusterletAddonConfigApplicationManagerSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigApplicationManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigCertPolicyControllerSpec) DeepCopyInto(out *KlusterletAddonConfigCertPolicyControllerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigCertPolicyControllerSpec.
func (in *KlusterletAddonConfigCertPolicyControllerSpec) DeepCopy() *KlusterletAddonConfigCertPolicyControllerSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigCertPolicyControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigIAMPolicyControllerSpec) DeepCopyInto(out *KlusterletAddonConfigIAMPolicyControllerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigIAMPolicyControllerSpec.
func (in *KlusterletAddonConfigIAMPolicyControllerSpec) DeepCopy() *KlusterletAddonConfigIAMPolicyControllerSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigIAMPolicyControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigList) DeepCopyInto(out *KlusterletAddonConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KlusterletAddonConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigList.
func (in *KlusterletAddonConfigList) DeepCopy() *KlusterletAddonConfigList {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KlusterletAddonConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigPolicyControllerSpec) DeepCopyInto(out *KlusterletAddonConfigPolicyControllerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigPolicyControllerSpec.
func (in *KlusterletAddonConfigPolicyControllerSpec) DeepCopy() *KlusterletAddonConfigPolicyControllerSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigPolicyControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigSearchCollectorSpec) DeepCopyInto(out *KlusterletAddonConfigSearchCollectorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigSearchCollectorSpec.
func (in *KlusterletAddonConfigSearchCollectorSpec) DeepCopy() *KlusterletAddonConfigSearchCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigSearchCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigSpec) DeepCopyInto(out *KlusterletAddonConfigSpec) {
	*out = *in
	if in.ClusterLabels != nil {
		in, out := &in.ClusterLabels, &out.ClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.SearchCollectorConfig = in.SearchCollectorConfig
	out.PolicyController = in.PolicyController
	out.ApplicationManagerConfig = in.ApplicationManagerConfig
	out.CertPolicyControllerConfig = in.CertPolicyControllerConfig
	out.IAMPolicyControllerConfig = in.IAMPolicyControllerConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigSpec.
func (in *KlusterletAddonConfigSpec) DeepCopy() *KlusterletAddonConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigStatus) DeepCopyInto(out *KlusterletAddonConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigStatus.
func (in *KlusterletAddonConfigStatus) DeepCopy() *KlusterletAddonConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletAddonConfigWorkManagerSpec) DeepCopyInto(out *KlusterletAddonConfigWorkManagerSpec) {
	*out = *in
	if in.ClusterLabels != nil {
		in, out := &in.ClusterLabels, &out.ClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletAddonConfigWorkManagerSpec.
func (in *KlusterletAddonConfigWorkManagerSpec) DeepCopy() *KlusterletAddonConfigWorkManagerSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletAddonConfigWorkManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletPrometheusIntegrationSpec) DeepCopyInto(out *KlusterletPrometheusIntegrationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletPrometheusIntegrationSpec.
func (in *KlusterletPrometheusIntegrationSpec) DeepCopy() *KlusterletPrometheusIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletPrometheusIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyController) DeepCopyInto(out *PolicyController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyController.
func (in *PolicyController) DeepCopy() *PolicyController {
	if in == nil {
		return nil
	}
	out := new(PolicyController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyControllerList) DeepCopyInto(out *PolicyControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyControllerList.
func (in *PolicyControllerList) DeepCopy() *PolicyControllerList {
	if in == nil {
		return nil
	}
	out := new(PolicyControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyControllerSpec) DeepCopyInto(out *PolicyControllerSpec) {
	*out = *in
	in.GlobalValues.DeepCopyInto(&out.GlobalValues)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyControllerSpec.
func (in *PolicyControllerSpec) DeepCopy() *PolicyControllerSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyControllerStatus) DeepCopyInto(out *PolicyControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyControllerStatus.
func (in *PolicyControllerStatus) DeepCopy() *PolicyControllerStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchCollector) DeepCopyInto(out *SearchCollector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchCollector.
func (in *SearchCollector) DeepCopy() *SearchCollector {
	if in == nil {
		return nil
	}
	out := new(SearchCollector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchCollector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchCollectorList) DeepCopyInto(out *SearchCollectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SearchCollector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchCollectorList.
func (in *SearchCollectorList) DeepCopy() *SearchCollectorList {
	if in == nil {
		return nil
	}
	out := new(SearchCollectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchCollectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchCollectorSpec) DeepCopyInto(out *SearchCollectorSpec) {
	*out = *in
	in.GlobalValues.DeepCopyInto(&out.GlobalValues)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchCollectorSpec.
func (in *SearchCollectorSpec) DeepCopy() *SearchCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(SearchCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchCollectorStatus) DeepCopyInto(out *SearchCollectorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchCollectorStatus.
func (in *SearchCollectorStatus) DeepCopy() *SearchCollectorStatus {
	if in == nil {
		return nil
	}
	out := new(SearchCollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManager) DeepCopyInto(out *WorkManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManager.
func (in *WorkManager) DeepCopy() *WorkManager {
	if in == nil {
		return nil
	}
	out := new(WorkManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerList) DeepCopyInto(out *WorkManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerList.
func (in *WorkManagerList) DeepCopy() *WorkManagerList {
	if in == nil {
		return nil
	}
	out := new(WorkManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerSpec) DeepCopyInto(out *WorkManagerSpec) {
	*out = *in
	if in.ClusterLabels != nil {
		in, out := &in.ClusterLabels, &out.ClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.GlobalValues.DeepCopyInto(&out.GlobalValues)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerSpec.
func (in *WorkManagerSpec) DeepCopy() *WorkManagerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerStatus) DeepCopyInto(out *WorkManagerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerStatus.
func (in *WorkManagerStatus) DeepCopy() *WorkManagerStatus {
	if in == nil {
		return nil
	}
	out := new(WorkManagerStatus)
	in.DeepCopyInto(out)
	return out
}
