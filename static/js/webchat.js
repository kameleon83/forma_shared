window.onload = function () {
    var conn;
    var msg = document.getElementById("webchat-msg");
    var log = document.getElementById("webchat-log");
    var name = document.getElementById("firstname");

    function timeString() {
        let heure = new Date(Date.now()).getHours().toLocaleString();
        let minute = new Date(Date.now()).getMinutes().toLocaleString();
        let second = new Date(Date.now()).getSeconds().toLocaleString();
        return heure + ":" + minute + ":" + second
    }

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        hashtag(item.innerHTML);
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("webchat").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        msg.value = urlCheck(msg.value);
        conn.send(timeString() + " - " + name.dataset.name + " : " + msg.value);
        msg.value = "";
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + location.hostname + ":9001/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerHTML = messages[i];
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }

    function urlCheck(message) {
        let link = /@(https?:\/\/)[a-z0-9.\/-_]+/gi;
        if (link.test(message)) {
            message = message.replace(link, "<a href='" + message.substring(1) + "'>" + message + "</a>");
        }
        return message
    }

    function hashtag(message) {
        let tag = /#[a-z0-9.\/-_]+/i;
        if (tag.test(message)) {
            let nameSend = name.innerText.substring(8).toLowerCase();
            if (nameSend === message.match(tag)[0].substring(1).toLowerCase()) {
                notifyChat(nameSend, "Chat Message")
            }
        }
    }

    function notifyChat(nameSend, nameCall) {
        console.log("notifyChat");
        if (!("Notification" in window)) {
            alert("This browser does not support desktop notification");
        } else if (Notification.permission === "granted") {
            let options = {
                body: nameSend + " t'appelle sur le chat",
                icon: "/images/chat-black.png",
                dir: "ltr",
                requireInteraction: true
            };
            let notification = new Notification(nameCall, options);
            notification.onclick = function () {
                notification.close();
                window.open().close();
                return window.focus();
            };
        } else if (Notification.permission !== 'denied') {
            Notification.requestPermission(function (permission) {
                if (!('permission' in Notification)) {
                    Notification.permission = permission;
                }

                if (permission === "granted") {
                    let options = {
                        body: nameSend + " t'appelle sur le chat",
                        icon: "/images/chat-black.png",
                        dir: "ltr",
                        requireInteraction: true
                    };
                    let notification = new Notification(nameCall, options);
                    notification.onclick = function () {
                        notification.close();
                        window.open().close();
                        return window.focus();
                    };
                }
            });
        }
    }
};

