package compression

import (
	"testing"

	ingressv1alpha1 "github.com/Abdiramen/kubernetes-operator-rename/api/ingress/v1alpha1"
	"github.com/Abdiramen/kubernetes-operator-rename/internal/annotations/parser"
	"github.com/Abdiramen/kubernetes-operator-rename/internal/annotations/testutil"
	"github.com/Abdiramen/kubernetes-operator-rename/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestCompressionWhenNotSupplied(t *testing.T) {
	ing := testutil.NewIngress()
	ing.SetAnnotations(map[string]string{})
	parsed, err := NewParser().Parse(ing)

	assert.Nil(t, parsed)
	assert.Error(t, err)
	assert.True(t, errors.IsMissingAnnotations(err))
}

func TestCompressionWhenSuppliedAndTrue(t *testing.T) {
	ing := testutil.NewIngress()
	annotations := map[string]string{}
	annotations[parser.GetAnnotationWithPrefix("https-compression")] = "true"
	ing.SetAnnotations(annotations)

	parsed, err := NewParser().Parse(ing)
	assert.NoError(t, err)

	compression, ok := parsed.(*ingressv1alpha1.EndpointCompression)
	if !ok {
		t.Fatalf("expected *ingressv1alpha1.EndpointCompression, got %T", parsed)
	}
	assert.Equal(t, true, compression.Enabled)
}

func TestCompressionWhenSuppliedAndFalse(t *testing.T) {
	ing := testutil.NewIngress()
	annotations := map[string]string{}
	annotations[parser.GetAnnotationWithPrefix("https-compression")] = "false"
	ing.SetAnnotations(annotations)

	parsed, err := NewParser().Parse(ing)
	assert.NoError(t, err)

	compression, ok := parsed.(*ingressv1alpha1.EndpointCompression)
	if !ok {
		t.Fatalf("expected *ingressv1alpha1.EndpointCompression, got %T", parsed)
	}
	assert.Equal(t, false, compression.Enabled)
}
