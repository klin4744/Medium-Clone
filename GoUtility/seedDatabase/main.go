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
	name    string
	imgurl  string
	pageurl string
}
type article struct {
	title          string
	userid         int
	content        string
	topic          string
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
		test for standard implementation`, "JavaScript", 3, "https://miro.medium.com/max/1718/1*qVKB83GozPulENVR-rMV6Q.jpeg", "February 20, 2019", 811}, {
			"Why goroutines are not lightweight threads?", 1, `GoLang is gaining incredible popularity these days.
			One of the main reasons for that is the simple and lightweight concurrency in the form of goroutines and channels that it offers to the developers.
			Concurrency has existed since long ago in the form of Threads which are used in almost all the applications these days.
			To understand why goroutines are not lightweight threads, we should first understand how Thread works in OS.
			If you are already familiar with Threads, you can directly skip here.
			What are Threads?
			A thread is just a sequence of instructions that can be executed independently by a processor. Threads are lighter than the process and so you can spawn a lot of them.
			A real life application would be a web server.
			A webserver typically is designed to handle multiple requests at the same time. And these requests normally don’t depend on each other.
			So a Thread can be created (or taken from a Thread pool) and requests can be delegated, to achieve concurrency.
			Modern processors can executed multiple threads at once (multi-threading) and also switch between threads to achieve parallelism.
			Are threads lighter than processes?
			Yes and No.
			In concept,
			Threads share memory and don’t need to create a new virtual memory space when they are created and thus don’t require a MMU (memory management unit) context switch
			Communication between threads is simpler as they have a shared memory while processes requires various modes of IPC (Inter-Process Communications) like semaphores, message queues, pipes etc.
			That being said, this doesn’t always guarantee a better performance than processes in this multi-core processor world.
			e.g. Linux doesn’t distinguish between threads and processes and both are called tasks. Each task can have a minimum to maximum level of sharing when cloned.
			When you call fork(), a new task is created with no shared file descriptors, PIDs and memory space. When you call pthread_create(), a new task is created with all of the above shared.
			Also, synchronising data in shared memory as well as in L1 cache of tasks running on multiple cores takes a bigger toll than running different processes on isolated memory.
			Linux developers have tried to minimise the cost between task switch and have succeeded at it. Creating a new task is still a bigger overhead than a new thread but switching is not.
			What can be improved in Threads?
			There are three things which make threads slow:
			Threads consume a lot of memory due to their large stack size (≥ 1MB). So creating 1000s of thread means you already need 1GB of memory.
			Threads need to restore a lot of registers some of which include AVX( Advanced vector extension), SSE (Streaming SIMD Ext.), Floating Point registers, Program Counter (PC), Stack Pointer (SP) which hurts the application performance.
			Threads setup and teardown requires call to OS for resources (such as memory) which is slow.
			Goroutines
			Goroutines exists only in the virtual space of go runtime and not in the OS.
			Hence, a Go Runtime scheduler is needed which manages their lifecycle.
			Go Runtime maintains three C structs for this purpose:
			The G Struct : This represents a single go routine with it’s properties such as stack pointer, base of stack, it’s ID, it’s cache and it’s status
			The M Struct : This represents an OS thread. It also contains a pointer to the global queue of runnable goroutines, the current running goroutine and the reference to the scheduler
			The Sched Struct : It is a global struct and contains the queues free and waiting goroutines as well as threads.
			So, on startup, go runtime starts a number of goroutines for GC, scheduler and user code. An OS Thread is created to handle these goroutines. These threads can be at most equal to GOMAXPROCS.
			Start from the bottom!
			A goroutine is created with initial only 2KB of stack size. Each function in go already has a check if more stack is needed or not and the stack can be copied to another region in memory with twice the original size. This makes goroutine very light on resources.
			Blocking is fine!
			If a goroutine blocks on system call, it blocks it’s running thread. But another thread is taken from the waiting queue of Scheduler (the Sched struct) and used for other runnable goroutines.
			However, if you communicate using channels in go which exists only in virtual space, the OS doesn’t block the thread. Such goroutines simply go in the waiting state and other runnable goroutine (from the M struct) is scheduled in it’s place.
			Don’t interrupt!
			The go runtime scheduler does cooperative scheduling, which means another goroutine will only be scheduled if the current one is blocking or done. Some of these cases are:
			Channel send and receive operations, if those operations would block.
			The Go statement, although there is no guarantee that new goroutine will be scheduled immediately.
			Blocking syscalls like file and network operations.
			After being stopped for a garbage collection cycle.
			This is better than pre-emptive scheduling which uses timely system interrupts (e.g. every 10 ms) to block and schedule a new thread which may lead a task to take longer than needed to finish when number of threads increases or when a higher priority tasks need to be scheduled while a lower priority task is running.
			Another advantage is that, since it is invoked implicitly in the code e.g. during sleep or channel wait, the compile only needs to safe/restore the registers which are alive at these points. In Go, this means only 3 registers i.e. PC, SP and DX (Data Registers) being updated during context switch rather than all registers (e.g. AVX, Floating Point, MMX).
			If you want to explore more about go concurrency you can refer to the links below:
			Concurrency is not parallelism by Rob Pike
			Analysis of Go runtime Scheduler
			Five things that make Go fast by Dave Cheney
			Discussion in golang-nuts mailing list`, "JavaScript", 2, "https://miro.medium.com/max/3840/1*_MqLBkRmPSp3MUV6c1FEkQ.jpeg", "March 23, 2018", 3600,
		}, {"Using Pre-Commit and Pre-Push Git Hooks in a React Project", 3, `One topic I have gotten more and more excited about throughout my software development career is quality! Perhaps I’ve been burned one too many times. Alas, I decided to test adding git hooks to a React project using the husky package. My goal was to make it so that, prior to either committing code or pushing to a git repository, both the eslint linter and jest test suite must run.
		`, "JavaScript", 1, "https://miro.medium.com/max/4200/1*xO6RQvFmAsRQQgod1Tk1Kg.jpeg", "March 23, 2018", 1200},
	}
	organizations := []organization{
		{"Level Up Coding", "https://cdn-images-1.medium.com/max/952/1*txRrggvTHssimaGlCMZ2mg@2x.png", "https://levelup.gitconnected.com/"},
		{"Codeburst.ip", "https://cdn-images-1.medium.com/max/392/1*LC0hwOq4FY2CG5F9W7R34Q@2x.png", "https://codeburst.io/"},
		{"Quick Code", "https://ph-files.imgix.net/7e0946c8-1b54-4403-b82e-57f5a59bf1ec?auto=format&auto=compress&codec=mozjpeg&cs=strip&w=692.1428571428571&h=380&fit=max&dpr=2", "https://medium.com/quick-code"},
	}
	authorauthors := []authorauthor{
		{1, 2}, {2, 3},
	}
	for _, author := range authors {
		sqlStatement := fmt.Sprintf(`
INSERT INTO author (name, emailaddress, imgurl)
VALUES ('%s', '%s', '%s')`, author.name, author.emailaddress, author.imgurl)
		_, err = db.Exec(sqlStatement)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, article := range articles {
		sqlStatement := fmt.Sprintf(`
INSERT INTO article (title, userid, content,organizationid,imgurl,dateposted,claps)
VALUES ('%s', %d, '%s', %d, '%s', '%s', %d)`, article.title, article.userid, article.content, article.organizationid, article.imgurl, article.dateposted, article.claps)
		_, err = db.Exec(sqlStatement)
		if err != nil {
			log.Fatalln(err)
		}
	}
	for _, org := range organizations {
		sqlStatement := fmt.Sprintf(`
INSERT INTO organization (name, imgurl, pageurl)
VALUES ('%s', '%s', '%s')`, org.name, org.imgurl, org.pageurl)
		_, err = db.Exec(sqlStatement)
		if err != nil {
			log.Fatalln(err)
		}
	}
	for _, rel := range authorauthors {
		sqlStatement := fmt.Sprintf(`
INSERT INTO authorauthor (user1id,user2id)
VALUES (%d, %d)`, rel.userid1, rel.userid2)
		_, err = db.Exec(sqlStatement)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
