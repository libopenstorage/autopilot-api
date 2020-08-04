package v1alpha1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// ActionApprovalResourcePlural is the name of the plural ActionApproval objects
	ActionApprovalResourcePlural = "actionapprovals"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ActionApproval ia used for approval of actions to be taken by Autopilot as part of a rule
type ActionApproval struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`
	Spec            ActionApprovalSpec   `json:"spec"`
	Status          ActionApprovalStatus `json:"status"`
}

// ActionApprovalSpec is the spec for an autopilot action approval
type ActionApprovalSpec struct {
	// ApprovalState is the current approval state of the action
	ApprovalState ActionApprovalState `json:"approvalState,omitempty"`
	// Actions are the actions that needs/needed approval
	Actions []*RuleAction `json:"actions,omitempty"`
}

// ActionApprovalStatus is current status of an autopilot action approval
type ActionApprovalStatus struct {
	// Rule is the parent autopilot rule that resulted in this action approval
	Rule types.NamespacedName
	// LastProcessTimestamp was the last time this approval was processed
	LastProcessTimestamp meta.Time `json:"lastProcessTimestamp"`
	// ActionPreviews provides a dry-run preview of the side-effects of the actions
	ActionPreviews []*AutopilotActionPreview `json:"actionPreviews,omitempty"`
}

// AutopilotActionPreview is a preview of an action and the expected result of it before it gets executed
type AutopilotActionPreview struct {
	// Action is the action spec that's going to get executed
	Action *RuleAction `json:"action,omitempty"`
	// ExpectedResult is a user friendly description of the outcome of executing the action
	ExpectedResult string `json:"expectedResult,omitempty"`
	// InvolvedObjects are the objects that are directly relevant to the action
	InvolvedObjects []*ActionApprovalInvolvedObject `json:"involvedObjects,omitempty"`
}

// ActionApprovalInvolvedObject represents a particular object that needs an action approval for it's
// action to proceed
type ActionApprovalInvolvedObject struct {
	Name            string                `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Namespace       string                `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	UID             types.UID             `json:"uid,omitempty" protobuf:"bytes,5,opt,name=uid,casttype=k8s.io/kubernetes/pkg/types.UID"`
	Kind            string                `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	APIVersion      string                `json:"apiVersion" protobuf:"bytes,5,opt,name=apiVersion"`
	ClusterName     string                `json:"clusterName,omitempty"`
	OwnerReferences []meta.OwnerReference `json:"ownerReferences,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ActionApprovalList is a list of ActionApproval objects in Kubernetes
type ActionApprovalList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ActionApproval `json:"items"`
}
