module github.com/kolikons/label-watch

go 1.16

replace k8s.io/client-go => k8s.io/client-go v0.19.4

replace k8s.io/apimachinery => k8s.io/apimachinery v0.19.4

require (
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.0.0-00010101000000-000000000000
)
