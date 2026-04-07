# --------------------------------------------------------------------------- #
# --- Main Terraform config to create the whole infra for the {env} environment
# --------------------------------------------------------------------------- #
locals {
  config = {
    # The current env: dev, qua, or prd, for instance
    env = "dev"
    # ex: northeurope
    location = "???"
    # The name of the subscription hosting identity-related resources
    identity_sub_name = "???"
    # The name of the subscription hosting resources related to tools & supervision
    management_sub_name = "???"
    # The name of the subscription hosting all the resources needed for this particular environment
    environment_sub_name = "???"
    # The name of the container registry we use
    acr_name = "???"
    # The resource group the container registry belongs to
    acr_rg = "???"
    # The domain name, something like: companyname.com
    domain_name = "???"
    # The port the API will listen to
    port = "{port}"
  }
}

module "infra_dev" {
  source = "../"
  config = local.config
}
