// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the set of dimensions that are defined for your Amazon Web Services
// accounts. Requires permission to access the ListDimensions
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListDimensions(ctx context.Context, params *ListDimensionsInput, optFns ...func(*Options)) (*ListDimensionsOutput, error) {
	if params == nil {
		params = &ListDimensionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDimensions", params, optFns, c.addOperationListDimensionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDimensionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDimensionsInput struct {

	// The maximum number of results to retrieve at one time.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDimensionsOutput struct {

	// A list of the names of the defined dimensions. Use DescribeDimension to get
	// details for a dimension.
	DimensionNames []string

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDimensionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDimensions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDimensions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDimensions(options.Region), middleware.Before); err != nil {
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

// ListDimensionsAPIClient is a client that implements the ListDimensions
// operation.
type ListDimensionsAPIClient interface {
	ListDimensions(context.Context, *ListDimensionsInput, ...func(*Options)) (*ListDimensionsOutput, error)
}

var _ ListDimensionsAPIClient = (*Client)(nil)

// ListDimensionsPaginatorOptions is the paginator options for ListDimensions
type ListDimensionsPaginatorOptions struct {
	// The maximum number of results to retrieve at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDimensionsPaginator is a paginator for ListDimensions
type ListDimensionsPaginator struct {
	options   ListDimensionsPaginatorOptions
	client    ListDimensionsAPIClient
	params    *ListDimensionsInput
	nextToken *string
	firstPage bool
}

// NewListDimensionsPaginator returns a new ListDimensionsPaginator
func NewListDimensionsPaginator(client ListDimensionsAPIClient, params *ListDimensionsInput, optFns ...func(*ListDimensionsPaginatorOptions)) *ListDimensionsPaginator {
	if params == nil {
		params = &ListDimensionsInput{}
	}

	options := ListDimensionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDimensionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDimensionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDimensions page.
func (p *ListDimensionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDimensionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListDimensions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListDimensions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListDimensions",
	}
}
