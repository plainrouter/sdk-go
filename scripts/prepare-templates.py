#!/usr/bin/env python3
"""Extend pinned upstream templates without forking their unrelated contents."""
import argparse
from pathlib import Path
from zipfile import ZipFile


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("jar", type=Path)
    parser.add_argument("output", type=Path)
    args = parser.parse_args()
    with ZipFile(args.jar) as archive:
        configuration = archive.read("go/configuration.mustache").decode()
        api = archive.read("go/api.mustache").decode()
    field = "\tHTTPClient       *http.Client\n"
    assert configuration.count(field) == 1, "upstream configuration template changed"
    configuration = configuration.replace(field, field + '''    // RequestValidator runs before sending a generated request. The serverURL
    // is the selected operation server, including context overrides and variables.
    // Root constructors install the signed-contract guard here.
    RequestValidator func(request *http.Request, serverURL string) error `json:"-"`
''')
    call = "\tlocalVarHTTPResponse, err := a.client.callAPI(req)"
    assert api.count(call) == 1, "upstream API template changed"
    api = api.replace(call, '''    if validate := a.client.cfg.RequestValidator; validate != nil {
        if err := validate(req, localBasePath); err != nil {
            return {{#returnType}}localVarReturnValue, {{/returnType}}nil, err
        }
    }

''' + call)
    (args.output / "configuration.mustache").write_text(configuration)
    (args.output / "api.mustache").write_text(api)


if __name__ == "__main__":
    main()
