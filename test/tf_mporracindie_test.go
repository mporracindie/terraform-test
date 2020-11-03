package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestTerraformAwsHelloWorldExample(t *testing.T) {
	t.Parallel()

	s3_bucket := "mporracindie-tf-gotest"
	file_name := []string{"gotest1.txt", "gotest2.txt"}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// VarFiles:     []string{"vars.tfvars"},
		Vars: map[string]interface{}{"s3_bucket": s3_bucket, "s3_files": file_name},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created.
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the bucket and files info
	bucket_info := terraform.Output(t, terraformOptions, "bucket_info")
	files_info := terraform.Output(t, terraformOptions, "files_info")

	// Assert bucket name
	assert.Equal(t, s3_bucket, bucket_info)

	// Assert each file name
	for _, file := range file_name {
		assert.Contains(t, files_info, file)
	}
}
