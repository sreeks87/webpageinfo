$(document).ready(function () {
    
    $("form").submit(function (event) {
        event.preventDefault();
        console.log($("#website").val())
      var formData = {
        url: $("#website").val(),
        };
    console.log(formData)
      $.ajax({
        type: "POST",
        url: "/webpageinfo",
        data:  JSON.stringify(formData),
        dataType: "json",
        encode: true,
        contentType: "application/json",
      }).done(function (data) {
        console.log(data);
        headStr="H1 : "+data.headings.h1count +"<br>"
        headStr+="H2 : "+data.headings.h2count+"<br>"
        headStr+="H3 : "+data.headings.h3count+"<br>"
        headStr+="H4 : "+data.headings.h4count+"<br>"
        headStr+="H5 : "+data.headings.h5count+"<br>"
        headStr+="H6 : "+data.headings.h6count+"<br>"

        $("#html").html(data.htmlversion)
        $("#title").html(data.pagetitle)
        $("#headingcount").html(headStr)
        $("#internal").html(data.links.internallinks)
        $("#external").html(data.links.externallinks)
        $("#broken").html(data.links.inaccessiblelinks)
        $("#login").html(data.loginform.toString())
      });
  
      
    });
  });