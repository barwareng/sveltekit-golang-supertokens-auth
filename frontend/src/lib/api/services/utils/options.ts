export interface SendOptions extends RequestInit {
	[key: string]: any; // for backward compatibility

	/**
	 * Optional custom fetch function to use for sending the request.
	 */
	fetch?: (url: RequestInfo | URL, config?: RequestInit) => Promise<Response>;

	/**
	 * Custom headers to send with the requests.
	 */
	headers?: { [key: string]: string };

	/**
	 * The body of the request (serialized automatically for json requests).
	 */
	body?: any;

	/**
	 * Query params that will be appended to the request url.
	 */
	query?: { [key: string]: any };
}
