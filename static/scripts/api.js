class Ajax {
    constructor() {
    }

    static encode(data) {
        const ret = [];
        for (let d in data) {
            ret.push(encodeURIComponent(d) + '=' + encodeURIComponent(data[d]));
        }
        return ret.join('&');
    }

    static async get(url, data) {
        try {
            url += '?' + this.encode(data);
            const response = await fetch(url);

            if (!response.ok) {
                const err = response.status + (response.statusText === '' ? '' : " = " + response.statusText);
                throw new Error(err);
            }

            const rdata = await response.json();
            return rdata;
        } catch (error) {
            Alert.show(false, error);
            console.error(error);
            return error;
        }
    }

    static async getw(url, body) {
        try {
            const response = await fetch(url, {
                method: 'get',
                headers: {
                    'Content-Type': 'Text'
                },
                body: body
            });

            if (!response.ok) {
                const err = response.status + (response.statusText === '' ? '' : " = " + response.statusText);
                throw new Error(err);
            }

            const rdata = await response.json();
            return rdata;
        } catch (error) {
            Alert.show(false, error);
            console.error(error);
            return error;
        }
    }

    static async post(url, data) {
        try {
            const response = await fetch(url, {
                method: 'post',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: data
            });

            if (!response.ok) {
                const err = response.status + (response.statusText === '' ? '' : " = " + response.statusText);
                throw new Error(err);
            }

            const rdata = await response.json();
            Alert.show(true, "Done.");
            return rdata;
        } catch (error) {
            Alert.show(false, error);
            console.error(error);
            return error;
        }
    }
}
