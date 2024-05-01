export const publicRoutes = new Set([
	'/signin',
	'/signin/verify-magic-link',
	'/signup',
	'/verify-email',
	'/verify-email/request-verification',
	'/verify-email/failed',
	'/verify-email',
	'/oauth-callback/google',
	'/oauth-callback/github',
	'/reset-password',
	'/reset-password/link-expired',
	'/reset-password/link-sent',
	'/reset-password/new',
	'/reset-password/success',
	'/refresh-session'
]);

export const isPublicRoute = (route: string) => {
	return publicRoutes.has(route);
};

export const onboardingAllowedRoutes = new Set(['/settings', '/support', '/onboarding']);
