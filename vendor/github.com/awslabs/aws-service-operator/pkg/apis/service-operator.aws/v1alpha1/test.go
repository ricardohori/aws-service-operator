package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Test defines the base resource
type Test struct {
	metav1.TypeMeta     `json:",inline"`
	metav1.ObjectMeta   `json:"metadata"`
	Spec                TestSpec                `json:"spec"`
	Status              TestStatus              `json:"status"`
	Output              TestOutput              `json:"output"`
	AdditionalResources TestAdditionalResources `json:"additionalResources"`
}
type TestSpec struct {
	CloudFormationTemplateName      string `json:"cloudFormationTemplateName"`
	CloudFormationTemplateNamespace string `json:"cloudFormationTemplateNamespace"`
	RollbackCount                   int    `json:"rollbackCount"`
}

// TestOutput defines the output resource for Test
type TestOutput struct {
	TestARN string `json:"testARN"`
}

// TestStatus holds the status of the Cloudformation template
type TestStatus struct {
	ResourceStatus       string `json:"resourceStatus"`
	ResourceStatusReason string `json:"resourceStatusReason"`
	StackID              string `json:"stackID"`
}

// TestAdditionalResources holds the additional resources
type TestAdditionalResources struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestList defines the list attribute for the Test type
type TestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Test `json:"items"`
}

func init() {
	localSchemeBuilder.Register(addTestTypes)
}

func addTestTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Test{},
		&TestList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
