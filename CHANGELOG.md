## [0.24.1](https://github.com/nestoca/joy/compare/v0.24.0...v0.24.1) (2024-02-06)


### Bug Fixes

* **PL-2353:** CODEOWNERS ([#77](https://github.com/nestoca/joy/issues/77)) ([f66bf24](https://github.com/nestoca/joy/commit/f66bf2458760641e26f562a8f98b7bef1738ea2e))

## [0.24.0](https://github.com/nestoca/joy/compare/v0.23.1...v0.24.0) (2024-02-05)


### Features

* **PL-2314:** add joy release render command ([eb577b0](https://github.com/nestoca/joy/commit/eb577b0c9e24823024e68f745567d1d8f3ccd33c))

## [0.23.1](https://github.com/nestoca/joy/compare/v0.23.0...v0.23.1) (2024-02-02)


### Bug Fixes

* **PL-2290:** do not emit newline at end of sealed secret in non-ttl mode ([24efcf1](https://github.com/nestoca/joy/commit/24efcf1dc8bdd1fbf21aa9d35995cd4977994ea6))

## [0.23.0](https://github.com/nestoca/joy/compare/v0.22.0...v0.23.0) (2024-01-29)


### Features

* **PL-2290:** allow secrets to be sealed from stdin ([a461604](https://github.com/nestoca/joy/commit/a4616049b06e0cd12a66f0a696571b0a09fa4e8e))

## [0.22.0](https://github.com/nestoca/joy/compare/v0.21.0...v0.22.0) (2024-01-19)


### Features

* **PL-2250:** add --auto-merge flag for joy release promote ([79b5ccb](https://github.com/nestoca/joy/commit/79b5ccb7596faf6c21aeea340c00d1198e8802b5))


## [0.21.0](https://github.com/nestoca/joy/compare/v0.20.5...v0.21.0) (2024-01-04)


### Features

* **PL-2051:** improve joy release promote UX for promotable changes ([70af0dc](https://github.com/nestoca/joy/commit/70af0dcfd10e44cebf54145b7c1f248ae86c9a71))

## [0.20.5](https://github.com/nestoca/joy/compare/v0.20.4...v0.20.5) (2024-01-03)


### Bug Fixes

* **PL-1906:** gha, fix handling of BREAKING ([7f21db3](https://github.com/nestoca/joy/commit/7f21db390c258380193fd7c9df6843b972b457e0))

## [0.20.4](https://github.com/nestoca/joy/compare/v0.20.3...v0.20.4) (2023-12-21)


### Bug Fixes

* **PL-1894:** fix ci tests ([fa9672e](https://github.com/nestoca/joy/commit/fa9672e1e42bceeff66c6e698c907675cbc91327))



## [0.20.3](https://github.com/nestoca/joy/compare/v0.20.2...v0.20.3) (2023-12-14)


### Bug Fixes

* **PL-2156:** fix promoted filepath to be the same relative to source ([6a77474](https://github.com/nestoca/joy/commit/6a77474084da4c0cdeedd7e0647211cd66439744))



## [0.20.2](https://github.com/nestoca/joy/compare/v0.20.1...v0.20.2) (2023-12-13)


### Bug Fixes

* check kind difference after source locking ([ae0748c](https://github.com/nestoca/joy/commit/ae0748c9d1004d308c42a2d18e548f285a2d62e3))
* locked values of seq or map type should not overwrite empty but non-nil dst values ([2f83575](https://github.com/nestoca/joy/commit/2f8357597ce5c470ae5b6705c991474a7af21ca7))
* merged mapping and sequence top level node ([1571d9c](https://github.com/nestoca/joy/commit/1571d9cd69d3dfc91985ca74d217fa835ae40f68))


### Reverts

* Revert "proposal: new sequence merging strategy" ([376de2b](https://github.com/nestoca/joy/commit/376de2ba84c0b3f737730c9a763ba19fc996e38c))



## [0.20.1](https://github.com/nestoca/joy/compare/v0.20.0...v0.20.1) (2023-11-27)


### Bug Fixes

* **PL-2110:** add lock migrator yaml file extension check ([093d09f](https://github.com/nestoca/joy/commit/093d09fae706a38c3d9b8cc18a61f69aec4ad19b))



# [0.20.0](https://github.com/nestoca/joy/compare/v0.19.1...v0.20.0) (2023-11-22)


### Features

* **PL-2009:** Add annotations and labels for all joy resources ([adbcb4e](https://github.com/nestoca/joy/commit/adbcb4eeed3a3e1dcdc1632c33ca0c638120c9d8))



## [0.19.1](https://github.com/nestoca/joy/compare/v0.19.0...v0.19.1) (2023-11-20)


### Bug Fixes

* goreleaser tagging ([83303cf](https://github.com/nestoca/joy/commit/83303cf6c5089c394014623080acc9f6078f82e7))



# [0.19.0](https://github.com/nestoca/joy/compare/v0.18.3...v0.19.0) (2023-11-20)


### Features

* **PL-2017:** joy diagnose cmd ([586796e](https://github.com/nestoca/joy/commit/586796ef7b7a07db175642306ba436c70ec5437e))



## [0.18.3](https://github.com/nestoca/joy/compare/v0.18.2...v0.18.3) (2023-11-15)


### Bug Fixes

* **help:** run help without invoking pre run checks ([c470a3f](https://github.com/nestoca/joy/commit/c470a3f85ef49673f66d492ebd1a285750a4a24f))



## [0.18.2](https://github.com/nestoca/joy/compare/v0.18.1...v0.18.2) (2023-11-14)


### Bug Fixes

* **README.md:** config file location and name ([17688fe](https://github.com/nestoca/joy/commit/17688fe498f66d10d4b73d62eb72ea91b4065953))



## [0.18.1](https://github.com/nestoca/joy/compare/v0.18.0...v0.18.1) (2023-11-14)


### Bug Fixes

* **PL-2090:** validate yaml tags in release files ([734b036](https://github.com/nestoca/joy/commit/734b03685200c396a653e09e29bc0d8956b65555))



# [0.18.0](https://github.com/nestoca/joy/compare/v0.17.1...v0.18.0) (2023-11-13)


### Features

* **PL-2089:** add min version to joyrc and root validations ([6c4963a](https://github.com/nestoca/joy/commit/6c4963abba55310465e0178512cbe1c5a7ae206f))



## [0.17.1](https://github.com/nestoca/joy/compare/v0.17.0...v0.17.1) (2023-11-08)


### Bug Fixes

* **PL-2050:** bugfix locked scalar value merging not writing its value as TODO ([#45](https://github.com/nestoca/joy/issues/45)) ([baad615](https://github.com/nestoca/joy/commit/baad615a9368ea3d26cb0da683471ae13c65b07e))



# [0.17.0](https://github.com/nestoca/joy/compare/v0.16.3...v0.17.0) (2023-11-08)


### Features

* **PL-1917:** add lock tag support ([20ba644](https://github.com/nestoca/joy/commit/20ba64474d4640fb8894bc9e988c72a0c8606d2e))
* **PL-1917:** add lock tag support ([#43](https://github.com/nestoca/joy/issues/43)) ([9de4a85](https://github.com/nestoca/joy/commit/9de4a85b317848e0ef233e98c5d8c556f3424f03))



## [0.16.3](https://github.com/nestoca/joy/compare/v0.16.2...v0.16.3) (2023-11-06)


### Bug Fixes

* **gh:** not setting proper gh work directory ([#39](https://github.com/nestoca/joy/issues/39)) ([0144aea](https://github.com/nestoca/joy/commit/0144aea0d2f213f1fd721ca9203858a3d9e6e422))



## [0.16.2](https://github.com/nestoca/joy/compare/v0.16.1...v0.16.2) (2023-11-03)


### Bug Fixes

* pin build dependencies ([a25361d](https://github.com/nestoca/joy/commit/a25361da3239659c7c9089d19addff5ff0d0a54a))
* use same import convention as other projects ([39c24c0](https://github.com/nestoca/joy/commit/39c24c0e034d45fde9885a80ff08c306afc8501e)), closes [/github.com/nestoca/joy/pull/35#discussion_r1377915933](https://github.com//github.com/nestoca/joy/pull/35/issues/discussion_r1377915933)



## [0.16.1](https://github.com/nestoca/joy/compare/v0.16.0...v0.16.1) (2023-10-30)


### Bug Fixes

* promoting release to an environment where it's missing ([#34](https://github.com/nestoca/joy/issues/34)) ([992495a](https://github.com/nestoca/joy/commit/992495a1b99222d69b9699ded2d6cd2bc669a33f))



# [0.16.0](https://github.com/nestoca/joy/compare/v0.15.1...v0.16.0) (2023-10-27)


### Features

* **devops-1938:** add joy setup command ([#33](https://github.com/nestoca/joy/issues/33)) ([ba1e134](https://github.com/nestoca/joy/commit/ba1e134db8c41f120f3dfdb5ab4c287371f58728))



## [0.15.1](https://github.com/nestoca/joy/compare/v0.15.0...v0.15.1) (2023-10-27)


### Bug Fixes

* joy sealed-secret seal not resolving/loading catalog properly ([7102123](https://github.com/nestoca/joy/commit/710212385e0a12bdcecd65e51bc1d1fc51555c63))



# [0.15.0](https://github.com/nestoca/joy/compare/v0.14.0...v0.15.0) (2023-10-26)


### Features

* **devops-1926:** move promotion environments selection to `joy release promote` command ([#29](https://github.com/nestoca/joy/issues/29)) ([db1fac5](https://github.com/nestoca/joy/commit/db1fac579ae8507b89634eadceb912baddf1f64a))



# [0.14.0](https://github.com/nestoca/joy/compare/v0.13.0...v0.14.0) (2023-10-24)


### Features

* **devops-1934:** Rename `joy proj/rel people` to `joy proj/rel owners/own` ([#31](https://github.com/nestoca/joy/issues/31)) ([28c9af1](https://github.com/nestoca/joy/commit/28c9af112b95dc9fac8ab8163fb56cd32180eb44))



# [0.13.0](https://github.com/nestoca/joy/compare/v0.12.1...v0.13.0) (2023-10-24)


### Features

* **devops-1920:** add `joy version` command ([#32](https://github.com/nestoca/joy/issues/32)) ([bc729a2](https://github.com/nestoca/joy/commit/bc729a28239f9739f575864364f8a6f38c43900c))



## [0.12.1](https://github.com/nestoca/joy/compare/v0.12.0...v0.12.1) (2023-10-13)


### Bug Fixes

* preserve order of locked elements ([#28](https://github.com/nestoca/joy/issues/28)) ([372806d](https://github.com/nestoca/joy/commit/372806dd381e04458d862c7f30e6d045c745b430))



# [0.12.0](https://github.com/nestoca/joy/compare/v0.11.0...v0.12.0) (2023-10-10)


### Features

* lock, handle both # lock and ## lock ([6b8e53d](https://github.com/nestoca/joy/commit/6b8e53d3af61365429e4a00e02c0e87ceebd32bd))
* make lock case insensitive ([6448ad1](https://github.com/nestoca/joy/commit/6448ad17d680efb05c581fd8d631acc6a2e729fa))



# [0.11.0](https://github.com/nestoca/joy/compare/v0.10.0...v0.11.0) (2023-09-28)


### Features

* **devops-1262:** add `joy pr promote` command ([#22](https://github.com/nestoca/joy/issues/22)) ([d3de039](https://github.com/nestoca/joy/commit/d3de0399da39b5502a12e756e518a906482a2bca))



# [0.10.0](https://github.com/nestoca/joy/compare/v0.9.1...v0.10.0) (2023-09-26)


### Features

* add repository details to project ([#24](https://github.com/nestoca/joy/issues/24)) ([03350d7](https://github.com/nestoca/joy/commit/03350d78658579dbed4510858c80795d9065deb5)), closes [#discussion_r13262817](https://github.com/nestoca/joy/issues/discussion_r13262817)



## [0.9.1](https://github.com/nestoca/joy/compare/v0.9.0...v0.9.1) (2023-09-25)


### Bug Fixes

* **config:** handle missing config file gracefully ([#25](https://github.com/nestoca/joy/issues/25)) ([ae73f97](https://github.com/nestoca/joy/commit/ae73f976acebc123f4e37d38be775cf134c9bcf0))



# [0.9.0](https://github.com/nestoca/joy/compare/v0.8.2...v0.9.0) (2023-09-08)


### Features

* pretending to add a feature ([#21](https://github.com/nestoca/joy/issues/21)) ([62dd5af](https://github.com/nestoca/joy/commit/62dd5af86e34b7971db17fe3b53838b4bce3d44c))



## [0.8.2](https://github.com/nestoca/joy/compare/v0.8.1...v0.8.2) (2023-09-07)


### Bug Fixes

* force version bump to create new release and test ci ([8b23453](https://github.com/nestoca/joy/commit/8b234534f30c53d5cbbe8c66c7e28b998a5a120d))



## [0.8.1](https://github.com/nestoca/joy/compare/v0.8.0...v0.8.1) (2023-09-07)


### Bug Fixes

* **DEVOPS-1879:** Mixed up source and target environments in `joy release promote` ([ab9b9a4](https://github.com/nestoca/joy/commit/ab9b9a4953093fde083a85e4a05937d2d8228d31))
* Selection messages were involuntarily affected by Goland function rename refactoring ([7797ad6](https://github.com/nestoca/joy/commit/7797ad6a6c79f959a291a9797d00fd7f7fdd9f80))



## [0.8.1](https://github.com/nestoca/joy/compare/v0.8.0...v0.8.1) (2023-08-30)


### Bug Fixes

* Selection messages were involuntarily affected by Goland function rename refactoring ([7797ad6](https://github.com/nestoca/joy/commit/7797ad6a6c79f959a291a9797d00fd7f7fdd9f80))



# Changelog

