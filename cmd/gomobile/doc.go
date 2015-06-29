// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DO NOT EDIT. GENERATED BY 'gomobile help documentation'.

/*
Gomobile is a tool for building and running mobile apps written in Go.

To install:

	$ go get golang.org/x/mobile/cmd/gomobile
	$ gomobile init

At least Go 1.5 is required. Until it is released, build tip from
source: http://golang.org/doc/install/source

Initialization rebuilds the standard library and may download
the Android NDK compiler.

Usage:

	gomobile command [arguments]

Commands:

	bind        build a shared library for android APK and iOS app
	build       compile android APK and iOS app
	init        install android compiler toolchain
	install     compile android APK and iOS app and install on device

Use 'gomobile help [command]' for more information about that command.

NOTE: iOS support is not ready yet.


Build a shared library for android APK and iOS app

Usage:

	gomobile bind [-target android|ios] [-o output] [build flags] [package]

Bind generates language bindings for the package named by the import
path, and compiles a library for the named target system.

The -target flag takes a target system name, either android (the
default) or ios.

For -target android, the bind command produces an AAR (Android ARchive)
file that archives the precompiled Java API stub classes, the compiled
shared libraries, and all asset files in the /assets subdirectory under
the package directory. The output is named '<package_name>.aar' by
default. This AAR file is commonly used for binary distribution of an
Android library project and most Android IDEs support AAR import. For
example, in Android Studio (1.2+), an AAR file can be imported using
the module import wizard (File > New > New Module > Import .JAR or
.AAR package), and setting it as a new dependency
(File > Project Structure > Dependencies).  This requires 'javac'
(version 1.7+) and Android SDK (API level 9 or newer) to build the
library for Android. The environment variable ANDROID_HOME must be set
to the path to Android SDK.

For -target ios, gomobile must be run on an OS X machine with Xcode
installed. Support is not complete.

The -v flag provides verbose output, including the list of packages built.

The build flags -a, -i, -n, -x, and -tags are shared with the build command.
For documentation, see 'go help build'.


Compile android APK and iOS app

Usage:

	gomobile build [-target android|ios] [-o output] [build flags] [package]

Build compiles and encodes the app named by the import path.

The named package must define a main function.

The -target flag takes a target system name, either android (the
default) or ios.

For -target android, if an AndroidManifest.xml is defined in the
package directory, it is added to the APK output. Otherwise, a default
manifest is generated.

For -target ios, gomobile must be run on an OS X machine with Xcode
installed. Support is not complete.

If the package directory contains an assets subdirectory, its contents
are copied into the output.

The -o flag specifies the output file name. If not specified, the
output file name depends on the package built.

The -v flag provides verbose output, including the list of packages built.

The build flags -a, -i, -n, -x, and -tags are shared with the build command.
For documentation, see 'go help build'.


Install android compiler toolchain

Usage:

	gomobile init [-u]

Init downloads and installs the Android C++ compiler toolchain.

The toolchain is installed in $GOPATH/pkg/gomobile.
If the Android C++ compiler toolchain already exists in the path,
it skips download and uses the existing toolchain.

The -u option forces download and installation of the new toolchain
even when the toolchain exists.


Compile android APK and iOS app and install on device

Usage:

	gomobile install [-target android] [build flags] [package]

Install compiles and installs the app named by the import path on the
attached mobile device.

Only -target android is supported. The 'adb' tool must be on the PATH.

The build flags -a, -i, -n, -x, and -tags are shared with the build command.
For documentation, see 'go help build'.
*/
package main
