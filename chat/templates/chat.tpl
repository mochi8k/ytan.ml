<html>
  <head>
    <title>チャット</title>
    <link rel="stylesheet"
          href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">
    <style>
      ul#messages { list-style: none; }
      ul#messages li { margin-bottom: 2px; }
      ul#messages li img { margin-right: 10px; }
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
          <textarea id="message" class="form-control"></textarea>
        </div>
        <input type="submit" value="送信" class="btn btn-default" />
      </form>
      <span id="log"></span>
    </div>
  </body>
  <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
  <script>
   var socket = new WebSocket('ws://{{.Host}}/chat')

   socket.onopen = function() {
     console.log('connected')
     $('#log').text($('#log').text() + '\nConnected');
   }
   
   socket.onclose = function() {
     console.log('connection time out')
     $('#log').text($('#log').text() + '\nDisConnected');
   }

  </script>
</html>
