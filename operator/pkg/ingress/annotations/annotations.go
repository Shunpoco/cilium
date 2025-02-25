// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package annotations

import (
	"strconv"

	"github.com/cilium/cilium/pkg/annotation"
	slim_corev1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/api/core/v1"
	slim_networkingv1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/api/networking/v1"
)

const (
	LBModeAnnotation           = annotation.Prefix + ".ingress" + "/loadbalancer-mode"
	ServiceTypeAnnotation      = annotation.Prefix + ".ingress" + "/service-type"
	InsecureNodePortAnnotation = annotation.Prefix + ".ingress" + "/insecure-node-port"
	SecureNodePortAnnotation   = annotation.Prefix + ".ingress" + "/secure-node-port"

	TCPKeepAliveEnabledAnnotation          = annotation.Prefix + "/tcp-keep-alive"
	TCPKeepAliveIdleAnnotation             = annotation.Prefix + "/tcp-keep-alive-idle"
	TCPKeepAliveProbeIntervalAnnotation    = annotation.Prefix + "/tcp-keep-alive-probe-interval"
	TCPKeepAliveProbeMaxFailuresAnnotation = annotation.Prefix + "/tcp-keep-alive-probe-max-failures"
	WebsocketEnabledAnnotation             = annotation.Prefix + "/websocket"
)

const (
	enabled                          = "enabled"
	defaultTCPKeepAliveEnabled       = 1  // 1 - Enabled, 0 - Disabled
	defaultTCPKeepAliveInitialIdle   = 10 // in seconds
	defaultTCPKeepAliveProbeInterval = 5  // in seconds
	defaultTCPKeepAliveMaxProbeCount = 10
	defaultWebsocketEnabled          = 0 // 1 - Enabled, 0 - Disabled
)

// GetAnnotationIngressLoadbalancerMode returns the loadbalancer mode for the ingress if possible.
func GetAnnotationIngressLoadbalancerMode(ingress *slim_networkingv1.Ingress) string {
	return ingress.GetAnnotations()[LBModeAnnotation]
}

// GetAnnotationServiceType returns the service type for the ingress if possible.
// Defaults to LoadBalancer
func GetAnnotationServiceType(ingress *slim_networkingv1.Ingress) string {
	val, exists := ingress.GetAnnotations()[ServiceTypeAnnotation]
	if !exists {
		return string(slim_corev1.ServiceTypeLoadBalancer)
	}
	return val
}

// GetAnnotationSecureNodePort returns the secure node port for the ingress if possible.
func GetAnnotationSecureNodePort(ingress *slim_networkingv1.Ingress) (*uint32, error) {
	val, exists := ingress.GetAnnotations()[SecureNodePortAnnotation]
	if !exists {
		return nil, nil
	}
	intVal, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return nil, err
	}
	res := uint32(intVal)
	return &res, nil
}

// GetAnnotationInsecureNodePort returns the insecure node port for the ingress if possible.
func GetAnnotationInsecureNodePort(ingress *slim_networkingv1.Ingress) (*uint32, error) {
	val, exists := ingress.GetAnnotations()[InsecureNodePortAnnotation]
	if !exists {
		return nil, nil
	}
	intVal, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return nil, err
	}
	res := uint32(intVal)
	return &res, nil
}

// GetAnnotationTCPKeepAliveEnabled returns 1 if enabled (default), 0 if disabled
func GetAnnotationTCPKeepAliveEnabled(ingress *slim_networkingv1.Ingress) int64 {
	val, exists := ingress.GetAnnotations()[TCPKeepAliveEnabledAnnotation]
	if !exists {
		return defaultTCPKeepAliveEnabled
	}
	if val == enabled {
		return 1
	}
	return 0
}

// GetAnnotationTCPKeepAliveIdle returns the time (in seconds) the connection needs to
// remain idle before TCP starts sending keepalive probes. Defaults to 10s.
// Related references:
//   - https://man7.org/linux/man-pages/man7/tcp.7.html
func GetAnnotationTCPKeepAliveIdle(ingress *slim_networkingv1.Ingress) int64 {
	val, exists := ingress.GetAnnotations()[TCPKeepAliveIdleAnnotation]
	if !exists {
		return defaultTCPKeepAliveInitialIdle
	}
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultTCPKeepAliveInitialIdle
	}
	return intVal
}

// GetAnnotationTCPKeepAliveProbeInterval returns the time (in seconds) between individual
// keepalive probes. Defaults to 5s.
// Related references:
//   - https://man7.org/linux/man-pages/man7/tcp.7.html
func GetAnnotationTCPKeepAliveProbeInterval(ingress *slim_networkingv1.Ingress) int64 {
	val, exists := ingress.GetAnnotations()[TCPKeepAliveProbeIntervalAnnotation]
	if !exists {
		return defaultTCPKeepAliveProbeInterval
	}
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultTCPKeepAliveProbeInterval
	}
	return intVal
}

// GetAnnotationTCPKeepAliveProbeMaxFailures returns the maximum number of keepalive probes TCP
// should send before dropping the connection. Defaults to 10.
// Related references:
//   - https://man7.org/linux/man-pages/man7/tcp.7.html
func GetAnnotationTCPKeepAliveProbeMaxFailures(ingress *slim_networkingv1.Ingress) int64 {
	val, exists := ingress.GetAnnotations()[TCPKeepAliveProbeMaxFailuresAnnotation]
	if !exists {
		return defaultTCPKeepAliveMaxProbeCount
	}
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultTCPKeepAliveMaxProbeCount
	}
	return intVal
}

// GetAnnotationWebsocketEnabled returns 1 if enabled (default), 0 if disabled
func GetAnnotationWebsocketEnabled(ingress *slim_networkingv1.Ingress) int64 {
	val, exists := ingress.GetAnnotations()[WebsocketEnabledAnnotation]
	if !exists {
		return defaultWebsocketEnabled
	}
	if val == enabled {
		return 1
	}
	return 0
}
