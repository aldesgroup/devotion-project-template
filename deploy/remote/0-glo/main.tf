# --------------------------------------------------------------------------- #
# --- Main Terraform config to the global resources, like the app reg
# --------------------------------------------------------------------------- #

# --------------------------------------------------------------------------- #
# --- App Registration = definition of the app
# - Defines a name
# - Scopes (access_as_user)
# - The audience (api://...)
# - Expose permissions
# --------------------------------------------------------------------------- #

resource "azuread_application" "app_???_devotionprojecttemplate" {
  display_name = "app-???-devotionprojecttemplate"

  # IMPORTANT : URI lisible mais conforme à la policy du tenant
  identifier_uris = ["api://???/app-???-devotionprojecttemplate"]

  api {
    requested_access_token_version = 2

    oauth2_permission_scope {
      admin_consent_description  = "Allow access to DevotionProjectTemplate API"
      admin_consent_display_name = "Access DevotionProjectTemplate API"
      enabled                    = true
      id                         = "???"
      type                       = "User"
      value                      = "access_as_user"
    }
  }
}

# --------------------------------------------------------------------------- #
# --- Service principal = operationnal instance of the app
# - Usable identity
# - Visible in Enterprise Apps
# - Can receive RBAC roles
# --------------------------------------------------------------------------- #

resource "azuread_service_principal" "sp_???_devotionprojecttemplate" {
  client_id = azuread_application.app_???_devotionprojecttemplate.client_id
}
