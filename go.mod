module kope.io/etcd-manager

go 1.12

// Use kops in the CloudNatix fork
replace k8s.io/kops => github.com/cloudnatix/kops v1.18.0-alpha.2.0.20201007160252-ae3cddf69a08

// Version kubernetes-1.15.3
//replace k8s.io/kubernetes => k8s.io/kubernetes v1.15.3
//replace k8s.io/api => k8s.io/api kubernetes-1.15.3
//replace k8s.io/apimachinery => k8s.io/apimachinery kubernetes-1.15.3
//replace k8s.io/client-go => k8s.io/client-go kubernetes-1.15.3
//replace k8s.io/cloud-provider => k8s.io/cloud-provider kubernetes-1.15.3
//replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers kubernetes-1.15.3

replace k8s.io/api => k8s.io/api v0.18.1

replace k8s.io/apimachinery => k8s.io/apimachinery v0.18.1

replace k8s.io/client-go => k8s.io/client-go v0.18.1

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.1

replace k8s.io/kubectl => k8s.io/kubectl v0.18.1

// Dependencies we don't really need, except that kubernetes specifies them as v0.0.0 which confuses go.mod
//replace k8s.io/apiserver => k8s.io/apiserver kubernetes-1.15.3
//replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver kubernetes-1.15.3
//replace k8s.io/kube-scheduler => k8s.io/kube-scheduler kubernetes-1.15.3
//replace k8s.io/kube-proxy => k8s.io/kube-proxy kubernetes-1.15.3
//replace k8s.io/cri-api => k8s.io/cri-api kubernetes-1.15.3
//replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib kubernetes-1.15.3
//replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers kubernetes-1.15.3
//replace k8s.io/component-base => k8s.io/component-base kubernetes-1.15.3
//replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap kubernetes-1.15.3
//replace k8s.io/metrics => k8s.io/metrics kubernetes-1.15.3
//replace k8s.io/sample-apiserver => k8s.io/sample-apiserver kubernetes-1.15.3
//replace k8s.io/kube-aggregator => k8s.io/kube-aggregator kubernetes-1.15.3
//replace k8s.io/kubelet => k8s.io/kubelet kubernetes-1.15.3
//replace k8s.io/cli-runtime => k8s.io/cli-runtime kubernetes-1.15.3
//replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager kubernetes-1.15.3
//replace k8s.io/code-generator => k8s.io/code-generator kubernetes-1.15.3
//replace k8s.io/cli-runtime => k8s.io/cli-runtime kubernetes-1.15.3

replace k8s.io/apiserver => k8s.io/apiserver v0.18.1

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.1

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.1

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.1

replace k8s.io/cri-api => k8s.io/cri-api v0.18.1

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.1

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.1

replace k8s.io/component-base => k8s.io/component-base v0.18.1

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.1

replace k8s.io/metrics => k8s.io/metrics v0.18.1

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.1

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.1

replace k8s.io/kubelet => k8s.io/kubelet v0.18.1

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.1

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.1

replace k8s.io/code-generator => k8s.io/code-generator v0.18.1

// Pick up a fix for gRPC build error (https://github.com/etcd-io/etcd/issues/11563)
replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible


require (
	cloud.google.com/go v0.38.0
	github.com/Azure/azure-sdk-for-go v46.3.0+incompatible
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.2
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Azure/go-autorest/autorest/validation v0.3.0 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.60.290
	github.com/aws/aws-sdk-go v1.30.16
	github.com/blang/semver v3.5.1+incompatible
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/digitalocean/godo v1.19.0
	github.com/golang/protobuf v1.3.3
	github.com/gophercloud/gophercloud v0.7.1-0.20200116011225-46fdd1830e9a
	github.com/pkg/sftp v0.0.0-20180127012644-738e088bbd93 // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.22.0
	google.golang.org/grpc v1.27.0
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.2.8
	honnef.co/go/tools v0.0.1-2019.2.3
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v1.0.0
	k8s.io/kops v0.0.0-00010101000000-000000000000
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
)
