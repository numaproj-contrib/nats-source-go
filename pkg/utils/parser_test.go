package utils

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"

	"github.com/numaproj/numaflow/pkg/apis/numaflow/v1alpha1"
)

func TestConfigParser_UnParseThenParse(t *testing.T) {
	var parsers = []Parser{
		&JSONConfigParser{},
		&YAMLConfigParser{},
	}
	for _, parser := range parsers {
		c := &v1alpha1.NatsSource{
			URL:     "nats://localhost:4222",
			Subject: "test",
			Queue:   "test",
			TLS: &v1alpha1.TLS{
				InsecureSkipVerify: true,
			},
			Auth: &v1alpha1.NatsAuth{
				Basic: &v1alpha1.BasicAuth{
					User: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "test",
						},
						Key: "test",
					},
				},
			},
		}
		configStr, err := parser.UnParse(c)
		assert.NoError(t, err)
		config, err := parser.Parse(configStr)
		assert.NoError(t, err)
		assert.Equal(t, c, config)
	}
}

func TestConfigParser_ParseErrScenarios(t *testing.T) {
	var parsers = []Parser{
		&JSONConfigParser{},
		&YAMLConfigParser{},
	}
	for _, parser := range parsers {
		_, err := parser.Parse("invalid config string")
		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "failed to parse config string"))
		_, err = parser.UnParse(nil)
		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "config cannot be nil"))
	}
}

func TestConfigParser_YAML(t *testing.T) {
	yamlStr := `
url: nats://localhost:4222
subject: test
queue: test
tls:
  insecureskipverify: true
auth:
  basic:
    user:
      localobjectreference:
        name: test
      key: test
`
	parser := &YAMLConfigParser{}
	config, err := parser.Parse(yamlStr)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(&v1alpha1.NatsSource{
		URL:     "nats://localhost:4222",
		Subject: "test",
		Queue:   "test",
		TLS: &v1alpha1.TLS{
			InsecureSkipVerify: true,
		},
		Auth: &v1alpha1.NatsAuth{
			Basic: &v1alpha1.BasicAuth{
				User: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "test",
					},
					Key: "test",
				},
			},
		},
	}, config))
}

func TestConfigParser_JSON(t *testing.T) {
	jsonStr := `
{
   "url":"nats://localhost:4222",
   "subject":"test",
   "queue":"test",
   "tls":{
      "insecureSkipVerify":true
   },
   "auth":{
      "basic":{
         "user":{
            "name":"test",
            "key":"test"
         }
      }
   }
}
`
	parser := &JSONConfigParser{}
	config, err := parser.Parse(jsonStr)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(&v1alpha1.NatsSource{
		URL:     "nats://localhost:4222",
		Subject: "test",
		Queue:   "test",
		TLS: &v1alpha1.TLS{
			InsecureSkipVerify: true,
		},
		Auth: &v1alpha1.NatsAuth{
			Basic: &v1alpha1.BasicAuth{
				User: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "test",
					},
					Key: "test",
				},
			},
		},
	}, config))
}
