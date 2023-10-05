# [2.0.0](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.18...v2.0.0) (2023-10-05)

### Other

- Merge pull request #46 from Seagate/feat/new-apiv2 ([7629499](https://github.com/Seagate/seagate-exos-x-api-go/commit/76294997a1616a7fe0e7590538c1181d04215d5f)), closes [#46](https://github.com/Seagate/seagate-exos-x-api-go/issues/46)
- feat!(apiv2): MC API v2 using OpenAPI Specification ([7141b6b](https://github.com/Seagate/seagate-exos-x-api-go/commit/7141b6b1f6932d7bf26511933197f27fac7c9741))

### BREAKING CHANGE

- Requires changes to all clients to use v1 or v2 of this library.

## [1.0.18](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.17...v1.0.18) (2023-10-04)

### Bug Fixes

- Add a workaround in the client.go template for XML/JSON content-type header bug ([96e9faf](https://github.com/Seagate/seagate-exos-x-api-go/commit/96e9faf91bd1f2419d5e8fd25b40a2e9a7ccceaf))
- include client.go with json/xml templated fix ([ad41c18](https://github.com/Seagate/seagate-exos-x-api-go/commit/ad41c182ec891594527dce85b7c61d416c77a876))
- use 'snapshot' instead of 'snapshots' for object.Name (#44) ([2ce9dce](https://github.com/Seagate/seagate-exos-x-api-go/commit/2ce9dce108216b9b9f47b30a7dd6ab8a2191804e)), closes [#44](https://github.com/Seagate/seagate-exos-x-api-go/issues/44)

### Chores

- add debug logging for forced json contentType ([6b48ee3](https://github.com/Seagate/seagate-exos-x-api-go/commit/6b48ee3ad24b7d407d9cf9a9607b62d9ff744d54))
- **deps:** bump github.com/prometheus/client_golang ([ae4da77](https://github.com/Seagate/seagate-exos-x-api-go/commit/ae4da77b4ac62a2846e7168f6113a63d4f009e7a))
- **deps:** bump google.golang.org/grpc from 1.26.0 to 1.53.0 ([6d18737](https://github.com/Seagate/seagate-exos-x-api-go/commit/6d1873767cbfbcc96cfbf0cef47d409d0448b8f2))

### Other

- Merge pull request #45 from Seagate/feat/final-apiv2 ([3376cdd](https://github.com/Seagate/seagate-exos-x-api-go/commit/3376cdd7e9716d09f003c2a1eb8ff4da4c4b8241)), closes [#45](https://github.com/Seagate/seagate-exos-x-api-go/issues/45)
- feat!: MC API v2 using OpenAPI Specification ([d53b770](https://github.com/Seagate/seagate-exos-x-api-go/commit/d53b77000486b2eeac03e8840f85e4b02a37108d))
- Merge pull request #38 from Seagate/dependabot/go_modules/google.golang.org/grpc-1.53.0 ([e8c4f8c](https://github.com/Seagate/seagate-exos-x-api-go/commit/e8c4f8c48af15d717fd6317eec3b2d87eddee67c)), closes [#38](https://github.com/Seagate/seagate-exos-x-api-go/issues/38)
- Merge branch 'main' into dependabot/go_modules/google.golang.org/grpc-1.53.0 ([e1d170f](https://github.com/Seagate/seagate-exos-x-api-go/commit/e1d170f64d399cc7bfa936b0685cc63e6dc09cfe))
- Resolved merge conflict by updating google.golang.org/grpc v1.53.0 ([2a2678f](https://github.com/Seagate/seagate-exos-x-api-go/commit/2a2678fd9688d9559b8a2c037da89f82f7d04c58))
- Merge pull request #39 from Seagate/dependabot/go_modules/github.com/prometheus/client_golang-1.11.1 ([e56bf7e](https://github.com/Seagate/seagate-exos-x-api-go/commit/e56bf7e479dd12d36896a837bf1b637fe26eef3a)), closes [#39](https://github.com/Seagate/seagate-exos-x-api-go/issues/39)
- Merge pull request #43 from Seagate/bug/correct-status-handling-for-info ([410c4c9](https://github.com/Seagate/seagate-exos-x-api-go/commit/410c4c9b14dca75e4d10120efd5fe4721a5f4f0d)), closes [#43](https://github.com/Seagate/seagate-exos-x-api-go/issues/43)
- remove redundant CreateCommonStatusFromStatus function ([585d5fd](https://github.com/Seagate/seagate-exos-x-api-go/commit/585d5fdb7ab6884238a8e2bb183bd4647d5fe122))
- correct regression test suite based on corrected status responses ([c2d9797](https://github.com/Seagate/seagate-exos-x-api-go/commit/c2d9797b6a199b7cac5ec179ead939ab2eead52f))
- update api v2 to handle status array ([f6ca79c](https://github.com/Seagate/seagate-exos-x-api-go/commit/f6ca79caca3f4a5fb7b7b97a3af5734914b5ccd2))
- generated new openapi client library ([a0b490c](https://github.com/Seagate/seagate-exos-x-api-go/commit/a0b490c633217a4caccadea122bbc80899c568fa))
- update validator to display all status responses ([6b266e3](https://github.com/Seagate/seagate-exos-x-api-go/commit/6b266e31e7c31cb1585735b1e30de0f3f0a91e9c))
- update spec generator and yaml to correct statusResource ([201d011](https://github.com/Seagate/seagate-exos-x-api-go/commit/201d011fc6b879f7c8981d22d1ce8d8a0091dcdf))
- Merge pull request #40 from Seagate/feat/docker-run-generate ([edc62ea](https://github.com/Seagate/seagate-exos-x-api-go/commit/edc62eae29fea168b24f7db2109479b04424b4b5)), closes [#40](https://github.com/Seagate/seagate-exos-x-api-go/issues/40)
- use GENERATOR_VERSION := v6.5.0 ([dafe034](https://github.com/Seagate/seagate-exos-x-api-go/commit/dafe0345e82f7a42810ea9ba465eeb88b08be918))
- use docker image 6.5.0 for openapi-generator-cli ([99a982f](https://github.com/Seagate/seagate-exos-x-api-go/commit/99a982faad0c92c4046b12fc74ca4f752342cd17))
- use docker run openapi-generator-cli:v6.2.1 ([6950b94](https://github.com/Seagate/seagate-exos-x-api-go/commit/6950b9491a4c103b2e7a922edc47036acd54a17c))
- Merge pull request #37 from Seagate/feat/mcapiv2 ([576d4eb](https://github.com/Seagate/seagate-exos-x-api-go/commit/576d4ebd5f505a026d78adc9eb0d00d434f2b44b)), closes [#37](https://github.com/Seagate/seagate-exos-x-api-go/issues/37)
- force Wwn to lowercase in CreateVolume ([f078a6b](https://github.com/Seagate/seagate-exos-x-api-go/commit/f078a6b78650d3357316cd2eab058e44124f428b))
- allow for ip address, protocol, and hostname variations ([f58b70e](https://github.com/Seagate/seagate-exos-x-api-go/commit/f58b70e901c5ca41ae8ccc7f6c3d868579ac2e60))
- accept protocol+ipaddress ([24a0769](https://github.com/Seagate/seagate-exos-x-api-go/commit/24a0769c04610bbcba21f5f482314d1dc3d8192b))
- correct CreateVolume for Linear ([f815102](https://github.com/Seagate/seagate-exos-x-api-go/commit/f815102536de6e845b548f19da3a64b8165e35ca))
- review changes for PR 37 ([ea9d97c](https://github.com/Seagate/seagate-exos-x-api-go/commit/ea9d97c8e914b7aa2db26b926808979a21a7f7be))
- apiv2 with openapi spec, generator, regression, sdk ([5b7fa5c](https://github.com/Seagate/seagate-exos-x-api-go/commit/5b7fa5c0a68439949341e89b6b97a4e57a37ba0f))

## [1.0.17](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.16...v1.0.17) (2023-06-06)

### Chores

- **deps:** bump github.com/prometheus/client_golang ([6c7da1b](https://github.com/Seagate/seagate-exos-x-api-go/commit/6c7da1b6c87a089dea7fa6c20454e61a282c4bf3))

### Other

- Merge pull request #32 from Seagate/dependabot/go_modules/github.com/prometheus/client_golang-1.11.1 ([2fa76d4](https://github.com/Seagate/seagate-exos-x-api-go/commit/2fa76d4777789344e92145cee0653fefced6dcf4)), closes [#32](https://github.com/Seagate/seagate-exos-x-api-go/issues/32)
- Merge branch 'main' into dependabot/go_modules/github.com/prometheus/client_golang-1.11.1 ([e587349](https://github.com/Seagate/seagate-exos-x-api-go/commit/e5873496bf1b9658fe388027aa172275aa1852a6))

## [1.0.16](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.15...v1.0.16) (2023-06-06)

### Chores

- **deps:** bump express from 4.17.1 to 4.17.3 in /mock ([f42a18f](https://github.com/Seagate/seagate-exos-x-api-go/commit/f42a18ff54016c3ccad280a869198015f8952e8b))

### Other

- Merge pull request #30 from Seagate/dependabot/npm_and_yarn/mock/express-4.17.3 ([300dae8](https://github.com/Seagate/seagate-exos-x-api-go/commit/300dae83beef342f809104834e4cf822c32a9530)), closes [#30](https://github.com/Seagate/seagate-exos-x-api-go/issues/30)
- Merge branch 'main' into dependabot/npm_and_yarn/mock/express-4.17.3 ([788c783](https://github.com/Seagate/seagate-exos-x-api-go/commit/788c783083e161bd9ae7e88c6239788e0c5835af))

## [1.0.15](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.14...v1.0.15) (2023-06-06)

### Chores

- **deps:** bump yaml and semantic-release ([209afcd](https://github.com/Seagate/seagate-exos-x-api-go/commit/209afcd0b40c44274b8ea27b7e8a4f81db75d2b0))

### Other

- Merge pull request #34 from Seagate/dependabot/npm_and_yarn/yaml-and-semantic-release--removed ([ff849de](https://github.com/Seagate/seagate-exos-x-api-go/commit/ff849def3ce7e15ed7ab1ce7a98192905cdb648a)), closes [#34](https://github.com/Seagate/seagate-exos-x-api-go/issues/34)

## [1.0.14](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.13...v1.0.14) (2023-06-06)

### Chores

- **deps:** bump golang.org/x/text from 0.3.3 to 0.3.8 ([3e8e462](https://github.com/Seagate/seagate-exos-x-api-go/commit/3e8e4620dd502cee584f52d324f464df31d4e66c))

### Other

- Merge pull request #33 from Seagate/dependabot/go_modules/golang.org/x/text-0.3.8 ([008bb09](https://github.com/Seagate/seagate-exos-x-api-go/commit/008bb09cae4aa6f3b81d7d4fac5a2e7775bafe24)), closes [#33](https://github.com/Seagate/seagate-exos-x-api-go/issues/33)

## [1.0.13](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.12...v1.0.13) (2023-06-06)

### Chores

- **deps:** bump qs and express in /mock ([5a99f7a](https://github.com/Seagate/seagate-exos-x-api-go/commit/5a99f7a0ee3cb6c5f4fd394ebce5804cc1ddf9f6))

### Other

- Merge pull request #28 from Seagate/dependabot/npm_and_yarn/mock/qs-and-express-6.11.0 ([da43e7b](https://github.com/Seagate/seagate-exos-x-api-go/commit/da43e7b6cf2c76c60bf3035c0a6393fc1cc8706e)), closes [#28](https://github.com/Seagate/seagate-exos-x-api-go/issues/28)

## [1.0.12](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.11...v1.0.12) (2022-12-22)

### Bug Fixes

- avoid recursive call to Login() (#27) ([3030dd2](https://github.com/Seagate/seagate-exos-x-api-go/commit/3030dd2fdea1423a997120fb8653a92b89f0f32b)), closes [#27](https://github.com/Seagate/seagate-exos-x-api-go/issues/27)

### Other

- Merge pull request #31 from Seagate/fix_recursive_login ([f617f52](https://github.com/Seagate/seagate-exos-x-api-go/commit/f617f5221b1d436ffa37e9f365d84c83fde04886)), closes [#31](https://github.com/Seagate/seagate-exos-x-api-go/issues/31) [#27](https://github.com/Seagate/seagate-exos-x-api-go/issues/27)

## [1.0.11](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.10...v1.0.11) (2022-12-12)

### Chores

- Update logging to klog2 for compatibility with csi driver ([815552d](https://github.com/Seagate/seagate-exos-x-api-go/commit/815552d29c5d1b445b9295165fb7c6d9b7c210ec))

## [1.0.10](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.9...v1.0.10) (2022-11-23)

### Bug Fixes

- **ShowHostMaps:** Return host maps when initiator has a nickname or is in a group ([9b72e0f](https://github.com/Seagate/seagate-exos-x-api-go/commit/9b72e0ff3de8c30b15c39cd6d46b294dfabd2b3e))

### Chores

- add workflow_dispatch to enable manual workflow run ([35cf7d9](https://github.com/Seagate/seagate-exos-x-api-go/commit/35cf7d9492fbf3bc60d9122773a4f9637066619b))
- remove release.yml changes ([d7686d1](https://github.com/Seagate/seagate-exos-x-api-go/commit/d7686d161cee7ab5ed2d998f02dfefd4b2f79f6f))
- update workflow_dispatch ([b3f01be](https://github.com/Seagate/seagate-exos-x-api-go/commit/b3f01be4189dd22dccdedbd9fde0070d235802fd))

### Other

- Merge pull request #26 from Seagate/chore/manualrelease2 ([c266540](https://github.com/Seagate/seagate-exos-x-api-go/commit/c266540f044528c470a60413ef263f3e21a67f7a)), closes [#26](https://github.com/Seagate/seagate-exos-x-api-go/issues/26)
- Merge pull request #25 from Seagate/chore/manualrelease ([b6d4a6e](https://github.com/Seagate/seagate-exos-x-api-go/commit/b6d4a6e5a31adb828470a8fb0b0dbe65a3f59a67)), closes [#25](https://github.com/Seagate/seagate-exos-x-api-go/issues/25)

## [1.0.8](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.7...v1.0.8) (2022-05-31)

### Bug Fixes

- ChooseLUN uses next highest lun before wrapping ([80883af](https://github.com/Seagate/seagate-exos-x-api-go/commit/80883af1a22a1043f83fe461343eaed57564711f))
- Correctly handle a volumes slice of zero length ([242b3fa](https://github.com/Seagate/seagate-exos-x-api-go/commit/242b3fa540319d8ec1d21200d0a390a5fed275ca))

### Other

- Merge pull request #19 from Seagate/fix/chooselun ([3d1a38b](https://github.com/Seagate/seagate-exos-x-api-go/commit/3d1a38b18ac6440960897e2ae08f06a62a52d86f)), closes [#19](https://github.com/Seagate/seagate-exos-x-api-go/issues/19)

## [1.0.7](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.6...v1.0.7) (2022-05-26)

### Bug Fixes

- CheckVolumeExists() optimization and new GetVolumeWwn() ([c73083a](https://github.com/Seagate/seagate-exos-x-api-go/commit/c73083a51c34135523b83c97b10c79da810fa950))

### Code Style

- GetVolumeWwn() improvement ([5e2424f](https://github.com/Seagate/seagate-exos-x-api-go/commit/5e2424f40cf0fb6561a620c602e490df436f3b94))

### Other

- Merge pull request #18 from Seagate/fix/wwn ([4bbb191](https://github.com/Seagate/seagate-exos-x-api-go/commit/4bbb191320d7abe0f53fc3c9e73eab711788c774)), closes [#18](https://github.com/Seagate/seagate-exos-x-api-go/issues/18)

## [1.0.6](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.5...v1.0.6) (2022-05-26)

### Chores

- **deps:** bump minimist from 1.2.5 to 1.2.6 ([a58ea6f](https://github.com/Seagate/seagate-exos-x-api-go/commit/a58ea6fc323c0eb20349658486b63389a3042460))
- **deps:** bump node-fetch from 2.6.5 to 2.6.7 ([a7bfda6](https://github.com/Seagate/seagate-exos-x-api-go/commit/a7bfda62b438d89ac2c88e378722faa71585b019))

### Other

- Merge pull request #17 from Seagate/dependabot/npm_and_yarn/node-fetch-2.6.7 ([b38d199](https://github.com/Seagate/seagate-exos-x-api-go/commit/b38d199d806d0e2cb257145ec6b8731de816cb08)), closes [#17](https://github.com/Seagate/seagate-exos-x-api-go/issues/17)
- Merge pull request #16 from Seagate/dependabot/npm_and_yarn/minimist-1.2.6 ([a8414ff](https://github.com/Seagate/seagate-exos-x-api-go/commit/a8414ff7b5e04d1065c7cc4d51e992472624e1c6)), closes [#16](https://github.com/Seagate/seagate-exos-x-api-go/issues/16)

## [1.0.5](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.4...v1.0.5) (2021-11-24)

### Bug Fixes

- **GetTargetId:** Retrieve id by type, such as iscsi, fc, etc ([5c8772b](https://github.com/Seagate/seagate-exos-x-api-go/commit/5c8772b9e7eb57b91976c329d037225c45e444ec))

### Other

- Merge pull request #15 from Seagate/fix/get-targetid ([a4f9042](https://github.com/Seagate/seagate-exos-x-api-go/commit/a4f9042a8fb9dea9ef08f903f9af5653fb7786a9)), closes [#15](https://github.com/Seagate/seagate-exos-x-api-go/issues/15)

## [1.0.4](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.3...v1.0.4) (2021-11-22)

### Bug Fixes

- ShowSnapshots() now takes two parameters, snapshot id and source volume id ([539fafe](https://github.com/Seagate/seagate-exos-x-api-go/commit/539fafe16eb06ca33391170986b7d9ec0c20dde0))

### Other

- Merge pull request #14 from Seagate/test/csi-sanity-other ([150a719](https://github.com/Seagate/seagate-exos-x-api-go/commit/150a719963981c13ac19ac2cea37028e558356e3)), closes [#14](https://github.com/Seagate/seagate-exos-x-api-go/issues/14)
- Merge pull request #13 from Seagate/test/csi-sanity-volumes ([42401a3](https://github.com/Seagate/seagate-exos-x-api-go/commit/42401a3c2c03e7b45642ca73ca0e0dce7daf6025)), closes [#13](https://github.com/Seagate/seagate-exos-x-api-go/issues/13)

### Tests

- correct ControllerPublishVolume csi-sanity issues ([7e4e2e5](https://github.com/Seagate/seagate-exos-x-api-go/commit/7e4e2e5dd1f5f23ab2b3b3c7e78b295d0aab2016))

## [1.0.3](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.2...v1.0.3) (2021-10-04)

### Bug Fixes

- correct go.mod path ([b470e32](https://github.com/Seagate/seagate-exos-x-api-go/commit/b470e328841368ab49213bf9159f043b5c14cb09))

### Other

- Merge branch 'main' of github.com:Seagate/seagate-exos-x-api-go into main ([b539b0d](https://github.com/Seagate/seagate-exos-x-api-go/commit/b539b0d76e5358f8b689559b0decd8ecfc3615c8))

## [1.0.2](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.1...v1.0.2) (2021-10-03)

### Chores

- trim changelog ([6fdf232](https://github.com/Seagate/seagate-exos-x-api-go/commit/6fdf232ba2ad7bbaa1012be5e9926c0ad7491aa0))

## [1.0.1](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.0...v1.0.1) (2021-10-03)

### Chores

- correct license in package.json ([5d6b648](https://github.com/Seagate/seagate-exos-x-api-go/commit/5d6b648cd8f63675fbc6dcd63c1a6727ca8b180b))

# 1.0.0 (2021-10-03)

### Bug Fixes

- **fix:** Discover iSCSI IQN and Portals from storage appliance, removes StorageClass requirement ([3e8bcc4](https://github.com/Seagate/seagate-exos-x-api-go/commit/3e8bcc4755fc411100511f596237680193d1fa34))
- **fix:** handle session timeout error code (2) ([dad3e24](https://github.com/Seagate/seagate-exos-x-api-go/commit/dad3e240b25060ccda74b9b36f01fd759d0346ed))
