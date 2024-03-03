# k6-jsonnet
k6 jsonnet extension


```js
import jsonnet from 'k6/x/jsonnet';

export default function () {
  const jsonnetTemplate = '{ "hello": "world" }';
  const jsonStr = jsonnet.evaluateJsonnet(jsonnetTemplate);
  console.log(jsonStr);
}

```
## Build it

```cmd
xk6 build --with github.com/VikashChauhan51/k6-jsonnet@v0.1.0
```

## Run sample script

```cmd
./k6 run examples/main.js

```