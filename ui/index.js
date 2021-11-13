$(document).ready(function () {
    $("form").submit(function (event) {
      var formData = {
        weburl: $("#website").val(),
        };
  
      $.ajax({
        type: "POST",
        url: "/webpageinfo",
        data: formData,
        dataType: "json",
        encode: true,
      }).done(function (data) {
        console.log(data);
      });
      event.preventDefault();
    });
  });