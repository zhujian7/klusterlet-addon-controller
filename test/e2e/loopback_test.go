// Copyright (c) Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package e2e_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"

	agentv1 "github.com/open-cluster-management/klusterlet-addon-controller/pkg/apis/agent/v1"
)

var _ = Describe("Loopback test", func() {
	gvr := schema.GroupVersionResource{
		Group:    "agent.open-cluster-management.io",
		Version:  "v1",
		Resource: "klusterletaddonconfigs",
	}

	It("install", func() {
		testKlusterletAddonConfig := &agentv1.KlusterletAddonConfig{
			TypeMeta: metav1.TypeMeta{
				APIVersion: agentv1.SchemeGroupVersion.String(),
				Kind:       "KlusterletAddonConfig",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      managedclusterName,
				Namespace: managedclusterName,
			},
			Spec: agentv1.KlusterletAddonConfigSpec{
				ApplicationManagerConfig: agentv1.KlusterletAddonConfigApplicationManagerSpec{
					Enabled: true,
				},
				CertPolicyControllerConfig: agentv1.KlusterletAddonConfigCertPolicyControllerSpec{
					Enabled: true,
				},
				IAMPolicyControllerConfig: agentv1.KlusterletAddonConfigIAMPolicyControllerSpec{
					Enabled: true,
				},
				PolicyController: agentv1.KlusterletAddonConfigPolicyControllerSpec{
					Enabled: true,
				},
				SearchCollectorConfig: agentv1.KlusterletAddonConfigSearchCollectorSpec{
					Enabled: true,
				},

				ClusterName:      managedclusterName,
				ClusterNamespace: "cluster1-test",
				ClusterLabels: map[string]string{
					"author": "tester",
				},
				Version: "2.0.0",
			},
		}

		raw, err := json.Marshal(testKlusterletAddonConfig)
		Expect(err).ToNot(HaveOccurred())
		obj := &unstructured.Unstructured{}
		err = json.Unmarshal(raw, obj)
		Expect(err).ToNot(HaveOccurred())

		By("Create klusterletaddonconfigs")
		_, err = hubDynamicClient.Resource(gvr).Namespace(managedclusterName).Create(context.TODO(), obj, metav1.CreateOptions{})
		if err != nil && !errors.IsAlreadyExists(err) {
			Expect(BeFalse()).To(BeTrue())
		}

		labels := []string{
			"name=klusterlet-addon-operator",
			"app=application-manager",
			"app=cert-policy-controller",
			"app=klusterlet-addon-iampolicyctrl",
			"app=policy-config-policy",
			"app=policy-framework",
			"app=search",
			"app=work-manager",
		}

		By("Check agent addon pod numbers, will not check the pod Status as the image pull secret do not provided")
		Eventually(func() bool {
			nsList, err := spokeClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				return false
			}

			for _, ns := range nsList.Items {
				podList, err := spokeClient.CoreV1().Pods(ns.Name).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					return false
				}
				pods := make([]string, 0)
				for _, p := range podList.Items {
					pods = append(pods, p.Name)
				}
				By(fmt.Sprintf("check pod for namespace %s, pods list: %v", ns.Name, pods))
				if ns.Name == agentAddonNamespace {
					By(fmt.Sprintf("=========================== namespace %v created ===========================", agentAddonNamespace))
					return true
				}
			}

			return false
		}, 180*time.Second, 3*time.Second).Should(BeTrue())

		for _, label := range labels {
			Eventually(func() bool {
				podList, err := spokeClient.CoreV1().Pods(agentAddonNamespace).List(context.TODO(), metav1.ListOptions{
					LabelSelector: label,
				})
				if err != nil {
					return false
				}

				By(fmt.Sprintf("check pod for %s, pod item: %d", label, len(podList.Items)))
				return len(podList.Items) > 0
			}, 600*time.Second, 1*time.Second).Should(BeTrue())
		}

		By("Delete klusterletaddonconfigs")
		err = hubDynamicClient.Resource(gvr).Namespace(managedclusterName).Delete(context.TODO(), managedclusterName, metav1.DeleteOptions{})
		Expect(err).ToNot(HaveOccurred())
	})
})
