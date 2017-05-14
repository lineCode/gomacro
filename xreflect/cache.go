/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * cache.go
 *
 *  Created on May 14, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

type Cache struct {
	RebuildDepth int
	TryResolve   func(name, pkgpath string) Type
	ReflectTypes map[reflect.Type]Type
	Pkgs         map[string]*Package
	Importer     types.ImporterFrom
}

func (cfg *Cache) rebuild() bool {
	return cfg.RebuildDepth >= 0
}

func (cfg *Cache) cache(rt reflect.Type, t Type) Type {
	if cfg.ReflectTypes == nil {
		cfg.ReflectTypes = make(map[reflect.Type]Type)
	}
	cfg.ReflectTypes[rt] = t
	return t
}

func (cfg *Cache) NewPackage(path, name string) *Package {
	pkg := cfg.Pkgs[path]
	if pkg == nil {
		pkg = NewPackage(path, name)
		if cfg.Pkgs == nil {
			cfg.Pkgs = make(map[string]*Package)
		}
		cfg.Pkgs[path] = pkg
	}
	return pkg
}

func (cfg *Cache) namedTypeFromImport(rtype reflect.Type) Type {
	t := cfg.namedTypeFromPackageCache(rtype)
	if t != nil {
		return t
	}
	if cfg.Importer == nil {
		cfg.Importer = DefaultImporter()
	}
	pkgpath := rtype.PkgPath()
	pkg, err := cfg.Importer.Import(pkgpath)
	if err != nil || pkg == nil {
		return nil
	}
	if cfg.Pkgs == nil {
		cfg.Pkgs = make(map[string]*Package)
	}
	cfg.Pkgs[pkgpath] = (*Package)(pkg)

	return cfg.namedTypeFromPackage(rtype, pkg)
}

func (cfg *Cache) namedTypeFromPackageCache(rtype reflect.Type) Type {
	pkgpath := rtype.PkgPath()
	pkg := (*types.Package)(cfg.Pkgs[pkgpath])
	if pkg == nil || !pkg.Complete() {
		return nil
	}
	return cfg.namedTypeFromPackage(rtype, pkg)
}

func (cfg *Cache) namedTypeFromPackage(rtype reflect.Type, pkg *types.Package) Type {
	scope := pkg.Scope()
	if scope == nil {
		return nil
	}
	obj := scope.Lookup(rtype.Name())
	if obj == nil {
		return nil
	}
	gtype := obj.Type()
	if gtype == nil {
		return nil
	}
	// debugf("imported named type %v for %v", gtype, rtype)
	t := maketype(gtype, rtype)
	return cfg.cache(rtype, t)
}