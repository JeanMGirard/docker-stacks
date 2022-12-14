CoreDNS
=======

CoreDNS is a DNS server that chains plugins and provides DNS Services

TL;DR;
------

```console
$ helm install --name coredns --namespace=kube-system stable/coredns
```

Introduction
------------

This chart bootstraps a [CoreDNS](https://github.com/coredns/coredns) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager. This chart will provide DNS Services and can be deployed in multiple configuration to support various scenarios listed below:

 - CoreDNS as a cluster dns service and a drop-in replacement for Kube/SkyDNS. This is the default mode and CoreDNS is deployed as cluster-service in kube-system namespace. This mode is chosen by setting `isClusterService` to true.
 - CoreDNS as an external dns service. In this mode CoreDNS is deployed as any kubernetes app in user specified namespace. The CoreDNS service can be exposed outside the cluster by using using either the NodePort or LoadBalancer type of service. This mode is chosen by setting `isClusterService` to false.
 - CoreDNS as an external dns provider for kubernetes federation. This is a sub case of 'external dns service' which uses etcd plugin for CoreDNS backend. This deployment mode as a dependency on `etcd-operator` chart, which needs to be pre-installed.

Prerequisites
-------------

-	Kubernetes 1.4+ with Beta APIs enabled

Installing the Chart
--------------------

The chart can be installed as follows:

```console
$ helm install --name coredns --namespace=kube-system stable/coredns
```

The command deploys CoreDNS on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists various ways to override default configuration during deployment.

> **Tip**: List all releases using `helm list`

Uninstalling the Chart
----------------------

To uninstall/delete the `my-release` deployment:

```console
$ helm delete coredns
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

Configuration
-------------

See `values.yaml` for configuration notes. Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
$ helm install --name coredns \
  --set rbac.create=false \
    stable/coredns
```

The above command disables automatic creation of RBAC rules.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. For example,

```console
$ helm install --name coredns -f values.yaml stable/coredns
```

> **Tip**: You can use the default [values.yaml](values.yaml)


Caveats
-------

The chart will automatically determine which protocols to listen on based on
the protocols you define in your zones. This means that you could potentially
use both "TCP" and "UDP" on a single port.
Some cloud environments like "GCE" or "Azure container service" cannot
create external loadbalancers with both "TCP" and "UDP" protocols. So
When deploying CoreDNS with `serviceType="LoadBalancer"` on such cloud
environments, make sure you do not attempt to use both protocols at the same
time.
