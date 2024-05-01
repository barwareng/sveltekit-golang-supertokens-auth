import { BaseService } from '$lib/api/services/utils/BaseService';

export class TeamService extends BaseService {
	create(body: any) {
		return this.client.send('/teams/create', {
			method: 'POST',
			body
		});
	}
}
