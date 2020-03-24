package geojson_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApMapFactory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GeoJson Suite")
}
