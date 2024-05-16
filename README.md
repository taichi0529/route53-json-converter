# Route53 Json Converter

A JSON conversion tool for migrating hosted zones in AWS Route 53 to another account.

## Build Instructions
Run the following command in an environment where you can build Go:

```bash
go build convert.go
```

## Hosted Zone Migration Steps

1. Export the records of the existing hosted zone as JSON.
2. Convert the exported JSON into an importable format.
3. Create a hosted zone in the new account.
4. Import the JSON converted in step 2.

While steps 1, 3, and 4 can be done using AWS CLI or the console, step 2 requires this tool for the conversion. Although the conversion itself is not a complex task, it can be cumbersome if there are many records, so this program was created to automate the process.

For detailed instructions on the conversion, refer to the following:
https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/hosted-zones-migrating.html#hosted-zones-migrating-edit-records

```bash
aws route53 list-resource-record-sets --hosted-zone-id ZXXXXXXXXXXXXXXXXXXX > xxxxxxxx.json
./convert xxxxxxxx.json > xxxxxxxx-converted.json
```

After this, create a hosted zone in the new account and import the JSON converted in step 2.

```bash
aws route53 change-resource-record-sets --hosted-zone-id ZYYYYYYYYYYYYYYYYYYYY --change-batch file://xxxxxxxx-converted.json
```