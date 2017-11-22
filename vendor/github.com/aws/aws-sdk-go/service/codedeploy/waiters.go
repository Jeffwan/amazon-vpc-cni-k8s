// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilDeploymentSuccessful uses the CodeDeploy API operation
// GetDeployment to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CodeDeploy) WaitUntilDeploymentSuccessful(input *GetDeploymentInput) error {
	return c.WaitUntilDeploymentSuccessfulWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDeploymentSuccessfulWithContext is an extended version of WaitUntilDeploymentSuccessful.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CodeDeploy) WaitUntilDeploymentSuccessfulWithContext(ctx aws.Context, input *GetDeploymentInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilDeploymentSuccessful",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "deploymentInfo.status",
				Expected: "Succeeded",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "deploymentInfo.status",
				Expected: "Failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "deploymentInfo.status",
				Expected: "Stopped",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetDeploymentInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetDeploymentRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
