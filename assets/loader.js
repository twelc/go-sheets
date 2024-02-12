window.onload = loadData

function loadData() {
    var url="https://docs.google.com/spreadsheets/d/141maOrpeeFsydVAWP-kIaziMCHn_fI8nQv0mFB78TVk/edit#gid=0";
    xmlhttp=new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
      if(xmlhttp.readyState == 4 && xmlhttp.status==200){
        document.getElementById("table_wrapper").innerHTML = xmlhttp.responseText;
      }
      else{
        alert(xmlhttp.status)
      }
    };
    xmlhttp.open("GET",url,true);
    xmlhttp.send(null);
  }

function getValues(spreadsheetId, range, callback) {
    try {
      gapi.client.sheets.spreadsheets.values.get({
        spreadsheetId: spreadsheetId,
        range: range,
      }).then((response) => {
        const result = response.result;
        const numRows = result.values ? result.values.length : 0;
        console.log(`${numRows} rows retrieved.`);
        if (callback) callback(response);
      });
    } catch (err) {
      document.getElementById('content').innerText = err.message;
      return;
    }
  }