
/**
 * Module dependencies.
 */

var express = require('express');
var routes = require('./routes');
var http = require('http');
var path = require('path');

var app = express();

// all environments
app.set('port', process.env.PORT || 3000);
app.use(express.favicon(__dirname + '/public/images/favicon.ico'));
app.use(express.logger('dev'));
app.use(express.json());
app.use(express.urlencoded());
app.use(express.methodOverride());
app.use(app.router);
app.use(express.static(path.join(__dirname, 'public')));

// development only
if ('development' == app.get('env')) {
  app.use(express.errorHandler());
}

app.get('/', routes.index);
app.get('/random', routes.random);

//funny 404 :D
app.get('*', function(req, res){
    res.sendfile('public/images/404.gif') ;;
});
//change the ip
http.createServer(app).listen(app.get('port'), function(){
  console.log('Express server listening on port ' +  app.get('port'));
});
