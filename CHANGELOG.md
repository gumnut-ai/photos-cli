# Changelog

## 0.2.0 (2026-04-18)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/gumnut-ai/photos-cli/compare/v0.1.0...v0.2.0)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([681f93f](https://github.com/gumnut-ai/photos-cli/commit/681f93f94d1446505d9f1a95a4be958a8a4cbd09))
* **api:** api update ([87549bc](https://github.com/gumnut-ai/photos-cli/commit/87549bc9657bba0b91f8cf4cec26db56a9f9fc90))
* **api:** api update ([5871806](https://github.com/gumnut-ai/photos-cli/commit/5871806f9e27b51d6a648e9ee2f3e5dde6dba376))
* **api:** api update ([a45e1c5](https://github.com/gumnut-ai/photos-cli/commit/a45e1c5b6398c3d269507d11d7359949db1908f0))
* **api:** api update ([bcb76ce](https://github.com/gumnut-ai/photos-cli/commit/bcb76ce5afed5d664b7191a341290dbf50923f0b))
* **api:** api update ([e3059bc](https://github.com/gumnut-ai/photos-cli/commit/e3059bc66535272f6faf9e3325ed043e5586b243))
* **api:** api update ([b329502](https://github.com/gumnut-ai/photos-cli/commit/b329502ffc6ffb11f4c99d49e2d622a432e93533))
* **api:** api update ([ed560a0](https://github.com/gumnut-ai/photos-cli/commit/ed560a0266e8b8883c6d0960c9cfd9a0852498ad))
* **api:** api update ([20ad50f](https://github.com/gumnut-ai/photos-cli/commit/20ad50f6daa748c59c497c6c0023ceefa283620c))
* **api:** api update ([fb2b190](https://github.com/gumnut-ai/photos-cli/commit/fb2b190f8c6285bf374b13c4fe580bb9e67ccc9e))
* **api:** api update ([e844231](https://github.com/gumnut-ai/photos-cli/commit/e844231c022395a0ec239b1f8cb1189cf420120d))
* **api:** api update ([0b8068b](https://github.com/gumnut-ai/photos-cli/commit/0b8068b3bcc47c113f9432bfedefaa62da5b665a))
* **api:** api update ([a70b12a](https://github.com/gumnut-ai/photos-cli/commit/a70b12a130e3980e6f8eb424c48a7fa0211dc674))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([c5223bf](https://github.com/gumnut-ai/photos-cli/commit/c5223bf94d73a715c54a07df8496a542a1559d8a))
* binary-only parameters become CLI flags that take filenames only ([1c16d39](https://github.com/gumnut-ai/photos-cli/commit/1c16d3956cdf2462ea3552522996794abb64e1b7))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([bf0a483](https://github.com/gumnut-ai/photos-cli/commit/bf0a483e6a15404dc423605a4c8319f149c2a3cf))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([d3bedfd](https://github.com/gumnut-ai/photos-cli/commit/d3bedfdde90d0ea5e2f9a9cf7ae8e3241538eb1b))
* **cli:** send filename and content type when reading input from files ([4f3197c](https://github.com/gumnut-ai/photos-cli/commit/4f3197cd7c078d55ee8b06b626d58076ba98596f))
* set CLI flag constant values automatically where `x-stainless-const` is set ([dfe4711](https://github.com/gumnut-ai/photos-cli/commit/dfe47110b410638a7783567fe79cbf549154ca79))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([3291dac](https://github.com/gumnut-ai/photos-cli/commit/3291dac5fb0db279501a57d9ea50d49e202302f5))
* better support passing client args in any position ([b6cd297](https://github.com/gumnut-ai/photos-cli/commit/b6cd297fff00867edbb25ab475f3b5359ce514e5))
* cli no longer hangs when stdin is attached to a pipe with empty input ([f69f33b](https://github.com/gumnut-ai/photos-cli/commit/f69f33b0bc47df60203c45a36f39cc80ba4fdc42))
* fall back to main branch if linking fails in CI ([756dbba](https://github.com/gumnut-ai/photos-cli/commit/756dbba5d2bea7e809d0022f556150af44d45264))
* fix for failing to drop invalid module replace in link script ([d99b4c2](https://github.com/gumnut-ai/photos-cli/commit/d99b4c2e5fc375ff3cdc0bcd2dbe6937ca05b942))
* fix for off-by-one error in pagination logic ([c48b8cb](https://github.com/gumnut-ai/photos-cli/commit/c48b8cbf28ffee9f081f4664c5ad9557198c1e6a))
* fix quoting typo ([d92c3e7](https://github.com/gumnut-ai/photos-cli/commit/d92c3e7e4bf3fb6f8d734a32d9f758ca75faa778))
* handle empty data set using `--format explore` ([87b0307](https://github.com/gumnut-ai/photos-cli/commit/87b03074eaa72ffca52ca5d0b2b5856f8dbc9fd1))
* improve linking behavior when developing on a branch not in the Go SDK ([39fbe1b](https://github.com/gumnut-ai/photos-cli/commit/39fbe1b06e336a0cc81bb2ee153db5489ea583cc))
* improved workflow for developing on branches ([98da762](https://github.com/gumnut-ai/photos-cli/commit/98da7625c2e9538b8f983f01f2d1790e57faf3ad))
* no longer require an API key when building on production repos ([e4c4eef](https://github.com/gumnut-ai/photos-cli/commit/e4c4eefa77522cf15c3f943dffc83ad59d05f511))
* only set client options when the corresponding CLI flag or env var is explicitly set ([e55ed16](https://github.com/gumnut-ai/photos-cli/commit/e55ed16be0f990804ba42ddc9b08ec2c1beab21f))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([c81f1cc](https://github.com/gumnut-ai/photos-cli/commit/c81f1cce262a40ed8c640a46502545f8ed9d0433))


### Chores

* add documentation for ./scripts/link ([02744e6](https://github.com/gumnut-ai/photos-cli/commit/02744e60342c447ab1986c3696d9866dfd93f103))
* **ci:** skip lint on metadata-only changes ([c5a6a17](https://github.com/gumnut-ai/photos-cli/commit/c5a6a17f280c683811b92ea926b5be4744f9fc85))
* **ci:** support manually triggering release workflow ([d8a4441](https://github.com/gumnut-ai/photos-cli/commit/d8a4441c74795b1b386ca062c27d0bd4cd986d5d))
* **cli:** additional test cases for `ShowJSONIterator` ([2ca504e](https://github.com/gumnut-ai/photos-cli/commit/2ca504e9bc5e9b6c1bd5d1bd9bfc9b5e9f3e0249))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([25bc4d7](https://github.com/gumnut-ai/photos-cli/commit/25bc4d7a4c58b1d3e8179e0de0888848d971b1cc))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([1d8722e](https://github.com/gumnut-ai/photos-cli/commit/1d8722e8b1a16e47bb3bde45709c6150aebcf5a7))
* **cli:** switch long lists of positional args over to param structs ([f9ebade](https://github.com/gumnut-ai/photos-cli/commit/f9ebadec9293d112036bbeaa759ad62564fa8b33))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([751097d](https://github.com/gumnut-ai/photos-cli/commit/751097d9db557f3612fc18b46d2fa7f81dcc8791))
* **internal:** codegen related update ([7bb0009](https://github.com/gumnut-ai/photos-cli/commit/7bb00090bd1cc53129b44b30f225b5c5d89ee4cb))
* **internal:** tweak CI branches ([a62b0ff](https://github.com/gumnut-ai/photos-cli/commit/a62b0ff116f9671e8563b224e3cff956f130b9b0))
* **internal:** update gitignore ([e73e727](https://github.com/gumnut-ai/photos-cli/commit/e73e7279f299f4b8ac2a287994067cd649abf648))
* mark all CLI-related tests in Go with `t.Parallel()` ([bb10e38](https://github.com/gumnut-ai/photos-cli/commit/bb10e3844a7a610837919b96a764cb2183d3bfef))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([e48036a](https://github.com/gumnut-ai/photos-cli/commit/e48036aeb7cdf8cbb3c2e1574e9e12fe33c6e4e2))
* omit full usage information when missing required CLI parameters ([4c44bf0](https://github.com/gumnut-ai/photos-cli/commit/4c44bf0ae8aed0890deeffd2f3588b86d0f9e2d7))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([95bc940](https://github.com/gumnut-ai/photos-cli/commit/95bc9403af65ef1b1c2fc1346065768521ef52f4))

## 0.1.0 (2026-03-13)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/gumnut-ai/photos-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([c7f6e5b](https://github.com/gumnut-ai/photos-cli/commit/c7f6e5b6550e24e790caa4bc9dfc7a4e9e94a64d))


### Chores

* configure new SDK language ([42d56ba](https://github.com/gumnut-ai/photos-cli/commit/42d56ba4ddd189163b0e93fd22579a3f1c6b2a20))
* configure new SDK language ([ba533e0](https://github.com/gumnut-ai/photos-cli/commit/ba533e00b1ef5e8bcfa098b14f754380567aa7cf))
* update SDK settings ([0b1551c](https://github.com/gumnut-ai/photos-cli/commit/0b1551c61c5b5694ac6af13b9cddb46628ecd249))
