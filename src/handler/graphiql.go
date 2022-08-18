package handler

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"html/template"
	"net/http"
)

// graphiqlData is the page data structure of the rendered GraphiQL page
type graphiqlData struct {
	QueryString     string
	VariablesString string
	OperationName   string
	ResultString    string
}

// renderGraphiQL renders the GraphiQL GUI
func renderGraphiQL(w http.ResponseWriter, params graphql.Params) {
	t := template.New("GraphiQL")
	t, err := t.Parse(graphiqlTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create variables string
	vars, err := json.MarshalIndent(params.VariableValues, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	varsString := string(vars)
	if varsString == "null" {
		varsString = ""
	}

	// Create result string
	var resString string
	if params.RequestString == "" {
		resString = ""
	} else {
		result, err := json.MarshalIndent(graphql.Do(params), "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resString = string(result)
	}

	d := graphiqlData{
		QueryString:     params.RequestString,
		ResultString:    resString,
		VariablesString: varsString,
		OperationName:   params.OperationName,
	}
	err = t.ExecuteTemplate(w, "index", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

// tmpl is the page template to render GraphiQL
const graphiqlTemplate = `
{{ define "index" }}
<!--
 *  Copyright (c) 2021 GraphQL Contributors
 *  All rights reserved.
 *
 *  This source code is licensed under the license found in the
 *  LICENSE file in the root directory of this source tree.
-->
<!DOCTYPE html>
<html>
  <head>
    <style>
      body {
        height: 100%;
        margin: 0;
        width: 100%;
        overflow: hidden;
      }

      #graphiql {
        height: 100vh;
      }
    </style>

    <!--
      This GraphiQL example depends on Promise and fetch, which are available in
      modern browsers, but can be "polyfilled" for older browsers.
      GraphiQL itself depends on React DOM.
      If you do not want to rely on a CDN, you can host these files locally or
      include them directly in your favored resource bundler.
    -->
    <script
      crossorigin
      src="https://unpkg.com/react@17/umd/react.development.js"
    ></script>
    <script
      crossorigin
      src="https://unpkg.com/react-dom@17/umd/react-dom.development.js"
    ></script>

    <!--
      These two files can be found in the npm module, however you may wish to
      copy them directly into your environment, or perhaps include them in your
      favored resource bundler.
     -->
    <link rel="stylesheet" href="https://unpkg.com/graphiql/graphiql.min.css" />
  </head>

  <body>
    <div id="graphiql">Loading...</div>
    <script
      src="https://unpkg.com/graphiql/graphiql.min.js"
      type="application/javascript"
    ></script>
    <script>
     	// Collect the URL parameters
	var parameters = {};

	window.location.search.substr(1).split('&').forEach(function(entry) {
	    var eq = entry.indexOf('=');
	    if (eq >= 0) {
	        parameters[decodeURIComponent(entry.slice(0, eq))] =
	            decodeURIComponent(entry.slice(eq + 1));
	    }
	});

	// Produce a Location query string from a parameter object.
	function locationQuery(params) {
	    return '?' + Object.keys(params).filter(function(key) {
	        return Boolean(params[key]);
	    }).map(function(key) {
	        return encodeURIComponent(key) + '=' +
	            encodeURIComponent(params[key]);
	    }).join('&');
	}

	// Derive a fetch URL from the current URL, sans the GraphQL parameters.
	var graphqlParamNames = {
	    query: true,
	    variables: true,
	    operationName: true
	};

	var otherParams = {};

	for (var k in parameters) {
	    if (parameters.hasOwnProperty(k) && graphqlParamNames[k] !== true) {
	        otherParams[k] = parameters[k];
	    }
	}

	var fetchURL = locationQuery(otherParams);

	// Defines a GraphQL fetcher using the fetch API.
	function graphQLFetcher(graphQLParams) {
	    return fetch(fetchURL, {
	        method: 'post',
	        headers: {
	            'Accept': 'application/json',
	            'Content-Type': 'application/json'
	        },
	        body: JSON.stringify(graphQLParams),
	        credentials: 'include',
	    }).then(function(response) {
	        return response.text();
	    }).then(function(responseBody) {
	        try {
	            return JSON.parse(responseBody);
	        } catch (error) {
	            return responseBody;
	        }
	    });
	}

	// When the query and variables string is edited, update the URL bar so
	// that it can be easily shared.
	function onEditQuery(newQuery) {
	    parameters.query = newQuery;
	    updateURL();
	}

	function onEditVariables(newVariables) {
	    parameters.variables = newVariables;
	    updateURL();
	}

	function onEditOperationName(newOperationName) {
	    parameters.operationName = newOperationName;
	    updateURL();
	}

	function updateURL() {
	    history.replaceState(null, null, locationQuery(parameters));
	}

	ReactDOM.render(
	    React.createElement(GraphiQL, {
	        fetcher: graphQLFetcher,
	        onEditQuery: onEditQuery,
	        onEditVariables: onEditVariables,
	        onEditOperationName: onEditOperationName,
	        query: {{.QueryString}},
	        response: {{.ResultString}},
	        variables: {{.VariablesString}},
	        operationName: {{.OperationName}},
	    }),
	    document.getElementById('graphiql'),
	);
    </script>
  </body>
</html>
{{ end }}
`
