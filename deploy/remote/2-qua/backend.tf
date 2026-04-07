# --------------------------------------------------------------------------- #
# -- Terraform state storage fdor the resources described in this folder
# --------------------------------------------------------------------------- #

terraform {
  backend "azurerm" {
    resource_group_name  = "rg-???-tfstates"
    storage_account_name = "st???tfstates"
    container_name       = "sc-???-tfstates"
    key                  = "tfstates.devotionprojecttemplate.qua"
  }
}
