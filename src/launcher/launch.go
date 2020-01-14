package launcher

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Configuration struct {
	Region string
}

type ServerSizes map[string]string

// LaunchServer creates an EC2 server instance relying on your
// shared configuration relying on your shared configuration
// (the one can be found in your .aws/credentials).
//
// This should be done in a more sophisticated way because currently the AMI
// for Amazon Linux 2 is hardcoded along with all the other parameters.
func LaunchServer() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2Svc := ec2.New(sess)

	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}

	runResult, err := ec2Svc.RunInstances(&ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the shared configuration region
		ImageId:      aws.String("ami-07cda0db070313c52"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("value: %v, Type: %T \n", runResult, runResult)

}
