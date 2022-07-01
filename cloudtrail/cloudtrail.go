package cloudtrail

import "time"

// Records is a representation of CloudTrail record as seen in S3 stored objects.
// It contains a list of Record objects.
type Records struct {
	Records []Record `json:"Records"`
}

// Record is a single CloudTrail event as stored in S3.
type Record struct {
	EventVersion string `json:"eventVersion"`
	UserIdentity struct {
		Type           string `json:"type"`
		PrincipalId    string `json:"principalId"`
		Arn            string `json:"arn"`
		AccountId      string `json:"accountId"`
		AccessKeyId    string `json:"accessKeyId"`
		SessionContext struct {
			SessionIssuer struct {
				Type        string `json:"type"`
				PrincipalId string `json:"principalId"`
				Arn         string `json:"arn"`
				AccountId   string `json:"accountId"`
				UserName    string `json:"userName"`
			} `json:"sessionIssuer"`
			WebIdFederationData struct {
			} `json:"webIdFederationData"`
			Attributes struct {
				CreationDate     time.Time `json:"creationDate"`
				MfaAuthenticated string    `json:"mfaAuthenticated"`
			} `json:"attributes"`
		} `json:"sessionContext"`
	} `json:"userIdentity"`
	EventTime         time.Time `json:"eventTime"`
	EventSource       string    `json:"eventSource"`
	EventName         string    `json:"eventName"`
	AwsRegion         string    `json:"awsRegion"`
	SourceIPAddress   string    `json:"sourceIPAddress"`
	UserAgent         string    `json:"userAgent"`
	RequestParameters struct {
		ResolveConflicts   string `json:"resolveConflicts"`
		AddonName          string `json:"addonName"`
		ClientRequestToken string `json:"clientRequestToken"`
		Name               string `json:"name"`
		AddonVersion       string `json:"addonVersion"`
	} `json:"requestParameters"`
	ResponseElements struct {
		Addon struct {
			AddonName    string `json:"addonName"`
			ClusterName  string `json:"clusterName"`
			Status       string `json:"status"`
			AddonVersion string `json:"addonVersion"`
			Health       struct {
				Issues []interface{} `json:"issues"`
			} `json:"health"`
			AddonArn   string  `json:"addonArn"`
			CreatedAt  float64 `json:"createdAt"`
			ModifiedAt float64 `json:"modifiedAt"`
			Tags       struct {
			} `json:"tags"`
		} `json:"addon"`
	} `json:"responseElements"`
	RequestID          string `json:"requestID"`
	EventID            string `json:"eventID"`
	ReadOnly           bool   `json:"readOnly"`
	EventType          string `json:"eventType"`
	ManagementEvent    bool   `json:"managementEvent"`
	RecipientAccountId string `json:"recipientAccountId"`
	EventCategory      string `json:"eventCategory"`
}
