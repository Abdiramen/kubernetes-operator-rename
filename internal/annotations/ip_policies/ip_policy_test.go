package ip_policies

import (
	"testing"

	ingressv1alpha1 "github.com/Abdiramen/kubernetes-operator-rename/api/ingress/v1alpha1"
	"github.com/Abdiramen/kubernetes-operator-rename/internal/annotations/parser"
	"github.com/Abdiramen/kubernetes-operator-rename/internal/annotations/testutil"
	"github.com/stretchr/testify/assert"
)

func TestParsesIPPolicies(t *testing.T) {
	ing := testutil.NewIngress()
	annotations := map[string]string{}
	annotations[parser.GetAnnotationWithPrefix("ip-policies")] = "abcd1234,some-test-policy"
	ing.SetAnnotations(annotations)

	policies, err := NewParser().Parse(ing)
	assert.NoError(t, err)
	assert.NotNil(t, policies)
	assert.Equal(t, policies, &ingressv1alpha1.EndpointIPPolicy{
		IPPolicies: []string{"abcd1234", "some-test-policy"},
	})
}
