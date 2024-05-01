import { browser } from '$app/environment';
import { getTeamCookie, setTeamCookie } from '$utils';
import type { ITeam } from '$utils/interfaces/teams.interface';
import type { LayoutLoad } from './$types';
import Session from 'supertokens-web-js/recipe/session';
export const load = (async () => {
	let teams: Partial<ITeam>[] = [];
	if (browser && (await Session.doesSessionExist())) {
		const accessTokenPayload = await Session.getAccessTokenPayloadSecurely();

		teams = accessTokenPayload.teams;
		if (!getTeamCookie() && teams.length > 0) {
			console.log('No team');
			setTeamCookie(teams[0].id!);
		}
	}
	return { teams };
}) satisfies LayoutLoad;
