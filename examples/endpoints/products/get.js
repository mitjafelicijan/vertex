(function () {

  var requestedQueryParams = JSON.parse(queryParams);
  console.log(requestedQueryParams.id, queryParams);

  var products = [
    {
      "title": "Keyboard",
      "price": 22.99
    }, {
      "title": "Mouse",
      "price": 12.99
    }
  ];

  return JSON.stringify(products);

})();