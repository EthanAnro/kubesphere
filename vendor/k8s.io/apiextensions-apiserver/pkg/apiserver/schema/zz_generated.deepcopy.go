//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package schema

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Extensions) DeepCopyInto(out *Extensions) {
	*out = *in
	if in.XListMapKeys != nil {
		in, out := &in.XListMapKeys, &out.XListMapKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.XListType != nil {
		in, out := &in.XListType, &out.XListType
		*out = new(string)
		**out = **in
	}
	if in.XMapType != nil {
		in, out := &in.XMapType, &out.XMapType
		*out = new(string)
		**out = **in
	}
	if in.XValidations != nil {
		in, out := &in.XValidations, &out.XValidations
		*out = make(v1.ValidationRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extensions.
func (in *Extensions) DeepCopy() *Extensions {
	if in == nil {
		return nil
	}
	out := new(Extensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Generic) DeepCopyInto(out *Generic) {
	*out = *in
	out.Default = in.Default.DeepCopy()
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(StructuralOrBool)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Generic.
func (in *Generic) DeepCopy() *Generic {
	if in == nil {
		return nil
	}
	out := new(Generic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedValueValidation) DeepCopyInto(out *NestedValueValidation) {
	*out = *in
	in.ValueValidation.DeepCopyInto(&out.ValueValidation)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = new(NestedValueValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]NestedValueValidation, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.ForbiddenGenerics.DeepCopyInto(&out.ForbiddenGenerics)
	in.ForbiddenExtensions.DeepCopyInto(&out.ForbiddenExtensions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedValueValidation.
func (in *NestedValueValidation) DeepCopy() *NestedValueValidation {
	if in == nil {
		return nil
	}
	out := new(NestedValueValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Structural) DeepCopyInto(out *Structural) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = new(Structural)
		(*in).DeepCopyInto(*out)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]Structural, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.Generic.DeepCopyInto(&out.Generic)
	in.Extensions.DeepCopyInto(&out.Extensions)
	if in.ValueValidation != nil {
		in, out := &in.ValueValidation, &out.ValueValidation
		*out = new(ValueValidation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Structural.
func (in *Structural) DeepCopy() *Structural {
	if in == nil {
		return nil
	}
	out := new(Structural)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StructuralOrBool) DeepCopyInto(out *StructuralOrBool) {
	*out = *in
	if in.Structural != nil {
		in, out := &in.Structural, &out.Structural
		*out = new(Structural)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StructuralOrBool.
func (in *StructuralOrBool) DeepCopy() *StructuralOrBool {
	if in == nil {
		return nil
	}
	out := new(StructuralOrBool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueValidation) DeepCopyInto(out *ValueValidation) {
	*out = *in
	if in.Maximum != nil {
		in, out := &in.Maximum, &out.Maximum
		*out = new(float64)
		**out = **in
	}
	if in.Minimum != nil {
		in, out := &in.Minimum, &out.Minimum
		*out = new(float64)
		**out = **in
	}
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(int64)
		**out = **in
	}
	if in.MinLength != nil {
		in, out := &in.MinLength, &out.MinLength
		*out = new(int64)
		**out = **in
	}
	if in.MaxItems != nil {
		in, out := &in.MaxItems, &out.MaxItems
		*out = new(int64)
		**out = **in
	}
	if in.MinItems != nil {
		in, out := &in.MinItems, &out.MinItems
		*out = new(int64)
		**out = **in
	}
	if in.MultipleOf != nil {
		in, out := &in.MultipleOf, &out.MultipleOf
		*out = new(float64)
		**out = **in
	}
	if in.Enum != nil {
		in, out := &in.Enum, &out.Enum
		*out = make([]JSON, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaxProperties != nil {
		in, out := &in.MaxProperties, &out.MaxProperties
		*out = new(int64)
		**out = **in
	}
	if in.MinProperties != nil {
		in, out := &in.MinProperties, &out.MinProperties
		*out = new(int64)
		**out = **in
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllOf != nil {
		in, out := &in.AllOf, &out.AllOf
		*out = make([]NestedValueValidation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = make([]NestedValueValidation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AnyOf != nil {
		in, out := &in.AnyOf, &out.AnyOf
		*out = make([]NestedValueValidation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(NestedValueValidation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueValidation.
func (in *ValueValidation) DeepCopy() *ValueValidation {
	if in == nil {
		return nil
	}
	out := new(ValueValidation)
	in.DeepCopyInto(out)
	return out
}