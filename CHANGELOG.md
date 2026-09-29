# Changelog

## 1.13.1-0.5.0 (upcoming)

* [PLT-4891] Rebase the fork onto upstream `v1.13.1` (reads the CAPI core `Cluster`/`MachinePool` as `v1beta2`, requires CAPI core >= v1.11). Stratio behaviours re-applied: GKE NetworkPolicy (Calico), boot disk CMEK, `clusterIpv4Cidr` and `ipAllocationPolicy` (incl. secondary range names), Managed Prometheus and per-component logging, Workload Identity on update, node pool resize with autoscaling enabled, `CLUSTER_ALREADY_HAS_OPERATION` handling, replica write-back to the `MachinePool`. Private cluster, `extended` release channel, Workload Identity on create and the taints/authorized-networks diff fixes now come from upstream
* [PLT-4891] Keep the authorized networks of `masterAuthorizedNetworksConfig` when creating a private-endpoint GKE cluster (upstream replaced them with an empty config and the follow-up `UpdateCluster` delayed the kubeconfig past cloud-provisioner's wait); skip the `LinuxNodeConfig` diff when it is not set (upstream `eedd22f5`), which recreated every node pool right after creation
* [PLT-4891] Clear `status.infrastructureMachineKind` on `GCPManagedMachinePool`: GKE node pools have no MachinePool Machines, and with the kind set CAPI >= v1.11 reported 0 ready/available replicas, blocking any caller that waits for the `MachinePool`
* [PLT-4891] Keep the CVE fixes already shipped in `1.6.1-0.4.x` on the new base: `google.golang.org/grpc` v1.83.2, `golang.org/x/net` v0.58.0, `golang.org/x/crypto` v0.56.0, `go.opentelemetry.io/otel*` v1.44.0, Go 1.26 (builder `golang:1.26.8`)

## 1.6.1-0.4.0 (2025-10-07)

* [PLT-4748] Bump `golang.org/x/net` (0.43.0 → 0.58.0), `golang.org/x/crypto` (0.42.0 → 0.56.0), `google.golang.org/grpc` (1.67.3 → 1.83.1), `go.opentelemetry.io/otel*` (1.39.0 → 1.44.0) and Go toolchain (1.24.6 → 1.26.0) to close vulnerabilities
* Fix `Jenkinsfile`'s `@Library('libpipelines@master')` pointing at a branch that no longer exists in `Stratio/jenkins-bootstrap` (renamed to `main`), breaking CI on every PR regardless of code changes
* [PLT-2635] Fix golang vulnerabilities to max provider version 1.24.6

## Previous development

### Branched to branch-1.6.1-0.4 (2025-07-22)

* [PLT-1548] [GKE] Activar Workload Identity



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
