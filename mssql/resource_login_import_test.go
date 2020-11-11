package mssql

import (
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
  "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
  "testing"
)

func TestAccLogin_importBasic(t *testing.T) {
  resource.Test(t, resource.TestCase{
    PreCheck:          func() { testAccPreCheck(t) },
    ProviderFactories: testAccProviders,
    CheckDestroy:      func(state *terraform.State) error { return testAccCheckLoginDestroy(t, state) },
    Steps: []resource.TestStep{
      {
        Config: testAccCheckLogin(t, "test_import", map[string]string{"login_name": "login_import", "password": "valueIsH8kd$¡"}),
        Check: resource.ComposeTestCheckFunc(
          testAccCheckLoginExists("mssql_login.test_import"),
        ),
      },
      {
        ResourceName:            "mssql_login.test_import",
        ImportState:             true,
        ImportStateVerify:       true,
        ImportStateVerifyIgnore: []string{"password"},
      },
    },
  })
}
