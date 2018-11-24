var ok = true;
if (ok) {
    console.log('yes ok')
}


try {
    var array = ['john', 'jade', 'luke'];
    for (var index = 0; index < array.length; index++) {
        var element = array[index];
        console.log(element);
    }
} catch (err) {
    console.log(err.stack + '\n');
}


function getErrorObject() {
    try { throw Error('') } catch (err) { return err; }
}

var err = getErrorObject();
console.log(JSON.stringify(err, null, 2));
var caller_line = err.stack.split("\n")[4];
var index = caller_line.indexOf("at ");
var clean = caller_line.slice(index + 2, caller_line.length);