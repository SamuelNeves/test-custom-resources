package helper

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type taggingHelper struct {
}

type TaggingHelper interface {
	GenerateMandatoryTags(*pulumi.Context) (*pulumi.StringMap, error)
}

func (a taggingHelper) GenerateMandatoryTags(pulumiCtx *pulumi.Context) (*pulumi.StringMap, error) {

	tags := pulumi.StringMap{
		"user:Service":    pulumi.String("MyService"),
		"user:Project":    pulumi.String(pulumiCtx.Project()),
		"user:CostCenter": pulumi.String(config.Require(pulumiCtx, "costCenter")),
	}
	return &tags, nil
}

func NewTaggingHelper() TaggingHelper {
	return &taggingHelper{}
}
