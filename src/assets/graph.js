
function getGraphData(){
    coupert = $("html").attr("coupert-item")
    querry = $("#querry").text()
    $.ajax({
        type: "post",
        url: "/get-graph-data",
        data: {"querry":querry, "coupert":coupert},
        dataType: "json",
        success: function (response) {
            settings = {};
            let d = new liteChart("chart", settings);

            // Set labels
            d.setLabels(response.time);

            // Set legends and values
            d.addLegend({"name": "default", "stroke": "#CDDC39", "fill": "#fff", "values": response.data});

            // Inject chart into DOM object
            let div = document.getElementById("wrapper");
            d.inject(div);

            // Draw
            d.draw();
        }
    });
}