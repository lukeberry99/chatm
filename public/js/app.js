new Vue({
  el: '#app',

  data: {
    ws: null,
    newMsg: '',
    chatContent: '',
    email: null,
    username: null,
    joined: false
  },

  created: function() {
    var self = this;
    this.ws = new WebSocket('ws://' + window.location.host + '/ws');
    this.ws.addEventListener('message', function(e) {
      var msg = JSON.parse(e.data);
      self.chatContent += `${msg.username}: ${msg.message}<br />`
      var element = document.getElementById('chat-messages');
      element.scrollTop = element.scrollHeight;
    });
  },

  methods: {
    send: function () {
      if (this.newMsg != '') {
        this.ws.send(
          JSON.stringify({
            email: this.email,
            username: this.username,
            message: this.newMsg
          }
          ));
        this.newMsg = ''; 
      }
    },

    join: function () {
      if (!this.email) {
        return
      }
      if (!this.username) {
        return
      }
      this.joined = true;
    },
  }
});
