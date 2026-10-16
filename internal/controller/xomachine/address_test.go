package xomachine

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gofrs/uuid"
	"go.uber.org/mock/gomock"

	corev1 "k8s.io/api/core/v1"

	xoclient "github.com/vatesfr/xenorchestra-go-sdk/client"
	"github.com/vatesfr/xenorchestra-go-sdk/pkg/payloads"
	xok8scommon "github.com/vatesfr/xenorchestra-k8s-common"
	k8smocks "github.com/vatesfr/xenorchestra-k8s-common/mocks"

	infrastructurev1beta2 "github.com/vatesfr/cluster-api-provider-vates/api/v1beta2"
)

func TestNodeIPv4(t *testing.T) {
	tests := []struct {
		name string
		addr string
		want bool
	}{
		{name: "IPv4", addr: "10.30.1.42", want: true},
		{name: "empty", addr: "", want: false},
		{name: "IPv6 link-local", addr: "fe80::f0ed:f7ff:feaa:20fa", want: false},
		{name: "IPv6 global", addr: "2001:db8::1", want: false},
		{name: "IPv4 link-local", addr: "169.254.72.203", want: false},
		{name: "IPv6-mapped IPv4", addr: "::ffff:10.30.1.42", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodeIPv4(tt.addr); got != tt.want {
				t.Errorf("nodeIPv4(%q) = %v, want %v", tt.addr, got, tt.want)
			}
		})
	}
}

func TestSetNodeAddress(t *testing.T) {
	internal := func(addr string) corev1.NodeAddress {
		return corev1.NodeAddress{Type: corev1.NodeInternalIP, Address: addr}
	}

	tests := []struct {
		name      string
		addresses []corev1.NodeAddress
		addr      string
		want      []corev1.NodeAddress
		wantChg   bool
	}{
		{
			name:    "adds the address when there is none",
			addr:    "10.30.1.42",
			want:    []corev1.NodeAddress{internal("10.30.1.42")},
			wantChg: true,
		},
		{
			name:      "updates an address that moved",
			addresses: []corev1.NodeAddress{internal("10.30.1.42")},
			addr:      "10.30.1.221",
			want:      []corev1.NodeAddress{internal("10.30.1.221")},
			wantChg:   true,
		},
		{
			name:      "leaves an unchanged address alone",
			addresses: []corev1.NodeAddress{internal("10.30.1.42")},
			addr:      "10.30.1.42",
			want:      []corev1.NodeAddress{internal("10.30.1.42")},
			wantChg:   false,
		},
		{
			name:      "ignores an empty address",
			addresses: []corev1.NodeAddress{internal("10.30.1.42")},
			addr:      "",
			want:      []corev1.NodeAddress{internal("10.30.1.42")},
			wantChg:   false,
		},
		{
			name:    "empty address on an empty machine is a no-op",
			addr:    "",
			want:    nil,
			wantChg: false,
		},
		{
			name:      "keeps addresses of other types",
			addresses: []corev1.NodeAddress{{Type: corev1.NodeHostName, Address: "vates-mp-0"}},
			addr:      "10.30.1.42",
			want: []corev1.NodeAddress{
				{Type: corev1.NodeHostName, Address: "vates-mp-0"},
				internal("10.30.1.42"),
			},
			wantChg: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xoMachine := &infrastructurev1beta2.XOMachine{
				Status: infrastructurev1beta2.XOMachineStatus{Addresses: tt.addresses},
			}

			if got := SetNodeAddress(xoMachine, tt.addr); got != tt.wantChg {
				t.Errorf("SetNodeAddress() changed = %v, want %v", got, tt.wantChg)
			}
			if !reflect.DeepEqual(xoMachine.Status.Addresses, tt.want) {
				t.Errorf("addresses = %#v, want %#v", xoMachine.Status.Addresses, tt.want)
			}
		})
	}
}

func TestResolveVMIPNow(t *testing.T) {
	tests := []struct {
		name   string
		mainIP string
		v1VM   *xoclient.Vm
		v1Err  error
		want   string
	}{
		{
			name:   "returns a usable MainIpAddress",
			mainIP: "10.30.1.42",
			want:   "10.30.1.42",
		},
		{
			name: "falls back to the V1 addresses when MainIpAddress is empty",
			v1VM: &xoclient.Vm{Addresses: map[string]string{"0": "10.30.1.42"}},
			want: "10.30.1.42",
		},
		{
			name:   "rejects an IPv6 link-local MainIpAddress",
			mainIP: "fe80::1",
			v1VM:   &xoclient.Vm{Addresses: map[string]string{"0": "10.30.1.42"}},
			want:   "10.30.1.42",
		},
		{
			name:   "rejects a non-IPv4 MainIpAddress",
			mainIP: "2001:db8::1",
			v1VM:   &xoclient.Vm{Addresses: map[string]string{"0": "10.30.1.42"}},
			want:   "10.30.1.42",
		},
		{
			name: "skips unusable V1 addresses",
			v1VM: &xoclient.Vm{Addresses: map[string]string{
				"0": "fe80::1",
				"1": "169.254.1.1",
				"2": "10.30.1.42",
			}},
			want: "10.30.1.42",
		},
		{
			name: "no usable address yet",
			v1VM: &xoclient.Vm{Addresses: map[string]string{"0": "fe80::1"}},
			want: "",
		},
		{
			name:  "a V1 error is not an address",
			v1Err: errors.New("xo unavailable"),
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockV1 := NewMockXOClient(ctrl)
			mockLib := k8smocks.NewMockLibrary(ctrl)
			mockLib.EXPECT().V1Client().Return(mockV1).AnyTimes()
			mockV1.EXPECT().GetVm(gomock.Any()).Return(tt.v1VM, tt.v1Err).AnyTimes()

			vm := &payloads.VM{ID: uuid.Must(uuid.NewV4()), MainIpAddress: tt.mainIP}
			got := ResolveVMIPNow(&xok8scommon.XoClient{Client: mockLib}, vm)

			if got != tt.want {
				t.Errorf("ResolveVMIPNow() = %q, want %q", got, tt.want)
			}
		})
	}
}
