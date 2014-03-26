
/*
 * GET home page.
 */
var randomWords = require('random-words');


exports.index = function(req, res){
    res.sendfile('public/index.html') ;
};


exports.random = function (req,res){
    var hiren = randomWords({ exactly: 10, join: ' ' });
    res.json({ random : hiren }) ;
};


exports.about = function(req ,res){
  res.sendfile('public/about.html');
};