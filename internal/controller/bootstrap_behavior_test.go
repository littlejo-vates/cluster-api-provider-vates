package controller

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	infrastructurev1beta2 "github.com/vatesfr/cluster-api-provider-vates/api/v1beta2"
)

var _ = Describe("bootstrapBehavior", func() {
	It("is an empty format (pass-through) when bootstrap is unset", func() {
		Expect(bootstrapBehavior(&infrastructurev1beta2.XOMachine{}).Format).To(BeEmpty())
	})

	It("uses the declared format", func() {
		b := bootstrapBehavior(&infrastructurev1beta2.XOMachine{
			Spec: infrastructurev1beta2.XOMachineSpec{
				Bootstrap: &infrastructurev1beta2.BootstrapBehavior{
					Format: infrastructurev1beta2.BootstrapFormatCloudInit,
				},
			},
		})
		Expect(b.Format).To(Equal(infrastructurev1beta2.BootstrapFormatCloudInit))
	})

	It("derives cloud-init from the deprecated CloudInit flag", func() {
		b := bootstrapBehavior(&infrastructurev1beta2.XOMachine{
			Spec: infrastructurev1beta2.XOMachineSpec{
				Bootstrap: &infrastructurev1beta2.BootstrapBehavior{
					//nolint:staticcheck // Exercising the deprecated fallback, removed in v1.2.
					CloudInit: true,
				},
			},
		})
		Expect(b.Format).To(Equal(infrastructurev1beta2.BootstrapFormatCloudInit))
	})

	It("prefers the format over the deprecated CloudInit flag", func() {
		b := bootstrapBehavior(&infrastructurev1beta2.XOMachine{
			Spec: infrastructurev1beta2.XOMachineSpec{
				Bootstrap: &infrastructurev1beta2.BootstrapBehavior{
					Format: infrastructurev1beta2.BootstrapFormatPassthrough,
					//nolint:staticcheck // Exercising the deprecated fallback, removed in v1.2.
					CloudInit: true,
				},
			},
		})
		Expect(b.Format).To(Equal(infrastructurev1beta2.BootstrapFormatPassthrough))
	})
})
