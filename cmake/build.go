package cmake

// BuildConfig describe the configuration of a cmake build
var BuildConfig = struct {
	SourceDir string
	BuildDir  string
	Generator string
	Configure []string
	Env       []string
}{
	BuildDir:  "/tmp/c2b-build",
	SourceDir: "/Users/laurent/alice/o2-dev/O2",
	Generator: "Ninja",
	Configure: []string{"-DCMAKE_MODULE_PATH=/Users/laurent/alice/o2-dev/O2/cmake/modules;/Users/laurent/alice/sw/osx_x86-64/FairRoot/latest-clion-o2/share/fairbase/cmake/modules;/Users/laurent/alice/sw/osx_x86-64/FairRoot/latest-clion-o2/share/fairbase/cmake/modules_old",
		"-DFairRoot_DIR=/Users/laurent/alice/sw/osx_x86-64/FairRoot/latest-clion-o2/",
		"-DALICEO2_MODULAR_BUILD=ON",
		"-DPythia6_DIR=/Users/laurent/alice/sw/osx_x86-64/pythia6/latest",
		"-DBOOST_ROOT=/Users/laurent/alice/sw/osx_x86-64/boost/latest-clion-o2/",
		"-DFAIRROOTPATH=/Users/laurent/alice/sw/osx_x86-64/FairRoot/latest-clion-o2/",
		"-DDDS_PATH=/Users/laurent/alice/sw/osx_x86-64/DDS/latest-clion-o2/",
		"-DCMAKE_PREFIX_PATH=/Users/laurent/alice/sw/osx_x86-64/ROOT/latest-clion-o2/cmake",
		"-DVc_DIR=/Users/laurent/alice/sw/osx_x86-64/Vc/latest-clion-o2/lib/cmake/Vc",
		"-DO2_TPCCA_TRACKING_LIB_DIR=/Users/laurent/alice/sw/osx_x86-64/O2HLTCATracking/latest-clion-o2",
		"-DProtobuf_LIBRARY=/Users/laurent/alice/sw/osx_x86-64/protobuf/latest-clion-o2/lib/libprotobuf.dylib",
		"-DProtobuf_LITE_LIBRARY=/Users/laurent/alice/sw/osx_x86-64/protobuf/latest-clion-o2/lib/libprotobuf-lite.dylib",
		"-DProtobuf_PROTOC_LIBRARY=/Users/laurent/alice/sw/osx_x86-64/protobuf/latest-clion-o2/lib/libprotoc.dylib",
		"-DProtobuf_INCLUDE_DIR=/Users/laurent/alice/sw/osx_x86-64/protobuf/latest-clion-o2/include",
		"-DProtobuf_PROTOC_EXECUTABLE=/Users/laurent/alice/sw/osx_x86-64/protobuf/latest-clion-o2/bin/protoc",
		"-DPYTHIA8_INCLUDE_DIR=/Users/laurent/alice/sw/osx_x86-64/pythia/latest/include",
		"-DMS_GSL_INCLUDE_DIR=/Users/laurent/alice/sw/osx_x86-64/ms_gsl/latest-clion-o2/include",
		"-DMonitoring_ROOT=/Users/laurent/alice/sw/osx_x86-64/Monitoring/latest-clion-o2",
		"-DConfiguration_ROOT=/Users/laurent/alice/sw/osx_x86-64/Configuration/latest-clion-o2",
		"-DRAPIDJSON_INCLUDEDIR=/Users/laurent/alice/sw/osx_x86-64/RapidJSON/latest/include",
		"-Dbenchmark_DIR=/Users/laurent/alice/sw/osx_x86-64/googlebenchmark/latest/lib/cmake/benchmark",
		"-DCMAKE_LIBRARY_PATH=/Users/laurent/alice/sw/osx_x86-64/GEANT3/latest/lib;/Users/laurent/alice/sw/osx_x86-64/GEANT4/latest/lib;/Users/laurent/alice/sw/osx_x86-64/GEANT4_VMC/latest/lib",
	},
	Env: []string{"ROOTSYS=/Users/laurent/alice/sw/osx_x86-64/ROOT/latest-clion-o2/",
		"PYTHIA_ROOT=/Users/laurent/alice/sw/osx_x86-64/pythia/latest",
		"PYTHIA8DATA=/Users/laurent/alice/sw/osx_x86-64/pythia/latest/share/Pythia8/xmldoc",
		"LD_LIBRARY_PATH=${HOME}/alice/sw/osx_x86-64/Geant3/latest/lib:${HOME}/alice/sw/osx_x86-64/Geant4/latest/lib:${HOME}/alice/sw/osx_x86-64/Geant4_VMC/latest/lib",
		"ALIBUILD_O2_TESTS=OFF",
	},
}
