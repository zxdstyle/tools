    (function () {
        var message_btn_htm = '<div class="pvr_chat_button">' +
            '        <svg fill="#fff" width="32" height="32" t="1606891142512" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2313" width="200" height="200"><path d="M928 736c-4 0-8.032-0.736-11.904-2.272-58.528-23.424-115.072-61.056-168.032-111.808a32 32 0 1 1 44.32-46.176c26.048 24.96 52.928 46.304 80.384 63.808a1113.728 1113.728 0 0 1-28.416-121.92 32.032 32.032 0 0 1 10.016-29.152C901.856 445.44 928 391.328 928 336 928 203.648 784.448 96 608 96S288 203.648 288 336c0 3.904 0.128 7.744 0.384 11.584a32 32 0 0 1-29.856 34.016c-17.824 0.512-32.864-12.224-34.016-29.856A233.088 233.088 0 0 1 224 336C224 168.384 396.256 32 608 32s384 136.384 384 304c0 68.16-28.832 134.08-81.568 187.392 18.048 94.624 47.008 168 47.296 168.736a31.936 31.936 0 0 1-7.136 34.496A31.936 31.936 0 0 1 928 736z" p-id="2314"></path><path d="M96 992a32 32 0 0 1-29.76-43.84c0.32-0.736 29.248-74.112 47.296-168.736C60.832 726.048 32 660.16 32 592 32 424.384 204.256 288 416 288s384 136.384 384 304S627.744 896 416 896c-47.296 0-93.504-6.848-137.76-20.352-53.664 51.936-110.88 90.272-170.368 114.048A31.36 31.36 0 0 1 96 992z m320-640c-176.448 0-320 107.648-320 240 0 55.296 26.144 109.44 73.632 152.48 8.096 7.36 11.904 18.336 10.016 29.152a1108.832 1108.832 0 0 1-28.416 121.952c32.768-20.896 64.672-47.264 95.456-78.784a32 32 0 0 1 33.6-7.808c43.136 15.264 88.8 23.008 135.712 23.008 176.448 0 320-107.648 320-240S592.448 352 416 352z" p-id="2315"></path><path d="M192 624a32 32 0 0 1-32-32c0-84.992 102.88-176 256-176a32 32 0 1 1 0 64c-117.216 0-192 66.336-192 112a32 32 0 0 1-32 32z" p-id="2316"></path></svg>' +
            '    </div>'

        function CHAT() {
            this.ws = ""
            this.init()
        }

        CHAT.prototype.init = function () {
            this.loadStyleSheet('./style.css')
            // this.createContainer()
            this.connect()

            this.onInputMessage()
        }

        // 创建根容器
        CHAT.prototype.createContainer = function () {
            var container = document.createElement("div")
            container.innerHTML = message_btn_htm;
            container.setAttribute("id", "chat-_container")
            document.body.appendChild(container)
        }

        // 加载样式文件
        CHAT.prototype.loadStyleSheet = function (url) {
            var link = document.createElement("link")
            link.setAttribute("rel", "stylesheet")
            link.setAttribute("type", "text/css")
            link.setAttribute("href", url)

            var heads = document.getElementsByTagName("head")
            if(heads.length)
                heads[0].appendChild(link);
            else
                document.documentElement.appendChild(link);
        }

        // 连接服务
        CHAT.prototype.connect = function () {
            var ws = new WebSocket("ws://127.0.0.1:8199/ws/handle")
            this.ws = ws

            ws.onopen = function () {
                console.log("WebSocket Server 连接成功！");
            };
            // ws连接关闭
            ws.onclose = function () {
                if (ws) {
                    ws.close();
                    ws = null;
                }
                console.log("WebSocket Server 连接关闭！");
            };
            // ws连接错误
            ws.onerror = function () {
                if (ws) {
                    ws.close();
                    ws = null;
                }
                console.log("WebSocket Server 连接关闭！");
            };
            // ws数据返回处理
            ws.onmessage = function (result) {
                console.log(" > " + result.data);
            };
        }

        CHAT.prototype.onInputMessage = function () {
            var that = this
            var inputs = document.getElementsByClassName("message-input")
            inputs[0].addEventListener('keypress', function (e) {
                that.ws.send(JSON.stringify({test:1}))
            })
        }

        window.CHATGGPP = CHAT
    })();

new CHATGGPP()


