# ...  golang webassembly experiments

# ... prerequisites

you need to have golang installed on your machine


# ... outline

below steps will create a simple webassembly program in golang
which is called from javascript in browser

after running below steps you will see 5 + 3 = 8 in browser console
to confirm that golang code is being called from javascript in browser


# ...  golang webassembly which will execute in browser

	cat main.go 

```

// main.go
package main

import (
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	a, b := args[0].Int(), args[1].Int()
	return a + b
}

func main() {
	js.Global().Set("add", js.FuncOf(add))
	select {} // Keep the program running
}

```

# ...  above golang code will add two numbers supplied by javascript in browser
#
# ...  so lets compile above code into webassembly file   main.wasm

```

GOOS=js GOARCH=wasm go build -o main.wasm main.go

```

# ...  make visible file   wasm_exec.js   
#      which lives at
#
#		$GOROOT/misc/wasm/wasm_exec.js 

	which can be shown by running command

	echo $(dirname $(dirname $(readlink -f $(which go))))/misc/wasm/wasm_exec.js

	so lets copy this file into our local dir to make visable to our code

```

cp  $(dirname $(dirname $(readlink -f $(which go))))/misc/wasm/wasm_exec.js  .

```

	NOTICE in above the command ends with a period to copy file into local dir

# ...  below is html which will run javascript to define input variables it sends to above golang code 

	cat index.html

```

<!-- index.html -->
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Go WebAssembly Example</title>
    <script src="wasm_exec.js"></script>
</head>
<body>
		    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);

            // Example usage of the Go function in JavaScript
            let sum = add(5, 3);
            console.log("5 + 3 =", sum);
        });
    </script>
</body>
</html>

```

#  ...   now lets spin up a server to render above index.html

	cat local_server.go

```

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}

```

# ...   execute our local server

```

go run local_server.go

```

# ...  finally point a browser at your local server ... running on port  8080


http://localhost:8080/


# ...  screen will be blank however below will be shown in browser console ( F12 )


5 + 3 = 8


#  ...  bingo  !!!!  its working



