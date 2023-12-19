Fork of [gorm](https://github.com/go-gorm/gorm) used by [cql](https://github.com/FrancoLiberali/cql). This version contains some extra features to the original library that may not be stable, so it should never be used outside of cql:

* Possibility to make a preload together with an update or delete statement ([pr](https://github.com/FrancoLiberali/gorm/pull/2)). The gorm contributors are not interested in adding this feature (see the [pr](https://github.com/go-gorm/gorm/pull/6583)).
* Possibility to make joins inside an update (only supported by mysql) ([pr](https://github.com/FrancoLiberali/gorm/pull/1)). The implementation of this feature is not complete and only works in cql use cases.

# GORM

The fantastic ORM library for Golang, aims to be developer friendly.

[![go report card](https://goreportcard.com/badge/github.com/go-gorm/gorm "go report card")](https://goreportcard.com/report/github.com/go-gorm/gorm)
[![test status](https://github.com/go-gorm/gorm/workflows/tests/badge.svg?branch=master "test status")](https://github.com/go-gorm/gorm/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)

## Overview

* Full-Featured ORM
* Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism, Single-table inheritance)
* Hooks (Before/After Create/Save/Update/Delete/Find)
* Eager loading with `Preload`, `Joins`
* Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
* Context, Prepared Statement Mode, DryRun Mode
* Batch Insert, FindInBatches, Find To Map
* SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, NamedArg, Search/Update/Create with SQL Expr
* Composite Primary Key
* Auto Migrations
* Logger
* Extendable, flexible plugin API: Database Resolver (Multiple Databases, Read/Write Splitting) / Prometheus…
* Every feature comes with tests
* Developer Friendly

## Getting Started

* GORM Guides [https://gorm.io](https://gorm.io)
* Gen Guides [https://gorm.io/gen/index.html](https://gorm.io/gen/index.html)

## Contributing

[You can help to deliver a better GORM, check out things you can do](https://gorm.io/contribute.html)

## Contributors

[Thank you](https://github.com/go-gorm/gorm/graphs/contributors) for contributing to the GORM framework!

## License

© Jinzhu, 2013~time.Now

Released under the [MIT License](https://github.com/go-gorm/gorm/blob/master/LICENSE)
