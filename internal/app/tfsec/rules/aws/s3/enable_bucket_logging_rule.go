package s3

import (
	"fmt"

	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/severity"

	"github.com/aquasecurity/tfsec/pkg/provider"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/hclcontext"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"

	"github.com/aquasecurity/tfsec/pkg/rule"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
)


func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID:   "AWS002",
		Service:   "s3",
		ShortCode: "enable-bucket-logging",
		Documentation: rule.RuleDocumentation{
			Summary:      "S3 Bucket does not have logging enabled.",
			Explanation:  `
Buckets should have logging enabled so that access can be audited. 
`,
			Impact:       "There is no way to determine the access to this bucket",
			Resolution:   "Add a logging block to the resource to enable access logging",
			BadExample:   `
resource "aws_s3_bucket" "bad_example" {

}
`,
			GoodExample:  `
resource "aws_s3_bucket" "good_example" {
	logging {
		target_bucket = "target-bucket"
	}
}
`,
			Links: []string{
				"https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html",
				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket",
			},
		},
		Provider:        provider.AWSProvider,
		RequiredTypes:   []string{"resource"},
		RequiredLabels:  []string{"aws_s3_bucket"},
		DefaultSeverity: severity.Medium,
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context) {
			if loggingBlock := resourceBlock.GetBlock("logging"); loggingBlock == nil {
				if resourceBlock.GetAttribute("acl") != nil && resourceBlock.GetAttribute("acl").Equals("log-delivery-write") {
					return
				}
				set.Add(
					result.New(resourceBlock).
						WithDescription(fmt.Sprintf("Resource '%s' does not have logging enabled.", resourceBlock.FullName())).
						WithRange(resourceBlock.Range()),
				)
			}
		},
	})
}