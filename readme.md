# Vertex

Create mock API's and enrich them with some basic logic and simplify prototyping.

**When to use Vertex?**

You need to hack together a prototype with a REST API and you need to have some logic in your API's (when stubs are not enough) and you don't want to invest time into making a Express app with all the bells and whistles that comes with it.

Just point Vertex into folder where you have your static HTML app and endpoints and that's it.

**Vertex provides developer with:**
- static file server for serving your prototype HTML and JS code,
- embedded JavaScript interpreter so you can write endpoints in Javascript,
- localStorage implementation that stores everything into JSON file.


**Table of contents**
- [Installation](#installation)
- [Usage](#usage)
- [Writing endpoints](#writing-endpoints)
- [Local storage](#local-storage)
- [Underscore.js](#underscorejs)
- [Limitations](#limitations)


## Installation

Download appropriate version from https://github.com/mitjafelicijan/vertex/releases and extract tarball.

If you want binary to be accessible from everywhere add path to the binary to your `$PATH`.

## Usage

By nature vertex needs configuration file called `vertex.yml`. If this file doesn't exist in folder where vertex is executed application creates boilerplate configuration file.

Boilerplate configuration file:

```yaml
vertex:

  # server settings
  host: 0.0.0.0
  port: 4001
  
  # api endpoint root
  prefix: /api/
  
  # database file location
  datastore: data.json
  
  # folder for html, css, js
  static: ./static
  
  # folder for rest api endpoints
  endpoints: ./endpoints
```

After file is found you can start Vertex with `./vertex` and start using it.

## Writing endpoints

Endpoints in Vertex are folders which contains:
- get.js,
- post.js,
- put.js,
- delete.js. 

You add just the ones you need.

Example of a products endpoint:
- ./endpoints
    - ./products
        - get.js
        - post.js

Basic endpoint `post.js`.

```js
(function () {

  var requestedQueryParams = JSON.parse(queryParams);
  console.log(requestedQueryParams.id, queryParams);

  var requestBody = JSON.parse(body);
  console.log(requestBody);

  return JSON.stringify(requestBody);

})();
```

For additional example you can download example project from release tab.

## Local storage

In `examples/test` folder you have example for localStorage and how to use it.

## Underscore.js

You can use [Underscore.js](https://underscorejs.org/) in your endpoints. It is embedded with Vertex.

## Limitations

- Targets ES5, no ES6 support currently
- Limited JS usage with no window object.