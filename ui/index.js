$(document).ready(function () {
    $("form").submit(function (event) {
        console.log("called")
        var url =$this.serialize()
        var payload={
            url:url
        }
        console.log(payload)
        $.ajax({
            type:"POST",
            url:"/webpageinfo",
            data:payload
        })
    });
  });