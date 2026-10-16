package plainrouter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/plainrouter/sdk-go/openapi"
)

// The guard inspects serialized JSON so AdditionalProperties cannot overwrite a
// typed integer or a closed object after model construction. It never coerces
// values: numbers and numeric strings retain their existing representation.
type contractTransport struct{ next http.RoundTripper }

type requestConstraint struct {
	Method  string                 `json:"method"`
	Path    string                 `json:"path"`
	Schema  map[string]interface{} `json:"schema"`
	pattern *regexp.Regexp
}

var requestConstraints = loadRequestConstraints()

func loadRequestConstraints() []requestConstraint {
	var routes []requestConstraint
	if err := json.Unmarshal([]byte(requestContractJSON), &routes); err != nil {
		panic(err)
	}
	for i := range routes {
		parts := strings.Split(routes[i].Path, "/")
		for j, part := range parts {
			if strings.HasPrefix(part, "{") {
				parts[j] = "[^/]+"
			} else {
				parts[j] = regexp.QuoteMeta(part)
			}
		}
		routes[i].pattern = regexp.MustCompile(strings.Join(parts, "/") + "$")
	}
	return routes
}

func guardConfiguration(configuration *openapi.Configuration) {
	client := configuration.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	next := client.Transport
	if next == nil {
		next = http.DefaultTransport
	}
	if _, guarded := next.(*contractTransport); guarded {
		return
	}
	// Preserve the caller's timeout, redirect policy, cookie jar and transport
	// without changing its shared http.Client.
	guarded := *client
	guarded.Transport = &contractTransport{next: next}
	configuration.HTTPClient = &guarded
}

func (transport *contractTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if request.Body != nil && strings.Contains(request.Header.Get("Content-Type"), "json") {
		for _, route := range requestConstraints {
			if request.Method != route.Method || !route.pattern.MatchString(request.URL.Path) {
				continue
			}
			raw, err := io.ReadAll(request.Body)
			request.Body.Close()
			if err != nil {
				return nil, err
			}
			var value interface{}
			decoder := json.NewDecoder(bytes.NewReader(raw))
			decoder.UseNumber()
			if err := decoder.Decode(&value); err != nil {
				return nil, err
			}
			if err := validateRequestValue(value, route.Schema, ""); err != nil {
				return nil, err
			}
			request.Body = io.NopCloser(bytes.NewReader(raw))
			break
		}
	}
	return transport.next.RoundTrip(request)
}

func constraintApplies(value interface{}, schema map[string]interface{}) bool {
	if choices, ok := schema["enum"].([]interface{}); ok {
		matched := false
		for _, choice := range choices {
			if choice == value {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	if object, ok := value.(map[string]interface{}); ok {
		properties, _ := schema["properties"].(map[string]interface{})
		for key, child := range properties {
			childSchema := child.(map[string]interface{})
			if _, discriminant := childSchema["enum"]; discriminant {
				if field, present := object[key]; present && !constraintApplies(field, childSchema) {
					return false
				}
			}
		}
	}
	kind, _ := schema["type"].(string)
	switch kind {
	case "null":
		return value == nil
	case "string":
		_, ok := value.(string)
		return ok
	case "object":
		_, ok := value.(map[string]interface{})
		return ok
	case "array":
		_, ok := value.([]interface{})
		return ok
	}
	return true
}

func validateRequestValue(value interface{}, schema map[string]interface{}, field string) error {
	if value == nil {
		return nil
	}
	if _, ok := value.(bool); ok {
		integer := schema["type"] == "integer"
		if kinds, ok := schema["type"].([]interface{}); ok {
			for _, kind := range kinds {
				integer = integer || kind == "integer"
			}
		}
		if integer {
			return fmt.Errorf("%s: boolean is not an integer", field)
		}
	}
	for _, union := range []string{"anyOf", "oneOf"} {
		if variants, ok := schema[union].([]interface{}); ok {
			var first error
			for _, variant := range variants {
				branch := variant.(map[string]interface{})
				if !constraintApplies(value, branch) {
					continue
				}
				err := validateRequestValue(value, branch, field)
				if err == nil {
					return nil
				}
				if first == nil {
					first = err
				}
			}
			if first != nil {
				return first
			}
			return fmt.Errorf("%s: no matching request variant", field)
		}
	}
	if variants, ok := schema["allOf"].([]interface{}); ok {
		for _, branch := range variants {
			if err := validateRequestValue(value, branch.(map[string]interface{}), field); err != nil {
				return err
			}
		}
	}
	if object, ok := value.(map[string]interface{}); ok {
		properties, _ := schema["properties"].(map[string]interface{})
		for key, child := range object {
			path := key
			if field != "" {
				path = field + "." + key
			}
			if property, known := properties[key]; known {
				if err := validateRequestValue(child, property.(map[string]interface{}), path); err != nil {
					return err
				}
			} else if schema["additionalProperties"] == false {
				return fmt.Errorf("%s: unknown field in closed request object", path)
			} else if extra, ok := schema["additionalProperties"].(map[string]interface{}); ok {
				if err := validateRequestValue(child, extra, path); err != nil {
					return err
				}
			}
		}
	}
	if array, ok := value.([]interface{}); ok {
		if items, ok := schema["items"].(map[string]interface{}); ok {
			for i, child := range array {
				if err := validateRequestValue(child, items, fmt.Sprintf("%s.%d", field, i)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
