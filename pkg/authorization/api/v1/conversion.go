package v1

import (
	"sort"

	"k8s.io/kubernetes/pkg/conversion"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/sets"

	oapi "github.com/openshift/origin/pkg/api"
	"github.com/openshift/origin/pkg/api/extension"
	newer "github.com/openshift/origin/pkg/authorization/api"
	uservalidation "github.com/openshift/origin/pkg/user/api/validation"
)

func Convert_v1_ResourceAccessReview_To_api_ResourceAccessReview(in *ResourceAccessReview, out *newer.ResourceAccessReview, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}
	if err := s.DefaultConvert(&in.AuthorizationAttributes, &out.Action, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	return nil
}

func Convert_api_ResourceAccessReview_To_v1_ResourceAccessReview(in *newer.ResourceAccessReview, out *ResourceAccessReview, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}
	if err := s.DefaultConvert(&in.Action, &out.AuthorizationAttributes, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	return nil
}

func Convert_v1_LocalResourceAccessReview_To_api_LocalResourceAccessReview(in *LocalResourceAccessReview, out *newer.LocalResourceAccessReview, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}
	if err := s.DefaultConvert(&in.AuthorizationAttributes, &out.Action, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	return nil
}

func Convert_api_LocalResourceAccessReview_To_v1_LocalResourceAccessReview(in *newer.LocalResourceAccessReview, out *LocalResourceAccessReview, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}
	if err := s.DefaultConvert(&in.Action, &out.AuthorizationAttributes, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	return nil
}

func Convert_v1_SubjectAccessReview_To_api_SubjectAccessReview(in *SubjectAccessReview, out *newer.SubjectAccessReview, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}
	if err := s.DefaultConvert(&in.AuthorizationAttributes, &out.Action, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	out.Groups = sets.NewString(in.GroupsSlice...)

	return nil
}

func Convert_api_SubjectAccessReview_To_v1_SubjectAccessReview(in *newer.SubjectAccessReview, out *SubjectAccessReview, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}
	if err := s.DefaultConvert(&in.Action, &out.AuthorizationAttributes, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	out.GroupsSlice = in.Groups.List()

	return nil
}

func Convert_v1_LocalSubjectAccessReview_To_api_LocalSubjectAccessReview(in *LocalSubjectAccessReview, out *newer.LocalSubjectAccessReview, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}
	if err := s.DefaultConvert(&in.AuthorizationAttributes, &out.Action, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	out.Groups = sets.NewString(in.GroupsSlice...)

	return nil
}

func Convert_api_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(in *newer.LocalSubjectAccessReview, out *LocalSubjectAccessReview, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}
	if err := s.DefaultConvert(&in.Action, &out.AuthorizationAttributes, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	out.GroupsSlice = in.Groups.List()

	return nil
}

func Convert_v1_ResourceAccessReviewResponse_To_api_ResourceAccessReviewResponse(in *ResourceAccessReviewResponse, out *newer.ResourceAccessReviewResponse, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	out.Users = sets.NewString(in.UsersSlice...)
	out.Groups = sets.NewString(in.GroupsSlice...)

	return nil
}

func Convert_api_ResourceAccessReviewResponse_To_v1_ResourceAccessReviewResponse(in *newer.ResourceAccessReviewResponse, out *ResourceAccessReviewResponse, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
		return err
	}

	out.UsersSlice = in.Users.List()
	out.GroupsSlice = in.Groups.List()

	return nil
}

func Convert_v1_PolicyRule_To_api_PolicyRule(in *PolicyRule, out *newer.PolicyRule, s conversion.Scope) error {
	SetDefaults_PolicyRule(in)
	if err := oapi.Convert_runtime_RawExtension_To_runtime_Object(&in.AttributeRestrictions, &out.AttributeRestrictions, s); err != nil {
		return err
	}

	out.APIGroups = in.APIGroups

	out.Resources = sets.String{}
	out.Resources.Insert(in.Resources...)

	out.Verbs = sets.String{}
	out.Verbs.Insert(in.Verbs...)

	out.ResourceNames = sets.NewString(in.ResourceNames...)

	out.NonResourceURLs = sets.NewString(in.NonResourceURLsSlice...)

	return nil
}

func Convert_api_PolicyRule_To_v1_PolicyRule(in *newer.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	if err := oapi.Convert_runtime_Object_To_runtime_RawExtension(&in.AttributeRestrictions, &out.AttributeRestrictions, s); err != nil {
		return err
	}

	out.APIGroups = in.APIGroups

	out.Resources = []string{}
	out.Resources = append(out.Resources, in.Resources.List()...)

	out.Verbs = []string{}
	out.Verbs = append(out.Verbs, in.Verbs.List()...)

	out.ResourceNames = in.ResourceNames.List()

	out.NonResourceURLsSlice = in.NonResourceURLs.List()

	return nil
}

func Convert_v1_Policy_To_api_Policy(in *Policy, out *newer.Policy, s conversion.Scope) error {
	out.LastModified = in.LastModified
	out.Roles = make(map[string]*newer.Role)
	return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
}

func Convert_api_Policy_To_v1_Policy(in *newer.Policy, out *Policy, s conversion.Scope) error {
	out.LastModified = in.LastModified
	out.Roles = make([]NamedRole, 0, 0)
	return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
}

func Convert_v1_RoleBinding_To_api_RoleBinding(in *RoleBinding, out *newer.RoleBinding, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields|conversion.AllowDifferentFieldTypeNames); err != nil {
		return err
	}

	// if the users and groups fields are cleared, then respect only subjects.  The field was set in the DefaultConvert above
	if in.UserNames == nil && in.GroupNames == nil {
		return nil
	}

	out.Subjects = newer.BuildSubjects(in.UserNames, in.GroupNames, uservalidation.ValidateUserName, uservalidation.ValidateGroupName)

	return nil
}

func Convert_api_RoleBinding_To_v1_RoleBinding(in *newer.RoleBinding, out *RoleBinding, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields|conversion.AllowDifferentFieldTypeNames); err != nil {
		return err
	}

	out.UserNames, out.GroupNames = newer.StringSubjectsFor(in.Namespace, in.Subjects)

	return nil
}

func Convert_v1_PolicyBinding_To_api_PolicyBinding(in *PolicyBinding, out *newer.PolicyBinding, s conversion.Scope) error {
	out.LastModified = in.LastModified
	out.RoleBindings = make(map[string]*newer.RoleBinding)
	return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
}

func Convert_api_PolicyBinding_To_v1_PolicyBinding(in *newer.PolicyBinding, out *PolicyBinding, s conversion.Scope) error {
	out.LastModified = in.LastModified
	out.RoleBindings = make([]NamedRoleBinding, 0, 0)
	return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
}

// and now the globals
func Convert_v1_ClusterPolicy_To_api_ClusterPolicy(in *ClusterPolicy, out *newer.ClusterPolicy, s conversion.Scope) error {
	out.LastModified = in.LastModified
	out.Roles = make(map[string]*newer.ClusterRole)
	return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
}

func Convert_api_ClusterPolicy_To_v1_ClusterPolicy(in *newer.ClusterPolicy, out *ClusterPolicy, s conversion.Scope) error {
	out.LastModified = in.LastModified
	out.Roles = make([]NamedClusterRole, 0, 0)
	return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
}

func Convert_v1_ClusterRoleBinding_To_api_ClusterRoleBinding(in *ClusterRoleBinding, out *newer.ClusterRoleBinding, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields|conversion.AllowDifferentFieldTypeNames); err != nil {
		return err
	}

	// if the users and groups fields are cleared, then respect only subjects.  The field was set in the DefaultConvert above
	if in.UserNames == nil && in.GroupNames == nil {
		return nil
	}

	out.Subjects = newer.BuildSubjects(in.UserNames, in.GroupNames, uservalidation.ValidateUserName, uservalidation.ValidateGroupName)

	return nil
}

func Convert_api_ClusterRoleBinding_To_v1_ClusterRoleBinding(in *newer.ClusterRoleBinding, out *ClusterRoleBinding, s conversion.Scope) error {
	if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields|conversion.AllowDifferentFieldTypeNames); err != nil {
		return err
	}

	out.UserNames, out.GroupNames = newer.StringSubjectsFor(in.Namespace, in.Subjects)

	return nil
}

func Convert_v1_ClusterPolicyBinding_To_api_ClusterPolicyBinding(in *ClusterPolicyBinding, out *newer.ClusterPolicyBinding, s conversion.Scope) error {
	out.LastModified = in.LastModified
	out.RoleBindings = make(map[string]*newer.ClusterRoleBinding)
	return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
}

func Convert_api_ClusterPolicyBinding_To_v1_ClusterPolicyBinding(in *newer.ClusterPolicyBinding, out *ClusterPolicyBinding, s conversion.Scope) error {
	out.LastModified = in.LastModified
	out.RoleBindings = make([]NamedClusterRoleBinding, 0, 0)
	return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
}

func Convert_v1_NamedRoleArray_to_api_RoleArray(in *[]NamedRole, out *map[string]*newer.Role, s conversion.Scope) error {
	for _, curr := range *in {
		newRole := &newer.Role{}
		if err := s.Convert(&curr.Role, newRole, 0); err != nil {
			return err
		}
		(*out)[curr.Name] = newRole
	}

	return nil
}
func Convert_api_NamedRoleArray_to_v1_RoleArray(in *map[string]*newer.Role, out *[]NamedRole, s conversion.Scope) error {
	allKeys := make([]string, 0, len(*in))
	for key := range *in {
		allKeys = append(allKeys, key)
	}
	sort.Strings(allKeys)

	for _, key := range allKeys {
		newRole := (*in)[key]
		oldRole := &Role{}
		if err := s.Convert(newRole, oldRole, 0); err != nil {
			return err
		}

		namedRole := NamedRole{key, *oldRole}
		*out = append(*out, namedRole)
	}

	return nil
}

func Convert_v1_NamedRoleBindingArray_to_api_RoleBindingArray(in *[]NamedRoleBinding, out *map[string]*newer.RoleBinding, s conversion.Scope) error {
	for _, curr := range *in {
		newRoleBinding := &newer.RoleBinding{}
		if err := s.Convert(&curr.RoleBinding, newRoleBinding, 0); err != nil {
			return err
		}
		(*out)[curr.Name] = newRoleBinding
	}

	return nil
}
func Convert_api_RoleBindingArray_to_v1_NamedRoleBindingArray(in *map[string]*newer.RoleBinding, out *[]NamedRoleBinding, s conversion.Scope) error {
	allKeys := make([]string, 0, len(*in))
	for key := range *in {
		allKeys = append(allKeys, key)
	}
	sort.Strings(allKeys)

	for _, key := range allKeys {
		newRoleBinding := (*in)[key]
		oldRoleBinding := &RoleBinding{}
		if err := s.Convert(newRoleBinding, oldRoleBinding, 0); err != nil {
			return err
		}

		namedRoleBinding := NamedRoleBinding{key, *oldRoleBinding}
		*out = append(*out, namedRoleBinding)
	}

	return nil
}

func Convert_v1_NamedClusterRoleArray_to_api_ClusterRoleArray(in *[]NamedClusterRole, out *map[string]*newer.ClusterRole, s conversion.Scope) error {
	for _, curr := range *in {
		newRole := &newer.ClusterRole{}
		if err := s.Convert(&curr.Role, newRole, 0); err != nil {
			return err
		}
		(*out)[curr.Name] = newRole
	}

	return nil
}
func Convert_api_ClusterRoleArray_to_v1_NamedClusterRoleArray(in *map[string]*newer.ClusterRole, out *[]NamedClusterRole, s conversion.Scope) error {
	allKeys := make([]string, 0, len(*in))
	for key := range *in {
		allKeys = append(allKeys, key)
	}
	sort.Strings(allKeys)

	for _, key := range allKeys {
		newRole := (*in)[key]
		oldRole := &ClusterRole{}
		if err := s.Convert(newRole, oldRole, 0); err != nil {
			return err
		}

		namedRole := NamedClusterRole{key, *oldRole}
		*out = append(*out, namedRole)
	}

	return nil
}
func Convert_v1_NamedClusterRoleBindingArray_to_ClusterRoleBindingArray(in *[]NamedClusterRoleBinding, out *map[string]*newer.ClusterRoleBinding, s conversion.Scope) error {
	for _, curr := range *in {
		newRoleBinding := &newer.ClusterRoleBinding{}
		if err := s.Convert(&curr.RoleBinding, newRoleBinding, 0); err != nil {
			return err
		}
		(*out)[curr.Name] = newRoleBinding
	}

	return nil
}
func Convert_api_ClusterRoleBindingArray_to_v1_NamedClusterRoleBindingArray(in *map[string]*newer.ClusterRoleBinding, out *[]NamedClusterRoleBinding, s conversion.Scope) error {
	allKeys := make([]string, 0, len(*in))
	for key := range *in {
		allKeys = append(allKeys, key)
	}
	sort.Strings(allKeys)

	for _, key := range allKeys {
		newRoleBinding := (*in)[key]
		oldRoleBinding := &ClusterRoleBinding{}
		if err := s.Convert(newRoleBinding, oldRoleBinding, 0); err != nil {
			return err
		}

		namedRoleBinding := NamedClusterRoleBinding{key, *oldRoleBinding}
		*out = append(*out, namedRoleBinding)
	}

	return nil
}

func addConversionFuncs(scheme *runtime.Scheme) {
	err := scheme.AddConversionFuncs(
		Convert_v1_NamedRoleArray_to_api_RoleArray,
		Convert_api_NamedRoleArray_to_v1_RoleArray,
		Convert_v1_NamedRoleBindingArray_to_api_RoleBindingArray,
		Convert_api_RoleBindingArray_to_v1_NamedRoleBindingArray,
		Convert_v1_NamedClusterRoleArray_to_api_ClusterRoleArray,
		Convert_api_ClusterRoleArray_to_v1_NamedClusterRoleArray,
		Convert_v1_NamedClusterRoleBindingArray_to_ClusterRoleBindingArray,
		Convert_api_ClusterRoleBindingArray_to_v1_NamedClusterRoleBindingArray,

		Convert_v1_SubjectAccessReview_To_api_SubjectAccessReview,
		Convert_api_SubjectAccessReview_To_v1_SubjectAccessReview,
		Convert_v1_LocalSubjectAccessReview_To_api_LocalSubjectAccessReview,
		Convert_api_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview,
		Convert_v1_ResourceAccessReview_To_api_ResourceAccessReview,
		Convert_api_ResourceAccessReview_To_v1_ResourceAccessReview,
		Convert_v1_LocalResourceAccessReview_To_api_LocalResourceAccessReview,
		Convert_api_LocalResourceAccessReview_To_v1_LocalResourceAccessReview,
		Convert_v1_ResourceAccessReviewResponse_To_api_ResourceAccessReviewResponse,
		Convert_api_ResourceAccessReviewResponse_To_v1_ResourceAccessReviewResponse,
		Convert_v1_PolicyRule_To_api_PolicyRule,
		Convert_api_PolicyRule_To_v1_PolicyRule,
		Convert_v1_Policy_To_api_Policy,
		Convert_api_Policy_To_v1_Policy,
		Convert_v1_RoleBinding_To_api_RoleBinding,
		Convert_api_RoleBinding_To_v1_RoleBinding,
		Convert_v1_PolicyBinding_To_api_PolicyBinding,
		Convert_api_PolicyBinding_To_v1_PolicyBinding,
		Convert_v1_ClusterPolicy_To_api_ClusterPolicy,
		Convert_api_ClusterPolicy_To_v1_ClusterPolicy,
		Convert_v1_ClusterRoleBinding_To_api_ClusterRoleBinding,
		Convert_api_ClusterRoleBinding_To_v1_ClusterRoleBinding,
		Convert_v1_ClusterPolicyBinding_To_api_ClusterPolicyBinding,
		Convert_api_ClusterPolicyBinding_To_v1_ClusterPolicyBinding,
	)
	if err != nil {
		// If one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}

	if err := scheme.AddFieldLabelConversionFunc("v1", "ClusterPolicy",
		oapi.GetFieldLabelConversionFunc(newer.ClusterPolicyToSelectableFields(&newer.ClusterPolicy{}), nil),
	); err != nil {
		panic(err)
	}

	if err := scheme.AddFieldLabelConversionFunc("v1", "ClusterPolicyBinding",
		oapi.GetFieldLabelConversionFunc(newer.ClusterPolicyBindingToSelectableFields(&newer.ClusterPolicyBinding{}), nil),
	); err != nil {
		panic(err)
	}

	if err := scheme.AddFieldLabelConversionFunc("v1", "Policy",
		oapi.GetFieldLabelConversionFunc(newer.PolicyToSelectableFields(&newer.Policy{}), nil),
	); err != nil {
		panic(err)
	}

	if err := scheme.AddFieldLabelConversionFunc("v1", "PolicyBinding",
		oapi.GetFieldLabelConversionFunc(newer.PolicyBindingToSelectableFields(&newer.PolicyBinding{}), nil),
	); err != nil {
		panic(err)
	}

	if err := scheme.AddFieldLabelConversionFunc("v1", "Role",
		oapi.GetFieldLabelConversionFunc(newer.RoleToSelectableFields(&newer.Role{}), nil),
	); err != nil {
		panic(err)
	}

	if err := scheme.AddFieldLabelConversionFunc("v1", "RoleBinding",
		oapi.GetFieldLabelConversionFunc(newer.RoleBindingToSelectableFields(&newer.RoleBinding{}), nil),
	); err != nil {
		panic(err)
	}
}

var _ runtime.NestedObjectDecoder = &PolicyRule{}
var _ runtime.NestedObjectEncoder = &PolicyRule{}

func (c *PolicyRule) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	extension.DecodeNestedRawExtensionOrUnknown(d, &c.AttributeRestrictions)
	return nil
}
func (c *PolicyRule) EncodeNestedObjects(e runtime.Encoder) error {
	return extension.EncodeNestedRawExtension(e, &c.AttributeRestrictions)
}

var _ runtime.NestedObjectDecoder = &SelfSubjectRulesReview{}
var _ runtime.NestedObjectEncoder = &SelfSubjectRulesReview{}

func (c *SelfSubjectRulesReview) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Status.Rules {
		c.Status.Rules[i].DecodeNestedObjects(d)
	}
	return nil
}
func (c *SelfSubjectRulesReview) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Status.Rules {
		if err := c.Status.Rules[i].EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}

var _ runtime.NestedObjectDecoder = &ClusterRole{}
var _ runtime.NestedObjectEncoder = &ClusterRole{}

func (c *ClusterRole) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Rules {
		c.Rules[i].DecodeNestedObjects(d)
	}
	return nil
}
func (c *ClusterRole) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Rules {
		if err := c.Rules[i].EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}

var _ runtime.NestedObjectDecoder = &Role{}
var _ runtime.NestedObjectEncoder = &Role{}

func (c *Role) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Rules {
		c.Rules[i].DecodeNestedObjects(d)
	}
	return nil
}
func (c *Role) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Rules {
		if err := c.Rules[i].EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}

var _ runtime.NestedObjectDecoder = &ClusterRoleList{}
var _ runtime.NestedObjectEncoder = &ClusterRoleList{}

func (c *ClusterRoleList) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Items {
		c.Items[i].DecodeNestedObjects(d)
	}
	return nil
}
func (c *ClusterRoleList) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Items {
		if err := c.Items[i].EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}

var _ runtime.NestedObjectDecoder = &RoleList{}
var _ runtime.NestedObjectEncoder = &RoleList{}

func (c *RoleList) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Items {
		c.Items[i].DecodeNestedObjects(d)
	}
	return nil
}
func (c *RoleList) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Items {
		if err := c.Items[i].EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}

var _ runtime.NestedObjectDecoder = &ClusterPolicy{}
var _ runtime.NestedObjectEncoder = &ClusterPolicy{}

func (c *ClusterPolicy) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Roles {
		c.Roles[i].Role.DecodeNestedObjects(d)
	}
	return nil
}
func (c *ClusterPolicy) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Roles {
		if err := c.Roles[i].Role.EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}

var _ runtime.NestedObjectDecoder = &Policy{}
var _ runtime.NestedObjectEncoder = &Policy{}

func (c *Policy) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Roles {
		c.Roles[i].Role.DecodeNestedObjects(d)
	}
	return nil
}
func (c *Policy) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Roles {
		if err := c.Roles[i].Role.EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}

var _ runtime.NestedObjectDecoder = &ClusterPolicyList{}
var _ runtime.NestedObjectEncoder = &ClusterPolicyList{}

func (c *ClusterPolicyList) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Items {
		c.Items[i].DecodeNestedObjects(d)
	}
	return nil
}
func (c *ClusterPolicyList) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Items {
		if err := c.Items[i].EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}

var _ runtime.NestedObjectDecoder = &PolicyList{}
var _ runtime.NestedObjectEncoder = &PolicyList{}

func (c *PolicyList) DecodeNestedObjects(d runtime.Decoder) error {
	// decoding failures result in a runtime.Unknown object being created in Object and passed
	// to conversion
	for i := range c.Items {
		c.Items[i].DecodeNestedObjects(d)
	}
	return nil
}
func (c *PolicyList) EncodeNestedObjects(e runtime.Encoder) error {
	for i := range c.Items {
		if err := c.Items[i].EncodeNestedObjects(e); err != nil {
			return err
		}
	}
	return nil
}
