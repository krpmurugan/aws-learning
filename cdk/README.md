# Welcome to your CDK TypeScript project!

You should explore the contents of this project. It demonstrates a CDK app with an instance of a stack (`CdkWorkshopStack`)
which contains an Amazon SQS queue that is subscribed to an Amazon SNS topic.

The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

- `npm run build` compile typescript to js
- `npm run watch` watch for changes and compile
- `npm run test` perform the jest unit tests
- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template

## Useful commands for my reference

- In order to work cdk MFA need to do the following steps

`aws sts assume-role --role-arn arn:aws:iam::xxxxxxx:role/DevAccountAccessRole --role-session-name "SharingSandbox" --profile xxxxxxx`

If run the above command will return the credentials details need to update the AccessKeyId, SecretAccessKey and SessionToken those in your aws credentials SharingSandbox section and trun the below commands will work or need to enable aws-mfa.

- `cdk synth --profile sharingsandbox`
- `cdk diff --profile sharingsandbox`
- `cdk deploy --profile sharingsandbox`
