//go:build !ignore_autogenerated

/*
Copyright  The cert-manager Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicy) DeepCopyInto(out *CertificateRequestPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicy.
func (in *CertificateRequestPolicy) DeepCopy() *CertificateRequestPolicy {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRequestPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyAllowed) DeepCopyInto(out *CertificateRequestPolicyAllowed) {
	*out = *in
	if in.CommonName != nil {
		in, out := &in.CommonName, &out.CommonName
		*out = new(CertificateRequestPolicyAllowedString)
		(*in).DeepCopyInto(*out)
	}
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.URIs != nil {
		in, out := &in.URIs, &out.URIs
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.EmailAddresses != nil {
		in, out := &in.EmailAddresses, &out.EmailAddresses
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.IsCA != nil {
		in, out := &in.IsCA, &out.IsCA
		*out = new(bool)
		**out = **in
	}
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = new([]v1.KeyUsage)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.KeyUsage, len(*in))
			copy(*out, *in)
		}
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(CertificateRequestPolicyAllowedX509Subject)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyAllowed.
func (in *CertificateRequestPolicyAllowed) DeepCopy() *CertificateRequestPolicyAllowed {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyAllowed)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyAllowedString) DeepCopyInto(out *CertificateRequestPolicyAllowedString) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
	if in.Validations != nil {
		in, out := &in.Validations, &out.Validations
		*out = make([]ValidationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyAllowedString.
func (in *CertificateRequestPolicyAllowedString) DeepCopy() *CertificateRequestPolicyAllowedString {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyAllowedString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyAllowedStringSlice) DeepCopyInto(out *CertificateRequestPolicyAllowedStringSlice) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
	if in.Validations != nil {
		in, out := &in.Validations, &out.Validations
		*out = make([]ValidationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyAllowedStringSlice.
func (in *CertificateRequestPolicyAllowedStringSlice) DeepCopy() *CertificateRequestPolicyAllowedStringSlice {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyAllowedStringSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyAllowedX509Subject) DeepCopyInto(out *CertificateRequestPolicyAllowedX509Subject) {
	*out = *in
	if in.Organizations != nil {
		in, out := &in.Organizations, &out.Organizations
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.Countries != nil {
		in, out := &in.Countries, &out.Countries
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.OrganizationalUnits != nil {
		in, out := &in.OrganizationalUnits, &out.OrganizationalUnits
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.Localities != nil {
		in, out := &in.Localities, &out.Localities
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.Provinces != nil {
		in, out := &in.Provinces, &out.Provinces
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.StreetAddresses != nil {
		in, out := &in.StreetAddresses, &out.StreetAddresses
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.PostalCodes != nil {
		in, out := &in.PostalCodes, &out.PostalCodes
		*out = new(CertificateRequestPolicyAllowedStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.SerialNumber != nil {
		in, out := &in.SerialNumber, &out.SerialNumber
		*out = new(CertificateRequestPolicyAllowedString)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyAllowedX509Subject.
func (in *CertificateRequestPolicyAllowedX509Subject) DeepCopy() *CertificateRequestPolicyAllowedX509Subject {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyAllowedX509Subject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyCondition) DeepCopyInto(out *CertificateRequestPolicyCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyCondition.
func (in *CertificateRequestPolicyCondition) DeepCopy() *CertificateRequestPolicyCondition {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyConstraints) DeepCopyInto(out *CertificateRequestPolicyConstraints) {
	*out = *in
	if in.MinDuration != nil {
		in, out := &in.MinDuration, &out.MinDuration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.MaxDuration != nil {
		in, out := &in.MaxDuration, &out.MaxDuration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(CertificateRequestPolicyConstraintsPrivateKey)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyConstraints.
func (in *CertificateRequestPolicyConstraints) DeepCopy() *CertificateRequestPolicyConstraints {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyConstraintsPrivateKey) DeepCopyInto(out *CertificateRequestPolicyConstraintsPrivateKey) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(v1.PrivateKeyAlgorithm)
		**out = **in
	}
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyConstraintsPrivateKey.
func (in *CertificateRequestPolicyConstraintsPrivateKey) DeepCopy() *CertificateRequestPolicyConstraintsPrivateKey {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyConstraintsPrivateKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyList) DeepCopyInto(out *CertificateRequestPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateRequestPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyList.
func (in *CertificateRequestPolicyList) DeepCopy() *CertificateRequestPolicyList {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRequestPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyPluginData) DeepCopyInto(out *CertificateRequestPolicyPluginData) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyPluginData.
func (in *CertificateRequestPolicyPluginData) DeepCopy() *CertificateRequestPolicyPluginData {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyPluginData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicySelector) DeepCopyInto(out *CertificateRequestPolicySelector) {
	*out = *in
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(CertificateRequestPolicySelectorIssuerRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(CertificateRequestPolicySelectorNamespace)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicySelector.
func (in *CertificateRequestPolicySelector) DeepCopy() *CertificateRequestPolicySelector {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicySelectorIssuerRef) DeepCopyInto(out *CertificateRequestPolicySelectorIssuerRef) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicySelectorIssuerRef.
func (in *CertificateRequestPolicySelectorIssuerRef) DeepCopy() *CertificateRequestPolicySelectorIssuerRef {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicySelectorIssuerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicySelectorNamespace) DeepCopyInto(out *CertificateRequestPolicySelectorNamespace) {
	*out = *in
	if in.MatchNames != nil {
		in, out := &in.MatchNames, &out.MatchNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicySelectorNamespace.
func (in *CertificateRequestPolicySelectorNamespace) DeepCopy() *CertificateRequestPolicySelectorNamespace {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicySelectorNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicySpec) DeepCopyInto(out *CertificateRequestPolicySpec) {
	*out = *in
	if in.Allowed != nil {
		in, out := &in.Allowed, &out.Allowed
		*out = new(CertificateRequestPolicyAllowed)
		(*in).DeepCopyInto(*out)
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = new(CertificateRequestPolicyConstraints)
		(*in).DeepCopyInto(*out)
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make(map[string]CertificateRequestPolicyPluginData, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicySpec.
func (in *CertificateRequestPolicySpec) DeepCopy() *CertificateRequestPolicySpec {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyStatus) DeepCopyInto(out *CertificateRequestPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateRequestPolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyStatus.
func (in *CertificateRequestPolicyStatus) DeepCopy() *CertificateRequestPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRule) DeepCopyInto(out *ValidationRule) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRule.
func (in *ValidationRule) DeepCopy() *ValidationRule {
	if in == nil {
		return nil
	}
	out := new(ValidationRule)
	in.DeepCopyInto(out)
	return out
}
