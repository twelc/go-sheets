
function getGraphData(){
    querry = $("#querry").text()
    console.log("parsing")
    $.ajax({
        type: "post",
        url: "/get-graph-data",
        data: {"querry":querry},
        dataType: "json",
        success: function (response) {
            console.log(response.data)
            $("#wait").hide()
            new Chart(document.getElementById("line-chart"), {
                type: 'line',
                data: {
                  labels: response.time,
                  datasets: [{ 
                      data: response.data,
                      label: querry,
                      borderColor: "#3e95cd",
                      fill: false
                    }]
                },
                options: {
                  title: {
                    display: false
                  }
                }
              });
        }
    });
}
$(window).on('load', getGraphData());