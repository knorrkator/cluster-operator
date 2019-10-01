package nfsserver

import (
	"context"
	"reflect"
	"testing"

	storageosapis "github.com/storageos/cluster-operator/pkg/apis"
	storageosv1 "github.com/storageos/cluster-operator/pkg/apis/storageos/v1"
	fakeStosClientset "github.com/storageos/cluster-operator/pkg/client/clientset/versioned/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// getTestCluster returns a StorageOSCluster object with the given properties.
func getTestCluster(
	name string, namespace string,
	spec storageosv1.StorageOSClusterSpec,
	status storageosv1.StorageOSClusterStatus) *storageosv1.StorageOSCluster {

	return &storageosv1.StorageOSCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec:   spec,
		Status: status,
	}
}

// getTestNFSServer returns a NFSServer object with the given properties.
func getTestNFSServer(
	name string, namespace string,
	spec storageosv1.NFSServerSpec,
	status storageosv1.NFSServerStatus) *storageosv1.NFSServer {

	return &storageosv1.NFSServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec:   spec,
		Status: status,
	}
}

func TestGetCurrentStorageOSCluster(t *testing.T) {
	emptySpec := storageosv1.StorageOSClusterSpec{}
	emptyStatus := storageosv1.StorageOSClusterStatus{}

	testcases := []struct {
		name            string
		clusters        []*storageosv1.StorageOSCluster
		wantClusterName string
		wantErr         error
	}{
		{
			name: "multiple clusters with one ready",
			clusters: []*storageosv1.StorageOSCluster{
				getTestCluster("cluster1", "default", emptySpec, emptyStatus),
				getTestCluster("cluster2", "foo", emptySpec,
					storageosv1.StorageOSClusterStatus{
						Phase: storageosv1.ClusterPhaseRunning,
					}),
				getTestCluster("cluster3", "default", emptySpec, emptyStatus),
			},
			wantClusterName: "cluster2",
		},
		{
			name: "multiple clusters with none ready",
			clusters: []*storageosv1.StorageOSCluster{
				getTestCluster("cluster1", "default", emptySpec, emptyStatus),
				getTestCluster("cluster2", "default", emptySpec,
					storageosv1.StorageOSClusterStatus{
						Phase: storageosv1.ClusterPhaseInitial,
					}),
				getTestCluster("cluster3", "default", emptySpec, emptyStatus),
			},
			wantErr: ErrNoCluster,
		},
		{
			name: "single cluster not ready",
			clusters: []*storageosv1.StorageOSCluster{
				getTestCluster("cluster1", "default", emptySpec, emptyStatus),
			},
			wantClusterName: "cluster1",
		},
		{
			name:    "no cluster",
			wantErr: ErrNoCluster,
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Create fake storageos client.
			stosClient := fakeStosClientset.NewSimpleClientset()

			// Create the clusters.
			for _, c := range tc.clusters {
				_, err := stosClient.StorageosV1().StorageOSClusters(c.Namespace).Create(c)
				if err != nil {
					t.Fatalf("failed to create StorageOSCluster: %v", err)
				}
			}

			// Create a reconciler.
			reconciler := ReconcileNFSServer{
				stosClientset: stosClient,
			}

			cc, err := reconciler.getCurrentStorageOSCluster()
			if err != nil {
				if !reflect.DeepEqual(tc.wantErr, err) {
					t.Fatalf("unexpected error while getting current cluster: %v", err)
				}
			} else {
				if tc.wantClusterName != cc.Name {
					t.Errorf("unexpected current cluster selection:\n\t(WNT) %s\n\t(GOT) %s", tc.wantClusterName, cc.Name)
				}
			}
		})
	}
}

func TestUpdateSpec(t *testing.T) {
	emptyClusterSpec := storageosv1.StorageOSClusterSpec{}
	emptyClusterStatus := storageosv1.StorageOSClusterStatus{}
	emptyNFSSpec := storageosv1.NFSServerSpec{}
	emptyNFSStatus := storageosv1.NFSServerStatus{}

	testcases := []struct {
		name          string
		cluster       *storageosv1.StorageOSCluster
		nfsServer     *storageosv1.NFSServer
		wantNFSServer *storageosv1.NFSServer
		wantUpdate    bool
		wantErr       error
	}{
		{
			name:      "inherit attributes from cluster",
			cluster:   getTestCluster("cluster1", "default", emptyClusterSpec, emptyClusterStatus),
			nfsServer: getTestNFSServer("nfs1", "default", emptyNFSSpec, emptyNFSStatus),
			wantNFSServer: getTestNFSServer("nfs1", "default",
				storageosv1.NFSServerSpec{
					StorageClassName: "fast",
					NFSContainer:     storageosv1.DefaultNFSContainerImage,
				},
				emptyNFSStatus,
			),
			wantUpdate: true,
		},
		{
			// Check if the overridden cluster level defaults are inherited to
			// the NFS Server.
			name: "update the default properties in cluster",
			cluster: getTestCluster(
				"cluster1", "default",
				storageosv1.StorageOSClusterSpec{
					StorageClassName: "testsc",
					Images: storageosv1.ContainerImages{
						NFSContainer: "test-image",
					},
				}, emptyClusterStatus),
			nfsServer: getTestNFSServer("nfs1", "default", emptyNFSSpec, emptyNFSStatus),
			wantNFSServer: getTestNFSServer("nfs1", "default", storageosv1.NFSServerSpec{
				StorageClassName: "testsc",
				NFSContainer:     "test-image",
			}, emptyNFSStatus),
			wantUpdate: true,
		},
		{
			// Check that there's no update when the NFS Server CR is already
			// up-to-date.
			name:    "no new attributes to update",
			cluster: getTestCluster("cluster1", "default", emptyClusterSpec, emptyClusterStatus),
			nfsServer: getTestNFSServer(
				"nfs1", "default",
				storageosv1.NFSServerSpec{
					StorageClassName: "fast",
					NFSContainer:     storageosv1.DefaultNFSContainerImage,
				}, emptyNFSStatus),
			wantUpdate: false,
		},
		{
			// When the attributes are defined in NFS Server CR, no CR update
			// should happen.
			name:    "override default attributes",
			cluster: getTestCluster("cluster1", "default", emptyClusterSpec, emptyClusterStatus),
			nfsServer: getTestNFSServer("nfs1", "default", storageosv1.NFSServerSpec{
				StorageClassName: "testsc",
				NFSContainer:     "test-image",
			}, emptyNFSStatus),
			wantUpdate: false,
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Create a new scheme and add StorageOS APIs to it. Pass this to the
			// k8s client so that it can create StorageOS resources.
			testScheme := runtime.NewScheme()
			storageosapis.AddToScheme(testScheme)

			client := fake.NewFakeClientWithScheme(testScheme, tc.cluster, tc.nfsServer)

			reconciler := ReconcileNFSServer{
				client: client,
			}

			// Update NFSServer instance with the StorageOS Cluster and check the
			// results.
			result, err := reconciler.updateSpec(tc.nfsServer, tc.cluster)
			if err != nil {
				t.Fatalf("error while updating spec: %v", err)
			}

			if result != tc.wantUpdate {
				t.Errorf("unexpected update spec result:\n\t(WNT) %t\n\t(GOT) %t", tc.wantUpdate, result)
			}

			// If there was an update, get the NFS Server and check if it's as
			// expected.
			if tc.wantUpdate {
				namespacedNameNFS := types.NamespacedName{Name: tc.nfsServer.Name, Namespace: tc.nfsServer.Namespace}
				nfsServer := &storageosv1.NFSServer{}

				if err := client.Get(context.TODO(), namespacedNameNFS, nfsServer); err != nil {
					t.Fatalf("failed to get NFS Server: %v", err)
				}

				if !reflect.DeepEqual(nfsServer, tc.wantNFSServer) {
					t.Errorf("unexpected NFS Server:\n\t(WNT) %v\n\t(GOT) %v", tc.wantNFSServer, nfsServer)
				}
			}
		})
	}
}
