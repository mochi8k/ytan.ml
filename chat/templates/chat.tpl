<html>
  <head>
    <title>Chat</title>
    <link rel="stylesheet"
          href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">
    <style>
      ul#messages { list-style: none; }
      ul#messages li { margin-bottom: 2px; }
      ul#messages li img { margin-right: 10px; }
	  .panel { margin-top: 20px; }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="panel panel-default">
        <div class="panel-body">
          <ul id="messages"></ul>
        </div>
      </div>
      <form id="chatbox" role="form">
        <div class="form-group">
          <input type="text" id="message" class="form-control" placeholder="Please enter the message here"></input>
        </div>
        <input type="submit" value="Submit" class="btn btn-default" />
      </form>
      <span id="log"></span>
    </div>
  </body>
  <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
  <script>
   var socket = new WebSocket('ws://{{.Host}}/chat/room');

   socket.onmessage = function(e) {
	 $('#messages').append($('<li>').text(e.data));
   }

   $('#chatbox').submit(function() {
	 var message = $('#message').val();
	 socket.send(message);
	 $('#message').val('');
	 
	 return false;
   });

  </script>
</html>
