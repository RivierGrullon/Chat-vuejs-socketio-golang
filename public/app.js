const socket = io()

new Vue({
    el: '#chat-app',
    data: {
        message: '',
        messages: []
    },
    created() {
        const self = this;
        socket.on('chat-message', function (msg) {
            self.messages.push(msg)
        })
    },
    methods: {
        sendMessage() {
            socket.emit('chat-message', this.message);
            this.message = '';
        }
    }
});