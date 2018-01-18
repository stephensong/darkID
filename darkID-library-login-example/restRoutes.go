package main

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Signup",
		"POST",
		"/signup",
		Signup,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
}
