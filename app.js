const express = require('express'),
      logger = require('morgan'),
      bodyParser = require('body-parser'),
      top = require('./routes/top'),
      methodOverride = require('method-override'),
      os = require('os'),
      mysql = require('mysql'),
      domain = require('express-domain-middleware'),
      app = express();

app.set('views', __dirname + '/views');
app.set('view_engine', 'ejs');

const connection = mysql.createConnection({
  host: 'localhost',
  user: 'riku',
  password: '241828iqS',
  database: 'typing_game'
});

//middleware
app.use(bodyParser.urlencoded({extended: true}));
app.use(methodOverride(function(req, res){
  if (req.body && typeof req.body === 'object' && '_method' in req.body) {
    var method = req.body._method;
    delete req.body._method;
    return method;
  }
}));
app.use(logger('dev'));
app.use(function(err, req, res, next){
  res.send(err.message);
});

app.use(express.static(__dirname + '/public'));
app.use(domain);

app.get('/', top.home);
app.get('/index', top.index);
app.get('/show/:id', top.show);
app.get('/new', top.new);
app.post('/create', top.create);
app.get('/edit/:id', top.edit);
app.post('/edit/:id/update', top.update);
app.post('/delete/:id', top.delete);
app.get('/practice/:id', top.practice);

// 8080ポートで開く
const PORT = 8080;
const HOST = '0.0.0.0';

// start server on the specified port and binding host
const server = app.listen(PORT, HOST, function() {
  // print a message when the server starts listening
  console.log(`server starting on http://${HOST}:${PORT}`);
});
exports.server = server;
