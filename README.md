

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
