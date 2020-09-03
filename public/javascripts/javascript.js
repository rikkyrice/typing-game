function add_list(){
  $('.form-lists').append('<li class="form-list"><input type="text" name="wordsList[]" class="input-list"><button type="button" class="remove-button" onClick="remove_list(this)">- 削除</button></li>');
}

function remove_list(elem){
  elem.parentNode.remove();
}
