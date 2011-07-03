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

Full function reference available at http://gopkgdoc.appspot.com/pkg/github.com/crazy2be/ini