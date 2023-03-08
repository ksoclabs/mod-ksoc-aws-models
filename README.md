# KSOC AWS Models

This repository contains some models which are missing from official AWS Go SDK.

## Codd Tour

package cloudtrail

```go
package cloudtrail
```
This declares the package name as cloudtrail.

```go
import "time"
```
The time package provides types for representing time in Go.

```go
// Records is a representation of CloudTrail record as seen in S3 stored objects.
// It contains a list of Record objects.
type Records struct {
	Records []Record `json:"Records"`
}
```

This declares a struct named `Records` which represents CloudTrail `records` as seen in S3 stored objects. It contains a field named `Records` which is a list of `Record` objects.

The `json:"Records"` tag indicates that this field should be marshalled to JSON as `Records`.

```go
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
```
This code is a Go package named "cloudtrail" that provides a data model for representing Amazon Web Services (AWS) CloudTrail records.

The package defines two structs, "Records" and "Record". The "Records" struct contains a slice of "Record" objects and represents CloudTrail records as seen in S3 stored objects. The "Record" struct is a single CloudTrail event as stored in S3.

The "Record" struct contains the following fields:

* "EventVersion": A string representing the event version.
* "UserIdentity": A struct containing information about the user who initiated the event.
* "EventTime": A time.Time object representing the time the event occurred.
* "EventSource": A string representing the AWS service that the event pertains to.
* "EventName": A string representing the name of the event that occurred.
* "AwsRegion": A string representing the AWS region where the event occurred.
* "SourceIPAddress": A string representing the IP address of the requester that caused the event.
* "UserAgent": A string representing the user agent of the requester that caused the event.
* "RequestParameters": A struct containing parameters associated with the request that caused the event.
* "ResponseElements": A struct containing elements of the response that was returned for the request that caused the event.
* "RequestID": A string representing the ID of the request that caused the event.
* "EventID": A string representing the ID of the event.
* "ReadOnly": A boolean indicating whether the event is a read-only event.
* "EventType": A string representing the type of the event.
* "ManagementEvent": A boolean indicating whether the event is a management event.
* "RecipientAccountId": A string representing the ID of the account that received the event.
* "EventCategory": A string representing the category of the event.

The `"UserIdentity"` field is itself a struct containing the following fields:

* "Type": A string representing the type of identity associated with the event.
* "PrincipalId": A string representing the principal ID associated with the event.
* "Arn": A string representing the Amazon Resource Name (ARN) associated with the identity.
* "AccountId": A string representing the AWS account ID associated with the identity.
* "AccessKeyId": A string representing the access key ID associated with the identity.
* "SessionContext": A struct containing information about the session context associated with the identity.

The `"SessionContext"` field is itself a struct containing the following fields:

* "SessionIssuer": A struct containing information about the entity that created the session.
* "WebIdFederationData": A struct containing information about the web identity federation data associated with the session.
* "Attributes": A struct containing session attributes associated with the session.

The `"SessionIssuer"` field is itself a struct containing the following fields:

* "Type": A string representing the type of entity that created the session.
* "PrincipalId": A string representing the principal ID associated with the entity.
* "Arn": A string representing the ARN associated with the entity.
* "AccountId": A string representing the AWS account ID associated with the entity.
* "UserName": A string representing the name of the user associated with the entity.

The "Attributes" field is itself a struct containing the following fields:

* "CreationDate": A time.Time object representing the date the session was created.
* "MfaAuthenticated": A string representing whether multi-factor authentication (MFA) was used to authenticate the session.

Overall, this package provides a useful data model for representing CloudTrail records and can be used by Go developers who need to work with these records in their applications.


## QuickStart

To use this package in your Go project, you can import it by including the following line in your code:

```go
import "github.com/ksoclabs/mod-ksoc-aws-models/cloudtrail"
```

Once you have imported the package, you can create a new Records object by unmarshaling JSON data from an S3 object into it. Here's an example of how to do this:

```go
import (
    "encoding/json"
    "github.com/ksoclabs/mod-ksoc-aws-models/cloudtrail"
)

func main() {
    jsonString := `{
        "Records": [
            {
                "eventVersion": "1.08",
                "userIdentity": {
                    "type": "IAMUser",
                    "principalId": "EXAMPLE",
                    "arn": "arn:aws:iam::123456789012:user/JaneDoe",
                    "accountId": "123456789012",
                    "accessKeyId": "EXAMPLE",
                    "sessionContext": {
                        "sessionIssuer": {
                            "type": "Role",
                            "principalId": "EXAMPLE",
                            "arn": "arn:aws:iam::123456789012:role/RoleA",
                            "accountId": "123456789012",
                            "userName": "RoleA"
                        },
                        "webIdFederationData": {},
                        "attributes": {
                            "creationDate": "2022-03-08T18:50:25Z",
                            "mfaAuthenticated": "false"
                        }
                    }
                },
                "eventTime": "2022-03-08T18:52:40Z",
                "eventSource": "s3.amazonaws.com",
                "eventName": "PutObject",
                "awsRegion": "us-west-2",
                "sourceIPAddress": "192.168.0.2.0",
                "userAgent": "aws-cli/2.5.56 Python/3.8.8 Linux/5.4.0-91-generic exe/x86_64.ubuntu.20",
                "requestParameters": {
                    "bucketName": "example-bucket",
                    "key": "example-object"
                },
                "responseElements": {
                    "x-amz-version-id": "1"
                },
                "requestID": "EXAMPLE",
                "eventID": "EXAMPLE",
                "readOnly": false,
                "eventType": "AwsApiCall",
                "managementEvent": false,
                "recipientAccountId": "123456789012",
                "eventCategory": "Data"
            }
        ]
    }`

    var records cloudtrail.Records
    err := json.Unmarshal([]byte(jsonString), &records)
    if err != nil {
        panic(err)
    }

    // Do something with the records...
}
```

In this example, the JSON data is stored in a string variable called `jsonString`. The `json.Unmarshal` function is used to unmarshal the JSON data into a new `Records` object called `records`. Once the data is unmarshaled, you can do something with the `records` object, such as iterate through the `Records` slice and perform some action for each `Record` object.

That's it! You should now be able to use this package in your Go project to work with CloudTrail records.
