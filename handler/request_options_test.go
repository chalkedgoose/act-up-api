package handler_test

import (
	"bytes"
	"fmt"
	"github.com/chalkedgoose/act-up-api/handler"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql/testutil"
)

func TestRequestOptions_GET_BasicQueryString(t *testing.T) {
	queryString := "query=query RebelsShipsQuery { rebels { name } }"
	expected := &handler.RequestOptions{
		Query:     "query RebelsShipsQuery { rebels { name } }",
		Variables: make(map[string]interface{}),
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("/graphql?%v", queryString), nil)
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_GET_ContentTypeApplicationGraphQL(t *testing.T) {
	body := []byte(`query RebelsShipsQuery { rebels { name } }`)
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("GET", "/graphql", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/graphql")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_GET_ContentTypeApplicationJSON(t *testing.T) {
	body := `
	{
		"query": "query RebelsShipsQuery { rebels { name } }"
	}`
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("GET", "/graphql", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_GET_ContentTypeApplicationUrlEncoded(t *testing.T) {
	data := url.Values{}
	data.Add("query", "query RebelsShipsQuery { rebels { name } }")

	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("GET", "/graphql", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestRequestOptions_POST_BasicQueryString_WithNoBody(t *testing.T) {
	queryString := "query=query RebelsShipsQuery { rebels { name } }"
	expected := &handler.RequestOptions{
		Query:     "query RebelsShipsQuery { rebels { name } }",
		Variables: make(map[string]interface{}),
	}

	req, _ := http.NewRequest("POST", fmt.Sprintf("/graphql?%v", queryString), nil)
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationGraphQL(t *testing.T) {
	body := []byte(`query RebelsShipsQuery { rebels { name } }`)
	expected := &handler.RequestOptions{
		Query: "query RebelsShipsQuery { rebels { name } }",
	}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/graphql")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationGraphQL_WithNonGraphQLQueryContent(t *testing.T) {
	body := []byte(`not a graphql query`)
	expected := &handler.RequestOptions{
		Query: "not a graphql query",
	}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/graphql")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationGraphQL_EmptyBody(t *testing.T) {
	body := []byte(``)
	expected := &handler.RequestOptions{
		Query: "",
	}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/graphql")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationGraphQL_NilBody(t *testing.T) {
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("POST", "/graphql", nil)
	req.Header.Add("Content-Type", "application/graphql")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestRequestOptions_POST_ContentTypeApplicationJSON(t *testing.T) {
	body := `
	{
		"query": "query RebelsShipsQuery { rebels { name } }"
	}`
	expected := &handler.RequestOptions{
		Query: "query RebelsShipsQuery { rebels { name } }",
	}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestRequestOptions_GET_WithVariablesAsObject(t *testing.T) {
	variables := url.QueryEscape(`{ "a": 1, "b": "2" }`)
	query := url.QueryEscape("query RebelsShipsQuery { rebels { name } }")
	queryString := fmt.Sprintf("query=%s&variables=%s", query, variables)
	expected := &handler.RequestOptions{
		Query: "query RebelsShipsQuery { rebels { name } }",
		Variables: map[string]interface{}{
			"a": float64(1),
			"b": "2",
		},
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("/graphql?%v", queryString), nil)
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestRequestOptions_POST_ContentTypeApplicationJSON_WithVariablesAsObject(t *testing.T) {
	body := `
	{
		"query": "query RebelsShipsQuery { rebels { name } }",
		"variables": { "a": 1, "b": "2" }
	}`
	expected := &handler.RequestOptions{
		Query: "query RebelsShipsQuery { rebels { name } }",
		Variables: map[string]interface{}{
			"a": float64(1),
			"b": "2",
		},
	}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationJSON_WithVariablesAsString(t *testing.T) {
	body := `
	{
		"query": "query RebelsShipsQuery { rebels { name } }",
		"variables": "{ \"a\": 1, \"b\": \"2\" }"
	}`
	expected := &handler.RequestOptions{
		Query: "query RebelsShipsQuery { rebels { name } }",
		Variables: map[string]interface{}{
			"a": float64(1),
			"b": "2",
		},
	}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationJSON_WithInvalidJSON(t *testing.T) {
	body := `INVALIDJSON{}`
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationJSON_WithNilBody(t *testing.T) {
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("POST", "/graphql", nil)
	req.Header.Add("Content-Type", "application/json")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestRequestOptions_POST_ContentTypeApplicationUrlEncoded(t *testing.T) {
	data := url.Values{}
	data.Add("query", "query RebelsShipsQuery { rebels { name } }")

	expected := &handler.RequestOptions{
		Query:     "query RebelsShipsQuery { rebels { name } }",
		Variables: make(map[string]interface{}),
	}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationUrlEncoded_WithInvalidData(t *testing.T) {
	data := "Invalid Data"

	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_POST_ContentTypeApplicationUrlEncoded_WithNilBody(t *testing.T) {

	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("POST", "/graphql", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestRequestOptions_PUT_BasicQueryString(t *testing.T) {
	queryString := "query=query RebelsShipsQuery { rebels { name } }"
	expected := &handler.RequestOptions{
		Query:     "query RebelsShipsQuery { rebels { name } }",
		Variables: make(map[string]interface{}),
	}

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/graphql?%v", queryString), nil)
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_PUT_ContentTypeApplicationGraphQL(t *testing.T) {
	body := []byte(`query RebelsShipsQuery { rebels { name } }`)
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("PUT", "/graphql", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/graphql")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_PUT_ContentTypeApplicationJSON(t *testing.T) {
	body := `
	{
		"query": "query RebelsShipsQuery { rebels { name } }"
	}`
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("PUT", "/graphql", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_PUT_ContentTypeApplicationUrlEncoded(t *testing.T) {
	data := url.Values{}
	data.Add("query", "query RebelsShipsQuery { rebels { name } }")

	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("PUT", "/graphql", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestRequestOptions_DELETE_BasicQueryString(t *testing.T) {
	queryString := "query=query RebelsShipsQuery { rebels { name } }"
	expected := &handler.RequestOptions{
		Query:     "query RebelsShipsQuery { rebels { name } }",
		Variables: make(map[string]interface{}),
	}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/graphql?%v", queryString), nil)
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_DELETE_ContentTypeApplicationGraphQL(t *testing.T) {
	body := []byte(`query RebelsShipsQuery { rebels { name } }`)
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("DELETE", "/graphql", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/graphql")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_DELETE_ContentTypeApplicationJSON(t *testing.T) {
	body := `
	{
		"query": "query RebelsShipsQuery { rebels { name } }"
	}`
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("DELETE", "/graphql", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
func TestRequestOptions_DELETE_ContentTypeApplicationUrlEncoded(t *testing.T) {
	data := url.Values{}
	data.Add("query", "query RebelsShipsQuery { rebels { name } }")

	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("DELETE", "/graphql", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestRequestOptions_POST_UnsupportedContentType(t *testing.T) {
	body := `<xml>query{}</xml>`
	expected := &handler.RequestOptions{}

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/xml")
	result := handler.NewRequestOptions(req)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
