const CONTENT_URL = "http://localhost:1061";


(() => {
    function load_includes(){
        const includes = document.getElementsByTagName('include');
        [].forEach.call(includes, i => {
            let filePath = i.getAttribute('src');
            fetch(filePath).then(file => {
                file.text().then(content => {
                    i.insertAdjacentHTML('afterend', content);
                    i.remove();
                });
            });
        });
    }

    function load_content() {
        const images = document.getElementsByTagName('img');
        [].forEach.call(images, i => {
            let filePath = i.getAttribute('src');
            i.setAttribute('src', CONTENT_URL + "/cdn" + filePath);
        });
    }

    load_includes();
    load_content();
})();




