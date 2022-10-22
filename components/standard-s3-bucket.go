package components

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StandardS3Bucket struct {
	pulumi.ResourceState
}
type StandardBucketArgs struct {
	Name                      string
	IntelligentTieringEnabled bool
	BlockPublicAccessEnabled  bool
	Region                    string
}

func NewStandardS3Bucket(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*StandardS3Bucket, error) {
	myComponent := &StandardS3Bucket{}
	err := ctx.RegisterComponentResource("pkg:index:StandardS3Bucket", name, myComponent, opts...)
	if err != nil {
		return nil, err
	}

	bucket, err := s3.NewBucket(ctx, fmt.Sprintf("%s-bucket", name),
		&s3.BucketArgs{ /*...*/ }, pulumi.Parent(myComponent))
	ctx.RegisterResourceOutputs(myComponent, pulumi.Map{
		"bucketDnsName": bucket.BucketDomainName,
		"bucketName":    bucket.ID(),
	})
	return myComponent, nil
}
