package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/awstesting/mock"
	"github.com/aws/aws-sdk-go/service/sts"
)

// AWSSessionClient is a wrapper for AWS session exported funcs
type AWSSessionClient struct {
	Must       func(sess *session.Session, err error) *session.Session
	NewSession func() (*session.Session, error)
}

// SessionClient returns a default AWSSessionClient
var SessionClient = &AWSSessionClient{
	session.Must,
	func() (*session.Session, error) {
		sess, err := session.NewSession()
		if err != nil {
			return sess, err
		}

		// determine if we are authorized to access AWS with the credentials provided.
		_, err = sts.New(sess).GetCallerIdentity(&sts.GetCallerIdentityInput{})
		if err != nil {
			return nil, err
		}

		return sess, err
	},
}

// MockAWSSessionClient returns a mock AWSSessionClient
var MockAWSSessionClient = &AWSSessionClient{
	Must: func(sess *session.Session, err error) *session.Session {
		return sess
	},
	NewSession: func() (*session.Session, error) {
		// mock session with fake credentials
		mockSession := mock.Session
		mockSession.Config.Region = aws.String("us-east-1")
		mockSession.Config.Credentials = credentials.NewStaticCredentials("123", "xyz", "")

		return mockSession, nil
	},
}
