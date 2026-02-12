package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StackProps struct {
	awscdk.StackProps
}

func NewFredagsbotenStack(scope constructs.Construct, id string, props *StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	table := awsdynamodb.NewTable(stack, jsii.String(os.Getenv("DYNAMO_TABLE_NAME")), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("image_url"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		ReadCapacity:  jsii.Number(1),
		WriteCapacity: jsii.Number(1),
	})

	function := awslambda.NewFunction(stack, jsii.String("FredagsbotenFunction"), &awslambda.FunctionProps{
		Architecture: awslambda.Architecture_ARM_64(),
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./app/build/"), nil),
		Environment: &map[string]*string{
			"DYNAMO_TABLE_NAME": jsii.String(*table.TableName()),
			"RUN_AT_TIME":       jsii.String("08:00"),
			"SLACK_WEBHOOK_URL": jsii.String(os.Getenv("SLACK_WEBHOOK_URL")),
		},
	})

	table.Grant(function, jsii.String("dynamodb:Scan"))

	awsevents.NewRule(stack, jsii.String("FredagsbotenScheduleRule"), &awsevents.RuleProps{
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			Minute:  jsii.String("0"),
			Hour:    jsii.String("*"),
			WeekDay: jsii.String("FRI"),
			Month:   jsii.String("*"),
			Year:    jsii.String("*"),
		}),
		Targets: &[]awsevents.IRuleTarget{
			awseventstargets.NewLambdaFunction(function, nil),
		},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewFredagsbotenStack(app, "FredagsbotenStack", &StackProps{})

	app.Synth(nil)
}
