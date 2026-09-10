# Changelog

## 1.6.1-0.5.0 (upcoming)

## 1.6.1-0.4.0 (2025-10-07)

* [PLT-4748] Bump `golang.org/x/net` (0.43.0 → 0.58.0), `golang.org/x/crypto` (0.42.0 → 0.56.0), `google.golang.org/grpc` (1.67.3 → 1.83.1), `go.opentelemetry.io/otel*` (1.39.0 → 1.44.0) and Go toolchain (1.24.6 → 1.26.0) to close vulnerabilities
* Fix `Jenkinsfile`'s `@Library('libpipelines@master')` pointing at a branch that no longer exists in `Stratio/jenkins-bootstrap` (renamed to `main`), breaking CI on every PR regardless of code changes
* [PLT-2635] Fix golang vulnerabilities to max provider version 1.24.6

## Previous development

### Branched to branch-1.6.1-0.4 (2025-07-22)

 * [PLT-1548] -  [GKE] Activar Workload Identity  - [`#44`](https://github.com/Stratio/cluster-api-provider-gcp/pull/44)
 * [PLT-1330] -  CMEK, SA & CIDRs  - [`#37`](https://github.com/Stratio/cluster-api-provider-gcp/pull/37)

### Branched to branch-1.6.1-0.3 (2024-12-09)

* [PLT-1330] CMEK - Service accounts & Secondary CIDR ranges adaption to R4.7



## 1.6.1-0.2.1 (2024-12-05)

* [PLT-1313] Support Secondary CIDR ranges
* [PLT-1246] CMEK Support

## 1.6.1-0.2.0 (2024-10-11)

* [PLT-965] Disable managed Monitoring and Logging
* [PLT-806] Add GKE Private cluster support
* [PLT-563] Fix autoscaling issues
* [PLT-326] First approach to manage taints addition, update and deletion on GKE
* [PLT-327] After creating a GKE cluster, it takes ~20 minutes for its status to be READY
* [PLT-911] Disable external endpoint

### Branched to branch-1.6.1-0.1 (2024-08-22)

* Add cluster api provider for GCP
