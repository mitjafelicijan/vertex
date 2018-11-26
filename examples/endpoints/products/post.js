(function () {

  var requestedQueryParams = JSON.parse(queryParams);
  console.log(requestedQueryParams.id, queryParams);

  var requestBody = JSON.parse(body);
  console.log(requestBody);

  return JSON.stringify(requestBody);

})();