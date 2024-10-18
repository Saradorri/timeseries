# InfluxDB User and Token Creation Guide
This document provides a step-by-step guide on how to create a user and generate a token in InfluxDB.
## Prerequisites
- Ensure you have InfluxDB installed and running.
- Access to the InfluxDB HTTP API.

## Step 1: Access the InfluxDB Shell
Open your terminal and access the InfluxDB shell with the following command:
```bash
influx
```
## Step 2: Create an Organization
First, create an organization if you haven't already. Replace `<org_name>` with your desired organization name:
```bash
influx org create --name <org_name>
```
## Step 3: Create a Bucket
Next, create a bucket within your organization. Replace `<bucket_name>` with your desired bucket name:
```bash
influx bucket create --name <bucket_name> --org <org_name>
```
## Step 4: Create a User
Create a user and specify the desired username and password. Replace `<username>` and `<password>` with your desired credentials:
```bash
influx user create --name <username> --password <password> --org <org_name>
```
## Step 5: Generate an API Token
You can generate an API token for your user using the following command. Replace `<username>` with the user you just created:
```bash
influx auth create --user <username> --read-bucket <bucket_name> --write-bucket <bucket_name>
```
## Step 6: Note the Token
The command will output the generated token. Make sure to save it securely as you will need it to authenticate your API requests.

## Example of Using the Token
When making API requests, include the token in the `Authorization` header. Here's an example using `curl`:
```bash
curl -G http://localhost:8086/api/v2/buckets --header "Authorization: Token <your_token>"
```
Replace `<your_token>` with the token you generated.
## Conclusion
You have successfully created a user and generated a token in InfluxDB. You can now use this token to authenticate API requests and interact with your InfluxDB instance.
