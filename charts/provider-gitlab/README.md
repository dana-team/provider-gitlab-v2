# provider-gitlab

![Version: 0.0.0](https://img.shields.io/badge/Version-0.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: latest](https://img.shields.io/badge/AppVersion-latest-informational?style=flat-square)

A Helm chart for Crossplane provider-gitlab

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| deploymentRuntimeConfig | object | `{"container":{"args":["--debug"],"name":"package-runtime"},"name":"gitlab-config"}` | Configuration to be added to the provider deployment via the DeploymentRuntimeConfig resource |
| fullnameOverride | string | `""` |  |
| image.repository | string | `"ghcr.io/dana-team/provider-gitlab"` | The repository of the provider container image. |
| image.tag | string | `""` | The tag of the manager container image. |
| nameOverride | string | `""` |  |
| provider.name | string | `"provider-gitlab"` | Name of the provider |
| provider.runtimeConfigRef.name | string | `"gitlab-config"` | Name of the DeploymentRuntimeConfig object to use |
| providerConfig | object | `{"credentials":{"secretRef":{"key":"credentials","name":"gitlab-creds","namespace":"crossplane-system"},"source":"Secret"},"name":"gitlab-default"}` | Provider authentication configuration |
| secret | object | `{"name":"gitlab-creds","token":"gitlab-token","type":"Opaque"}` | Secret values for the provider authentication. |
| secret.name | string | `"gitlab-creds"` | Name of the secret. |
| secret.token | string | `"gitlab-token"` | The gitlab access token to authenticate with. |
| secret.type | string | `"Opaque"` | Type of the secret. |

