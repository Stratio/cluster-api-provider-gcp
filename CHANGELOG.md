# Changelog

## 1.13.1-0.1.0 (upcoming)

* [PLT-4891] Rebase the fork onto upstream `v1.13.1` (reads the CAPI core `Cluster`/`MachinePool` as `v1beta2`, requires CAPI core >= v1.11). Stratio behaviours re-applied: GKE NetworkPolicy (Calico), boot disk CMEK, `clusterIpv4Cidr` and `ipAllocationPolicy` (incl. secondary range names), Managed Prometheus and per-component logging, Workload Identity on update, node pool resize with autoscaling enabled, `CLUSTER_ALREADY_HAS_OPERATION` handling, replica write-back to the `MachinePool`. Private cluster, `extended` release channel, Workload Identity on create and the taints/authorized-networks diff fixes now come from upstream

## 1.6.1-0.4.2 (upcoming)

* [PLT-4830] Bump `google.golang.org/grpc` (1.83.1 → 1.83.2) to close a vulnerability; the rebuild also picks up `libssl3`/`libcrypto3` 3.5.8 from the `alpine:latest` runtime base

## 1.6.1-0.4.1 (2026-09-10)

* [PLT-4748] Bump `golang.org/x/net` (0.43.0 → 0.58.0), `golang.org/x/crypto` (0.42.0 → 0.56.0), `google.golang.org/grpc` (1.67.3 → 1.83.1), `go.opentelemetry.io/otel*` (1.39.0 → 1.44.0) and Go toolchain (1.24.6 → 1.26.0) to close vulnerabilities
* Fix `Jenkinsfile`'s `@Library('libpipelines@master')` pointing at a branch that no longer exists in `Stratio/jenkins-bootstrap` (renamed to `main`), breaking CI on every PR regardless of code changes

## 1.6.1-0.4.0 (2025-10-07)

* [PLT-2635] Fix golang vulnerabilities to max provider version 1.24.6
* [PLT-1548] [GKE] Activar Workload Identity

## 1.6.1-0.3.1 (2025-02-26)

* [PLT-1496] Use extended release channel in GKE by default

## Previous development

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
