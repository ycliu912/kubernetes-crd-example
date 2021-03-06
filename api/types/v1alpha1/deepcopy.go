package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInfo is an autogenerated deepcopy functions, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInfo(out *Project) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = ProjectSpec{
		Replicas: in.Spec.Replicas,
	}
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	out := Project{}
	in.DeepCopyInfo(&out)

	return &out
}

// DeepGopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object
func (in *ProjectList) DeepCopyObject() runtime.Object {
	out := ProjectList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Project, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInfo(&out.Items[i])
		}
	}

	return &out
}
