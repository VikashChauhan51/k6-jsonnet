# k6-jsonnet
k6 [jsonnet](https://jsonnet.org/) extension.

## Usage
```js
import jsonnet from 'k6/x/jsonnet';

export default function () {
  const jsonnetTemplate = '{ "hello": "world" }';
  const jsonStr = jsonnet.evaluateJsonnet(jsonnetTemplate);
  console.log(jsonStr);
}

```

## API
Functions:
----
### evaluateJsonnet
- evaluateJsonnet(`jsonnetTemplate`: *string*): *string*
    - It will take `jsonnet` template string and return `json` object string after evaluate the template.


## Build it
To build a `k6` binary with this extension, first ensure you have the prerequisites:
  - [Go toolchain](https://go101.org/article/go-toolchain.html).
  - [Git](https://git-scm.com/).
  

  Then:

1. Install xk6:
    ```cmd
    go install go.k6.io/xk6/cmd/xk6@latest
    ```
2. Build the binary:
    ```cmd
    # specific tag version
    xk6 build --with github.com/VikashChauhan51/k6-jsonnet@v0.1.0
    or
    # latest version
     xk6 build --with github.com/VikashChauhan51/k6-jsonnet@latest
    ```

## Run sample script

```cmd
./k6 run examples/main.js

```