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