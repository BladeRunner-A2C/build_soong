// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"fmt"
	"os"
	"path/filepath"
)

// ModuleOutDir returns the path to the module-specific output directory.
func ModuleOutDir(ctx AndroidModuleContext) string {
	return filepath.Join(ctx.AConfig().IntermediatesDir(),
		ctx.ModuleDir(), ctx.ModuleName(), ctx.ModuleSubDir())
}

// ModuleSrcDir returns the path of the directory that all source file paths are
// specified relative to.
func ModuleSrcDir(ctx AndroidModuleContext) string {
	return filepath.Join(ctx.AConfig().SrcDir(), ctx.ModuleDir())
}

// ModuleBinDir returns the path to the module- and architecture-specific binary
// output directory.
func ModuleBinDir(ctx AndroidModuleContext) string {
	return filepath.Join(ModuleOutDir(ctx), "bin")
}

// ModuleLibDir returns the path to the module- and architecture-specific
// library output directory.
func ModuleLibDir(ctx AndroidModuleContext) string {
	return filepath.Join(ModuleOutDir(ctx), "lib")
}

// ModuleGenDir returns the module directory for generated files
// path.
func ModuleGenDir(ctx AndroidModuleContext) string {
	return filepath.Join(ModuleOutDir(ctx), "gen")
}

// ModuleObjDir returns the module- and architecture-specific object directory
// path.
func ModuleObjDir(ctx AndroidModuleContext) string {
	return filepath.Join(ModuleOutDir(ctx), "obj")
}

// ModuleGoPackageDir returns the module-specific package root directory path.
// This directory is where the final package .a files are output and where
// dependent modules search for this package via -I arguments.
func ModuleGoPackageDir(ctx AndroidModuleContext) string {
	return filepath.Join(ModuleOutDir(ctx), "pkg")
}

// ModuleIncludeDir returns the module-specific public include directory path.
func ModuleIncludeDir(ctx AndroidModuleContext) string {
	return filepath.Join(ModuleOutDir(ctx), "include")
}

// ModuleProtoDir returns the module-specific public proto include directory path.
func ModuleProtoDir(ctx AndroidModuleContext) string {
	return filepath.Join(ModuleOutDir(ctx), "proto")
}

func ModuleJSCompiledDir(ctx AndroidModuleContext) string {
	return filepath.Join(ModuleOutDir(ctx), "js")
}

// CheckModuleSrcDirsExist logs an error on a property if any of the directories relative to the
// Blueprints file don't exist.
func CheckModuleSrcDirsExist(ctx AndroidModuleContext, dirs []string, prop string) {
	for _, dir := range dirs {
		fullDir := filepath.Join(ModuleSrcDir(ctx), dir)
		if _, err := os.Stat(fullDir); err != nil {
			if os.IsNotExist(err) {
				ctx.PropertyErrorf(prop, "module source directory %q does not exist", dir)
			} else {
				ctx.PropertyErrorf(prop, "%s", err.Error())
			}
		}
	}
}

// CheckModuleSrcDirsExist logs an error on a property if any of the directories relative to the
// top of the source tree don't exist.
func CheckSrcDirsExist(ctx AndroidModuleContext, dirs []string, prop string) {
	for _, dir := range dirs {
		fullDir := filepath.Join(ctx.AConfig().SrcDir(), dir)
		if _, err := os.Stat(fullDir); err != nil {
			if os.IsNotExist(err) {
				ctx.PropertyErrorf(prop, "top-level source directory %q does not exist", dir)
			} else {
				ctx.PropertyErrorf(prop, "%s", err.Error())
			}
		}
	}
}

// Returns a path relative to the top level source directory.  Panics if path is not inside the
// top level source directory.
func SrcDirRelPath(ctx AndroidModuleContext, path string) string {
	srcDir := ctx.AConfig().SrcDir()
	relPath, err := filepath.Rel(srcDir, path)
	if err != nil {
		panic(fmt.Errorf("%q is not inside %q: %s", path, srcDir, err.Error()))
	}

	return relPath
}