#### Step 1 : Create an Bucket
#### Step 2 : Create an User
#### Step 3 : Attach a Policy - AmazonS3FullAccess
#### Step 4 : Create an Access key

### Setup Notification When upload file into s3 Bucket 
A.  1. Add Policy to SNS 
    2. Add Policy to SQS 
B.  Bucket[Properties] --> 
    Create Event Notification
        1.Give event name 
        2.Select Event Type
        3.Destination : SNS Topic --> Choose SNS Topic

C. Create an Subscription inside SNS To Send notification to SQS
    Create Subscription --> Topic ARN -->  Select protocol--> Amazon SQS --> Create subscription
Same for HTTP --> Select protocol as HTTP/S

D. Policy need to be updated : Check policy-example.json


Webhook :

Make Public URL for your local app
https://dashboard.ngrok.com/get-started/setup/macos

