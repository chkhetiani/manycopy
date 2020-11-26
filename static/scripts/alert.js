class Alert {
    static show(success, text) {
        Alert.remove();
        clearTimeout(this._z)
        document.body.insertAdjacentHTML('beforeend',
            `<div id="__alert" class="alert alert-${success ? 'success' : 'danger'}">
				<div class="alert-text">${text}</div>
			</div>`);
        setTimeout(()=> {
            document.querySelector('#__alert').style.bottom = '30px';
            document.querySelector('#__alert').style.opacity = '1';
        }, 200);


        this._z = setTimeout(Alert.remove, 2500);
    }

    static remove() {
        if (document.querySelector(`#__alert`)) {
            document.querySelector(`#__alert`).remove();
        }
    }
}