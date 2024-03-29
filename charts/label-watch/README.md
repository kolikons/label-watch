# label-watch

![Version: 0.0.2](https://img.shields.io/badge/Version-0.0.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.2](https://img.shields.io/badge/AppVersion-0.0.2-informational?style=flat-square)

label-watch checks a specific label on worker node then create an label

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| kolikons |  | https://github.com/kolikons/label-watch |

## Source Code

* <https://github.com/kolikons/label-watch>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Anti-affinity to disallow deploying client and master nodes on the same worker node |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"kolikons/label-watch"` |  |
| image.tag | string | `""` | Overrides the image tag whose default is the chart appVersion. |
| imagePullSecrets | list | `[]` |  |
| label_watch.interval | string | `"30m"` | Supports format: 's', 'm', 'h' |
| label_watch.labels | string | `"node-type"` | Label that's checking on worker nodes then set label in format node-role.kubernetes.io/VALUE_FROM_LABEL=true. Supports multiple labels via coma separator. Example:  node-type,type,etc   |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` | Node labels for pod assignment ref: https://kubernetes.io/docs/user-guide/node-selection/ |
| podAnnotations | object | `{}` | Key/value pairs that are attached to pods. |
| podLabels | object | `{}` | Key/value pairs that are attached to pods. |
| podSecurityContext | object | `{}` |  |
| rbac | object | `{"create":true}` | Create Cluster Role to allow modify nodes |
| replicaCount | int | `1` | count of POD |
| resources | object | `{}` | We usually recommend not to specify default resources and to leave this as a conscious choice for the user. This also increases chances charts run on environments with little resources, such as Minikube. If you do want to specify resources, uncomment the following lines, adjust them as necessary, and remove the curly braces after 'resources:'. |
| securityContext | object | `{}` |  |
| serviceAccount | object | `{"annotations":{},"create":true,"name":""}` | Specifies whether a service account should be created |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.name | string | `""` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |
| tolerations | list | `[]` | Tolerations for pod assignment ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.7.0](https://github.com/norwoodj/helm-docs/releases/v1.7.0)
