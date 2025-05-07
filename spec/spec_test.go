package spec_test

import (
	"testing"

	"github.com/quiqupltd/wheniwork-client/spec"
	"github.com/stretchr/testify/assert"
)

func TestSpec(t *testing.T) {
	spec := spec.Document()

	assert.NotNil(t, spec)
}
