// var wordsList = ['apples', 'bananas', 'oranges', 'mangoes', 'grapes'];  //単語リスト
var wordsList = [];
var restWords = [];                                                     //残りの単語
var targetWord = [];                                                    //タイプする単語
var clearWords = [];                                                    //クリアした文字
var misTypeNum = 0;                                                     //ミスしたタイプ数
var clearTypeNum = 0;                                                   //クリアしたタイプ数
var s_Time = null;                                                      //開始時刻
var e_Time = null;                                                      //終了時刻
var passSec = 0;                                                        //経過時間
var count = 3;                                                          //最初のカウントダウン

//コールバック関数のインスタンス
var passageID;
var countID;
var timeID;
var startID;

//キー操作を行うファンクション
window.addEventListener('keyup', function(){
  //スタート判定キーイベント
  if(event.keyCode == '32' && restWords.length == 0){
    start();
    event.returnValue = false;
  }

  //restart判定キーイベント
  else if(event.keyCode == '27'){
    init();
  }

  //タイプ判定キーイベント
  else if(targetWord.length > 0){
    check(event.keyCode);
  }
});

//リスタート処理
function init(){
  clearAll();
  clearInterval(countID);
  clearInterval(passageID);
  clearTimeout(timeID);
  clearTimeout(startID);
  $('.targetArea').text('ここに単語が表示されます.');
  $('.startExplanation').css('visibility', 'visible');
  $('.showTime').css('display', 'none');
}

//スタート処理
function start(){
  clearAll();

  var list = document.getElementById('hiddenArray').children;
  for(var i=0; i<list.length; i++){
    wordsList.push(list[i].textContent);
  }

  for(var i=0; i<wordsList.length; i++){
    restWords.push(wordsList[i]);
  }
  startCount();
  timeID = setTimeout('startTime();', 4000);
  startID = setTimeout('showWords(restWords);', 4000);
}

//初期化処理
function clearAll(){
  $('.targetArea').css('color', 'black');
  $('.startExplanation').css('visibility', 'hidden');
  wordsList = [];
  restWords = [];
  clearWords = [];
  misTypeNum = 0;
  clearTypeNum = 0;
  s_Time = null;
  e_Time = null;
  passSec = 0;
  count = 3;
  $('.result').css('display', 'none');
}

//単語の出現処理
function showWords(restWords){
  if(restWords.length > 0){
    // $('h1').text(restWords[0]);
    var word = restWords[Math.floor(Math.random() * restWords.length)];
    targetWord = word.split("");
    $('.targetArea').text(word);
    $('.hiddenArea').text(word);
  }else{
    endTime();
    $('.startExplanation').css('visibility', 'visible');
    showResult();
  }
}

//キー入力によるタイプ正確判断処理
function check(keyCode){
  if(keyCode == String(convertToKeyCode(targetWord[0]))){
    clearTypeNum ++;
    clearWords.push(targetWord[0]);
    targetWord.splice(0, 1);
    $('.targetArea').html("<font color='#CCCCCC'>"+clearWords.join('')+"</font>"+targetWord.join(''));
  }else{
    misTypeNum ++; //これはログイン情報ごとに変更しないといけないかも
  }
  if(targetWord.length == 0){
    for(i=0; i<restWords.length; i++){
      if(restWords[i] == $('.hiddenArea').text()){
        restWords.splice(i, 1);
      }
    }
    clearWords = [];
    showWords(restWords);
  }
}

//結果の表示
function showResult(){
  $('.targetArea').text('clear!');
  $('.targetArea').css('color', 'red');
  $('.result').css('display', 'block');
  $('.clearType').text(clearTypeNum);
  $('.misType').text(misTypeNum);
  $('.rate').text(Math.round(1000*clearTypeNum/(clearTypeNum+misTypeNum))/10);
  $('.time').text(Math.round((e_Time.getTime() - s_Time.getTime())/1000*10)/10);
}

//経過時間の表示
function showPassage(){
  passSec += 0.1;
  $('.passTime').text(Math.round(passSec*10)/10);
}

//開始時間の取得
function startTime(){
  $('.showTime').css('display', 'block');
  s_Time = new Date();
  passageID = setInterval('showPassage()', 100);
}

//終了時間の取得
function endTime(){
  $('.showTime').css('display', 'none');
  e_Time = new Date();
  clearInterval(passageID);
}

//スタート時のカウントダウンを行う処理
function startCount(){
  $('.targetArea').text(count);
  countID = setInterval('countDown()', 1000);
}

//スタート時のカウントダウンを行う処理
function countDown(){
  if(count > 1){
    count--;
    $('.targetArea').text(count);
  }else if(count == 1){
    $('.targetArea').text('START!');
    count--;
  }else{
    clearInterval(countID);
  }
}

//単語をキーコードに変換する関数
function convertToKeyCode(word){
  if(!isNaN(word)){
    return word+48;
  }else{
    return getAlphabetIndex(word)+65;
  }
}

//アルファベットの番号を返す処理
function getAlphabetIndex(str){
  var alphabet = "abcdefghijklmnopqrstuvwxyz";
  return (alphabet.indexOf(str) === -1) ? (alphabet.indexOf(str.toLowerCase())) : (alphabet.indexOf(str));
}
