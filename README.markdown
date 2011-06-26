INI File Reader
===============

Install it:

	goinstall github.com/crazy2be/ini

Import it:

	import "github.com/crazy2be/ini"

Use it:

	settings, err := ini.Load("somefile.ini")
	
	if err != nil {
		log.Fatal("Could not load settings :(")
		os.Exit(1)
	}

	if settings["foo"] == "bar" {
		doBar()
	} else {
		doFoo()
	}


Summary
-----
A simple package to read ini files (actually just name=value files, [sections] are not supported yet because i've had no need for them. Only reads values as strings, if you want to read different types and all, goconf will probably be better for your needs.


Functions
---------

### func Load(filename string) (settings map[string]string, err os.Error)
Load an ini file. Pass a filename, returns a map of all of the name=value pairs within the file, and an error if applicable.
TODO: Read by line, rather than buffering the whole file into memory.

### func Save(filename string, settings map[string]string) (err os.Error)
Save a map of settings to filename in name=value format. Returns nil on success, error otherwise. Creates any necessary directories to save the file according to the path given.
