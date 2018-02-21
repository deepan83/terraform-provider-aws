package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSOrganization_basic(t *testing.T) {
	var organization organizations.Organization

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSOrganizationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSOrganizationConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSOrganizationExists("aws_organization.test", &organization),
					resource.TestCheckResourceAttr("aws_organization.test", "feature_set", organizations.OrganizationFeatureSetAll),
					resource.TestCheckResourceAttrSet("aws_organization.test", "arn"),
					resource.TestCheckResourceAttrSet("aws_organization.test", "master_account_arn"),
					resource.TestCheckResourceAttrSet("aws_organization.test", "master_account_email"),
					resource.TestCheckResourceAttrSet("aws_organization.test", "feature_set"),
				),
			},
		},
	})
}

func TestAccAWSOrganization_consolidatedBilling(t *testing.T) {
	var organization organizations.Organization

	feature_set := organizations.OrganizationFeatureSetConsolidatedBilling

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSOrganizationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSOrganizationConfigConsolidatedBilling(feature_set),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSOrganizationExists("aws_organization.test", &organization),
					resource.TestCheckResourceAttr("aws_organization.test", "feature_set", feature_set),
				),
			},
		},
	})
}

func testAccCheckAWSOrganizationDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).organizationsconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_organization" {
			continue
		}

		params := &organizations.DescribeOrganizationInput{}

		resp, err := conn.DescribeOrganization(params)

		if err != nil {
			if isAWSErr(err, organizations.ErrCodeAWSOrganizationsNotInUseException, "") {
				return nil
			}
			return err
		}

		if resp != nil && resp.Organization != nil {
			return fmt.Errorf("Bad: Organization still exists: %q", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckAWSOrganizationExists(n string, a *organizations.Organization) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Organization ID not set")
		}

		conn := testAccProvider.Meta().(*AWSClient).organizationsconn
		params := &organizations.DescribeOrganizationInput{}

		resp, err := conn.DescribeOrganization(params)

		if err != nil {
			return err
		}

		if resp == nil || resp.Organization == nil {
			return fmt.Errorf("Organization %q does not exist", rs.Primary.ID)
		}

		a = resp.Organization

		return nil
	}
}

const testAccAWSOrganizationConfig = "resource \"aws_organization\" \"test\" {}"

func testAccAWSOrganizationConfigConsolidatedBilling(feature_set string) string {
	return fmt.Sprintf(`
resource "aws_organization" "test" {
  feature_set = "%s"
}
`, feature_set)
}
