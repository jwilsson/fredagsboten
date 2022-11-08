# fredagsboten

Fredagsboten is a Slack bot to automatically let your colleagues know it's Friday by posting an randomly selected image to a channel of your choice.

## Prerequisites
* A Slack workspace.
* An AWS account.

## Setup
1. Start by creating a [Slack app and setting up Incoming Webhooks](https://slack.com/intl/en-se/help/articles/115005265063-incoming-webhooks-for-slack).
2. Configure your [AWS account to work with Serverless](https://serverless.com/framework/docs/providers/aws/guide/credentials/).
3. Configure a `DYNAMO_TABLE_NAME` environmental variable with the name of your DynamoDB table.
4. Configure a `SLACK_WEBHOOK_URL` environmental variable with your Slack Webhook URL.
5. After deploying the first time, add some images to your DynamoDB table, see the item structure below.
6. Change any other values in `serverless.yml` to fit your needs.
7. Profit!

### Example image item
```json
{
    "image_url": "https://grodanboll.azurewebsites.net/img/fredagsgrodan/fredagsgrodan512.png",
}
```

* `image_url` - points to the image file to use
