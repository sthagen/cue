// Code generated by gocode.Generate; DO NOT EDIT.

package pkg1

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
	_ "cuelang.org/go/pkg"
)

var cuegenvalMyStruct = cuegenMake("MyStruct", &MyStruct{})

// Validate validates x.
func (x *MyStruct) Validate() error {
	return cuegenCodec.Validate(cuegenvalMyStruct, x)
}

// Complete completes x.
func (x *MyStruct) Complete() error {
	return cuegenCodec.Complete(cuegenvalMyStruct, x)
}

var cuegenvalOtherStruct = cuegenMake("OtherStruct", &OtherStruct{})

// Validate validates x.
func (x *OtherStruct) Validate() error {
	return cuegenCodec.Validate(cuegenvalOtherStruct, x)
}

var cuegenvalString = cuegenMake("String", nil)

// ValidateCUE validates x.
func (x String) ValidateCUE() error {
	return cuegenCodec.Validate(cuegenvalString, x)
}

var cuegenvalSpecialString = cuegenMake("SpecialString", nil)

// ValidateSpecialString validates x.
func ValidateSpecialString(x string) error {
	return cuegenCodec.Validate(cuegenvalSpecialString, x)
}

var cuegenvalPtr = cuegenMake("Ptr", Ptr(nil))

// ValidatePtr validates x.
func ValidatePtr(x Ptr) error {
	return cuegenCodec.Validate(cuegenvalPtr, x)
}

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenInstance.Value().FieldByName(name, true)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 601 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\x94R\xdfo\xd3:\x18\xb5\xd3^\xe9\xd6\u06bd\x12\u007f\x00\u0487\x9fZ4\xd2\x1f\x12B\x8a\x16`\x1b\x03\xedaka\x80\x10\x88\a/\xf9\x9aZs\xed\x908c\x15l\x120\xc6\xfe\xea\x059M\xb6\xc2\xdb\xf2\x92\xa3/9\xdf9>>\xff\x95\xbf<\ua557\x84\x96\xdf\tyT~kQ\xba&un\x85\x8e\xf0\x99\xb0\xc2\xcdi\x8b\xb6_\x19c\xa9Gh{\"\uc32e\x11\xfa\xcfs\xa90\xa7\xe5\x05!\xe4n\xf9\u04e3\xf4\xff\x0f\x1f\xa3\x02\xfd\xa9T5\xf3\x82\xd0\xf2\x9c\x90n\xf9\xa3E\xe9\xbf7\xf3sB=\xda\xde\x17st\x8b\xda\u0550\x11B\xaeZ/\xcbK\xe2QJ\u05e3\x02\x95\u0409o\xb2\xa4\x9f\x98>\xea\xc8\xc4R;\x1c\x99\x18\xfb\x16s\x1b\v+\xfa\xe9Q2\xa4\x94\xdeq\xef~\xe3\u06cf\n\xa4W\xde\xfbTDG\"Ap\x1f\x19\x93\xf3\xd4d\x16\xba\xac\xc3o\xb1}\xc4Y\x87\xe76\x93:\xc9\x1d\x9c\v;\xe3\xac\xc7\xd8\xde\xe2\xc0fEd\x03\xf8\xc2:\x9b\x01\xc0F8\x1c\xb0\xceV\x00\x10\x9e\xf1HX\x0e_\xe1>\x8fM\xc2Yg\xfc$\x80\xb1\x9da\xb6\xe4\xb0\xcen\x00\xce\xd6\xc8\u07ed\\\xed!;\x85\xa7\x89\xe9\xaeGf\x9e*\xb4\x18n\u05e0\xc7V\x88\x8dXm\xc8\xdf6\xda\n\xa9\xf3M\xbd\xe8\xf2w\xbc\xc7:\x93`\xb9w\"\xa3#\xb7\x95\x1dT\xbf\x06P?\xf7B\xce\x1b\\\t\x1e\v%ca1|[\x83\xed7;=v\x90b$\x85j\xc8\xe1\x19\u03d7\x13\xbed\xd9E\x8a\xe1\xd2E\x8f\xed&\xdad\xf8z&\xf3J&<\xe3Sc8\x1b\u03e5\xbd\xd6\x05\x90\xdaV\xdc\a=\xc6\xfa}\xd87z\xe7D\xe6V\xea\x04>K\xa5\xe0\x10\xc1\u0325\xb5\x18\x83\xc8\xc1\x9d\x19A\xe6\xa0\r\xe0\xa7B\x1e\v\x85\xda\xc2\v\x03N\xdag+\xf4*\x94\xad&\x94:\xc8ZEV\u05a0\xd0x\xe2r\xc6\x18\n\xad0\xcf\x01OR%#i\xd5\x02P\x8bC\x85\xb1\u03e6\xc6\x04\xce&\x9b\u062cI\xda\u0777\xbfW(+S\x85\xe3iw8\xe8\xb1SF\x88w\x9b\x8a\x8e\ua28e\xfe\xac\xa8X)\xe8\u8ea07mcM9\x1a3\x1b\xc3\xc1`\xe5\xa8\u007f\u077f8\x8c\xb83\xb7\xbc\xfa\x00\x1e?d\x84\xfc\x0e\x00\x00\xff\xff\xdb\r\u0320\xe1\x03\x00\x00")
