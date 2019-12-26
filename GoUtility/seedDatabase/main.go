package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type author struct {
	name         string
	emailaddress string
	imgurl       string
}
type authorauthor struct {
	userid1 int
	userid2 int
}
type organization struct {
	imgurl  string
	pageurl string
}
type article struct {
	title          string
	userid         int
	content        string
	organizationid int
	imgurl         string
	dateposted     string
	claps          int
}

func main() {
	connStr := "user=kevin2 password=1234 dbname=medium sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	authors := []author{
		{"Kartik Khare", "ktikkhare@fakeemail.com", "https://cdn-images-1.medium.com/fit/c/200/200/1*mm1eWBY4TZbncymiNckSqw.jpeg"},
		{"Bob", "Bob@fakeemail.com", "https://miro.medium.com/fit/c/256/256/1*7AmulBADt7u-Hc1CMlg8_Q.jpeg"},
		{"Nick Scialli", "NickScialli@fakeemail.com", "https://miro.medium.com/fit/c/96/96/1*tZXMb_y-vvbhULjm2OBD8Q.jpeg"},
	}
	articles := []article{
		{"Write a Web Service with Go Plug-Ins", 2, `The target of this article is to inspect why and how to build a modular web service.
		In an automation driven environment, there are many conditions that bring project evolution to Hell. Some of those arise from a non-modular environment.
		Even if you are on a microservices architecture, with a compiled language all this case leads to rebuilding the whole service project:
		Edit how a single part of the project work.
		Optimize the scaling up of a single endpoint.
		And also, you will face some of these uncomfortable situations:
		Reuse the same behaviour you have in a service into another service.
		You probably can use only a single programming language.
		You need to share the whole project if you have a team (this little thing in an off-shoring project can be very bad).
		With a plugin architecture, you can develop the service as a set of components that enable to:
		Have function-dedicated teams
		Detached deployment for the components with a rolling update.
		Reuse a component in several projects.
		Customize how to scale a service by configuring winch component you have to include.
		Last but not least, you can develop the component in each language that builds a library!
		Architecture
		Taking an example web service, we can split our project into the following parts:
		Core: that will read the configurations, load the component, listen and serve the HTTP.
		Controller: an HTTP Handler function.
		Middlewares: a component that checks if the request is authorized to access the controller.

		In this way, if you update the core to HTTPS, you can redeploy only the core file and only for the services that need to be HTTPS compliant. In the same way, if you update the JWT plugin to use a new hash method, you have only to redeploy the plugin and reload the core.
		Architecture components
		Controller
		A controller is an HTTP handler function that will apply logic to the request to build and return the response.
		Middlewares
		This kind of objects arises from the needs to do things like filter, trace or log, the incoming request, without involving all the project to manage it.
		A kind of middleware can be an IP filter: you want to enable only certain IP to access your service. When the request reaches your server, you can directly check if the source IP is in your whitelist. In this way, you don’t load other parts of the service.
		One of the advantages can be a reduction of the risk to expose detail in case of malicious intent.

		Example: let’s create another middleware that checks the access headers. Now, chaining it with the previous IP middleware, only the users authorized in our IP pool can access the service.
		So, if the first step in the controller is to query your DB for user data, you’re safe from SQL injection. That’s because you don’t start communicating with the DB if all the preconditions are fine.
		Core
		As explained, for this project we use the Plugin functionality to separate the components. For our scope, the Core component will care about loading Controllers and Middlewares Plugins, map it to the endpoints using a configuration file and finally start the HTTP listener.
		All other routes will return “404 not found”.

		Mapping middlewares and controller to an endpoint
		This level of independence can be easily loved for cloud use-case:

		Implementation
		Let’s start from a basic core with a controller (HTTP function handler) that responds to our home:

		On top of this, add the constructs needed to implement the middlewares:

		The middleware type

		…and a Chain function that provide a mechanism to concatenate middlewares
		Now build a sample middleware: method middleware checks the match between the HTTP method and the one passed as argument. If not, returns a 400 Bad Request.
		The arguments in this case, are a sequence of approved HTTP methods that we need to split and check:

		It’s time to attach the middleware inside the Chain, passing what HTTP Method you want to allow in function arguments.
		Then pass the Chain to the HandlerFunc inside our basic HTTP service:

		Full example source code here
		Plugins
		N.B. it works only on Linux, but containers provides a great workaround.
		The package of a plugin needs to be “Main”. Unlike that, the package can’t see the entities such as types and functions in the “real” main package. So, as a suggestion, maintain plugins dumber as possible.
		Inside the Plugin, you must export a variable or a function as the symbol to load

		To build the plugin, we need to use the -buildmode=plugin flag and specify the result name
		$go build -buildmode=plugin -o first.so first.go
		You have built your first plugin!
		Search in your path ad you will see first.so file that is a standard library that you can import in any language that supports this.
		I’m proud of you my little padawan… let’s use it in GO:
		Make a new file “main.go” and load the library file you created before:

		Add the symbol search code:

		Now you can use your function from the plugin inside “main.go”:

		RUN
		$ go run main.go
		you will see
		Hello FROM PLUGIN!!!
		Example source code here
		It’s time to move on GO ‘Oop-style’ to get more than a single function from the Plugins.
		Evolve the first.go plugin using a type through which you can attach functions as methods, then export a variable symbol that refers to the type as an object:

		Now you have to change the kind of import in a more secure way in the main.go:

		Plugins Implementation
		At this moment you know:
		what is the project target
		what is the target architecture
		what is the core
		what is a controller
		what is a middleware
		how to build a plugin and use it
		you’re ready to build the “controller” plugin!
		In our repository create a plugin folder:
		$mkdir plugins
		Inside we create two folders, one for middlewares, one for controllers
		$cd plugins
		$mkdir controller
		$mkdir middlewares
		Build the Controllers
		Inside the plugins/controllers folder create general.go:

		Build the Middlewares
		Now export the method middleware you have done before, inside plugin.
		Under plugins/middlewaresfolder create method.so:

		build the plugins:
		$go build -buildmode=plugin -o plugins/middlewares/method.so plugins/middlewares/method.go
		$go build -buildmode=plugin -o plugins/controllers/genearal.so plugins/controllers/genearal.go
		Import the Plugins
		To import the Plugins you will load it from a configuration file that map endpoints to middlewares and controller.
		Sample routes.jsonlooks like this:

		Creating the file in this way, you can attach several middlewares to a route and use a middleware in several routes.
		Read the configurations
		Now you can proceed on reading the configuration and map it to a struct (with this tool is very simple to translate json to go struct):

		And let’s make a function to read from JSON:

		Load the Plugins
		As you call an exported type method from the plugin, we need to adopt some conventions, I opted for:
		Controller type with method Fire()
		Middleware type with method Pass()
		Walking into the configuration we can dynamically link the libraries:
		From “plugin.Open” documentation: If a path has already been opened, then the existing *Plugin is returned It is safe for concurrent use by multiple goroutines.
		Load Controller plugin:

		Load middleware modules to attach on the route:

		Now we can put all together to work starting the web server and test our service.
		$go build -o start -v
		You can see in the complete Repository with other features:
		script to create plugins scaffold
		makefile to build all and clean all
		test for standard implementation`, 1, "https://miro.medium.com/max/1718/1*qVKB83GozPulENVR-rMV6Q.jpeg", "February 20, 2019", 811},
	}
	for _, author := range authors {
		sqlStatement := fmt.Sprintf(`
INSERT INTO author (name, emailaddress, imgurl)
VALUES (%v, %v, %v)`, author.name, author.emailaddress, author.imgurl)
		_, err = db.Exec(sqlStatement)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, article := range articles {
		sqlStatement := fmt.Sprintf(`
INSERT INTO article (title, userid, content,organizationid,imgurl,dateposted,claps)
VALUES (%v, %v, %v, %v, %v, %v, %v)`, article.title, article.userid, article.content, article.organizationid, article.imgurl, article.dateposted, article.claps)
		_, err = db.Exec(sqlStatement)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
