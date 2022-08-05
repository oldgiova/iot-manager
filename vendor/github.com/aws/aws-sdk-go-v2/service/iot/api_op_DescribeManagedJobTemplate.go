// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// View details of a managed job template.
func (c *Client) DescribeManagedJobTemplate(ctx context.Context, params *DescribeManagedJobTemplateInput, optFns ...func(*Options)) (*DescribeManagedJobTemplateOutput, error) {
	if params == nil {
		params = &DescribeManagedJobTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeManagedJobTemplate", params, optFns, c.addOperationDescribeManagedJobTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeManagedJobTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeManagedJobTemplateInput struct {

	// The unique name of a managed job template, which is required.
	//
	// This member is required.
	TemplateName *string

	// An optional parameter to specify version of a managed template. If not
	// specified, the pre-defined default version is returned.
	TemplateVersion *string

	noSmithyDocumentSerde
}

type DescribeManagedJobTemplateOutput struct {

	// The unique description of a managed template.
	Description *string

	// The document schema for a managed job template.
	Document *string

	// A map of key-value pairs that you can use as guidance to specify the inputs for
	// creating a job from a managed template. documentParameters can only be used when
	// creating jobs from Amazon Web Services managed templates. This parameter can't
	// be used with custom job templates or to create jobs from them.
	DocumentParameters []types.DocumentParameter

	// A list of environments that are supported with the managed job template.
	Environments []string

	// The unique Amazon Resource Name (ARN) of the managed template.
	TemplateArn *string

	// The unique name of a managed template, such as AWS-Reboot.
	TemplateName *string

	// The version for a managed template.
	TemplateVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeManagedJobTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeManagedJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeManagedJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeManagedJobTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeManagedJobTemplate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeManagedJobTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeManagedJobTemplate",
	}
}
