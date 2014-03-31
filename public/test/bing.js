// JavaScript Document
function changeIframe() {
  document.getElementById('count').disabled = true;
  document.getElementById('count').style.backgroundColor = "#000000";
  document.getElementById('message').innerHTML = "<font color='red'><b><img src='loading.gif'> running...</b></font>";
  var control = document.getElementById('st');
  if(control.style.visibility == "visible" || control.style.visibility == ""){
      control.style.visibility = "hidden";}
  else{
      control.style.visibility = "visible";}
  runSearch();
}

function runSearch(){
  var count = document.getElementById('count').value;
  var searchQuery = createRandomWordx();
  document.getElementById('bingIframe').src = 'http://www.bing.com/search?q=' + searchQuery + '+&go=&qs=n&sk=&form=QBLH';
  count = count - 1;
  document.getElementById('count').value = count;
  if (count > 0){
    var timeToHold = Math.floor(Math.random() * 30 + 10);
    t = setTimeout("runSearch()", timeToHold * 500);
  } else {
    document.getElementById('bingIframe').src = 'http://www.bing.com/search?q=spoofee+deals&go=&qs=n&sk=&form=QBLH';
    document.getElementById('bingIframe').src = 'http://www.google.com/search?hl=en&q=spoofee+deals&oq=spoofee+deals';
    document.getElementById('bingIframe').src = 'http://www.bing.com/rewards/dashboard';
	document.getElementById('count').disabled = false;
	document.getElementById('count').style.backgroundColor = "#FFF";
	document.getElementById('message').innerHTML = "<font color='green'><b color='green'>Finished!</b></font>";
  var control = document.getElementById('st');
  if(control.style.visibility == "visible" || control.style.visibility == ""){
      control.style.visibility = "hidden";}
  else{
      control.style.visibility = "visible";}
  }
}



function createRandomWordx(){
    var xhReq = new XMLHttpRequest();
    xhReq.open("GET", "http://localhost:3000/random", false);
    xhReq.send(null);
    return JSON.parse(xhReq.responseText);
}
