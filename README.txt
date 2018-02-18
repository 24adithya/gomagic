1)Install Go in desired directory.
2)Set GOROOT env variable to the installed go location. Ensure that the %PATH% also includes %GOROOT%\bin.
3)Set GOPATH env variable to the directory where all your go files will be present. This path can change based on requirement.
4)By default eclipse creates project structure
    bin
    pkg
    src

Ensure that your src files go under src by creating another folder src : 

Working project structure  	
	bin
    pkg
    src
	  -src

5)Download gocode, guru and godef by running the following commands:
  gocode is useful for content assist. Eclipse will start gocode server automatically when content assist is required.
  guru gives information about methods being used i.e. source code information
  godef  - given an expression or a location in a source file, prints the location of the definition of the symbol referred to.

  
  go get -u github.com/nsf/gocode
  go get -u golang.org/x/tools/cmd/guru
  go get -u github.com/rogpeppe/godef

a) By default it will install under %GOPATH%\bin if %GOBIN% isn't specified (At least when i was testing). Copy them and paste them under %GOROOT%\bin to have 
   all executables at the same location
  
Specify the same in eclipse path variables under 
	Preferences -> Go -> Tools

Also specify gofmt by specifying the exact location i.e. %GOROOT%\bin. (Need to test whether this is needed/correct ?)

6)You will also need to specify Go installation location and Eclipse GOPATH under 
	Preferences -> Go
	

	
Nittygritties
a)pkg containing main should be named as main
	