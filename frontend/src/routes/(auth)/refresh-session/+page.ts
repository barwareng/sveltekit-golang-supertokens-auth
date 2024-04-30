import { browser } from '$app/environment';
import Session from 'supertokens-web-js/recipe/session';
import type { PageLoad } from './$types';
import { goto } from '$app/navigation';

export const load = (async ({ url }) => {
	if (browser) {
		Session.attemptRefreshingSession().then((success) => {
			if (success) {
				if (url.searchParams.has('redirectBack'))
					window.location.href = url.searchParams.get('redirectBack') as string;
				else window.location.href = '/';
			} else {
				goto('/signin');
			}
		});
	}
	return {};
}) satisfies PageLoad;
