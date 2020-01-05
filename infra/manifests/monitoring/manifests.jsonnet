local k = import 'ksonnet/ksonnet.beta.4/k.libsonnet';

local kp =
  (import 'kube-prometheus/kube-prometheus.libsonnet') +
  // Uncomment the following imports to enable its patches
  (import 'kube-prometheus/kube-prometheus-anti-affinity.libsonnet') +
  (import 'kube-prometheus/kube-prometheus-kubespray.libsonnet') +
  (import 'etcd-mixin/mixin.libsonnet') +
  // (import 'kube-prometheus/kube-prometheus-managed-cluster.libsonnet') +
  // (import 'kube-prometheus/kube-prometheus-node-ports.libsonnet') +
  // (import 'kube-prometheus/kube-prometheus-static-etcd.libsonnet') +
  // (import 'kube-prometheus/kube-prometheus-thanos-sidecar.libsonnet') +
  {
    _config+:: {
      namespace: 'monitoring',

      etcd+:: {
        ips: ['130.79.207.30', '130.79.207.31', '130.79.207.32'],
      },

      grafana+:: {
        config: {
          sections: {
            "auth.anonymous": {enabled: true},
          },
        },
      },
    },

    nodeExporter+: {
      podSecurityPolicy+:
        local policy = k.policy.v1beta1.podSecurityPolicy;
        policy.new() +
        policy.mixin.metadata.withName("psp-node-exporter") +
        policy.mixin.metadata.withLabels({ app: 'node-exporter' }) +
        policy.mixin.spec.withPrivileged(false) +
        policy.mixin.spec.withVolumes(['configMap', 'secret', 'hostPath']) +
        policy.mixin.spec.withHostNetwork(true) +
        policy.mixin.spec.withHostIpc(false) +
        policy.mixin.spec.withHostPid(true) +
        policy.mixin.spec.withHostPorts({min: 0, max: 65535}) +
        policy.mixin.spec.runAsUser.withRule('RunAsAny') +
        policy.mixin.spec.seLinux.withRule('RunAsAny') +
        policy.mixin.spec.supplementalGroups.withRule('MustRunAs') +
        policy.mixin.spec.supplementalGroups.withRanges({min: 1, max: 65535}) +
        policy.mixin.spec.fsGroup.withRule('MustRunAs') +
        policy.mixin.spec.fsGroup.withRanges({min: 1, max: 65535}) +
        policy.mixin.spec.withReadOnlyRootFilesystem(false),

      clusterRole+:
        local clusterRole = k.rbac.v1.clusterRole;
        local policyRule = clusterRole.rulesType;
        local podSecurityRole =
          policyRule.new() +
          policyRule.withApiGroups(['extensions']) +
          policyRule.withResourceNames(['psp-node-exporter']) +
          policyRule.withResources([
            'podsecuritypolicies',
          ]) +
          policyRule.withVerbs(['use']);
        clusterRole.withRulesMixin([podSecurityRole])
    },

    prometheus+:: {
      serviceEtcd:
        local service = k.core.v1.service;
        local servicePort = k.core.v1.service.mixin.spec.portsType;

        local etcdServicePort = servicePort.newNamed('metrics', 2381, 2381);

        service.new('etcd', null, etcdServicePort) +
        service.mixin.metadata.withNamespace('kube-system') +
        service.mixin.metadata.withLabels({ 'component': 'etcd' }) +
        service.mixin.spec.withClusterIp('None') +
        service.mixin.spec.withSelector({ 'component': 'etcd' }),
      serviceMonitorEtcd:
        {
          apiVersion: 'monitoring.coreos.com/v1',
          kind: 'ServiceMonitor',
          metadata: {
            name: 'etcd',
            namespace: 'kube-system',
            labels: {
              'component': 'etcd',
            },
          },
          spec: {
            jobLabel: 'component',
            endpoints: [
              {
                port: 'metrics',
                interval: '30s',
                scheme: 'http',
              },
            ],
            selector: {
              matchLabels: {
                'component': 'etcd',
              },
            },
          },
        },

    },
  };

{ ['setup/0namespace-' + name + '.json']: kp.kubePrometheus[name] for name in std.objectFields(kp.kubePrometheus) } +
{
  ['setup/prometheus-operator-' + name + '.json']: kp.prometheusOperator[name]
  for name in std.filter((function(name) name != 'serviceMonitor'), std.objectFields(kp.prometheusOperator))
} +
// serviceMonitor is separated so that it can be created after the CRDs are ready
{ 'prometheus-operator-serviceMonitor.json': kp.prometheusOperator.serviceMonitor } +
{ ['node-exporter-' + name + '.json']: kp.nodeExporter[name] for name in std.objectFields(kp.nodeExporter) } +
{ ['kube-state-metrics-' + name + '.json']: kp.kubeStateMetrics[name] for name in std.objectFields(kp.kubeStateMetrics) } +
{ ['alertmanager-' + name + '.json']: kp.alertmanager[name] for name in std.objectFields(kp.alertmanager) } +
{ ['prometheus-' + name + '.json']: kp.prometheus[name] for name in std.objectFields(kp.prometheus) } +
{ ['prometheus-adapter-' + name + '.json']: kp.prometheusAdapter[name] for name in std.objectFields(kp.prometheusAdapter) } +
{ ['grafana-' + name + '.json']: kp.grafana[name] for name in std.objectFields(kp.grafana) }
