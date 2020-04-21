package main

import (
	"testing"

	kurlv1beta1 "github.com/replicatedhq/kurl/kurlkinds/pkg/apis/cluster/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_parseBashFlags(t *testing.T) {
	tests := []struct {
		name            string
		oldInstaller    *kurlv1beta1.Installer
		mergedInstaller *kurlv1beta1.Installer
		bashFlags       string
		wantError       bool
	}{
		{
			name: "All proper flags and values new fields",
			oldInstaller: &kurlv1beta1.Installer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: kurlv1beta1.InstallerSpec{},
			},
			mergedInstaller: &kurlv1beta1.Installer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: kurlv1beta1.InstallerSpec{
					Docker: kurlv1beta1.Docker{
						DockerRegistryIP: "1.1.1.1",
					},
					Kubernetes: kurlv1beta1.Kubernetes{
						MasterAddress:      "1.1.1.1",
						HACluster:          true,
						ControlPlane:       true,
						KubeadmToken:       "token",
						KubeadmTokenCAHash: "hash",
						Version:            "1.18.1",
						CertKey:            "secret",
					},
					Kurl: kurlv1beta1.Kurl{
						Airgap: true,
					},
				},
			},
			bashFlags: "airgap cert-key=secret control-plane docker-registry-ip=1.1.1.1 ha kubeadm-token=token kubeadm-token-ca-hash=hash kubernetes-master-address=1.1.1.1 kubernetes-version=1.18.1 installer-spec-file=in.yaml",
			wantError: false,
		},
		{
			name: "All proper flags and values replace fields",
			oldInstaller: &kurlv1beta1.Installer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: kurlv1beta1.InstallerSpec{
					Docker: kurlv1beta1.Docker{
						DockerRegistryIP: "2.2.2.2",
					},
					Kubernetes: kurlv1beta1.Kubernetes{
						MasterAddress:      "2.2.2.2",
						HACluster:          false,
						ControlPlane:       false,
						KubeadmToken:       "badtoken",
						KubeadmTokenCAHash: "badhash",
						Version:            "1.15.0",
						CertKey:            "badsecret",
					},
					Kurl: kurlv1beta1.Kurl{
						Airgap: false,
					},
				},
			},
			mergedInstaller: &kurlv1beta1.Installer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: kurlv1beta1.InstallerSpec{
					Docker: kurlv1beta1.Docker{
						DockerRegistryIP: "1.1.1.1",
					},
					Kubernetes: kurlv1beta1.Kubernetes{
						MasterAddress:      "1.1.1.1",
						HACluster:          true,
						ControlPlane:       true,
						KubeadmToken:       "token",
						KubeadmTokenCAHash: "hash",
						Version:            "1.18.1",
						CertKey:            "secret",
					},
					Kurl: kurlv1beta1.Kurl{
						Airgap: true,
					},
				},
			},
			bashFlags: "airgap cert-key=secret control-plane docker-registry-ip=1.1.1.1 ha kubeadm-token=token kubeadm-token-ca-hash=hash kubernetes-master-address=1.1.1.1 kubernetes-version=1.18.1 installer-spec-file=in.yaml",
			wantError: false,
		},
		{
			name: "Proper flag with no value",
			oldInstaller: &kurlv1beta1.Installer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: kurlv1beta1.InstallerSpec{},
			},
			mergedInstaller: &kurlv1beta1.Installer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: kurlv1beta1.InstallerSpec{},
			},
			bashFlags: "certkey",
			wantError: true,
		},
		{
			name: "Improper flag",
			oldInstaller: &kurlv1beta1.Installer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: kurlv1beta1.InstallerSpec{},
			},
			mergedInstaller: &kurlv1beta1.Installer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: kurlv1beta1.InstallerSpec{},
			},
			bashFlags: "BaD FlAgS",
			wantError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)

			err := parseBashFlags(test.oldInstaller, test.bashFlags)
			if test.wantError {
				req.Error(err)
			} else {
				req.NoError(err)
			}

			assert.Equal(t, test.oldInstaller, test.mergedInstaller)
		})
	}
}