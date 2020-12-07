(function () {
    var html = '<div id="chat-btn" class="pvr_chat_button has_message">' +
        '        <svg fill="#fff" width="32" height="32" t="1606891142512" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2313" width="200" height="200"><path d="M928 736c-4 0-8.032-0.736-11.904-2.272-58.528-23.424-115.072-61.056-168.032-111.808a32 32 0 1 1 44.32-46.176c26.048 24.96 52.928 46.304 80.384 63.808a1113.728 1113.728 0 0 1-28.416-121.92 32.032 32.032 0 0 1 10.016-29.152C901.856 445.44 928 391.328 928 336 928 203.648 784.448 96 608 96S288 203.648 288 336c0 3.904 0.128 7.744 0.384 11.584a32 32 0 0 1-29.856 34.016c-17.824 0.512-32.864-12.224-34.016-29.856A233.088 233.088 0 0 1 224 336C224 168.384 396.256 32 608 32s384 136.384 384 304c0 68.16-28.832 134.08-81.568 187.392 18.048 94.624 47.008 168 47.296 168.736a31.936 31.936 0 0 1-7.136 34.496A31.936 31.936 0 0 1 928 736z" p-id="2314"></path><path d="M96 992a32 32 0 0 1-29.76-43.84c0.32-0.736 29.248-74.112 47.296-168.736C60.832 726.048 32 660.16 32 592 32 424.384 204.256 288 416 288s384 136.384 384 304S627.744 896 416 896c-47.296 0-93.504-6.848-137.76-20.352-53.664 51.936-110.88 90.272-170.368 114.048A31.36 31.36 0 0 1 96 992z m320-640c-176.448 0-320 107.648-320 240 0 55.296 26.144 109.44 73.632 152.48 8.096 7.36 11.904 18.336 10.016 29.152a1108.832 1108.832 0 0 1-28.416 121.952c32.768-20.896 64.672-47.264 95.456-78.784a32 32 0 0 1 33.6-7.808c43.136 15.264 88.8 23.008 135.712 23.008 176.448 0 320-107.648 320-240S592.448 352 416 352z" p-id="2315"></path><path d="M192 624a32 32 0 0 1-32-32c0-84.992 102.88-176 256-176a32 32 0 1 1 0 64c-117.216 0-192 66.336-192 112a32 32 0 0 1-32 32z" p-id="2316"></path></svg>' +
        '    </div>' +
        '    <div id="chat-window" style="display: none" class="pvr_chat_wrapper">' +
        '        <div class="pvr_chat_content">' +
        '            <div id="close-chat-window" class="close_chat">' +
        '                <svg width="14" height="14" t="1606895276516" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3061" width="200" height="200"><path d="M925.468404 822.294069 622.19831 512.00614l303.311027-310.331931c34.682917-27.842115 38.299281-75.80243 8.121981-107.216907-30.135344-31.369452-82.733283-34.259268-117.408013-6.463202L512.000512 399.25724 207.776695 87.993077c-34.675754-27.796066-87.272669-24.90625-117.408013 6.463202-30.178323 31.414477-26.560936 79.375815 8.121981 107.216907l303.311027 310.331931L98.531596 822.294069c-34.724873 27.820626-38.341237 75.846432-8.117888 107.195418 30.135344 31.43699 82.72919 34.326806 117.408013 6.485715l304.178791-311.219137 304.177767 311.219137c34.678824 27.841092 87.271646 24.951275 117.408013-6.485715C963.808618 898.140501 960.146205 850.113671 925.468404 822.294069z" p-id="3062"></path></svg>' +
        '            </div>' +
        '            <div class="chat_header">' +
        '                <div id="avatar-box" class="pvr-user-w with-status status-green">' +
        '                    <div class="pvr-user-avatar-w">' +
        '                        <div class="user-avatar">' +
        '                            <svg width="38" height="38" t="1606895679059" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="16436" width="200" height="200"><path d="M336 960H209.1l-1.1-33.39C208 794.83 344.11 688 512 688s304 106.83 304 238.61L814.9 960H336z" fill="#FFFFFF" p-id="16437"></path><path d="M822.64 968H201.36L200 926.61c0-66.24 32.71-128.4 92.1-175C350.92 705.43 429 680 512 680s161.08 25.43 219.9 71.59c59.39 46.62 92.1 108.78 92.1 175v0.26z m-605.8-16h590.32l0.84-25.52C807.9 799.38 675.16 696 512 696S216.1 799.38 216 926.48z" p-id="16438"></path><path d="M560 608h-96v90s0 86 48 86 48-86 48-86z" fill="#FFCEBF" p-id="16439"></path><path d="M512 792c-16.46 0-38.44-8.47-49.71-48.85A188.19 188.19 0 0 1 456 698v-98h112v98a188.19 188.19 0 0 1-6.29 45.15C550.44 783.53 528.46 792 512 792z m-40-176v82c0 0.78 0.55 78 40 78 15.88 0 27.42-12.5 34.29-37.15A173.86 173.86 0 0 0 552 698v-82z" fill="#222222" p-id="16440"></path><path d="M432 832l-16-144 48-16 48 112-80 48z" fill="#FFFFFF" p-id="16441"></path><path d="M432 840a8 8 0 0 1-7.95-7.12l-16-144a8 8 0 0 1 5.42-8.47l48-16a8 8 0 0 1 9.88 4.44l48 112a8 8 0 0 1-3.23 10l-80 48A8 8 0 0 1 432 840z m-7.33-146.46l13.9 125.19 63.33-38-42.35-98.81z" fill="#222222" p-id="16442"></path><path d="M592 832l16-144-48-16-48 112 80 48z" fill="#FFFFFF" p-id="16443"></path><path d="M592 840a8 8 0 0 1-4.12-1.14l-80-48a8 8 0 0 1-3.23-10l48-112a8 8 0 0 1 9.88-4.44l48 16a8 8 0 0 1 5.42 8.47l-16 144a8 8 0 0 1-8 7.12z m-69.9-59.27l63.33 38 13.9-125.19-34.88-11.62z" fill="#222222" p-id="16444"></path><path d="M720 960l-16-48 48-144-80-48-96 240H464L352 720l-80 48 48 144-16 48" fill="#231E37" p-id="16445"></path><path d="M720 968a8 8 0 0 1-7.59-5.47l-16-48a8 8 0 0 1 0-5.06l46-137.91-66.44-39.87L583.43 963a8 8 0 0 1-7.43 5H464a8 8 0 0 1-7.25-4.62l-108.25-232-66.88 40.13 46 137.91a8 8 0 0 1 0 5.06l-16 48a8 8 0 0 1-15.18-5.06L311.57 912l-47.16-141.47a8 8 0 0 1 3.47-9.39l80-48a8 8 0 0 1 11.37 3.48L469.09 952h101.49l94-235a8 8 0 0 1 11.55-3.89l80 48a8 8 0 0 1 3.47 9.39L712.43 912l15.16 45.47A8 8 0 0 1 720 968z" fill="#222222" p-id="16446"></path><path d="M251.81 395.81m-68.19 0a68.19 68.19 0 1 0 136.38 0 68.19 68.19 0 1 0-136.38 0Z" fill="#222222" p-id="16447"></path><path d="M741.85 395.84m-58.15 0a58.15 58.15 0 1 0 116.3 0 58.15 58.15 0 1 0-116.3 0Z" fill="#FFCEBF" p-id="16448"></path><path d="M741.85 462A66.15 66.15 0 1 1 808 395.84 66.23 66.23 0 0 1 741.85 462z m0-116.29A50.15 50.15 0 1 0 792 395.84a50.2 50.2 0 0 0-50.15-50.14z" fill="#222222" p-id="16449"></path><path d="M742 296.6l0.84 111.7c0 136.84-104.13 247.77-232.59 247.77S277.69 545.14 277.69 408.3l0.84-111.7S288 80 512 80s230 216.6 230 216.6z" fill="#FFCEBF" p-id="16450"></path><path d="M510.29 664.07c-64.37 0-124.85-26.67-170.29-75.07s-70.31-112.48-70.31-180.7l0.85-112c0.1-2.29 2.8-56.69 33.75-111.86 18.29-32.63 42.68-58.64 72.48-77.32C413.88 83.8 459.38 72 512 72s98 11.8 134.72 35.08c29.54 18.71 53.53 44.74 71.33 77.39 30.08 55.2 31.92 109.62 32 111.91v0.16l0.84 111.7c0 68.28-25 132.45-70.29 180.74s-105.95 75.09-170.31 75.09zM286.53 296.83l-0.84 111.53c0 132.15 100.75 239.71 224.6 239.71S734.88 540.51 734.88 408.3L734 296.76c-0.08-2.12-2.33-54.09-30.32-105.14C666 122.86 601.51 88 512 88s-154.65 34.89-193.47 103.71c-28.59 50.67-31.84 102.29-32 105.12z" fill="#222222" p-id="16451"></path><path d="M298.84 453.99a37 21.15 0 1 0 74 0 37 21.15 0 1 0-74 0Z" fill="#FFA096" p-id="16452"></path><path d="M647.73 453.99a37 21.15 0 1 0 74 0 37 21.15 0 1 0-74 0Z" fill="#FFA096" p-id="16453"></path><path d="M407.66 356.19l-8.2 8.2a31.89 31.89 0 0 0-8-3.34v-11.59a8 8 0 0 0-16 0v11.59a31.82 31.82 0 0 0-8 3.34l-8.2-8.2a8 8 0 0 0-11.32 11.32l8.21 8.2a31.72 31.72 0 1 0 54.71 0l8.2-8.2a8 8 0 0 0-11.31-11.32zM672.72 356.19a8 8 0 0 0-11.32 0l-8.2 8.2a31.67 31.67 0 0 0-8-3.34v-11.59a8 8 0 1 0-16 0v11.59a31.89 31.89 0 0 0-8 3.34l-8.2-8.2a8 8 0 0 0-11.31 11.32l8.2 8.2a31.72 31.72 0 1 0 54.72 0l8.2-8.2a8 8 0 0 0-0.09-11.32z" fill="#222222" p-id="16454"></path><path d="M446.85 560.92s21.15 10.57 63.44 10.57 63.43-10.57 63.43-10.57" fill="#FFCEBF" p-id="16455"></path><path d="M510.29 579.49c-43.68 0-66.09-11-67-11.42a8 8 0 0 1 7.13-14.32c0.28 0.13 20.48 9.74 59.89 9.74s59.69-9.65 59.89-9.75a8 8 0 0 1 7.12 14.33c-0.95 0.47-23.32 11.42-67.03 11.42z" fill="#222222" p-id="16456"></path><path d="M704 144m-112 0a112 112 0 1 0 224 0 112 112 0 1 0-224 0Z" fill="#AA6E50" p-id="16457"></path><path d="M704 264a120 120 0 1 1 120-120 120.13 120.13 0 0 1-120 120z m0-224a104 104 0 1 0 104 104A104.11 104.11 0 0 0 704 40z" fill="#222222" p-id="16458"></path><path d="M758.86 325.71c0-148.47-98.67-257.14-246.86-257.14S265.14 177.24 265.14 325.71c0 0 298.29-10.28 360-102.85 72 92.57 133.72 102.85 133.72 102.85z" fill="#AA6E50" p-id="16459"></path><path d="M758.86 333.71a8.38 8.38 0 0 1-1.32-0.1c-2.57-0.43-62-11.21-132.26-97.72-30.72 36-98.55 63.23-201.92 81-81.88 14.08-157.19 16.8-157.94 16.83a8 8 0 0 1-8.28-8c0-37.57 6.14-73.1 18.24-105.6 11.94-32.08 29.37-60.38 51.8-84.14a233.51 233.51 0 0 1 80.66-55.49C439.63 67.27 474.67 60.57 512 60.57s72.37 6.7 104.16 19.91A233.51 233.51 0 0 1 696.82 136c22.43 23.76 39.86 52.06 51.8 84.14 12.1 32.5 18.24 68 18.24 105.6a8 8 0 0 1-8 8zM512 76.57c-137.94 0-235.08 98.57-238.75 240.76 21.7-1.12 82.55-5.06 147.56-16.24 105.35-18.13 173.7-46.71 197.68-82.67a8 8 0 0 1 13-0.47c36 46.29 69.13 71 90.58 83.56a165.41 165.41 0 0 0 28.64 13.57C746.05 174.1 649.21 76.57 512 76.57zM512 592c-20.5 0-38 9.34-44.87 22.48C352.78 605.16 264 541.23 264 464a8 8 0 0 0-16 0c0 45.41 26.23 87.87 73.86 119.55 39 26 88.74 42.16 142.41 46.75 2.55 14.88 22.94 8.21 47.73 8.21 26.51 0 48 7.63 48-11.63S538.51 592 512 592zM512 968a8 8 0 0 1-8-8V784a8 8 0 0 1 16 0v176a8 8 0 0 1-8 8z" fill="#222222" p-id="16460"></path></svg>' +
        '                        </div>' +
        '                    </div>' +
        '                    <div class="user-name">' +
        '                        <h6 class="user-title">' +
        '                            <a href="http://bootstrapmb.com/" target="_blank">PVR Tech Studio</a>' +
        '                        </h6>' +
        '                        <div id="online-status" class="status online">在线</div>' +
        '                    </div>' +
        '                </div>' +
        '            </div>' +
        '            <div class="chat-messages theme_2">' +
        '                <div class="message">' +
        '                    <div class="message-content">Lorem Ipsum is simply dummy.</div>' +
        '                </div>' +
        '                <div class="date-break">Mon 10:20am</div>' +
        '                <div class="message">' +
        '                    <div class="message-content">Lorem Ipsum is simply dummy text of the printing.</div>' +
        '                </div>' +
        '                <div class="message self">' +
        '                    <div class="message-content">Lorem Ipsum is simply dummy text of the printing and typesetting industry.</div>' +
        '                </div>' +
        '            </div>' +
        '            <div class="chat-controls">' +
        '                <input class="message-input" id="message-input" placeholder="Type your message here..." type="text">' +
        '                <div class="chat-extra">' +
        '                    <a href="javascript:void(0)">' +
        '                        <span class="extra-tooltip">Attach Document</span>' +
        '                        <svg width="20" height="20" t="1606896766621" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="17330" width="200" height="200"><path d="M633.553 251.102c15.993-12.795 38.385-12.795 55.978 1.6 15.993 15.993 15.993 38.384 0 54.378L347.264 647.747c-22.39 20.792-22.39 57.577 0 81.568 20.792 22.391 57.578 22.391 81.568 0l401.444-403.042c55.978-55.979 55.978-148.742 0-204.72s-148.742-55.979-204.72 0l-47.982 47.98-12.795 12.796-369.455 369.455c-91.165 91.165-91.165 236.708 0 327.872 91.164 91.165 236.707 91.165 327.872 0L894.25 511.8c6.397-3.199 9.596-7.997 12.795-12.795 15.993-15.994 38.385-15.994 54.378 0s15.994 38.385 0 54.379l-3.198 3.199c-3.2 1.599-6.398 6.397-9.597 9.596L577.574 934.035c-119.953 119.953-316.676 119.953-436.63 0s-119.952-316.676 0-436.63l430.233-431.83c86.366-86.367 227.111-86.367 315.077 0 86.366 86.366 86.366 227.11 0 315.076L483.21 783.694c-52.78 52.78-139.145 52.78-190.325 0-52.78-52.78-52.78-139.146 0-190.326l340.667-342.266z m0 0" fill="#333333" p-id="17331"></path></svg>' +
        '                    </a>' +
        '                    <a href="javascript:void(0)">' +
        '                        <span class="extra-tooltip">Insert Photo</span>' +
        '                        <svg width="25" height="25" t="1606896950947" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="31870" width="200" height="200"><path d="M695.4 217.4c13.7 0.1 24.9-11.1 25-24.8l0.3-65.6c0-13.9-11.1-25.1-24.9-25.1-13.8-0.1-25 11.1-25.1 24.9l-0.2 65.6c0 13.8 11.1 25 24.9 25z" p-id="31871"></path><path d="M512 601.8m-99.6 0a99.6 99.6 0 1 0 199.2 0 99.6 99.6 0 1 0-199.2 0Z" p-id="31872"></path><path d="M598.4 223.7c9.7 9.7 25.5 9.8 35.3 0.1 9.9-9.7 9.9-25.5 0.2-35.3L587.6 142c-9.8-9.8-25.6-9.9-35.4-0.1-9.7 9.6-9.8 25.5 0 35.3l46.2 46.5zM792.2 224.2l46.6-46.3c9.8-9.8 9.8-25.6 0.1-35.3-9.8-9.8-25.6-9.8-35.4-0.1L757 188.8c-9.8 9.7-9.8 25.5-0.1 35.3 9.8 9.8 25.5 9.8 35.3 0.1zM794.9 379.1c-22.5 0-40.7 18.2-40.7 40.7s18.2 40.7 40.7 40.7c22.6 0 40.8-18.2 40.8-40.7s-18.2-40.8-40.8-40.7z" p-id="31873"></path><path d="M858.6 257.6l-416.4 0.2v-41.4c0-12.7-9.5-23.1-21.1-23.1H218.4c-11.6 0-21.1 10.4-21.1 23.1v41.4h-36.8l-0.8 0.1C106.4 261 64.1 305 64 359.1v461.5c0.1 56 45.5 101.4 101.5 101.5h693.1c56-0.1 101.3-45.4 101.4-101.5V359.1c-0.1-56-45.5-101.3-101.4-101.5z m42.8 563c0 23.4-19.4 42.9-42.8 42.9H165.5c-23.4 0-42.9-19.5-43-42.9V359.1c-0.1-22.3 17.7-41.1 39.7-42.7l696.3-0.2c23.4 0.1 42.9 19.5 42.9 42.9v461.5z" p-id="31874"></path><path d="M512 380.8c-121.4 0-219.9 98.5-219.9 219.9 0 121.5 98.5 219.9 219.9 219.9 121.5 0 219.9-98.5 219.9-219.9 0-121.5-98.5-219.9-219.9-219.9z m0 379.2c-88 0-159.3-71.3-159.3-159.3S424 441.4 512 441.4s159.3 71.3 159.3 159.3S600 760 512 760z" p-id="31875"></path></svg>                    </a>' +
        '                    <a href="javascript:void(0)" class="change_chat_theme">' +
        '                        <span class="extra-tooltip">Change chat theme</span>' +
        '                        <svg width="20" height="20" t="1606897004693" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="32613" width="200" height="200"><path d="M915.692308 78.769231h-59.076923c-15.753846 0-29.538462 13.784615-29.538462 29.538461v137.846154c0 17.723077-9.846154 25.6-23.630769 13.784616-5.907692-7.876923-11.815385-13.784615-19.692308-19.692308-98.461538-98.461538-236.307692-139.815385-378.092308-112.246154-49.230769 9.846154-96.492308 29.538462-137.846153 57.107692-120.123077 78.769231-189.046154 206.769231-191.015385 344.615385-1.969231 106.338462 39.384615 212.676923 114.215385 289.476923 78.769231 82.707692 185.107692 128 299.323077 128 100.430769 0 194.953846-35.446154 269.784615-98.461538 13.784615-11.815385 13.784615-31.507692 1.969231-43.323077l-41.353846-41.353847c-9.846154-9.846154-27.569231-11.815385-39.384616-1.96923-70.892308 59.076923-167.384615 82.707692-263.876923 59.076923-25.6-5.907692-51.2-17.723077-74.830769-31.507693C230.4 720.738462 177.230769 590.769231 208.738462 460.8c5.907692-25.6 17.723077-51.2 31.507692-74.830769C295.384615 289.476923 391.876923 236.307692 494.276923 236.307692c78.769231 0 153.6 31.507692 208.738462 86.646154 9.846154 7.876923 17.723077 17.723077 23.630769 27.569231 5.907692 15.753846-7.876923 23.630769-25.6 23.630769h-137.846154c-15.753846 0-29.538462 13.784615-29.538462 29.538462v61.046154c0 15.753846 11.815385 27.569231 27.569231 27.56923h360.369231c13.784615 0 25.6-11.815385 25.6-25.6V108.307692C945.230769 92.553846 931.446154 78.769231 915.692308 78.769231z" p-id="32614"></path></svg>' +
        '                    </a>' +
        '                </div>' +
        '            </div>' +
        '        </div>' +
        '    </div>'

    function CHAT() {
        this.ws = ""
        this.init()
    }

    CHAT.prototype.init = function () {
        this.loadStyleSheet('./style.css')
        this.createContainer()
        this.connect()

        this.onInputMessage()
    }

    // 创建根容器
    CHAT.prototype.createContainer = function () {
        var container = document.createElement("div")
        container.innerHTML = html;
        container.setAttribute("id", "chat-_container")
        document.body.appendChild(container)

        if (window.localStorage.getItem("CHAT_WINDOW_OPEN_STATUS")) {
            document.getElementById("chat-window").removeAttribute("style")
        }
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
        var that = this
        var ws = new WebSocket("ws://127.0.0.1:8199/ws/handle")
        this.ws = ws

        var offline = function() {
            var status = document.getElementById("online-status")
            status.setAttribute("class", "status offline")
            status.innerText = "离线"
            document.getElementById("avatar-box").setAttribute("class", "pvr-user-w with-status status-red")
        }

        var online = function() {
            var status = document.getElementById("online-status")
            status.setAttribute("class", "status online")
            status.innerText = "在线"
            document.getElementById("avatar-box").setAttribute("class", "pvr-user-w with-status status-green")
        }

        ws.onopen = function () {
            console.log("WebSocket Server 连接成功！");
            online()
        };
        // ws连接关闭
        ws.onclose = function () {
            if (that.ws) {
                that.ws.close();
                that.ws = null;
            }
            offline()
        };
        // ws连接错误
        ws.onerror = function () {
            if (that.ws) {
                that.ws.close();
                that.ws = null;
            }
            offline()
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

        document.getElementById("chat-btn").addEventListener("click", this.toggleWindow)
        document.getElementById("close-chat-window").addEventListener("click", this.toggleWindow)
    }

    CHAT.prototype.toggleWindow = function () {
        var window = document.getElementById("chat-window")
        if (window.hasAttribute("style")) {
            window.removeAttribute("style")
            window.localStorage.setItem("CHAT_WINDOW_OPEN_STATUS", true)
        } else {
            window.setAttribute("style", "display:none")
            window.localStorage.setItem("CHAT_WINDOW_OPEN_STATUS", false)
        }
    }

    window.CHATGGPP = CHAT
})();

new CHATGGPP()

