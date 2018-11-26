# Vertex

Create mock API's and enrich them with some basic logic and simplify prototyping.

Vertex provides developer with:
- static file server for serving your prototype HTML and JS code
- embedded JavaScript interpreter so you can write endpoints in Javascript
- localStorage implementation that stores everything into JSON file


**Table of contents**
- [Installation](#installation)
- [Usage](#usage)
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


## Limitations

- Targets ES5, no ES6 support currently