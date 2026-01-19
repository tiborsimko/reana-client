# Changelog

## [0.95.0](https://github.com/tiborsimko/reana-client/compare/v0.9.4...0.95.0) (2026-01-19)


### ⚠ BREAKING CHANGES

* **python:** drop support for Python 3.6 and 3.7

### Build

* **appimage:** upgrade to Python 3.8.18 ([#704](https://github.com/tiborsimko/reana-client/issues/704)) ([783c17a](https://github.com/tiborsimko/reana-client/commit/783c17a97c265d0d3cfe97857dc414c6bd7c8b11))
* **docker:** create `reana-client` container image ([#710](https://github.com/tiborsimko/reana-client/issues/710)) ([2c99c5d](https://github.com/tiborsimko/reana-client/commit/2c99c5d1bd36e4303885875375085f7d714e8732)), closes [#709](https://github.com/tiborsimko/reana-client/issues/709)
* **docker:** upgrade to Ubuntu 24.04 and Python 3.12 ([#721](https://github.com/tiborsimko/reana-client/issues/721)) ([4662ad6](https://github.com/tiborsimko/reana-client/commit/4662ad6ef875a7914772654f6be776ad91647e54))
* **python:** add minimal `pyproject.toml` ([#725](https://github.com/tiborsimko/reana-client/issues/725)) ([166ced7](https://github.com/tiborsimko/reana-client/commit/166ced7ef826564289e650c82fb5b33d0e6d205a))
* **python:** add support for Python 3.13 ([#732](https://github.com/tiborsimko/reana-client/issues/732)) ([80d88dc](https://github.com/tiborsimko/reana-client/commit/80d88dc1cb33057f1ce2354fc56e3517fd4d1259))
* **python:** add support for Python 3.13 ([#736](https://github.com/tiborsimko/reana-client/issues/736)) ([fd9b944](https://github.com/tiborsimko/reana-client/commit/fd9b9446d58f21cc6e57b343874d55433532c959))
* **python:** add support for Python 3.14 ([#753](https://github.com/tiborsimko/reana-client/issues/753)) ([191859d](https://github.com/tiborsimko/reana-client/commit/191859d83cdd953359ced997e631e55ca633fe5a))
* **python:** bump shared REANA packages as of 2024-11-28 ([#736](https://github.com/tiborsimko/reana-client/issues/736)) ([778df03](https://github.com/tiborsimko/reana-client/commit/778df037dbeb1340478060e7f913dfff7c0235e5))
* **python:** bump shared REANA packages as of 2025-01-09 ([#757](https://github.com/tiborsimko/reana-client/issues/757)) ([1748ef5](https://github.com/tiborsimko/reana-client/commit/1748ef5aa85e76017938d405c4967f4976da70f7))
* **python:** drop support for Python 3.6 and 3.7 ([#715](https://github.com/tiborsimko/reana-client/issues/715)) ([822a538](https://github.com/tiborsimko/reana-client/commit/822a5387acacb40d53ebe951b7f2332e26f0c5dc))
* **python:** relax pathspec version constraint ([#758](https://github.com/tiborsimko/reana-client/issues/758)) ([d131a25](https://github.com/tiborsimko/reana-client/commit/d131a25880a9f97578e6a2c57177bec6a5034d4e))
* **python:** remove deprecated `pytest-runner` ([#725](https://github.com/tiborsimko/reana-client/issues/725)) ([d6232a5](https://github.com/tiborsimko/reana-client/commit/d6232a552d81470803478dda36dd368d2b485ab5))
* **python:** use optional deps instead of `tests_require` ([#725](https://github.com/tiborsimko/reana-client/issues/725)) ([3397c7e](https://github.com/tiborsimko/reana-client/commit/3397c7e9f494180a8c15457084da0a700147d84c))


### Features

* **cli:** add test command ([#724](https://github.com/tiborsimko/reana-client/issues/724)) ([79d0483](https://github.com/tiborsimko/reana-client/commit/79d0483112480570664381b8ffa3c3642b8bc363)), closes [#719](https://github.com/tiborsimko/reana-client/issues/719)
* **etc:** add bash and zsh shell completion files ([#756](https://github.com/tiborsimko/reana-client/issues/756)) ([aa091f0](https://github.com/tiborsimko/reana-client/commit/aa091f0f28de46ba026040458408907ad2ed39e1))
* **worfklow:** enrich list command with share-related arguments ([#692](https://github.com/tiborsimko/reana-client/issues/692)) ([5e4d12b](https://github.com/tiborsimko/reana-client/commit/5e4d12b809542b15b20b00e475ffa22fef4f8de5)), closes [#687](https://github.com/tiborsimko/reana-client/issues/687)
* **workflow:** add `logs --follow` functionality ([#731](https://github.com/tiborsimko/reana-client/issues/731)) ([dda80d1](https://github.com/tiborsimko/reana-client/commit/dda80d15eedb5ac54bd5da112a0e6fb8a7749f26))
* **workflow:** add share-add command ([#692](https://github.com/tiborsimko/reana-client/issues/692)) ([f7b53e7](https://github.com/tiborsimko/reana-client/commit/f7b53e7dd4d5ddbb1e76378afd027271e14a4249)), closes [#680](https://github.com/tiborsimko/reana-client/issues/680)
* **workflow:** add share-remove command ([#692](https://github.com/tiborsimko/reana-client/issues/692)) ([20caca3](https://github.com/tiborsimko/reana-client/commit/20caca3028b2b57afae1f2546c00639c9c298995)), closes [#681](https://github.com/tiborsimko/reana-client/issues/681)
* **workflow:** add share-status command ([#692](https://github.com/tiborsimko/reana-client/issues/692)) ([807ef40](https://github.com/tiborsimko/reana-client/commit/807ef4089dd72a4c46f7f1a05cca0c316115d131)), closes [#686](https://github.com/tiborsimko/reana-client/issues/686)


### Bug fixes

* **cli:** update workflow sharing commands after changes ([#692](https://github.com/tiborsimko/reana-client/issues/692)) ([5d0aca7](https://github.com/tiborsimko/reana-client/commit/5d0aca7c7c14278cc2a67c70197b8f97e99b747f))
* **status:** display correct duration of stopped workflows ([#701](https://github.com/tiborsimko/reana-client/issues/701)) ([b53def8](https://github.com/tiborsimko/reana-client/commit/b53def8dd3246b10d4da0f2367710af0911e284c)), closes [#699](https://github.com/tiborsimko/reana-client/issues/699)


### Code refactoring

* **docs:** move from reST to Markdown ([#703](https://github.com/tiborsimko/reana-client/issues/703)) ([c9c4d53](https://github.com/tiborsimko/reana-client/commit/c9c4d530eb3e1e3d6996fe71821116815c8eaba3))


### Code style

* **black:** format with black v24 ([#702](https://github.com/tiborsimko/reana-client/issues/702)) ([02dc830](https://github.com/tiborsimko/reana-client/commit/02dc83009a6477c1ae045f4e1a6ea9f9e66640fb))
* **python:** adapt code base to the new black version ([#758](https://github.com/tiborsimko/reana-client/issues/758)) ([b6faeda](https://github.com/tiborsimko/reana-client/commit/b6faedab8a9089c3bbf42b91b90f34bf0afa2ffb))


### Test suite

* **snakemake:** allow running Snakemake 7 tests on Python 3.11+ ([#700](https://github.com/tiborsimko/reana-client/issues/700)) ([8ad7ff1](https://github.com/tiborsimko/reana-client/commit/8ad7ff19e98d1f9231af65bf608d408031546a3e)), closes [#655](https://github.com/tiborsimko/reana-client/issues/655)


### Continuous integration

* **actions:** pin setuptools 70 ([#728](https://github.com/tiborsimko/reana-client/issues/728)) ([0a4bcc7](https://github.com/tiborsimko/reana-client/commit/0a4bcc79af33dd00a6a03216be32d10000bb432b))
* **actions:** update GitHub actions due to Node 16 deprecation ([#714](https://github.com/tiborsimko/reana-client/issues/714)) ([3465eb0](https://github.com/tiborsimko/reana-client/commit/3465eb0496dba6b4fb64fd6edd7872ae59ec5a98))
* added flake8 linter ([2ea9df3](https://github.com/tiborsimko/reana-client/commit/2ea9df38ea6777893cd229901599fc4b66bdc573))
* added system dependencies for graphviz support ([9d6d46b](https://github.com/tiborsimko/reana-client/commit/9d6d46b3b9e6e0b4098fe9f0d7d7b17a8f3997b7))
* **commitlint:** addition of commit message linter ([#695](https://github.com/tiborsimko/reana-client/issues/695)) ([2de7d61](https://github.com/tiborsimko/reana-client/commit/2de7d61db96693e8ee9c3ac555aef9dbfd7bb4bc))
* **commitlint:** allow release commit style ([#708](https://github.com/tiborsimko/reana-client/issues/708)) ([f552752](https://github.com/tiborsimko/reana-client/commit/f55275296cd6cc72b4d21d89f51442842cb15d30))
* **commitlint:** check for the presence of concrete PR number ([#698](https://github.com/tiborsimko/reana-client/issues/698)) ([fa5b7c7](https://github.com/tiborsimko/reana-client/commit/fa5b7c76eb25bfb1591e6fae4a142d975e14b937))
* **commitlint:** fix local running of commit linter on macOS ([#746](https://github.com/tiborsimko/reana-client/issues/746)) ([ddf1eea](https://github.com/tiborsimko/reana-client/commit/ddf1eeacd5aa8bca8f6d9729fa21ebb3bf2fbf5d))
* **commitlint:** improve checking of merge commits ([#721](https://github.com/tiborsimko/reana-client/issues/721)) ([67e4da3](https://github.com/tiborsimko/reana-client/commit/67e4da31a42bf7bc20049aa9d32c8ba912413b1e))
* do not use `-I` when installing dependencies ([e717555](https://github.com/tiborsimko/reana-client/commit/e7175559735eb226662c47f67ec72993d16f573d))
* pin ubuntu version in GA jobs ([e68fb9f](https://github.com/tiborsimko/reana-client/commit/e68fb9fc7602d2d1e94b22386a13ce5a9af991df))
* **pytest:** install `tests` package variant instead of `all` ([#703](https://github.com/tiborsimko/reana-client/issues/703)) ([fe0b00a](https://github.com/tiborsimko/reana-client/commit/fe0b00af1ad7b79ec607de7b810f597a3d6df93a))
* **pytest:** invoke `pytest` directly instead of `setup.py test` ([#725](https://github.com/tiborsimko/reana-client/issues/725)) ([c755d8a](https://github.com/tiborsimko/reana-client/commit/c755d8a91abb39ce4ef547b3fa844eb49d0fccb1))
* **release-please:** initial configuration ([#695](https://github.com/tiborsimko/reana-client/issues/695)) ([5b278f1](https://github.com/tiborsimko/reana-client/commit/5b278f131b59d3ecfd3c7f129040a126cd01b60a))
* remove six from python-tests job ([2cb42c5](https://github.com/tiborsimko/reana-client/commit/2cb42c5a6c8c6bedc6c6f8321fc29f7e428ad5fb))
* revert back docs.reana.io CLI API check ([b8d893d](https://github.com/tiborsimko/reana-client/commit/b8d893dc8049f439cc060d78ba0527d333ac9cd7))
* **runners:** upgrade CI runners to Ubuntu 22.04 ([#741](https://github.com/tiborsimko/reana-client/issues/741)) ([22c74e5](https://github.com/tiborsimko/reana-client/commit/22c74e5955f60b3fca900462626d3b948e1ee010))
* **runners:** upgrade CI runners to Ubuntu 22.04 ([#741](https://github.com/tiborsimko/reana-client/issues/741)) ([639992a](https://github.com/tiborsimko/reana-client/commit/639992a3c13a3b601541f0a04c48afd33008ac62))
* **shellcheck:** fix exit code propagation ([#698](https://github.com/tiborsimko/reana-client/issues/698)) ([fe696ea](https://github.com/tiborsimko/reana-client/commit/fe696eae4cef119b29784ab80ec03d3f4cc089ea))
* **tox:** fix collecting code coverage information ([#727](https://github.com/tiborsimko/reana-client/issues/727)) ([cd45ad5](https://github.com/tiborsimko/reana-client/commit/cd45ad5ec13317f7796686b0de4c972ec8b8ca7f))
* update all actions ([c9cb332](https://github.com/tiborsimko/reana-client/commit/c9cb3328b9bd4d219f86907199cff432c619d87a))
* use Python 3.12 stable release ([dcd415f](https://github.com/tiborsimko/reana-client/commit/dcd415fe08b06b486f968086869fc91b0d58483d))


### Documentation

* add .readthedocs.yaml to migrate to RTD v2 ([f2efcc5](https://github.com/tiborsimko/reana-client/commit/f2efcc5196d20a9680c49b946f60f3df542729ee))
* add interactive session screenshot ([2564206](https://github.com/tiborsimko/reana-client/commit/256420656f8201ad5d81295c357b6cf35725b681))
* add script to generate cli api docs ([44811f7](https://github.com/tiborsimko/reana-client/commit/44811f7f8b3f4af5711fb3813182d4d9638437fb))
* added get_workflow_logs function under rest API docs section ([58ae79b](https://github.com/tiborsimko/reana-client/commit/58ae79bee36dd7ba51b6d609e0769d206955ccd9))
* added more functions under rest API docs section ([d36de00](https://github.com/tiborsimko/reana-client/commit/d36de00b2b2f748de0bcc19cbb163c9ed2ec00f2))
* adding commands summary from `--help` ([a72be53](https://github.com/tiborsimko/reana-client/commit/a72be53d54a1027aa0dd69a98a077d145f909459))
* adds serial partial execution example ([88df3fa](https://github.com/tiborsimko/reana-client/commit/88df3fa5554e6e96ed531523c2106bcde03525d6))
* **authors:** complete list of contributors ([#705](https://github.com/tiborsimko/reana-client/issues/705)) ([875997c](https://github.com/tiborsimko/reana-client/commit/875997c06e657d3e19e1af32324127caa2b1a9c5))
* autoload client methods and add missing parameter descriptions ([5fc10fb](https://github.com/tiborsimko/reana-client/commit/5fc10fbba5ab26d6070339e4172ac2847ddfabca))
* **cli:** fix `open` command documentation typo ([#716](https://github.com/tiborsimko/reana-client/issues/716)) ([9a4dfc2](https://github.com/tiborsimko/reana-client/commit/9a4dfc2d5dccf6d3bc0290719f0ad6bb3dc3735a))
* **cli:** fix `open` command documentation typo ([#728](https://github.com/tiborsimko/reana-client/issues/728)) ([c822dd6](https://github.com/tiborsimko/reana-client/commit/c822dd6570d5474e535be83d0ee4beb44ecee85b))
* cmd example fix for cwl partial execution ([3a26ab7](https://github.com/tiborsimko/reana-client/commit/3a26ab77627145014069878a562ecc9e9bc070c9))
* document serial CACHING parameter ([73dd36a](https://github.com/tiborsimko/reana-client/commit/73dd36a31765c4818f7b70e8fd03eb783e74583e))
* harmonise example file names ([0dac57a](https://github.com/tiborsimko/reana-client/commit/0dac57ac52cffaba475b642fc85a22bc73b39c56))
* harmonise REANA_WORKON name in help ([94cab3c](https://github.com/tiborsimko/reana-client/commit/94cab3cdfa7f6edf43a4ff1418336e4c0ad72de7))
* ineractive session userguide + kubernetes update. ([b2e1b1d](https://github.com/tiborsimko/reana-client/commit/b2e1b1d1112b0b9879558c1ab554c4ff0eab7638))
* more advice on access tokens ([d303c58](https://github.com/tiborsimko/reana-client/commit/d303c58d7c09f43aa8fe0c32d9638bc8004b3a27))
* reanmes HTCONDORCERN* vars to CERN* ([eee648b](https://github.com/tiborsimko/reana-client/commit/eee648b23559be56caa00a1c36f637e39e6577f8))
* remove non-autogenerated sections ([f15e4c7](https://github.com/tiborsimko/reana-client/commit/f15e4c705ba1e21ac878be401fe60fb6eaf5fb4e))
* remove reana-cluster references ([9d6976f](https://github.com/tiborsimko/reana-client/commit/9d6976f757093d3a95fce81c7907d6d8c4375b04))
* renaming workflows cmd to list ([ab17a8d](https://github.com/tiborsimko/reana-client/commit/ab17a8d15d42b93e25e6200d21d3822ed0fb3094))
* rendering fix & rephrasing of interactive sessions ([e9e1e2d](https://github.com/tiborsimko/reana-client/commit/e9e1e2dfa9d01690f98fda8e03740977eea248b9))
* rendering fixes ([070f206](https://github.com/tiborsimko/reana-client/commit/070f206f93dbb3ab641329dc61b39bf78b6edd99))
* revert cli api docs autogeneration ([d22d5f3](https://github.com/tiborsimko/reana-client/commit/d22d5f3178b84e0896c9e4191ba1a2a7ca072d86))
* secrets example for HTCondorCERN ([fcee100](https://github.com/tiborsimko/reana-client/commit/fcee100d90609c75b4f9beb2e12b1f41bb669718))
* section for upload and download wrappers ([6dfe48a](https://github.com/tiborsimko/reana-client/commit/6dfe48adb33a9c3dc2dfa6d8801820f01f0e30ff))
* set default language to English ([821931e](https://github.com/tiborsimko/reana-client/commit/821931ef3cc8c2508b3f989d009263c944da772e))
* single-page RTFD outline ([cd9045a](https://github.com/tiborsimko/reana-client/commit/cd9045af6b5a1b708545d9f0d755bdc37992ec6d))
* update changelog ([ac7867c](https://github.com/tiborsimko/reana-client/commit/ac7867c08c519d1a70ccd54a940b45076d3382a7))
* update instructions for latest rc version ([a6df2f2](https://github.com/tiborsimko/reana-client/commit/a6df2f27daf8978866d6a802ed32769029df5be0))
* **usage:** add shell completion instructions ([#756](https://github.com/tiborsimko/reana-client/issues/756)) ([03d2270](https://github.com/tiborsimko/reana-client/commit/03d22708f4ad831d2bc4058f4f75077b0db8988f))
* virtualenv instead of mkvirtualenv ([4422c17](https://github.com/tiborsimko/reana-client/commit/4422c1754fe624ceae5f7c50f904bfb413c22767))


### Chores

* **master:** release 0.95.0a1 ([9a410cc](https://github.com/tiborsimko/reana-client/commit/9a410cc030bef6fb39ea5f0c35a93bbe8730297a))

## [0.9.4](https://github.com/reanahub/reana-client/compare/0.9.3...0.9.4) (2024-11-29)


### Build

* **docker:** create `reana-client` container image ([#710](https://github.com/reanahub/reana-client/issues/710)) ([2c99c5d](https://github.com/reanahub/reana-client/commit/2c99c5d1bd36e4303885875375085f7d714e8732)), closes [#709](https://github.com/reanahub/reana-client/issues/709)
* **python:** add support for Python 3.13 ([#736](https://github.com/reanahub/reana-client/issues/736)) ([fd9b944](https://github.com/reanahub/reana-client/commit/fd9b9446d58f21cc6e57b343874d55433532c959))
* **python:** bump shared REANA packages as of 2024-11-28 ([#736](https://github.com/reanahub/reana-client/issues/736)) ([778df03](https://github.com/reanahub/reana-client/commit/778df037dbeb1340478060e7f913dfff7c0235e5))


### Continuous integration

* **actions:** pin setuptools 70 ([#728](https://github.com/reanahub/reana-client/issues/728)) ([0a4bcc7](https://github.com/reanahub/reana-client/commit/0a4bcc79af33dd00a6a03216be32d10000bb432b))


### Documentation

* **cli:** fix `open` command documentation typo ([#728](https://github.com/reanahub/reana-client/issues/728)) ([c822dd6](https://github.com/reanahub/reana-client/commit/c822dd6570d5474e535be83d0ee4beb44ecee85b))

## [0.9.3](https://github.com/reanahub/reana-client/compare/0.9.2...0.9.3) (2024-03-13)


### Build

* **appimage:** upgrade to Python 3.8.18 ([#704](https://github.com/reanahub/reana-client/issues/704)) ([783c17a](https://github.com/reanahub/reana-client/commit/783c17a97c265d0d3cfe97857dc414c6bd7c8b11))


### Bug fixes

* **status:** display correct duration of stopped workflows ([#701](https://github.com/reanahub/reana-client/issues/701)) ([b53def8](https://github.com/reanahub/reana-client/commit/b53def8dd3246b10d4da0f2367710af0911e284c)), closes [#699](https://github.com/reanahub/reana-client/issues/699)


### Code refactoring

* **docs:** move from reST to Markdown ([#703](https://github.com/reanahub/reana-client/issues/703)) ([c9c4d53](https://github.com/reanahub/reana-client/commit/c9c4d530eb3e1e3d6996fe71821116815c8eaba3))


### Code style

* **black:** format with black v24 ([#702](https://github.com/reanahub/reana-client/issues/702)) ([02dc830](https://github.com/reanahub/reana-client/commit/02dc83009a6477c1ae045f4e1a6ea9f9e66640fb))


### Test suite

* **snakemake:** allow running Snakemake 7 tests on Python 3.11+ ([#700](https://github.com/reanahub/reana-client/issues/700)) ([8ad7ff1](https://github.com/reanahub/reana-client/commit/8ad7ff19e98d1f9231af65bf608d408031546a3e)), closes [#655](https://github.com/reanahub/reana-client/issues/655)


### Continuous integration

* **commitlint:** addition of commit message linter ([#695](https://github.com/reanahub/reana-client/issues/695)) ([2de7d61](https://github.com/reanahub/reana-client/commit/2de7d61db96693e8ee9c3ac555aef9dbfd7bb4bc))
* **commitlint:** allow release commit style ([#708](https://github.com/reanahub/reana-client/issues/708)) ([f552752](https://github.com/reanahub/reana-client/commit/f55275296cd6cc72b4d21d89f51442842cb15d30))
* **commitlint:** check for the presence of concrete PR number ([#698](https://github.com/reanahub/reana-client/issues/698)) ([fa5b7c7](https://github.com/reanahub/reana-client/commit/fa5b7c76eb25bfb1591e6fae4a142d975e14b937))
* **pytest:** install `tests` package variant instead of `all` ([#703](https://github.com/reanahub/reana-client/issues/703)) ([fe0b00a](https://github.com/reanahub/reana-client/commit/fe0b00af1ad7b79ec607de7b810f597a3d6df93a))
* **release-please:** initial configuration ([#695](https://github.com/reanahub/reana-client/issues/695)) ([5b278f1](https://github.com/reanahub/reana-client/commit/5b278f131b59d3ecfd3c7f129040a126cd01b60a))
* **shellcheck:** fix exit code propagation ([#698](https://github.com/reanahub/reana-client/issues/698)) ([fe696ea](https://github.com/reanahub/reana-client/commit/fe696eae4cef119b29784ab80ec03d3f4cc089ea))


### Documentation

* **authors:** complete list of contributors ([#705](https://github.com/reanahub/reana-client/issues/705)) ([875997c](https://github.com/reanahub/reana-client/commit/875997c06e657d3e19e1af32324127caa2b1a9c5))

## 0.9.2 (2023-12-19)

- Changes `validate` command to show detailed errors when the specification file is not a valid YAML file.
- Changes the validation of specification files to show improved validation warnings, which also indicate where unexpected properties are located in the file.
- Fixes `create_workflow_from_json` API function to always load and send the workflow specification to the server.
- Fixes `list` command to accept case-insensitive column names when sorting the returned workflow runs via the `--sort` option.
- Fixes `run` wrapper command for workflows that do not contain `inputs` clause in their specification.

## 0.9.1 (2023-09-27)

- Adds support for Python 3.12.
- Adds `prune` command to delete all intermediate files of a given workflow. Use with care.
- Changes `open` command to inform user about the auto-closure of interactive sessions after a certain inactivity timeout.
- Changes `validate` command to display non-critical validation warnings when checking the REANA specification file.
- Fixes `list` command to correctly list workflows when sorting them by their run number or by the size of their workspace.
- Fixes `du` command help message typo.
- Fixes `validation --environments` command to correctly handle fully-qualified image names.

## 0.9.0 (2023-01-26)

- Adds support for Python 3.11.
- Adds support for `.gitignore` and `.reanaignore` to specify files that should not be uploaded to REANA.
- Adds `retention-rules-list` command to list the retention rules of a workflow.
- Changes REANA specification loading and validation functionalities by porting some of the logic to `reana-commons`.
- Changes `create` and `restart` commands to always upload the REANA specification file.
- Changes `delete` command to always delete the workflow's workspace.
- Changes `delete_workflow` Python API function to always delete the workflow's workspace.
- Changes `download` command to add the possibility to write files to the standard output via `-o -` option.
- Changes `list` command to hide deleted workflows by default.
- Changes `list` command to allow displaying deleted workflows via `--all` and `--show-deleted-runs` options.
- Changes `list` and `status` commands to allow displaying the duration of workflows with the `--include-duration` option.
- Changes `mv` command to allow moving files while a workflow is running.
- Changes `upload` command to prevent uploading symlinks.
- Changes `validation --environment` command to use Docker registry v2 APIs to check that a Docker image exists in DockerHub.
- Fixes `list` command to highlight the workflow specified in `REANA_WORKON` correctly.
- Fixes `secrets-delete` command error message when deleting non existing secrets.
- Fixes `start` command to report failed workflows as errors.
- Fixes `start` and `run` commands to correctly follow the execution of the workflow until termination.
- Fixes `status` command to respect output format provided by the `--format` option.
- Fixes `upload` command to report when input directories are listed under the `files` section in the REANA specification file and vice versa.
- Fixes `validate --environment` command to detect illegal whitespace characters in Docker image names.

## 0.8.1 (2022-02-15)

- Adds support for creating reana-client standalone AppImage executables.
- Adds support for Python 3.10.
- Adds workflow name validation for `create_workflow_from_json()` Python API function.
- Fixes formatting of error messages and sets appropriate exit status codes.

## 0.8.0 (2021-11-24)

- Adds support for running and validating Snakemake workflows.
- Adds support for `outputs.directories` in `reana.yaml` allowing to easily download output directories.
- Adds new command `quota-show` to retrieve information about total CPU and Disk usage and quota limits.
- Adds new command `info` that retrieves general information about the cluster, such as available workspace path settings.
- Changes `validate` command to add the possibility to check the workflow against server capabilities such as desired workspace path via `--server-capabilities` option.
- Changes `list` command to add the possibility to filter by workflow status and search by workflow name via `--filter` option.
- Changes `list` command to add the possibility to filter and display all the runs of a given workflow via `-w` option.
- Changes `list` command to stop including workflow progress and workspace size by default. Please use new options `--include-progress` and `--include-workspace-size` to show this information.
- Changes `list --sessions` command to display the status of interactive sessions.
- Changes `logs` command to display also the start and finish times of individual jobs.
- Changes `ls` command to add the possibility to filter by file name, size and last-modified values via `--filter` option.
- Changes `du` command to add the possibility filter by file name and size via `--filter` option.
- Changes `delete` command to prevent hard-deletion of workflows.
- Changes Yadage workflow specification loading to be done in `reana-commons`.
- Changes CWL workflow engine to `cwltool` version `3.1.20210628163208`.
- Removes support for Python 2.7. Please use Python 3.6 or higher from now on.

## 0.7.5 (2021-07-05)

- Changes workflow validation to display more granular output.
- Changes workflow parameters validation to warn about misused parameters for each step.
- Changes dependencies to unpin six so that client may be installed in more contexts.
- Fixes environment image validation not to test repetitively the same image.
- Fixes `upload_to_server()` Python API function to silently skip uploading in case of none-like inputs.

## 0.7.4 (2021-04-28)

- Adds support of wildcard patterns to `ls` command.
- Adds support of directory download and wildcard patterns to `download` command.
- Changes `list` command to include deleted workflows by default.
- Fixes environment image validation info message where UIDs were switched.

## 0.7.3 (2021-03-24)

- Adds validation of workflow input parameters to the `validate` command.
- Adds optional validation of workflow environment images (`--environments`) to the `validate` command.

## 0.7.2 (2021-01-15)

- Adds support for Python 3.9.
- Fixes exception handling when uploading files.
- Fixes minor code warnings.
- Fixes traling slash issue from user exported REANA_SERVER_URL.

## 0.7.1 (2020-11-10)

- Changes `ping` command output to include REANA client and server version information.
- Fixes `upload` command to properly display errors.

## 0.7.0 (2020-10-20)

- Adds option to `logs` command to filter job logs according to compute backend, docker image, status and step name.
- Adds new `restart` command to restart previously run or failed workflows.
- Adds possibility to specify operational options in the `reana.yaml` of the workflow.
- Fixes user experience by preventing dots as part of the workflow name to avoid confusion with restart runs.
- Changes `du` command output format.
- Changes file loading to optimise CLI performance.
- Changes `logs` command to enhance formatting using marks and colours.
- Changes from Bravado to requests to improve download performance.
- Changes `ping` command to perform user access token validation.
- Changes defaults to accept both `reana.yaml` and `reana.yml` filenames.
- Changes `diff` command to improve output formatting.
- Changes code formatting to respect `black` coding style.
- Changes documentation to single-page layout.

## 0.6.1 (2020-06-09)

- Fixes installation troubles for REANA 0.6.x release series by pinning several
  dependencies.

## 0.6.0 (2019-12-27)

- Introduces user secrets management commands `secrets-add`,
  `secrets-list` and `secrets-delete`.
- Enhances `run` and `create` commands to allow specifying
  workfow via the `--workflow` flag.
- Introduces new command `version` to report client version.
- Fixes `upload` command behaviour for uploading very large files.
- Simplifies `run` command by removing free upload parameters.
- Upgrades `cwltool` to 1.0.20191022103248.
- Disables SSL verification warnings when talking to self-signed server
  certificates.

## 0.5.0 (2019-04-24)

- Introduces new `resources` field in `reana.yaml` specification file
  allowing to declare computing resources needed for workflow runs, such as the
  CVMFS repositories via `cvmfs` subfield.
- Improves `reana-client` embedded command-line documentation (`-help`) by
  grouping commands and providing concrete usage examples for all commands.
- Enhances workflow `start` command allowing to override input parameters
  (`--parameter`) and to specify additional operational options
  (`--option`).
- Introduces new workflow `run` wrapper command that creates workflow, uploads
  its input data and code and starts its execution.
- Introduces new workflow `stop` command for stopping a running workflow.
- Enhances workflow `logs` command output capabilities via new `--json`
  option.
- Introduces new workflow `diff` command for comparing two workflow runs.
- Introduces new workflow `delete` command for deleting one or more workflow
  runs.
- Introduces new session `open` command allowing to run interactive sessions
  such as Jupyter notebook upon workflow workspace.
- Introduces new session `close` command for closing interactive sessions.
- Renames past `workflows` command to `list` allowing to list both workflow
  runs and interactive sessions.
- Introduces new workspace `du` command for checking workspace disk usage.
- Introduces new workspace `mv` command for moving files within workspace.
- Introduces new workspace `rm` command for removing files within workspace.
- Renames past workspace `list` command to `ls` allowing to list workspace
  files. Enhances its output capabilities via new `--format` option.
- Introduces new API function `create_workflow_from_json()` which allows
  developers and third-party systems to create workflows directly from JSON
  specification.

## 0.4.0 (2018-11-07)

- Enhances test suite and increases code coverage.
- Changes license to MIT.

## 0.3.1 (2018-09-25)

- Amends upload and download commands that will now upload/download all the
  files specified in `reana.yaml` in case no arguments are provided.
- Fixes `status` command's JSON output mode.
- Upgrades CWL reference implementation to version `1.0.20180912090223`.
- Renames Serial workflow operational parameter from ``` CACHING``to ``CACHE ```.
- Adds support for Python 3.7.

## 0.3.0 (2018-08-10)

- Adds support for
  [Serial workflows](http://reana-workflow-engine-serial.readthedocs.io/en/latest/).

- CLI refactored to a flat design:

  - `inputs`/`outputs`/`code` removed, everything is a file managed
    : with upload/download/list commands.
  - Removes `workflow` command, workflows are managed with
    : `create`/`start`/`status`.

- Removes `analyes` command, now `validate` is first level command.

- `status` now shows the selected workflow progress and current command on
  verbose mode.

- Requires the usage of an access token to talk to REANA Server.

- Fixes bug when uploading binary files.

- Supports addition of workflow engine parameters when using `start` for
  serial workflows.

- Improves error messages.

## 0.2.0 (2018-04-20)

- Adds support for Common Workflow Language workflows.
- Adds support for persistent user-selected workflow names.
- Enables file and directory input uploading using absolute paths.
- Adds new `status` command to display the current status of the client.
- Reduces verbosity level for commands and improves error messages.

## 0.1.0 (2018-01-30)

- Initial public release.
