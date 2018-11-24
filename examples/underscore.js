var mapping = _.map([1, 2, 3], function (num) { return num * 3; });
console.log('mapping:', mapping);

var reducing = _.reduce([1, 2, 3], function (memo, num) { return memo + num; }, 0);
console.log('reducing:', reducing);