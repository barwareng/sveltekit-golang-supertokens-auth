import { VITE_API_BASE_URL } from '$lib/env';
import { ClientResponseError } from './ClientResponseError';
import { TeamService } from './services/TeamService';
import { UserService } from './services/UserService';
import type { SendOptions } from './services/utils/options';

// list of known SendOptions keys (everything else is treated as query param)
const knownSendOptionsKeys = [
	'fetch',
	'headers',
	'body',
	'query',
	'params',
	'cache',
	'credentials',
	'headers',
	'integrity',
	'keepalive',
	'method',
	'mode',
	'redirect',
	'referrer',
	'referrerPolicy',
	'signal',
	'window'
];

export interface BeforeSendResult {
	[key: string]: any; // for backward compatibility
	url?: string;
	options?: { [key: string]: any };
}

export default class Client {
	fetchFunc: (url: RequestInfo | URL, config?: RequestInit) => Promise<Response>;
	baseUrl: string;
	beforeSend?: (url: string, options: SendOptions) => BeforeSendResult | Promise<BeforeSendResult>;

	afterSend?: (response: Response, data: any) => any;

	users: UserService;
	teams: TeamService;
	private cancelControllers: { [key: string]: AbortController } = {};
	private enableAutoCancellation = true;

	constructor(baseUrl = '/', fetchFunc = fetch) {
		this.fetchFunc = fetchFunc;
		this.baseUrl = baseUrl;
		this.users = new UserService(this);
		this.teams = new TeamService(this);
	}
	// Set the fetch function as Sveltekit's event.fetch or window.fetch
	setFetch(fetchFunc: typeof fetch): Client {
		this.fetchFunc = fetchFunc;
		return this;
	}
	/**
	 * Globally enable or disable auto cancellation for pending duplicated requests.
	 */
	autoCancellation(enable: boolean): Client {
		this.enableAutoCancellation = !!enable;

		return this;
	}

	/**
	 * Cancels single request by its cancellation key.
	 */
	cancelRequest(requestKey: string): Client {
		if (this.cancelControllers[requestKey]) {
			this.cancelControllers[requestKey].abort();
			delete this.cancelControllers[requestKey];
		}

		return this;
	}

	/**
	 * Cancels all pending requests.
	 */
	cancelAllRequests(): Client {
		for (const k in this.cancelControllers) {
			this.cancelControllers[k].abort();
		}

		this.cancelControllers = {};

		return this;
	}

	/**
	 * Builds a full client url by safely concatenating the provided path.
	 */
	buildUrl(path: string): string {
		let url = this.baseUrl;

		// construct an absolute base url if in a browser environment
		if (
			typeof window !== 'undefined' &&
			!!window.location &&
			!url.startsWith('https://') &&
			!url.startsWith('http://')
		) {
			url = window.location.origin?.endsWith('/')
				? window.location.origin.substring(0, window.location.origin.length - 1)
				: window.location.origin || '';

			if (!this.baseUrl.startsWith('/')) {
				url += window.location.pathname || '/';
				url += url.endsWith('/') ? '' : '/';
			}

			url += this.baseUrl;
		}

		// concatenate the path
		if (path) {
			url += url.endsWith('/') ? '' : '/'; // append trailing slash if missing
			url += path.startsWith('/') ? path.substring(1) : path;
		}

		return url;
	}

	/**
	 * Sends an api http request.
	 */
	async send<T = any>(path: string, options: SendOptions): Promise<T> {
		options = this.initSendOptions(path, options);

		// build url + path
		let url = this.buildUrl(path);

		if (this.beforeSend) {
			const result = Object.assign({}, await this.beforeSend(url, options));
			if (typeof result.url !== 'undefined' || typeof result.options !== 'undefined') {
				url = result.url || url;
				options = result.options || options;
			} else if (Object.keys(result).length) {
				// legacy behavior
				options = result as SendOptions;
				// console?.warn &&
				// 	console.warn(
				// 		'Deprecated format of beforeSend return: please use `return { url, options }`, instead of `return options`.'
				// 	);
			}
		}

		// serialize the query parameters
		if (typeof options.query !== 'undefined') {
			const query = this.serializeQueryParams(options.query);
			if (query) {
				url += (url.includes('?') ? '&' : '?') + query;
			}
			delete options.query;
		}

		// ensures that the json body is serialized
		if (
			this.getHeader(options.headers, 'Content-Type') == 'application/json' &&
			options.body &&
			typeof options.body !== 'string'
		) {
			options.body = JSON.stringify(options.body);
		}

		options.credentials = 'include';
		// send the request
		return this.fetchFunc(url, options)
			.then(async (response) => {
				let data: any = {};

				try {
					data = await response.json();
				} catch (_) {
					// all api responses are expected to return json
					// with the exception of the realtime event and 204
				}

				if (this.afterSend) {
					data = structuredClone(await this.afterSend(response, data));
				}

				if (response.status >= 400) {
					throw new ClientResponseError({
						url: response.url,
						status: response.status,
						data: data
					});
				}

				return data.data as T;
			})
			.catch((err) => {
				// wrap to normalize all errors
				throw new ClientResponseError(err);
			});
	}

	/**
	 * Shallow copy the provided object and takes care to initialize
	 * any options required to preserve the backward compatability.
	 *
	 * @param  {SendOptions} options
	 * @return {SendOptions}
	 */
	private initSendOptions(path: string, options: SendOptions): SendOptions {
		options = Object.assign({ method: 'GET' } as SendOptions, options);
		options.query = options.query || {};

		// auto convert the body to FormData, if needed
		options.body = this.convertToFormDataIfNeeded(options.body);

		// move unknown send options as query parameters
		for (const key in options) {
			if (knownSendOptionsKeys.includes(key)) {
				continue;
			}

			options.query[key] = options[key];
			delete options[key];
		}

		// add the json header, if not explicitly set
		// (for FormData body the Content-Type header should be skipped since the boundary is autogenerated)
		if (
			this.getHeader(options.headers, 'Content-Type') === null &&
			!this.isFormData(options.body)
		) {
			options.headers = Object.assign({}, options.headers, {
				'Content-Type': 'application/json'
			});
		}

		// handle auto cancelation for duplicated pending request
		if (this.enableAutoCancellation && options.requestKey !== null) {
			const requestKey = options.requestKey || (options.method || 'GET') + path;

			delete options.requestKey;

			// cancel previous pending requests
			this.cancelRequest(requestKey);

			const controller = new AbortController();
			this.cancelControllers[requestKey] = controller;
			options.signal = controller.signal;
		}

		return options;
	}

	/**
	 * Converts analyzes the provided body and converts it to FormData
	 * in case a plain object with File/Blob values is used.
	 */
	private convertToFormDataIfNeeded(body: any): any {
		if (
			typeof FormData === 'undefined' ||
			typeof body === 'undefined' ||
			typeof body !== 'object' ||
			body === null ||
			this.isFormData(body) ||
			!this.hasBlobField(body)
		) {
			return body;
		}

		const form = new FormData();

		for (const key in body) {
			const values = Array.isArray(body[key]) ? body[key] : [body[key]];
			for (const val of values) {
				form.append(key, val);
			}
		}

		return form;
	}

	/**
	 * Checks if the submitted body object has at least one Blob/File field.
	 */
	private hasBlobField(body: { [key: string]: any }): boolean {
		for (const key in body) {
			const values = Array.isArray(body[key]) ? body[key] : [body[key]];
			for (const v of values) {
				if (
					(typeof Blob !== 'undefined' && v instanceof Blob) ||
					(typeof File !== 'undefined' && v instanceof File)
				) {
					return true;
				}
			}
		}

		return false;
	}

	/**
	 * Extracts the header with the provided name in case-insensitive manner.
	 * Returns `null` if no header matching the name is found.
	 */
	private getHeader(headers: { [key: string]: string } | undefined, name: string): string | null {
		headers = headers || {};
		name = name.toLowerCase();

		for (const key in headers) {
			if (key.toLowerCase() == name) {
				return headers[key];
			}
		}

		return null;
	}

	/**
	 * Loosely checks if the specified body is a FormData instance.
	 */
	private isFormData(body: any): boolean {
		return (
			body &&
			// we are checking the constructor name because FormData
			// is not available natively in some environments and the
			// polyfill(s) may not be globally accessible
			(body.constructor.name === 'FormData' ||
				// fallback to global FormData instance check
				// note: this is needed because the constructor.name could be different in case of
				//       custom global FormData implementation, eg. React Native on Android/iOS
				(typeof FormData !== 'undefined' && body instanceof FormData))
		);
	}

	/**
	 * Serializes the provided query parameters into a query string.
	 */
	private serializeQueryParams(params: { [key: string]: any }): string {
		const result: Array<string> = [];
		for (const key in params) {
			if (params[key] === null) {
				// skip null query params
				continue;
			}

			const value = params[key];
			const encodedKey = encodeURIComponent(key);

			if (Array.isArray(value)) {
				// repeat array params
				for (const v of value) {
					result.push(encodedKey + '[]=' + encodeURIComponent(v));
				}
			} else if (value instanceof Date) {
				result.push(encodedKey + '=' + encodeURIComponent(value.toISOString()));
			} else if (typeof value !== null && typeof value === 'object') {
				result.push(encodedKey + '=' + encodeURIComponent(JSON.stringify(value)));
			} else {
				result.push(encodedKey + '=' + encodeURIComponent(value));
			}
		}

		return result.join('&');
	}
}

export const client = new Client(VITE_API_BASE_URL as string);
