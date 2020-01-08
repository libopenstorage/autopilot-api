package k8sutil

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	cgoscheme "k8s.io/client-go/kubernetes/scheme"
)

var (
	// scheme tracks the type registry for the sdk
	// This scheme is used to decode json data into the correct Go type based on the object's GVK
	// All types that the operator watches must be added to this scheme
	scheme      = runtime.NewScheme()
	codecs      = serializer.NewCodecFactory(scheme)
	decoderFunc = decoder
)

func init() {
	// Add the standard kubernetes [GVK:Types] type registry
	// e.g (v1,Pods):&v1.Pod{}
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	cgoscheme.AddToScheme(scheme)
}

// UtilDecoderFunc retrieve the correct decoder from a GroupVersion
// and the schemes codec factory.
type UtilDecoderFunc func(schema.GroupVersion, serializer.CodecFactory) runtime.Decoder

func decoder(gv schema.GroupVersion, codecs serializer.CodecFactory) runtime.Decoder {
	codec := codecs.UniversalDecoder(gv)
	return codec
}

type addToSchemeFunc func(*runtime.Scheme) error

// AddToSDKScheme allows CRDs to register their types with the sdk scheme
func AddToSDKScheme(addToScheme addToSchemeFunc) {
	addToScheme(scheme)
}
