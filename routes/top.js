const mysql = require('mysql');

const connection = mysql.createConnection({
  host: '127.0.0.1',
  user: 'riku',
  password: '241828iqS',
  database: 'typing_game'
});

exports.home = function(req, res){
  res.render('home.ejs');
}

exports.index = function(req, res){
  connection.query(
    'SELECT * FROM wordsList', (error, results) => {
      res.render('index.ejs', {wordsList: results});
    }
  );
}

exports.show = function(req, res){
  connection.query(
    'SELECT * FROM wordsList WHERE id = ?',
    [req.params.id],
    (error, results) => {
      var wordsName = results[0].title;
      connection.query(
        'SELECT * FROM words WHERE listID = ?',
        [req.params.id],
        (error, results) => {
          res.render('show.ejs', {name: wordsName, words: results});
        }
      );
    }
  );
}

exports.new = function(req, res){
  res.render('new.ejs');
}

exports.create = function(req, res){
  var listID = null;
  connection.query(
    'INSERT INTO wordsList (title) VALUES (?)',
    [req.body.wordsTitle],
    (error, results) => {
      listID = results.insertId;

      req.body.wordsList.forEach((list) => {
        connection.query(
          'INSERT INTO words (name, listID) VALUES (?, ?)',
          [list, listID]
        );
      });
      res.redirect('/index');
    }
  );

}

exports.edit = function(req, res){
  connection.query(
    'SELECT * FROM wordsList WHERE id = ?',
    [req.params.id],
    (error, results) => {
      var wordsName = results[0].title;
      connection.query(
        'SELECT * FROM words WHERE listID = ?',
        [req.params.id],
        (error, results) => {
          res.render('edit.ejs', {name: wordsName, words: results});
        }
      );
    }
  );
}

exports.update = function(req, res){
  var listID = req.params.id;
}

exports.delete = function(req, res){

}

exports.practice = function(req, res){
  connection.query(
    'SELECT name FROM words WHERE listID = ?',
    [req.params.id],
    (error, results) => {
      res.render('practice.ejs', {words: results});
    }
  )
}
