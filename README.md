## 2.1: Project setup and creating a module
### Creating a module:
 You can pick [almost any string](https://golang.org/ref/mod#go-mod-file-ident) as your module path, but the important thing to focus on is uniqueness. 
 To avoid potential import conflicts with other people’s projects or the standard library in the future, 
 you want to pick a module path that is globally unique and unlikely to be used by anything else. 
 In the Go community, a common convention is to base your module paths on a URL that you own. <br> <br>

 Setting up your project as a module has a number of advantages — including
 making it much easier to manage third-party dependencies, [avoid supply-chain attacks](https://go.dev/blog/supply-chain),
 and ensure reproducible builds of your application in the future.

### Module paths for downloadable packages:
If you’re creating a project which can be downloaded and used by other people and programs, 
then it’s good practice for your module path to equal the location that the code can be downloaded from.
<br>
For instance, if your package is hosted at https://github.com/mhbhuiyan99/snippetbox then the module
path for the project should be ```github.com/mhbhuiyan99/snippetbox```.

## 2.2: Web application basics

1. ${{\color{Greenyellow}\large{\textsf{Handler:}}}}\$ You can think of handlers as being a bit like controllers. They’re responsible
 for executing your application logic and for writing HTTP response headers and bodies.
2. ${{\color{Greenyellow}\large{\textsf{Router (or servemux in Go terminology):}}}}\$ This stores a mapping between the URL routing patterns for your application and the corresponding
 handlers. Usually you have one servemux for your application containing all your routes.
3. ${{\color{Greenyellow}\large{\textsf{Web Server:}}}}\$ One of the great things about Go is that you can establish a web server and listen for incoming requests as part of your application itself.
 You don’t need an external third-party server like Nginx, Apache or Caddy.

Use the ```http.ListenAndServe()``` function to start a new web server. We pass in two parameters: <br>
- the TCP network address to listen on (for example ":4000") and <br>
- the servemux we just created.
  
Handler function is just a regular Go function with two parameters.
- The ```http.ResponseWriter``` parameter provides methods for assembling a HTTP
 response and sending it to the user.
- The ```*http.Request parameter``` is a pointer to
 a struct which holds information about the current request (like the HTTP method
 and the URL being requested).

```
mux := http.NewServeMux()
mux.HandleFunc("/", home)
```
Go’s ```servemux``` treats the route pattern "/" like a catch-all. So at the moment all HTTP requests to our server will be 
handled by the home function, regardless of their URL path.

### Network addresses:
 The TCP network address that you pass to ```http.ListenAndServe()``` should be in the format ```"host:port"```. 
 If you omit the host (like we did with ":4000") then the server will listen on all your computer’s available network interfaces. 
 Generally, you only need to specify a host in the address if your computer has multiple network interfaces and you want to listen on just
 one of them.

## 2.3: Routing requests
### Trailing slashes in route patterns:
 Go’s servemux has different matching rules depending on whether a route pattern ends with a trailing slash or not. <br>

 - When a pattern doesn’t have a trailing slash, it will only be matched (and the corresponding handler called) when the request
URL path exactly matches the pattern in full.
 -  When a route pattern ends with a trailing slash — like "/" or "/static/" — it is known as a subtree path pattern.
Subtree path patterns are matched (and the corresponding handler called) whenever the start of a request URL path matches the subtree path.

### Restricting subtree paths:
To prevent subtree path patterns from acting like they have a wildcard at the end, you can append the special character sequence {$} 
to the end of the pattern — like ```"/{$}"``` or ```"/static/{$}"```.
``` mux.HandleFunc("/{$}", home) // Restrict this route to exact matches on / only.```
It’s only permitted to use {$} at the end of subtree path patterns (i.e. patterns that end with a trailing slash). 

## 2.4: Wildcard route patterns
Wildcard segments in a route pattern are denoted by an wildcard identifier inside {} brackets. Like this:<br>
``` mux.HandleFunc("/products/{category}/item/{itemID}", exampleHandler)``` <br>
 In this example, the route pattern contains two wildcard segments. The first segment has
 the identifier category and the second has the identifier itemID. <br>

 When defining a route pattern, each path segment (the bit between
 forward slash characters) can only contain one wildcard and the wildcard needs to fill
 the whole path segment. Patterns like ```"/products/c_{category}"```,```/date/{y}-{m}-{d}``` or ```/{slug}.html``` are not valid.<br>

Inside your handler, you can retrieve the corresponding value for a wildcard segment using
its identifier and the ```r.PathValue()``` method. The ```r.PathValue()``` method always returns a string value.

 


 
