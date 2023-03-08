# KSOC AWS Models

## CloudTrail Go Models

This repository contains Go models for CloudTrail events. Specifically, this code defines a Go struct for the CloudTrail event format as stored in Amazon S3.

### Files

* `cloudtrail.go`: This file contains the `Records` and `Record` structs. The `Records` struct represents a list of `Record` objects, where each `Record` object represents a single CloudTrail event as stored in S3.
* `README.md`: This file you are reading now provides an overview of this repository.

### Records struct

The `Records` struct contains a list of `Record` objects. Each `Record` object represents a single CloudTrail event as stored in S3. The struct is defined as follows:

```go
type Records struct {
    Records []Record `json:"Records"`
}
```

### Record struct

The `Record` struct represents a single CloudTrail event as stored in S3. It contains a variety of fields that correspond to the CloudTrail event schema. The struct is defined as follows:

```go
type Record struct {
    EventVersion       string      `json:"eventVersion"`
    UserIdentity       struct {...} `json:"userIdentity"`
    EventTime          time.Time   `json:"eventTime"`
    EventSource        string      `json:"eventSource"`
    EventName          string      `json:"eventName"`
    AwsRegion          string      `json:"awsRegion"`
    SourceIPAddress    string      `json:"sourceIPAddress"`
    UserAgent          string      `json:"userAgent"`
    RequestParameters  struct {...} `json:"requestParameters"`
    ResponseElements   struct {...} `json:"responseElements"`
    RequestID          string      `json:"requestID"`
    EventID            string      `json:"eventID"`
    ReadOnly           bool        `json:"readOnly"`
    EventType          string      `json:"eventType"`
    ManagementEvent    bool        `json:"managementEvent"`
    RecipientAccountId string      `json:"recipientAccountId"`
    EventCategory      string      `json:"eventCategory"`
}
```

### Conclusion

These Go structs can be used to deserialize CloudTrail events as stored in Amazon S3, which can be useful for analyzing CloudTrail logs or for building integrations with AWS services that consume CloudTrail logs.