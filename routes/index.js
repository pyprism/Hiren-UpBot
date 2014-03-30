
/*
 * GET home page.
 */
var randomWords = require('random-words');


exports.index = function(req, res){
    res.sendfile('public/index.html') ;
};


exports.random = function (req,res){
    res.header('Access-Control-Allow-Origin', "*" );
    var hiren = randomWords({ exactly: 5, join: ' ' });
    res.json( hiren ) ;
};

