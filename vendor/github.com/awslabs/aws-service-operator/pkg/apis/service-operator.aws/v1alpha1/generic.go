package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GenericTopic defines the base resource
type GenericTopic struct {
	metav1.TypeMeta     `json:",inline"`
	metav1.ObjectMeta   `json:"metadata"`
	Spec                GenericTopicSpec                `json:"spec"`
	Status              GenericTopicStatus              `json:"status"`
	Output              GenericTopicOutput              `json:"output"`
	AdditionalResources GenericTopicAdditionalResources `json:"additionalResources"`
}
type GenericTopicSpec struct {
	CloudFormationTemplateName      string `json:"cloudFormationTemplateName"`
	CloudFormationTemplateNamespace string `json:"cloudFormationTemplateNamespace"`
	RollbackCount                   int    `json:"rollbackCount"`
}

// GenericTopicOutput defines the output resource for GenericTopic
type GenericTopicOutput struct {
	GenericARN string `json:"genericARN"`
}

// GenericTopicStatus holds the status of the Cloudformation template
type GenericTopicStatus struct {
	ResourceStatus       string `json:"resourceStatus"`
	ResourceStatusReason string `json:"resourceStatusReason"`
	StackID              string `json:"stackID"`
}

// GenericTopicAdditionalResources holds the additional resources
type GenericTopicAdditionalResources struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GenericTopicList defines the list attribute for the GenericTopic type
type GenericTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GenericTopic `json:"items"`
}

func init() {
	localSchemeBuilder.Register(addGenericTopicTypes)
}

func addGenericTopicTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&GenericTopic{},
		&GenericTopicList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
