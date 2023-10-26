output "new_apikey" {
  value     = capella_apikey.new_apikey
  sensitive = true
}

resource "capella_apikey" "new_apikey" {
  organization_id    = var.organization_id
  name               = var.apikey.name
  organization_roles = var.apikey.organization_roles
  allowed_cidrs      = var.apikey.allowed_cidrs
  resources = [
    {
      id    = var.project_id
      roles = var.resource.roles
      type  = var.resource.type
    }
  ]
}
