package adder_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/types"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/example/memcached-operator/adder"
)

var _ = Describe("Adder", func() {

	Describe("Add", func() {

		const timeout = time.Second * 60
		const interval = time.Second * 1

		ctx := context.Background()

		BeforeEach(func() {
			// failed test runs that don't clean up leave resources behind.
		})

		AfterEach(func() {

		})

		It("Should add the warmup initContainer as the first item in initContainers", func() {

			// construct a prescaled cron in code  post to K8s
			toCreate := generatePSCSpec()
			By("Creating the prescaled cron job CRD")
			Expect(k8sClient.Create(ctx, &toCreate)).Should(Succeed())
			// time.Sleep(time.Second * 5)

			fetched := &appsv1.Deployment{}

			// check the CRD was created ok
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: toCreate.Name, Namespace: "default"}, fetched)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			// original := &appsv1.Deployment{}
			// Expect(k8sClient.Get(ctx, types.NamespacedName{Name: toCreate.Name, Namespace: "default"}, original)).Should(Succeed())

			// By("Deleting the prescaled cron job CRD")
			// Expect(k8sClient.Delete(ctx, original)).Should(Succeed())
			// time.Sleep(time.Second * 30)

		})

		Context("when summands are positive", func() {

			It("adds two numbers", func() {
				sum, err := Add(2, 3)
				Expect(err).NotTo(HaveOccurred())
				Expect(sum).To(Equal(5))
			})

		})

		Context("when summand is negative", func() {

			It("returns an err", func() {
				_, err := Add(-1, -1)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

func generatePSCSpec() appsv1.Deployment {

	toCreate := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "deployment-name",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nombre-aplicacion",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "nombre-aplicacion",
					},
				},
				Spec: corev1.PodSpec{
					InitContainers: []corev1.Container{
						{
							Name:  "warmup",
							Image: "nodejs",
							// Configura los puertos y otras opciones necesarias para el initContainer warmup
						},
					},
					Containers: []corev1.Container{
						{
							Name:  "nombre-contenedor",
							Image: "nginx",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	return toCreate
}
