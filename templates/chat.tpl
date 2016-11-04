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
    </div>
  </body>
</html>
