package restful

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// FilterChain is a request scoped object to process one or more filters before calling the target RouteFunction.
type FilterChain struct {
	Filters []FilterFunction // ordered list of FilterFunction
	Index   int              // index into filters that is currently in progress
	Target  RouteFunction    // function to call after passing all filters
}

// ProcessFilter passes the request,response pair through the next of Filters.
// Each filter can decide to proceed to the next Filter or handle the Response itself.
func (f *FilterChain) ProcessFilter(request *Request, response *Response) {
	if f.Index < len(f.Filters) {
		f.Index++
		f.Filters[f.Index-1](request, response, f)
	} else {
		go f.Target(request, response)
	}
}

// FilterFunction definitions must call ProcessFilter on the FilterChain to pass on the control and eventually call the RouteFunction
type FilterFunction func(*Request, *Response, *FilterChain)
