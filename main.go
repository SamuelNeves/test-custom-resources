package main

import (
	"github.com/SamuelNeves/test-custom-resources/components"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		_, err := components.NewStandardS3Bucket(ctx, "test-resources")
		if err != nil {
			return err
		}

		//println(" %s", pulumi.All(re))
		// Export the name of the bucket
		return nil
	})
}
