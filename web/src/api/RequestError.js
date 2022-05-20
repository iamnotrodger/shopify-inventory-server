class RequestError extends Error {
	constructor(message, statusCode = null) {
		super(message);

		this.constructor = RequestError;
		this.__proto__ = RequestError.prototype;
		this.message = message;
		this.statusCode = statusCode;
	}

	static async parseResponse(response) {
		const text = await response.text();
		return new RequestError(text, response.status);
	}
}

export default RequestError;
