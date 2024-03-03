import jsonnet from 'k6/x/jsonnet';

export default function () {
  const jsonnetTemplate = `{
    person1: {
      name: "Alice",
      welcome: "Hello " + self.name + "!",
    },
    person2: self.person1 { name: "Bob" },
  }`;
  const jsonStr = jsonnet.evaluateJsonnet(jsonnetTemplate);
  console.log(jsonStr);
}