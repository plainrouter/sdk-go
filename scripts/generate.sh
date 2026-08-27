#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUTPUT="${1:-$ROOT}"
GENERATED_OUTPUT="$OUTPUT/openapi"
PYTHON_BIN="${PYTHON_BIN:-python3}"
JAVA_BIN="${JAVA_BIN:-java}"
GOFMT_BIN="${GOFMT_BIN:-gofmt}"
GENERATOR_VERSION="7.25.0"
GENERATOR_SHA256="41ce4f6b07f196676439d710759fa1ced7a08066d06ff1bf314681470289efae"
GENERATOR_CACHE="${TMPDIR:-/tmp}/plainrouter-openapi-generator"
GENERATOR_JAR="$GENERATOR_CACHE/openapi-generator-cli-$GENERATOR_VERSION.jar"

mkdir -p "$GENERATOR_CACHE" "$GENERATED_OUTPUT"
cp "$ROOT/.openapi-generator-ignore" "$GENERATED_OUTPUT/.openapi-generator-ignore"
if [[ ! -f "$GENERATOR_JAR" ]]; then
  curl --fail --location --silent --show-error \
    "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$GENERATOR_VERSION/openapi-generator-cli-$GENERATOR_VERSION.jar" \
    --output "$GENERATOR_JAR"
fi

ACTUAL_SHA256="$(shasum -a 256 "$GENERATOR_JAR" | awk '{print $1}')"
if [[ "$ACTUAL_SHA256" != "$GENERATOR_SHA256" ]]; then
  echo "OpenAPI Generator checksum mismatch: $ACTUAL_SHA256" >&2
  exit 1
fi

WORK="$(mktemp -d "${TMPDIR:-/tmp}/plainrouter-go-generate.XXXXXX")"
trap 'rm -rf "$WORK"' EXIT
NORMALIZED_SPEC="$WORK/openapi.normalized.json"

"$PYTHON_BIN" "$ROOT/scripts/verify_spec.py"
"$PYTHON_BIN" "$ROOT/scripts/normalize_openapi.py" "$ROOT/spec/openapi.json" "$NORMALIZED_SPEC"
PACKAGE_VERSION="$("$PYTHON_BIN" -c 'import json,sys; print(json.load(open(sys.argv[1]))["info"]["version"])' "$ROOT/spec/openapi.json")"

"$JAVA_BIN" -jar "$GENERATOR_JAR" generate \
  -g go \
  -i "$NORMALIZED_SPEC" \
  -o "$GENERATED_OUTPUT" \
  --git-user-id plainrouter \
  --git-repo-id sdk-go/openapi \
  --ignore-file-override "$ROOT/.openapi-generator-ignore" \
  --global-property=models,apis,apiTests=false,modelTests=false,supportingFiles=configuration.go:client.go:response.go:utils.go \
  --additional-properties="packageName=openapi,packageVersion=$PACKAGE_VERSION,withGoMod=false,enumClassPrefix=true,enumUnknownDefaultCase=true,disallowAdditionalPropertiesIfNotPresent=false"

find "$GENERATED_OUTPUT" -maxdepth 1 -name '*.go' -print0 | xargs -0 "$GOFMT_BIN" -w
