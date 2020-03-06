// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAppPolicyDeploymentSummaryPerApp undocumented
type ManagedAppPolicyDeploymentSummaryPerApp struct {
	// Object is the base model of ManagedAppPolicyDeploymentSummaryPerApp
	Object
	// MobileAppIdentifier Deployment of an app.
	MobileAppIdentifier *MobileAppIdentifier `json:"mobileAppIdentifier,omitempty"`
	// ConfigurationAppliedUserCount Number of users the policy is applied.
	ConfigurationAppliedUserCount *int `json:"configurationAppliedUserCount,omitempty"`
}