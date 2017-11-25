package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	/*
		{
		InstanceId: "instanceId",
		InstanceType: {
			Value: "t2.small"
			}
		}
	*/
	svc := ec2.New(session.New())
	input := &ec2.DescribeInstanceAttributeInput{
		Attribute:  aws.String("instanceType"),
		InstanceId: aws.String(""), // instanceId
	}
	result, err := svc.DescribeInstanceAttribute(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
