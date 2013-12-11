echo Testing $1

GOPATH="`pwd`/go"
export PATH=$PATH:${GOPATH}/bin
GOCMD="/usr/local/go/bin/go"
SERVERPATH="${GOPATH}/src"
appName=depinject

case $1 in
	build)
		echo "Building IoC benchmarking server"
		cd "${SERVERPATH}/mns/depinject"
 		GOPATH=${GOPATH} ${GOCMD} build -o $appName
 	;;
 	get)
		echo "Get Application Packages"
		cd "${SERVERPATH}/mns/depinject"
		GOPATH=${GOPATH} ${GOCMD} get -d
	;;
 	*) 
       echo "usage: application_executable {build|get}" ;;
 esac
 exit 0