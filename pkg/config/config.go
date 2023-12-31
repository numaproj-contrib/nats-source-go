package config

import corev1 "k8s.io/api/core/v1"

/* Package config defines the configuration for the NATS user-defined source.
The configuration includes the URL to connect to NATS cluster, the subject onto which messages are published,
the queue for queue subscription, the TLS configuration for the NATS client and the authentication information.
*/

// Config represents the configuration for the NATS client.
type Config struct {
	// URL to connect to NATS cluster, multiple urls could be separated by comma.
	URL string `json:"url" protobuf:"bytes,1,opt,name=url"`
	// Subject holds the name of the subject onto which messages are published.
	Subject string `json:"subject" protobuf:"bytes,2,opt,name=subject"`
	// Queue is used for queue subscription.
	Queue string `json:"queue" protobuf:"bytes,3,opt,name=queue"`
	// TLS configuration for the NATS client.
	// +optional
	TLS *TLS `json:"tls" protobuf:"bytes,4,opt,name=tls"`
	// Auth information
	// +optional
	Auth *Auth `json:"auth,omitempty" protobuf:"bytes,5,opt,name=auth"`
}

// TLS defines the TLS configuration for the NATS client.
type TLS struct {
	// +optional
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty" protobuf:"bytes,1,opt,name=insecureSkipVerify"`
	// CACertSecret refers to the secret that contains the CA cert
	// +optional
	CACertSecret *corev1.SecretKeySelector `json:"caCertSecret,omitempty" protobuf:"bytes,2,opt,name=caCertSecret"`
	// CertSecret refers to the secret that contains the cert
	// +optional
	CertSecret *corev1.SecretKeySelector `json:"clientCertSecret,omitempty" protobuf:"bytes,3,opt,name=certSecret"`
	// KeySecret refers to the secret that contains the key
	// +optional
	KeySecret *corev1.SecretKeySelector `json:"clientKeySecret,omitempty" protobuf:"bytes,4,opt,name=keySecret"`
}

// BasicAuth represents the basic authentication approach which contains a username and a password.
type BasicAuth struct {
	// Secret for auth user
	// +optional
	User *corev1.SecretKeySelector `json:"user,omitempty" protobuf:"bytes,1,opt,name=user"`
	// Secret for auth password
	// +optional
	Password *corev1.SecretKeySelector `json:"password,omitempty" protobuf:"bytes,2,opt,name=password"`
}

// Auth represents the authentication information for the NATS client.
type Auth struct {
	// Basic auth, which contains a username and a password,
	// +optional
	Basic *BasicAuth `json:"basic,omitempty" protobuf:"bytes,1,opt,name=basic"`
	// Token auth
	// +optional
	Token *corev1.SecretKeySelector `json:"token,omitempty" protobuf:"bytes,2,opt,name=token"`
	// NKey auth
	// +optional
	NKey *corev1.SecretKeySelector `json:"nkey,omitempty" protobuf:"bytes,3,opt,name=nkey"`
}
